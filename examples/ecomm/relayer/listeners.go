package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/contracts/eth_auction"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/common"
)

// Check if any auction's status is Ending,
// Then publish related event
func runAuctionListener() {
	for {
		// listen each new creating auction
		// and create a go route for each
		a := listenNewAuction() // listen new tx posted on auction contract
		go onNewAuction(a)
	}
}

func onNewAuction(a *ecomm.Auction) {
	auctionResults[a.ID] = &ecomm.FinalizeAuctionArgs{
		AuctionID: a.ID,
	}

	// check target auction's status change
	for {
		time.Sleep(1 * time.Second)
		auction, err := assetClient.GetAuction(a.ID)
		check(err)
		if auction.Status == "Ending" {
			break
		}
	}

	authT, err := cclib.NewTransactor(root_key, password)
	check(err)

	// Change Auction Contract on Eth
	auction_addr := common.HexToAddress(a.EthAddr)
	auction_contract, err := eth_auction.NewEthAuction(auction_addr, ethClient)
	check(err)

	tx, _ := auction_contract.EndAuction(authT)
	receipt := ecomm.WaitTx(ethClient, tx, fmt.Sprintf("Change Auction %s status to 'ENDING'", auction_addr))

	payload, err := json.Marshal(&ecomm.Tx{
		Platform: "eth",
		Type:     "EndAuction",
		Receipt:  receipt,
	})
	check(err)

	t := time.Now()
	cclib.LogEventToFile(logInfoFile, ecomm.AuctionEndingEvent, payload, t)
	cclib.LastEventTimestamp.Set(t)

	// Change Auction Contract on Eth
	auction_addr = common.HexToAddress(a.QuorumAddr)
	auction_contract, err = eth_auction.NewEthAuction(auction_addr, quoClient)
	check(err)

	tx, _ = auction_contract.EndAuction(authT)
	receipt = ecomm.WaitTx(quoClient, tx, fmt.Sprintf("Change Auction %s status to 'ENDING'", auction_addr))

	payload, err = json.Marshal(&ecomm.Tx{
		Platform: "quo",
		Type:     "EndAuction",
		Receipt:  receipt,
	})
	check(err)

	b, _ := json.Marshal(a)
	ccsvc.Publish(ecomm.AuctionEndingEvent, b)

	t = time.Now()
	cclib.LogEventToFile(logInfoFile, ecomm.AuctionEndingEvent, payload, t)
	//cclib.LastEventTimestamp.Set(t)

}

func listenNewAuction() *ecomm.Auction {
	lastID, err := assetClient.GetLastAuctionID()
	check(err)

	for {
		time.Sleep(1 * time.Second) // check auction list per second
		auctionID, err := assetClient.GetLastAuctionID()
		check(err)
		if auctionID > lastID {
			auction, err := assetClient.GetAuction(auctionID)
			check(err)
			return auction
		}
	}
}
