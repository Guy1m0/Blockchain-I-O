package ecomm

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var (
	aucT              *bind.TransactOpts
	support_auc_types = []string{"english", "dutch", "cb1p", "cb2p", "all"}

	//bid_key string
	//auc_key  = "../../keys/key1"
)

// Use key 1 as default auctioneer
func create(asset_name string, auc_type string, usr_name string) {
	t := time.Now()
	//fmt.Println("Auc type:", auc_type)

	if !stringInSlice(support_auc_types, auc_type) {
		log.Println("[fabric] not support auction type")
		return
	}
	log.Println("[fabric] Adding asset")
	_, err := assetClient.AddAsset(asset_name, aucT.From.Hex(), auc_type)
	check(err)

	LogEvent(logInfoFile, AssetAddingEvent, asset_name, auc_type, t, "", 0)
}

func reveal(auctionID int) {
	t := time.Now()
	cclib.LastEventTimestamp.Set(t, timeInfoFile)

	a, err := assetClient.GetAuction(auctionID)
	check(err)

	if a.Status != "open" {
		err = fmt.Errorf("auction status error")
		check(err)
	}

	log.Println("[fabric] Reveal auction")
	_, err = assetClient.CloseAuction(auctionID)
	check(err)

	payload, _ := json.Marshal(a)
	t = time.Now()
	cclib.LogEventToFile(logInfoFile, AuctionClosingEvent, payload, t, timeInfoFile)

	//@reset
	cclib.LastEventTimestamp.Set(t, timeInfoFile)
}

func cancel(auctionID int) {
	t := time.Now()

	a, err := assetClient.GetAuction(auctionID)
	check(err)

	if a.Status != "open" {
		err = fmt.Errorf("auction status error")
		check(err)
	}

	log.Println("[fabric] Cancel auction")
	_, err = assetClient.CancelAuction(auctionID)
	check(err)

	LogEvent(logInfoFile, AuctionCancelingEvent, a.AssetID, a.AucType, t, "", 0)
}

func close(auctionID int) {
	t := time.Now()
	cclib.LastEventTimestamp.Set(t, timeInfoFile)

	a, err := assetClient.GetAuction(auctionID)
	check(err)

	if a.Status != "open" {
		err = fmt.Errorf("auction status error")
		check(err)
	}

	log.Println("[fabric] Conclude auction")
	_, err = assetClient.CloseAuction(auctionID)
	check(err)

	payload, _ := json.Marshal(a)
	t = time.Now()
	cclib.LogEventToFile(logInfoFile, AuctionClosingEvent, payload, t, timeInfoFile)

	//@reset
	cclib.LastEventTimestamp.Set(t, timeInfoFile)
}
