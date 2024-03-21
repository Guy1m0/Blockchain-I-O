package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/contracts/eth_auction"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// fabric relayer
func handleAddAssetEvent(eventPayload []byte) error {
	log.Println("[fabric] Get Asset")
	t := time.Now()

	var result ecomm.AssetAddingEventPayload
	err := json.Unmarshal(eventPayload, &result)
	check(err)

	assetID := result.AssetID
	asset, err := assetClient.GetAsset(assetID)
	check(err)

	auc_type := result.AucType
	fmt.Println("Auc Type:", auc_type)
	ecomm.LogEvent(logInfoFile, ecomm.AssetAddingEvent, assetID, auc_type, t, "", 0)

	payloadJSON, _ := json.Marshal(asset)
	wrapper := ecomm.EventWrapper{Type: "Asset", Result: payloadJSON}
	payload, _ := json.Marshal(wrapper)

	// @todo: change this payload to be AssetAddingEventPayload?
	ccsvc.Publish(ecomm.AssetAddingEvent, payload)

	log.Println("[fabric] Start auction")
	t = time.Now()

	var ethAddr, quoAddr common.Address
	switch auc_type {
	case "english":
		ethAddr = contract_info.EnglishAuction.EthAddr
		quoAddr = contract_info.EnglishAuction.QuoAddr
	case "dutch":
		ethAddr = contract_info.DutchAuction.EthAddr
		quoAddr = contract_info.DutchAuction.QuoAddr
	case "cb1p":
		ethAddr = contract_info.Cb1pAuction.EthAddr
		quoAddr = contract_info.Cb1pAuction.QuoAddr
	case "cb2p":
		ethAddr = contract_info.Cb2pAuction.EthAddr
		quoAddr = contract_info.Cb2pAuction.QuoAddr
	default:
		fmt.Println("Auction type error")
	}

	args := ecomm.StartAuctionArgs{
		AssetID:    asset.ID,
		AucType:    auc_type,
		EthAddr:    ethAddr.String(),
		QuorumAddr: quoAddr.String(),
	}

	_, err = assetClient.StartAuction(args)

	check(err)
	ecomm.LogEvent(logInfoFile, ecomm.AuctionStartingEvent, assetID, auc_type, t, "", 0)
	return nil
}

// fabric relayer
func handleStartAuctionEvent(eventPayload []byte) error {
	//t := time.Now()
	var result ecomm.StartAuctionEventPayload
	var tx *types.Transaction
	var receipt1, receipt2 *types.Receipt

	authT, _ := cclib.NewTransactor(root_key, password)

	err := json.Unmarshal(eventPayload, &result)
	check(err)

	auction, err := assetClient.GetAuction(result.ID)
	check(err)

	// ethAddr := auction.EthAddr
	// quoAddr := auction.QuorumAddr

	switch auction.AucType {
	case "english":
		tx, err = eth_english_auction_contract.Create(authT, big.NewInt(int64(auction.AuctionID)), auction.AssetID, result.Owner)
		check(err)
		receipt1 = ecomm.WaitTx(ethClient, tx, fmt.Sprintf("Create new auction with type: %s and ID: %d", auction.AucType, auction.AuctionID))

		tx, err = quo_english_auction_contract.Create(authT, big.NewInt(int64(auction.AuctionID)), auction.AssetID, result.Owner)
		check(err)
		receipt2 = ecomm.WaitTx(quoClient, tx, fmt.Sprintf("Create new auction with type: %s and ID: %d", auction.AucType, auction.AuctionID))
	case "dutch":

	case "cb1p":
		tx, err = eth_cb1p_contract.Create(authT, big.NewInt(int64(auction.AuctionID)), auction.AssetID, result.Owner)
		check(err)
		receipt1 = ecomm.WaitTx(ethClient, tx, fmt.Sprintf("Create new auction with type: %s and ID: %d", auction.AucType, auction.AuctionID))

		tx, err = quo_cb1p_contract.Create(authT, big.NewInt(int64(auction.AuctionID)), auction.AssetID, result.Owner)
		check(err)
		receipt2 = ecomm.WaitTx(quoClient, tx, fmt.Sprintf("Create new auction with type: %s and ID: %d", auction.AucType, auction.AuctionID))

	case "cb2p":

	default:
		fmt.Println("Auction type error")

	}

	cost := receipt1.GasUsed
	note := "ETH:" + strconv.FormatUint(cost, 10)
	cost += receipt2.GasUsed
	note += " QUO:" + strconv.FormatUint(receipt2.GasUsed, 10)

	t := time.Now()
	ecomm.LogEvent(logInfoFile, ecomm.AuctionStartingEvent, auction.AssetID, auction.AucType, t, note, cost)

	log.Println("[fabirc] Start Auction with ID: ", result.ID)
	log.Println("AuctionID", auction.AuctionID)

	payloadJSON, _ := json.Marshal(auction)
	wrapper := ecomm.EventWrapper{Type: "Start Auction", Result: payloadJSON}
	payload, _ := json.Marshal(wrapper)

	err = ccsvc.Publish(ecomm.AuctionStartingEvent, payload)

	return err
}

