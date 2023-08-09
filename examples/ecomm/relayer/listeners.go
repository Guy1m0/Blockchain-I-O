package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/contracts/eth_auction"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
)

const (
	timeInfoFile = "../timer"
)

// sets up the event listener, and handleEvent, which contains
// the logic to execute when an event is received.
func startListeningForEvents(client *ecomm.AssetClient) error {
	events := []string{"AddAsset", "StartAuction", "EndAuction", "FinAuction"}
	var regs []fab.Registration
	var notifiers []<-chan *fab.CCEvent

	for _, eventID := range events {
		reg, notifier, err := client.Register(eventID)
		if err != nil {
			return fmt.Errorf("failed to register chaincode event for %s: %v", eventID, err)
		}
		regs = append(regs, reg)
		notifiers = append(notifiers, notifier)
	}

	defer func() {
		for _, reg := range regs {
			client.Unregister(reg)
		}
	}()

	for {
		select {
		case event := <-notifiers[0]:
			err := handleAddAssetEvent(string(event.Payload))
			if err != nil {
				return err
			}
		case event := <-notifiers[1]:
			err := handleStartAuctionEvent(string(event.Payload))
			if err != nil {
				return err
			}
		case event := <-notifiers[2]:
			err := handleEndAuctionEvent(string(event.Payload))
			if err != nil {
				return err
			}
		case event := <-notifiers[3]:
			err := handleFinAuctionEvent(string(event.Payload))
			if err != nil {
				return err
			}
		}
	}
}

// // Check if any auction's status is Ending,
// // Then publish related event
// func runAuctionListener() {
// 	for {
// 		// listen each new creating auction
// 		// and create a go route for each
// 		a := listenNewAuction() // listen new tx posted on auction contract
// 		//fmt.Println("Find New Auction!")
// 		onNewAuction(a)
// 	}
// }

// func listenNewAuction() *ecomm.Auction {
// 	lastID, err := assetClient.GetLastAuctionID()
// 	check(err)

// 	//fmt.Println("Trying to find new Auction")
// 	for {
// 		time.Sleep(1 * time.Second) // check auction list per second
// 		auctionID, err := assetClient.GetLastAuctionID()
// 		//fmt.Println("find auction ID: ", auctionID)
// 		check(err)
// 		if auctionID > lastID {
// 			t := time.Now()
// 			auction, err := assetClient.GetAuction(auctionID)
// 			check(err)

// 			payload, _ := json.Marshal(auction)
// 			cclib.LogEventToFile(logInfoFile, ecomm.KafkaReceivedEvent, payload, t, timeInfoFile)
// 			return auction
// 		}
// 	}
// }

// fabric
func onNewAuction(a *ecomm.Auction) {
	// Initialize auctionResults with new auction as input
	auctionResultsMu.Lock()
	auctionResults[a.ID] = &ecomm.FinalizeAuctionArgs{
		AuctionID: a.ID,
	}
	auctionResultsMu.Unlock()

	//go fabric_listener(a)
	go eth_listener(a)
	go quo_listener(a)
	// check target auction's status change

	//cclib.LastEventTimestamp.Set(t)

}

// // Only for Auction Ending
// func fabric_listener(a *ecomm.Auction) {
// 	for {
// 		time.Sleep(1 * time.Second)
// 		auction, err := assetClient.GetAuction(a.ID)
// 		check(err)
// 		if auction.Status == "Ending" {
// 			break
// 		}
// 	}

// 	t := time.Now()

// 	payload, _ := json.Marshal(a)
// 	cclib.LogEventToFile(logInfoFile, ecomm.KafkaReceivedEvent, payload, t, timeInfoFile)

// 	// initialize
// 	authT, err := cclib.NewTransactor(root_key, password)
// 	check(err)

// 	eth_auction_addr := common.HexToAddress(a.EthAddr)
// 	eth_auction_contract, err := eth_auction.NewEthAuction(eth_auction_addr, ethClient)
// 	check(err)

// 	quo_auction_addr := common.HexToAddress(a.QuorumAddr)
// 	quo_auction_contract, err := eth_auction.NewEthAuction(quo_auction_addr, quoClient)
// 	check(err)

