package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
)

// Smart Contract handler
func handleHighestBidIncreasedEvent(eventPayload ecomm.HighestBidIncreased, bid ecomm.Bid, t time.Time) error {
	log.Printf("[%s] HighestBidIncreased Event", strings.ToUpper(bid.Platform))

	amount := new(big.Int).Div(eventPayload.BidAmount, ecomm.DecimalB).String()
	eventID := eventPayload.Id + "_" + bid.Platform + "_" + eventPayload.Bidder.String()[36:]
	ecomm.LogEvent(logInfoFile, ecomm.BidEvent, eventID, eventPayload.AuctionType, t, "", 0)

	bid.BidAmount = amount
	bid.Bidder = eventPayload.Bidder
	bid.AssetID = eventPayload.Id

	payloadJSON, _ := json.Marshal(bid)
	wrapper := ecomm.EventWrapper{Type: "Bid", Result: payloadJSON}
	payload, _ := json.Marshal(wrapper)

	// asset, _ := assetClient.GetAsset(eventPayload.Id)
	// auction, _ := assetClient.GetAuction(asset.PendingAuctionID)
	// fmt.Println("find auction in new bid: ", auction.ID, "status: ", auction.Status)

	ccsvc.Publish(ecomm.BidEvent, payload)
	return nil
}

// Smart Contract handler
func handleNewBidHashEvent(eventPayload ecomm.NewBidHash, bidHash ecomm.BidHash, t time.Time) error {
	log.Printf("[%s] NewBidHash Event", strings.ToUpper(bidHash.Platform))

	//amount := new(big.Int).Div(eventPayload.BidAmount, ecomm.DecimalB).String()
	eventID := eventPayload.Id + "_" + bidHash.Platform + "_" + eventPayload.Bidder.String()[36:]
	ecomm.LogEvent(logInfoFile, ecomm.BidEvent, eventID, "", t, "", 0)

	bidHash.BidHash = eventPayload.BidHash
	bidHash.Bidder = eventPayload.Bidder
	bidHash.AssetID = eventPayload.Id

	payloadJSON, _ := json.Marshal(bidHash)
	wrapper := ecomm.EventWrapper{Type: "BidHash", Result: payloadJSON}
	payload, _ := json.Marshal(wrapper)

	// asset, _ := assetClient.GetAsset(eventPayload.Id)
	// auction, _ := assetClient.GetAuction(asset.PendingAuctionID)
	// fmt.Println("find auction in new bid: ", auction.ID, "status: ", auction.Status)

	ccsvc.Publish(ecomm.BidEvent, payload)
	return nil
}

func handleWithdrawBidEvent(eventPayload ecomm.WithdrawBid, bid ecomm.Bid, t time.Time) error {
	log.Printf("[%s] WithdrawBid Event", strings.ToUpper(bid.Platform))

	amount := new(big.Int).Div(eventPayload.Amount, ecomm.DecimalB).String()
	eventID := eventPayload.Id + "_" + bid.Platform + "_" + eventPayload.Bidder.String()[36:]
	ecomm.LogEvent(logInfoFile, ecomm.WithdrawEvent, eventID, "", t, "", 0)

	bid.BidAmount = amount
	bid.Bidder = eventPayload.Bidder
	bid.AssetID = eventPayload.Id

	ecomm.UpdateLog(logInfoFile, ecomm.WithdrawEvent, eventID, 0, "Withdraw: MDAI "+amount)

	payloadJSON, _ := json.Marshal(bid)
	wrapper := ecomm.EventWrapper{Type: "Withdraw", Result: payloadJSON}
	payload, _ := json.Marshal(wrapper)

	ccsvc.Publish(ecomm.WithdrawEvent, payload)
	return nil
}

func handleDecisionMadeEvent(eventPayload ecomm.DecisionMade, t time.Time) error {
	payload := eventPayload.JsonString
	cclib.LogEventToFile(logInfoFile, ecomm.RelayerDetectedEvent, []byte(eventPayload.JsonString), t, timeInfoFile)

	var result ecomm.AuctionResult

	err := json.Unmarshal([]byte(payload), &result)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	log.Printf("[%s] Decesion Made Event", strings.ToUpper(result.Platform))

	proceed := eventPayload.Prcd

	_, err = assetClient.FinAuction(result, proceed)
	check(err)

	// transfer token to original owner
	if proceed {
		fabric_ERC20_contract := ecomm.NewErc20Client(fabric_ERC20)
		auction, _ := assetClient.GetAuction(result.AuctionID)
		asset, _ := assetClient.GetAsset(auction.AssetID)

		amt := result.HighestBid
		//amt, _ := new(big.Int).SetString(result.HighestBid, 10)
		//big.NewInt(int64(result.HighestBid))
		quotient := new(big.Int).Div(&amt, ecomm.DecimalB)
		fabric_ERC20_contract.Transfer(asset.Owner, quotient.String())
	}

	ccsvc.Publish(ecomm.AuctionFinalizingEvent, []byte(payload))
	//t := time.Now()
	return nil
}

func smartContractEvent(eventPayload []byte) {
	t := time.Now()
	var wrapper ecomm.EventWrapper
	var event, eventID string
	err := json.Unmarshal([]byte(eventPayload), &wrapper)
	check(err)

	switch wrapper.Type {
	case "Bid":
		var bid ecomm.Bid
		err = json.Unmarshal(wrapper.Result, &bid)
		check(err)

		event = ecomm.BidEvent
		eventID = bid.AssetID + "_" + bid.Platform + "_" + bid.Bidder.String()[36:]
		//fmt.Printf("Received Asset: %+v\n", asset)
	case "Withdraw":
		var bid ecomm.Bid
		err = json.Unmarshal(wrapper.Result, &bid)
		check(err)

		event = ecomm.WithdrawEvent
		eventID = bid.AssetID + "_" + bid.Platform + "_" + bid.Bidder.String()[36:]

	case "BidHash":
		var bidHash ecomm.BidHash
		err = json.Unmarshal(wrapper.Result, &bidHash)
		check(err)

		event = ecomm.WithdrawEvent
		eventID = bidHash.AssetID + "_" + bidHash.Platform + "_" + bidHash.Bidder.String()[36:]
	default:
		fmt.Printf("Unknown type: %s\n", wrapper.Type)
	}

	ecomm.LogEvent(logInfoFile, event, eventID, "", t, "", 0)
}
