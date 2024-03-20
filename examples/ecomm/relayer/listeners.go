package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/contracts/cb1p_auction"
	"github.com/Guy1m0/Blockchain-I-O/contracts/english_auction"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
)

const (
	timeInfoFile = "../timer"
)

// sets up the event listener, and handleEvent, which contains
// the logic to execute when an event is received.
// Fabric relayer
func startFabricListener(client *ecomm.AssetClient) error {
	events := []string{"AddAsset", "StartAuction", "CloseAuction", "CancelAuction", "AuctionClosed"}
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
			err := handleAddAssetEvent(event.Payload)
			if err != nil {
				return err
			}
		case event := <-notifiers[1]:
			err := handleStartAuctionEvent(event.Payload)
			if err != nil {
				return err
			}
		case event := <-notifiers[2]:
			err := handleCloseAuctionEvent(string(event.Payload))
			if err != nil {
				return err
			}
		case event := <-notifiers[3]:
			err := handleCancelAuctionEvent(string(event.Payload))
			if err != nil {
				return err
			}
		case event := <-notifiers[4]:
			err := handleAuctionClosedEvent(string(event.Payload))
			if err != nil {
				return err
			}
		}
	}
}

func startContractListener(contract_info ecomm.ContractInfo) {
	// For English Auction
	abi := english_auction.EnglishAuctionABI
	eth_english_auction_contract, _ = english_auction.NewEnglishAuction(contract_info.EnglishAuction.EthAddr, ethClient)
	go startListeningForAuctionEvents(contract_info.EnglishAuction.EthAddr, abi, "eth")

	quo_english_auction_contract, _ = english_auction.NewEnglishAuction(contract_info.EnglishAuction.QuoAddr, quoClient)
	go startListeningForAuctionEvents(contract_info.EnglishAuction.QuoAddr, abi, "quo")

	// // For Dutch Auction
	// abi = dutch_auction.DutchAuctionABI
	// eth_dutch_auction_contract, _ = dutch_auction.NewDutchAuction(contract_info.DutchAuction.EthAddr, ethClient)
	// go startListeningForAuctionEvents(contract_info.DutchAuction.EthAddr, abi, *ethClient)

	// quo_dutch_auction_contract, _ = dutch_auction.NewDutchAuction(contract_info.DutchAuction.QuoAddr, quoClient)
	// go startListeningForAuctionEvents(contract_info.DutchAuction.QuoAddr, abi, *quoClient)

	// For Close Bid 1st Price Auction
	abi = cb1p_auction.Cb1pAuctionABI
	eth_cb1p_contract, _ = cb1p_auction.NewCb1pAuction(contract_info.Cb1pAuction.EthAddr, ethClient)
	go startListeningForAuctionEvents(contract_info.Cb1pAuction.EthAddr, abi, "eth")

	quo_cb1p_contract, _ = cb1p_auction.NewCb1pAuction(contract_info.Cb1pAuction.QuoAddr, quoClient)
	go startListeningForAuctionEvents(contract_info.Cb1pAuction.QuoAddr, abi, "quo")

	// // For Close Bid 2nd Price Auction
	// abi = cb2p_auction.Cb2pAuctionABI
	// eth_cb2p_contract, _ = cb2p_auction.NewCb2pAuction(contract_info.Cb2pAuction.EthAddr, ethClient)
	// go startListeningForAuctionEvents(contract_info.Cb2pAuction.EthAddr, abi, *ethClient)

	// quo_cb2p_contract, _ = cb2p_auction.NewCb2pAuction(contract_info.Cb2pAuction.QuoAddr, quoClient)
	// go startListeningForAuctionEvents(contract_info.Cb2pAuction.QuoAddr, abi, *quoClient)
}

// client := ethclient

// contract, err := english_auction.NewEnglishAuction(english_auction_info.EthAddr, client)

// eth_english_auction_contract, err := english_auction.NewEnglishAuction(english_auction_info.EthAddr, ethClient)
// quo_english_auction_contract, err := english_auction.NewEnglishAuction(english_auction_info.QuoAddr, quoClient)

// eth_cb1p_contract, err  := cb1p_auction.NewCb1pAuction(cb1p_auction_info.EthAddr, ethClient)
// quo_cb1p_contract, err  := cb1p_auction.NewCb1pAuction(cb1p_auction_info.QuoAddr, quoClient)

func startListeningForAuctionEvents(auction_addr common.Address, auction_abi string, platform string) {
	ethclient.Dial(fmt.Sprintf("http://%s:8545", "localhost"))
	wsURL := "ws://localhost:8545"
	if platform == "quo" {
		wsURL = "ws://localhost:8546"
	}
	client, err := ethclient.Dial(wsURL)
	check(err)

	query := ethereum.FilterQuery{Addresses: []common.Address{auction_addr}}
	contractAbi, err := abi.JSON(strings.NewReader(string(auction_abi)))
	check(err)

	// Listen to "HighestBidIncreased" and "DecisionMade" events
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	check(err)

	for vLog := range logs {
		switch vLog.Topics[0].Hex() {
		case contractAbi.Events["HighestBidIncreased"].ID.Hex():
			t := time.Now()
			var event ecomm.HighestBidIncreased
			err := contractAbi.UnpackIntoInterface(&event, "HighestBidIncreased", vLog.Data)
			check(err)

			result := ecomm.Bid{
				Platform: platform,
				// AuctionID:   auction_id,
				AuctionAddr: auction_addr,
			}
			// call handler
			handleHighestBidIncreasedEvent(event, result, t)
			fmt.Printf("New highest bid for %s: %s by bidder %s\n", event.AuctionType, event.BidAmount.String(), event.Bidder.Hex())
		case contractAbi.Events["WithdrawBid"].ID.Hex():
			t := time.Now()
			var event ecomm.WithdrawBid
			err := contractAbi.UnpackIntoInterface(&event, "WithdrawBid", vLog.Data)
			check(err)

			result := ecomm.Bid{
				Platform: platform,
				// AuctionID:   auction_id,
				AuctionAddr: auction_addr,
			}
			// call handler
			handleWithdrawBidEvent(event, result, t)

		case contractAbi.Events["NewBidHash"].ID.Hex():
			t := time.Now()
			var event ecomm.NewBidHash
			err := contractAbi.UnpackIntoInterface(&event, "NewBidHash", vLog.Data)
			check(err)

			result := ecomm.BidHash{
				Platform: platform,
				// AuctionID:   auction_id,
				AuctionAddr: auction_addr,
			}
			// call handler
			handleNewBidHashEvent(event, result, t)

		case contractAbi.Events["DecisionMade"].ID.Hex():
			var event ecomm.DecisionMade
			t := time.Now()

			err := contractAbi.UnpackIntoInterface(&event, "DecisionMade", vLog.Data)
			check(err)

			// call handler
			handleDecisionMadeEvent(event, t)
			fmt.Printf("Decision Made: Winner %s, Amount %s, ID %s\n", event.Winner.Hex(), event.Amount.String(), event.Id)

			// Unsubscribe and break out of the loop
			sub.Unsubscribe()
			return
			// break
		}
	}
}