// 	// Check highest bid
// 	eth_highestBid, _ := eth_auction_contract.HighestBid(&bind.CallOpts{})
// 	eth_highestBidder, _ := eth_auction_contract.HighestBidder(&bind.CallOpts{})
// 	quo_highestBid, _ := quo_auction_contract.HighestBid(&bind.CallOpts{})
// 	quo_highestBidder, _ := quo_auction_contract.HighestBidder(&bind.CallOpts{})

// 	winner_platform := "eth"
// 	a.HighestBid = int(eth_highestBid.Int64())
// 	a.HighestBidder = eth_highestBidder.Hex()

// 	if eth_highestBid.Cmp(quo_highestBid) < 0 {
// 		winner_platform = "quo"
// 		a.HighestBid = int(quo_highestBid.Int64())
// 		a.HighestBidder = quo_highestBidder.Hex()
// 	}

// 	// Change Auction Contract on Eth
// 	tx, _ := eth_auction_contract.EndAuction(authT)
// 	receipt := ecomm.WaitTx(ethClient, tx, fmt.Sprintf("Change Auction %s status to 'ENDING'", eth_auction_addr))

// 	payload, err = json.Marshal(&ecomm.Tx{
// 		Platform: "eth",
// 		Type:     "EndAuction",
// 		Receipt:  receipt,
// 	})
// 	check(err)

// 	t = time.Now()
// 	cclib.LogEventToFile(logInfoFile, ecomm.TransactionMinedEvent, payload, t, timeInfoFile)
// 	cclib.LastEventTimestamp.Set(t, timeInfoFile)

// 	// Change Auction Contract on Quo
// 	tx, _ = quo_auction_contract.EndAuction(authT)
// 	receipt = ecomm.WaitTx(quoClient, tx, fmt.Sprintf("Change Auction %s status to 'ENDING'", quo_auction_addr))

// 	payload, err = json.Marshal(&ecomm.Tx{
// 		Platform: "quo",
// 		Type:     "EndAuction",
// 		Receipt:  receipt,
// 	})
// 	check(err)

// 	t = time.Now()
// 	cclib.LogEventToFile(logInfoFile, ecomm.TransactionMinedEvent, payload, t, timeInfoFile)

// 	// Abort one platform's Auction
// 	if winner_platform == "eth" {
// 		tx, _ = quo_auction_contract.Abort(authT)
// 		ecomm.WaitTx(quoClient, tx, fmt.Sprintf("Close Auction since winner on other platform %s", winner_platform))
// 		winner_listener(*eth_auction_contract, a)
// 	} else {
// 		tx, _ = eth_auction_contract.Abort(authT)
// 		ecomm.WaitTx(ethClient, tx, fmt.Sprintf("Close Auction since winner on other platform %s", winner_platform))
// 		winner_listener(*quo_auction_contract, a)
// 	}

// 	a.HighestBidPlatform = winner_platform
// 	payload, _ = json.Marshal(a)
// 	//ccsvc.Publish(ecomm.AuctionEndingEvent, payload)

// 	t = time.Now()
// 	cclib.LogEventToFile(logInfoFile, ecomm.AuctionEndingEvent, payload, t, timeInfoFile)

// }

