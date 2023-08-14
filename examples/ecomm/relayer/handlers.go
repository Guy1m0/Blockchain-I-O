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
)

// fabric relayer
func handleAddAssetEvent(eventPayload string) error {
	log.Println("[fabric] Creating auction")

	t := time.Now()
	payload, _ := json.Marshal(eventPayload)
	cclib.LogEventToFile(logInfoFile, ecomm.RelayerDetectedEvent, payload, t, timeInfoFile)

	parts := strings.SplitN(eventPayload, ": ", 2)
	if len(parts) != 2 {
		return fmt.Errorf("received unexpected event: %s", eventPayload)
	}

	eventDetail := parts[1]
	asset, err := assetClient.GetAsset(eventDetail)
	check(err)

	// @todo: publish events which handled by relayer on eth and quo

	log.Println("[ethereum] Deploying auction")

	ethAddr, receipt_eth := ecomm.DeployCrossChainAuction(ethClient, eth_ERC20, asset.ID, root_key)
	payload_eth, _ := json.Marshal(&ecomm.Tx{
		Platform: "eth",
		Type:     "newAuction",
		Hash:     receipt_eth.TxHash,
	})
	check(err)

	// elimiate the effects caused by some unknown reasons that relayer detects
	// events before the related tx mined in original platform
	cclib.LastEventTimestamp.Set(t, timeInfoFile)
	t = time.Now()
	cclib.LogEventToFile(logInfoFile, ecomm.TransactionMinedEvent, payload_eth, t, timeInfoFile)
	cclib.LastEventTimestamp.Set(t, timeInfoFile)

	log.Println("[quorum] Deploying auction")
	quoAddr, receipt_quo := ecomm.DeployCrossChainAuction(quoClient, quo_ERC20, asset.ID, root_key)
	payload_quo, err := json.Marshal(&ecomm.Tx{
		Platform: "quo",
		Type:     "newAuction",
		Hash:     receipt_quo.TxHash,
	})
	check(err)

	t = time.Now()
	cclib.LogEventToFile(logInfoFile, ecomm.TransactionMinedEvent, payload_quo, t, timeInfoFile)
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

	return err
}

// fabric relayer
func handleStartAuctionEvent(eventPayload string) error {
	log.Println("[kafka] Publish Auction Creating Event")

	//t := time.Now()
	payload, _ := json.Marshal(eventPayload)
	//cclib.LogEventToFile(logInfoFile, ecomm.RelayerDetectedEvent, payload, t, timeInfoFile)

	parts := strings.SplitN(eventPayload, ": ", 2)
	if len(parts) < 2 {
		return fmt.Errorf("received unexpected event: %s", eventPayload)
	}

	eventDetail := parts[1]
	auctionID, _ := strconv.Atoi(eventDetail)
	//fmt.Println("Auction ID:", auctionID, "eventPayload:", eventPayload)
	a, err := assetClient.GetAuction(auctionID)
	check(err)

	log.Println("[fabirc] Start Auction with ID: ", a.ID)
	// Listen new auction
	onNewAuction(a)

	//payload, _ = json.Marshal(a)

	// @reset timer
	t := time.Now()
	cclib.LastEventTimestamp.Set(t, timeInfoFile)
	err = ccsvc.Publish(ecomm.AuctionStartingEvent, payload)

	return err
}

func handleHighestBidIncreasedEvent(eventPayload ecomm.HighestBidIncreasedEvent, result ecomm.AuctionResult, t time.Time) error {
	log.Printf("[%s] HighestBid Increased Event", strings.ToUpper(result.Platform))

	result.HighestBid = int(eventPayload.Amount.Int64())
	result.HighestBidder = eventPayload.Bidder.Hex()

	//t := time.Now()
	payload, _ := json.Marshal(result)
	cclib.LogEventToFile(logInfoFile, ecomm.RelayerDetectedEvent, payload, t, timeInfoFile)

	// // Check later if really need this
	// auctionResultsMu.Lock()
	// auctionResults[a.ID].EthResult.AuctionResult = result
	// auctionResultsMu.Unlock()

	// @reset
	t = time.Now()
	cclib.LastEventTimestamp.Set(t, timeInfoFile)
	ccsvc.Publish(ecomm.BidEvent, payload)

	return nil
}

