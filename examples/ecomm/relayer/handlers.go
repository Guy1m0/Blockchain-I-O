package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
)

var mutex sync.Mutex

func handleAddAssetEvent(eventPayload string) error {

	t := time.Now()
	payload, _ := json.Marshal(eventPayload)
	cclib.LogEventToFile(logInfoFile, ecomm.RelayerDetectedEvent,
		payload, t, timeInfoFile)

	prefix := "Asset added: "

	if !strings.HasPrefix(eventPayload, prefix) {
		log.Fatalf("Received unexpected event: %s\n", eventPayload)
		return nil
	}

	assetID := strings.TrimPrefix(eventPayload, prefix)
	asset, err := assetClient.GetAsset(assetID)
	check(err)

	fmt.Println("Starting auction")
	fmt.Println("[ethereum] Deploying auction")

	t = time.Now()
	cclib.LastEventTimestamp.Set(t, timeInfoFile)

	ethAddr, receipt_eth := ecomm.DeployCrossChainAuction(ethClient, eth_ERC20, root_key)
	payload, _ = json.Marshal(&ecomm.Tx{
		Platform: "eth",
		Type:     "newAuction",
		Receipt:  receipt_eth,
	})
	check(err)

	t = time.Now()
	cclib.LogEventToFile(logInfoFile, ecomm.TransactionMinedEvent, payload, t, timeInfoFile)
	cclib.LastEventTimestamp.Set(t, timeInfoFile)

	fmt.Println("[quorum] Deploying auction")
	quoAddr, receipt_quo := ecomm.DeployCrossChainAuction(quoClient, quo_ERC20, root_key)
	payload, err = json.Marshal(&ecomm.Tx{
		Platform: "quo",
		Type:     "newAuction",
		Receipt:  receipt_quo,
	})
	check(err)

	t = time.Now()
	cclib.LogEventToFile(logInfoFile, ecomm.TransactionMinedEvent, payload, t, timeInfoFile)

	fmt.Println("[fabric] Creating auction")
	myAuction := ecomm.StartAuction(assetClient, asset.ID, ethAddr, quoAddr)

	// publish
	ccsvc_, err := cclib.NewEventService(strings.Split(zkNodes, ","), "auctioner")
	check(err)

	payload, _ = json.Marshal(myAuction)
	err = ccsvc_.Publish(ecomm.AuctionCreatingEvent, payload)

	return err
}

func handleSignedAuctionResult(payload []byte) {
	var result ecomm.SignedAuctionResult
	err := json.Unmarshal(payload, &result)
	check(err)

	mutex.Lock()
	defer mutex.Unlock()

	merged := auctionResults[result.HostAuctionID]
	if result.Platform == "ethereum" {
		merged.EthResult.AuctionResult = result.AuctionResult
		merged.EthResult.Signatures = append(merged.EthResult.Signatures, result.Signature)
	} else {
		merged.QuorumResult.AuctionResult = result.AuctionResult
		merged.QuorumResult.Signatures = append(merged.QuorumResult.Signatures, result.Signature)
	}

	if len(merged.EthResult.Signatures) >= 2 && len(merged.QuorumResult.Signatures) >= 2 {
		assetClient.FinalizeAuction(*merged)
	}
}

func handleAuctionCreating(payload []byte) {
	t := time.Now()
	cclib.LogEventToFile(logInfoFile, ecomm.KafkaReceivedEvent, payload, t, timeInfoFile)
	//fmt.Println("Received Auction Creating Event in Kafka!")
	// var result ecomm.Auction
	// err := json.Unmarshal(payload, &result)
	// check(err)

	// mutex.Lock()
	// defer mutex.Unlock()
	return
}

func handleAuctionEnding(payload []byte) {

	// var result ecomm.Auction
	// err := json.Unmarshal(payload, &result)
	// check(err)

	// mutex.Lock()
	// defer mutex.Unlock()
	return
}

func handleBiddingAuction(payload []byte) {
	var result ecomm.AuctionResult
	err := json.Unmarshal(payload, &result)
	check(err)
}

func handleProceedingAuctionResult(payload []byte) {

}

func handleAbortAuctionResult(payload []byte) {

}
