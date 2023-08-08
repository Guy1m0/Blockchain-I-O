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
	//defer cclib.LogEventToFile(logInfoFile, event string, payload []byte, t time.Time, timer_file string)
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

	ethAddr, receipt_eth := ecomm.DeployCrossChainAuction(ethClient, eth_ERC20)
	payload, err := json.Marshal(&ecomm.Tx{
		Platform: "eth",
		Type:     "newAuction",
		Receipt:  receipt_eth,
	})
	check(err)

	t = time.Now()
	cclib.LogEventToFile(logInfoFile, ecomm.TransactionMinedEvent, payload, t, timeInfoFile)

	fmt.Println("[quorum] Deploying auction")
	quoAddr, receipt_quo := ecomm.DeployCrossChainAuction(quoClient, quo_ERC20)
	payload, err = json.Marshal(&ecomm.Tx{
		Platform: "quo",
		Type:     "newAuction",
		Receipt:  receipt_quo,
	})
	check(err)

	t = time.Now()
	cclib.LogEventToFile(logInfoFile, ecomm.TransactionMinedEvent, payload, t, timeInfoFile)

	fmt.Println("[fabric] Creating auction")
	// a.EthAddr = ethAddr
	// a.QuorumAddr = quoAddr
	// _, err = assetClient.SetAuction(a)
	// check(err)
	myAuction := ecomm.StartAuction(assetClient, asset.ID, ethAddr, quoAddr)

	// publish
	ccsvc, err := cclib.NewEventService(strings.Split(zkNodes, ","), "auctioner")
	check(err)
	payload, _ = json.Marshal(myAuction)
	ccsvc.Publish(ecomm.AuctionCreatingEvent, payload)

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

func handleAuctionCreating(payload []byte) {
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