func handleCloseAuctionEvent(eventPayload string) error {
	log.Println("[ETH/QUO] Determin winner")

	t := time.Now()
	payload, _ := json.Marshal(eventPayload)
	cclib.LogEventToFile(logInfoFile, ecomm.RelayerDetectedEvent, payload, t, timeInfoFile)

	parts := strings.SplitN(eventPayload, ": ", 2)
	if len(parts) < 2 {
		return fmt.Errorf("received unexpected event: %s", eventPayload)
		// Now eventDetail contains the string after ": "
	}

	eventDetail := parts[1]
	auctionID, _ := strconv.Atoi(eventDetail)

	a, err := assetClient.GetAuction(auctionID)
	check(err)

	// load auction contract
	eth_auction_addr := common.HexToAddress(a.EthAddr)
	eth_auction_contract, err := eth_auction.NewEthAuction(eth_auction_addr, ethClient)
	check(err)

	quo_auction_addr := common.HexToAddress(a.QuorumAddr)
	quo_auction_contract, err := eth_auction.NewEthAuction(quo_auction_addr, quoClient)
	check(err)

	// Check highest bid
	eth_highestBid, _ := eth_auction_contract.HighestBid(&bind.CallOpts{})
	//eth_highestBidder, _ := eth_auction_contract.HighestBidder(&bind.CallOpts{})
	quo_highestBid, _ := quo_auction_contract.HighestBid(&bind.CallOpts{})
	//quo_highestBidder, _ := quo_auction_contract.HighestBidder(&bind.CallOpts{})

	eth_bool := false
	quo_bool := true
	a.HighestBidPlatform = "eth"
	//a.HighestBid = eth_highestBid
	if eth_highestBid.Cmp(quo_highestBid) < 0 {
		eth_bool = true
		quo_bool = false
		a.HighestBidPlatform = "quo"
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

	payload, err = json.Marshal(&ecomm.Tx{
		Platform: "eth",
		Type:     "CloseAuction",
		Hash:     receipt.TxHash,
	})
	check(err)

	t = time.Now()
	cclib.LogEventToFile(logInfoFile, ecomm.TransactionMinedEvent, payload, t, timeInfoFile)
	cclib.LastEventTimestamp.Set(t, timeInfoFile)

	// Change Auction Contract on Quo
	tx, _ = quo_auction_contract.CloseAuction(authT, quo_bool)
	receipt = ecomm.WaitTx(quoClient, tx, fmt.Sprintf("Change Auction %s status to 'ENDING'", quo_auction_addr))

	payload, err = json.Marshal(&ecomm.Tx{
		Platform: "quo",
		Type:     "CloseAuction",
		Hash:     receipt.TxHash,
	})
	check(err)

	t = time.Now()
	cclib.LogEventToFile(logInfoFile, ecomm.TransactionMinedEvent, payload, t, timeInfoFile)
	//cclib.LastEventTimestamp.Set(t, timeInfoFile)

	payload, _ = json.Marshal(a)
	ccsvc.Publish(ecomm.AuctionClosingEvent, payload)

	return nil
}

func handleDecisionMadeEvent(eventPayload ecomm.DecisionMadeEvent, t time.Time) error {
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

		amt := big.NewInt(int64(result.HighestBid))
		quotient := new(big.Int).Div(amt, ecomm.DecimalB)
		fabric_ERC20_contract.Transfer(asset.Owner, quotient.String())
	}

	ccsvc.Publish(ecomm.AuctionFinalizingEvent, []byte(payload))
	//t := time.Now()
	return nil
}

func handleCancelAuctionEvent(eventPayload string) error {
	log.Println("[ETH/QUO] Clase auctions on both platforms")

	t := time.Now()
	payload, _ := json.Marshal(eventPayload)
	cclib.LogEventToFile(logInfoFile, ecomm.RelayerDetectedEvent, payload, t, timeInfoFile)

	parts := strings.SplitN(eventPayload, ": ", 2)
	if len(parts) < 2 {
		return fmt.Errorf("received unexpected event: %s", eventPayload)
	}

	eventDetail := parts[1]
	auctionID, _ := strconv.Atoi(eventDetail)

	a, err := assetClient.GetAuction(auctionID)
	check(err)

	// load auction contract
	eth_auction_addr := common.HexToAddress(a.EthAddr)
	eth_auction_contract, err := eth_auction.NewEthAuction(eth_auction_addr, ethClient)
	check(err)

	quo_auction_addr := common.HexToAddress(a.QuorumAddr)
	quo_auction_contract, err := eth_auction.NewEthAuction(quo_auction_addr, quoClient)
	check(err)

	log.Println("[ETH/QUO] Change contract state")
	authT, err := cclib.NewTransactor(root_key, password)
	check(err)

	// @reset timer
	t = time.Now()
	cclib.LastEventTimestamp.Set(t, timeInfoFile)

	// Change Auction Contract on Eth
	tx, _ := eth_auction_contract.CloseAuction(authT, true)
	receipt := ecomm.WaitTx(ethClient, tx, fmt.Sprintf("Change Auction %s status to 'ENDED'", eth_auction_addr))

	payload, err = json.Marshal(&ecomm.Tx{
		Platform: "eth",
		Type:     "finAuction",
		Hash:     receipt.TxHash,
	})
	check(err)

	t = time.Now()
	cclib.LogEventToFile(logInfoFile, ecomm.TransactionMinedEvent, payload, t, timeInfoFile)
	cclib.LastEventTimestamp.Set(t, timeInfoFile)

	// Change Auction Contract on Quo
	tx, _ = quo_auction_contract.CloseAuction(authT, true)
	receipt = ecomm.WaitTx(quoClient, tx, fmt.Sprintf("Change Auction %s status to 'ENDED'", quo_auction_addr))

	payload, err = json.Marshal(&ecomm.Tx{
		Platform: "quo",
		Type:     "finAuction",
		Hash:     receipt.TxHash,
	})
	check(err)

	t = time.Now()
	cclib.LogEventToFile(logInfoFile, ecomm.TransactionMinedEvent, payload, t, timeInfoFile)
	//cclib.LastEventTimestamp.Set(t, timeInfoFile)

	payload, _ = json.Marshal(a)
	ccsvc.Publish(ecomm.AuctionClosingEvent, payload)

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

func logEvent(payload []byte) {
	t := time.Now()
	cclib.LogEventToFile(logInfoFile, ecomm.KafkaReceivedEvent, payload, t, timeInfoFile)
}