// Publish Bidding Event if new highest bid on Eth
func eth_listener(a *ecomm.Auction) {
	//t := time.Now()

	client := ethClient
	// authT, err := cclib.NewTransactor(root_key, password)
	// check(err)

	auction_addr := common.HexToAddress(a.EthAddr)
	auction_contract, err := eth_auction.NewEthAuction(auction_addr, client)
	check(err)

	// Check winner
	highestBid, err := auction_contract.HighestBid(&bind.CallOpts{})
	check(err)
	highestBidder, _ := auction_contract.HighestBidder(&bind.CallOpts{})

	result := ecomm.AuctionResult{
		Platform:      "eth",
		HostAuctionID: a.ID,
		AuctionAddr:   auction_addr.Hex(),
	}

	result.HighestBid = int(highestBid.Int64())
	result.HighestBidder = highestBidder.Hex()

	// if result.HighestBid > 0 {
	// 	payload, _ := json.Marshal(result)

	// }

	// Update Asset contract if find new highest
	for {
		// Check once per second
		time.Sleep(1 * time.Second)
		t := time.Now()

		// exit if auction is not accessable
		status, _ := auction_contract.Status(&bind.CallOpts{})
		if status != "running" {
			return
		}

		highestBid_, _ := auction_contract.HighestBid(&bind.CallOpts{})

		if highestBid_.Cmp(highestBid) > 0 {
			highestBidder, _ = auction_contract.HighestBidder(&bind.CallOpts{})
			highestBid = highestBid_

			result.HighestBid = int(highestBid.Int64())
			result.HighestBidder = highestBidder.Hex()

			// update Asset contract
			// if result.HighestBid > a.HighestBid{
			// 	a.HighestBid = result.HighestBid
			// 	a.HighestBidPlatform = result.Platform
			// 	a.HighestBidder = result.HighestBidder

			// }

			// Check later if really need this
			auctionResultsMu.Lock()
			auctionResults[a.ID].EthResult.AuctionResult = result
			auctionResultsMu.Unlock()

			payload, _ := json.Marshal(result)
			//ccsvc.Publish(ecomm.BiddingAuctionEvent, payload)
			cclib.LogEventToFile(logInfoFile, ecomm.BiddingAuctionEvent, payload, t, timeInfoFile)
			//cclib.LastEventTimestamp.Set(t)
		}
	}
}

// Publish Bidding Event if new highest bid on Quo
func quo_listener(a *ecomm.Auction) {
	//t := time.Now()

	client := quoClient
	// authT, err := cclib.NewTransactor(root_key, password)
	// check(err)

	auction_addr := common.HexToAddress(a.QuorumAddr)
	auction_contract, err := eth_auction.NewEthAuction(auction_addr, client)
	check(err)

	// Check winner
	highestBid, err := auction_contract.HighestBid(&bind.CallOpts{})
	check(err)
	highestBidder, _ := auction_contract.HighestBidder(&bind.CallOpts{})

	result := ecomm.AuctionResult{
		Platform:      "quo",
		HostAuctionID: a.ID,
		AuctionAddr:   auction_addr.Hex(),
	}

	result.HighestBid = int(highestBid.Int64())
	result.HighestBidder = highestBidder.Hex()

	// if result.HighestBid > 0 {
	// 	payload, _ := json.Marshal(result)

	// }

	// Update Asset contract if find new highest
	for {
		// Check once per second
		time.Sleep(1 * time.Second)
		t := time.Now()

		// exit if auction is not accessable
		status, _ := auction_contract.Status(&bind.CallOpts{})
		if status != "running" {
			return
		}

		highestBid_, _ := auction_contract.HighestBid(&bind.CallOpts{})

		if highestBid_.Cmp(highestBid) > 0 {
			highestBidder, _ = auction_contract.HighestBidder(&bind.CallOpts{})
			highestBid = highestBid_

			result.HighestBid = int(highestBid.Int64())
			result.HighestBidder = highestBidder.Hex()

			// Check later if really need this
			auctionResultsMu.Lock()
			auctionResults[a.ID].QuorumResult.AuctionResult = result
			auctionResultsMu.Unlock()

			payload, _ := json.Marshal(result)
			//ccsvc.Publish(ecomm.BiddingAuctionEvent, payload)
			cclib.LogEventToFile(logInfoFile, ecomm.BiddingAuctionEvent, payload, t, timeInfoFile)
			//cclib.LastEventTimestamp.Set(t)
		}
	}
}

func winner_listener(contract eth_auction.EthAuction, a *ecomm.Auction) {
	//client := ethClient

	// authT, err := cclib.NewTransactor(root_key, password)
	// check(err)

	for {
		time.Sleep(1 * time.Second)
		status, _ := contract.Status(&bind.CallOpts{})

		if status == "ended" {
			t := time.Now()
			payload, _ := json.Marshal(a)
			cclib.LogEventToFile(logInfoFile, ecomm.ProceedAuctionResultEvent,
				payload, t, timeInfoFile)
			return
			//assetClient.FinalizeAuction()
		}

	}

}
