package main

import (
	"encoding/json"
	"sync"

	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
)

var mutex sync.Mutex

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
