package main

import (
	"encoding/json"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
)

func runAuctionListener() {
	for {
		a := listenNewAuction()
		go onNewAuction(a)
	}
}

func onNewAuction(a *ecomm.Auction) {
	auctionResults[a.ID] = &ecomm.FinalizeAuctionArgs{
		AuctionID: a.ID,
	}
	onAuctionEnding(listenAuctionEnding(a))
}

func onAuctionEnding(a *ecomm.Auction) {
	b, _ := json.Marshal(a)
	ccsvc.Publish(ecomm.AuctionEndingEvent, b)
}

func listenNewAuction() *ecomm.Auction {
	lastID, err := assetClient.GetLastAuctionID()
	check(err)

	for {
		time.Sleep(1 * time.Second)
		auctionID, err := assetClient.GetLastAuctionID()
		check(err)
		if auctionID > lastID {
			auction, err := assetClient.GetAuction(auctionID)
			check(err)
			return auction
		}
	}
}

func listenAuctionEnding(a *ecomm.Auction) *ecomm.Auction {
	for {
		time.Sleep(1 * time.Second)
		auction, err := assetClient.GetAuction(a.ID)
		check(err)
		if auction.Status == "Ending" {
			return auction
		}
	}
}
