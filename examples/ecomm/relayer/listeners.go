package relayer

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/contracts/eth_auction"
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
func startListeningForAssetEvents(client *ecomm.AssetClient) error {
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

// fabric
func onNewAuction(a *ecomm.Auction) {
	// Initialize auctionResults with new auction as input
	auctionResultsMu.Lock()
	auctionResults[a.ID] = &ecomm.FinalizeAuctionArgs{
		AuctionID: a.ID,
	}
	auctionResultsMu.Unlock()

	go startListeningForAuctionEvents(a.ID, a.EthAddr, "eth")
	go startListeningForAuctionEvents(a.ID, a.QuorumAddr, "quo")

}

func startListeningForAuctionEvents(auction_id int, address string, platform string) {
	ethclient.Dial(fmt.Sprintf("http://%s:8545", "localhost"))
	wsURL := "ws://localhost:8545"
	if platform == "quo" {
		wsURL = "ws://localhost:8546"
	}
	client, err := ethclient.Dial(wsURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	auction_addr := common.HexToAddress(address)
	query := ethereum.FilterQuery{Addresses: []common.Address{auction_addr}}

	contractAbi, err := abi.JSON(strings.NewReader(string(eth_auction.EthAuctionABI)))
	check(err)

	// Listen to "HighestBidIncreased" and "DecisionMade" events
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	check(err)

	for vLog := range logs {
		switch vLog.Topics[0].Hex() {
		case contractAbi.Events["HighestBidIncreased"].ID.Hex():
			t := time.Now()
			var event ecomm.HighestBidIncreasedEvent
			err := contractAbi.UnpackIntoInterface(&event, "HighestBidIncreased", vLog.Data)
			check(err)

			result := ecomm.AuctionResult{
				Platform:    platform,
				AuctionID:   auction_id,
				AuctionAddr: address,
			}
			// call handler
			handleHighestBidIncreasedEvent(event, result, t)
			fmt.Printf("New highest bid: %s by %s\n", event.Amount.String(), event.Bidder.Hex())

		case contractAbi.Events["DecisionMade"].ID.Hex():
			var event ecomm.DecisionMadeEvent
			t := time.Now()

			err := contractAbi.UnpackIntoInterface(&event, "DecisionMade", vLog.Data)
			check(err)

			// call handler
			handleDecisionMadeEvent(event, t)
			fmt.Printf("Decision Made: Winner %s, Amount %s, ID %s\n", event.Winner.Hex(), event.Amount.String(), event.Id)

			// Unsubscribe and break out of the loop
			sub.Unsubscribe()
			break
		}
	}

}