func handleCancelAuctionEvent(eventPayload string) error {
	t := time.Now()
	parts := strings.SplitN(eventPayload, ": ", 2)
	if len(parts) < 2 {
		return fmt.Errorf("received unexpected event: %s", eventPayload)
	}

	auctionID, err := strconv.Atoi(parts[1])
	check(err)
	auction, err := assetClient.GetAuction(auctionID)
	check(err)
	ecomm.LogEvent(logInfoFile, ecomm.AuctionCancelingEvent, auction.AssetID, auction.AucType, t, "", 0)

	log.Println("[ETH/QUO] Close auctions on both platforms")
	// load auction contract
	eth_auction_addr := common.HexToAddress(auction.EthAddr)
	eth_auction_contract, err := eth_auction.NewEthAuction(eth_auction_addr, ethClient)
	check(err)

	quo_auction_addr := common.HexToAddress(auction.QuorumAddr)
	quo_auction_contract, err := eth_auction.NewEthAuction(quo_auction_addr, quoClient)
	check(err)

	log.Println("[ETH/QUO] Change contract state")
	authT, err := cclib.NewTransactor(root_key, password)
	check(err)

	// Change Auction Contract on Eth
	tx, _ := eth_auction_contract.CloseAuction(authT, true)
	receipt := ecomm.WaitTx(ethClient, tx, fmt.Sprintf("Change Auction %s status to 'ENDED'", eth_auction_addr))

	cost := receipt.GasUsed
	note := "ETH:" + strconv.FormatUint(cost, 10)

	// Change Auction Contract on Quo
	tx, _ = quo_auction_contract.CloseAuction(authT, true)
	receipt = ecomm.WaitTx(quoClient, tx, fmt.Sprintf("Change Auction %s status to 'ENDED'", quo_auction_addr))

	cost += receipt.GasUsed
	note += " QUO:" + strconv.FormatUint(receipt.GasUsed, 10)

	ecomm.UpdateLog(logInfoFile, ecomm.AuctionCancelingEvent, auction.AssetID, "", cost, note)
	//ccsvc.Publish(ecomm.AuctionClosingEvent, payload)

	payloadJSON, _ := json.Marshal(auction)
	wrapper := ecomm.EventWrapper{Type: "Cancel Auction", Result: payloadJSON}
	payload, _ := json.Marshal(wrapper)

	return ccsvc.Publish(ecomm.AuctionCancelingEvent, payload)
}

