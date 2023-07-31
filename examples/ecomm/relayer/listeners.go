package main

import (
	"encoding/json"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
)

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
