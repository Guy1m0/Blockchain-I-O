package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
)

var mutex sync.Mutex

func handleAddAssetEvent(eventPayload string) error {
	log.Println("[fabric] Creating auction")

	t := time.Now()
	payload, _ := json.Marshal(eventPayload)
	cclib.LogEventToFile(logInfoFile, ecomm.RelayerDetectedEvent, payload, t, timeInfoFile)

	parts := strings.SplitN(eventPayload, ": ", 2)
	if len(parts) < 2 {
		return fmt.Errorf("received unexpected event: %s", eventPayload)
		// Now eventDetail contains the string after ": "
	}

	eventDetail := parts[1]
	asset, err := assetClient.GetAsset(eventDetail)
	check(err)

	log.Println("[ethereum] Deploying auction")

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

	log.Println("[quorum] Deploying auction")
	quoAddr, receipt_quo := ecomm.DeployCrossChainAuction(quoClient, quo_ERC20, root_key)
	payload, err = json.Marshal(&ecomm.Tx{
		Platform: "quo",
		Type:     "newAuction",
		Receipt:  receipt_quo,
	})
	check(err)

	t = time.Now()
	cclib.LogEventToFile(logInfoFile, ecomm.TransactionMinedEvent, payload, t, timeInfoFile)
	// @reset
	cclib.LastEventTimestamp.Set(t, timeInfoFile)

	fmt.Println("[fabric] Creating auction")

	args := ecomm.StartAuctionArgs{
		AssetID:    asset.ID,
		EthAddr:    ethAddr,
		QuorumAddr: quoAddr,
	}
	payload, _ = json.Marshal(args)
	_, err = assetClient.StartAuction(args)
	cclib.LogEventToFile(logInfoFile, ecomm.AuctionStartingEvent, payload, t, timeInfoFile)

	//myAuction := ecomm.StartAuction(assetClient, asset.ID, ethAddr, quoAddr)

	return err
}

func handleStartAuctionEvent(eventPayload string) error {
	log.Println("[kafka] Publish Auction Creating Event")

	t := time.Now()
	payload, _ := json.Marshal(eventPayload)
	cclib.LogEventToFile(logInfoFile, ecomm.RelayerDetectedEvent, payload, t, timeInfoFile)

	parts := strings.SplitN(eventPayload, ": ", 2)
	if len(parts) < 2 {
		return fmt.Errorf("received unexpected event: %s", eventPayload)
	}

	eventDetail := parts[1]
	auctionID, _ := strconv.Atoi(eventDetail)
	//fmt.Println("Auction ID:", auctionID, "eventPayload:", eventPayload)
	myAuction, err := assetClient.GetAuction(auctionID)
	check(err)
	// publish
	//ccsvc_, _ := cclib.NewEventService(strings.Split(zkNodes, ","), "auctioner")

	payload, _ = json.Marshal(myAuction)
	err = ccsvc.Publish(ecomm.AuctionStartingEvent, payload)

	return err
}

func handleEndAuctionEvent(eventPayload string) error {
	return nil
}

func handleFinAuctionEvent(eventPayload string) error {
	return nil
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

func logEvent(payload []byte) {
	t := time.Now()
	cclib.LogEventToFile(logInfoFile, ecomm.KafkaReceivedEvent, payload, t, timeInfoFile)
	//fmt.Println("Received Auction Creating Event in Kafka!")
	// var result ecomm.Auction
	// err := json.Unmarshal(payload, &result)
	// check(err)

	// mutex.Lock()
	// defer mutex.Unlock()
	//return
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