func handleCloseAuctionEvent(eventPayload string) error {
	t := time.Now()
	parts := strings.SplitN(eventPayload, ": ", 2)
	if len(parts) < 2 {
		return fmt.Errorf("received unexpected event: %s", eventPayload)
	}

	auctionID, err := strconv.Atoi(parts[1])
	check(err)
	auction, err := assetClient.GetAuction(auctionID)
	check(err)
	ecomm.LogEvent(logInfoFile, ecomm.AuctionCancelingEvent, auction.AssetID, auction.AucType, t, "", 0)

	log.Println("[ETH/QUO] Determin winner")
	// load auction contract
	eth_auction_addr := common.HexToAddress(auction.EthAddr)
	eth_auction_contract, err := eth_auction.NewEthAuction(eth_auction_addr, ethClient)
	check(err)

	quo_auction_addr := common.HexToAddress(auction.QuorumAddr)
	quo_auction_contract, err := eth_auction.NewEthAuction(quo_auction_addr, quoClient)
	check(err)

	// Check highest bid
	eth_highestBid, _ := eth_auction_contract.HighestBid(&bind.CallOpts{})
	//eth_highestBidder, _ := eth_auction_contract.HighestBidder(&bind.CallOpts{})
	quo_highestBid, _ := quo_auction_contract.HighestBid(&bind.CallOpts{})
	//quo_highestBidder, _ := quo_auction_contract.HighestBidder(&bind.CallOpts{})

	eth_bool := false
	quo_bool := true

	if eth_highestBid.Cmp(quo_highestBid) < 0 {
		eth_bool = true
		quo_bool = false
	}

	log.Println("[ETH/QUO] Change contract state")
	authT, err := cclib.NewTransactor(root_key, password)
	check(err)

	// @reset timer
	t = time.Now()
	cclib.LastEventTimestamp.Set(t, timeInfoFile)

	// Change Auction Contract on Eth
	tx, _ := eth_auction_contract.CloseAuction(authT, eth_bool)
	receipt := ecomm.WaitTx(ethClient, tx, fmt.Sprintf("Change Auction %s status to 'ENDING'", eth_auction_addr))
	fmt.Println(receipt.TxHash.Hex())
	// payload, err = json.Marshal(&ecomm.Tx{
	// 	Platform: "eth",
	// 	Type:     "CloseAuction",
	// 	Hash:     receipt.TxHash,
	// })
	// check(err)

	// t = time.Now()
	// cclib.LogEventToFile(logInfoFile, ecomm.TransactionMinedEvent, payload, t, timeInfoFile)
	// cclib.LastEventTimestamp.Set(t, timeInfoFile)

	// Change Auction Contract on Quo
	tx, _ = quo_auction_contract.CloseAuction(authT, quo_bool)
	receipt = ecomm.WaitTx(quoClient, tx, fmt.Sprintf("Change Auction %s status to 'ENDING'", quo_auction_addr))
	fmt.Println(receipt.TxHash.Hex())
	// payload, err = json.Marshal(&ecomm.Tx{
	// 	Platform: "quo",
	// 	Type:     "CloseAuction",
	// 	Hash:     receipt.TxHash,
	// })
	// check(err)

	// t = time.Now()
	// cclib.LogEventToFile(logInfoFile, ecomm.TransactionMinedEvent, payload, t, timeInfoFile)
	// //cclib.LastEventTimestamp.Set(t, timeInfoFile)

	// payload, _ = json.Marshal(a)
	// ccsvc.Publish(ecomm.AuctionClosingEvent, payload)

	return nil
}

func handleAuctionClosedEvent(eventPayload string) error {
	return nil
}

// func handleSignedAuctionResult(payload []byte) {
// 	var result ecomm.SignedAuctionResult
// 	err := json.Unmarshal(payload, &result)
// 	check(err)

// 	mutex.Lock()
// 	defer mutex.Unlock()

// 	merged := auctionResults[result.AuctionID]
// 	if result.Platform == "ethereum" {
// 		merged.EthResult.AuctionResult = result.AuctionResult
// 		merged.EthResult.Signatures = append(merged.EthResult.Signatures, result.Signature)
// 	} else {
// 		merged.QuorumResult.AuctionResult = result.AuctionResult
// 		merged.QuorumResult.Signatures = append(merged.QuorumResult.Signatures, result.Signature)
// 	}

// 	if len(merged.EthResult.Signatures) >= 2 && len(merged.QuorumResult.Signatures) >= 2 {
// 		assetClient.FinalizeAuction(*merged)
// 	}
// }

func chainCodeEvent(eventPayload []byte) {
	t := time.Now()
	var wrapper ecomm.EventWrapper
	var event, eventID string
	err := json.Unmarshal([]byte(eventPayload), &wrapper)
	check(err)

	switch wrapper.Type {
	case "Asset":
		var asset ecomm.Asset
		err = json.Unmarshal(wrapper.Result, &asset)
		check(err)

		event = ecomm.AssetAddingEvent
		eventID = asset.ID
		//fmt.Printf("Received Asset: %+v\n", asset)
	case "Start Auction":
		var auction ecomm.Auction
		err = json.Unmarshal(wrapper.Result, &auction)
		check(err)

		event = ecomm.AuctionStartingEvent
		eventID = auction.AssetID
	case "Cancel Auction":
		var auction ecomm.Auction
		err = json.Unmarshal(wrapper.Result, &auction)
		check(err)

		event = ecomm.AuctionCancelingEvent
		eventID = auction.AssetID
		//fmt.Printf("Received Auction: %+v\n", auction)
	// case "Bid":
	// 	var bid ecomm.Bid
	// 	err = json.Unmarshal(wrapper.Result, &bid)
	// 	check(err)

	// 	event = ecomm.BidEvent
	// 	eventID = bid.AssetID
	//fmt.Printf("Received Bid: %+v\n", bid)
	default:
		fmt.Printf("Unknown type: %s\n", wrapper.Type)
	}
	//log.Println("Kafka received event:", event, "with ID:", eventID)
	//cclib.LogEventToFile(logInfoFile, ecomm.KafkaReceivedEvent, payload, t, timeInfoFile)
	ecomm.LogEvent(logInfoFile, event, eventID, "", t, "", 0)
}
