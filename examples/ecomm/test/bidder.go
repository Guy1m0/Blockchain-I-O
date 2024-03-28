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
	"github.com/Guy1m0/Blockchain-I-O/contracts/cb1p_auction"
	"github.com/Guy1m0/Blockchain-I-O/contracts/cb2p_auction"
	"github.com/Guy1m0/Blockchain-I-O/contracts/dutch_auction"
	"github.com/Guy1m0/Blockchain-I-O/contracts/english_auction"
	"github.com/Guy1m0/Blockchain-I-O/contracts/stable_coin"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func load_bidder_key(name string) string {
	users, err := ecomm.ReadUsersFromFile(userInfoFile)
	check(err)

	for _, user := range users {
		//fmt.Println("Find ", name, "in", user.UserID)
		if name == user.UserID {
			//fmt.Println("Find!", user.KeyFile)
			return user.KeyFile
		}
	}

	return "../../keys/key2"
}

// also no relayer involved, 'locally' make bid
func bidAuction(auction_id int, amount *big.Int, bid_key string) {
	t := time.Now()

	var contract_info ecomm.ContractInfo
	ecomm.ReadJsonFile(contractInfoFile, &contract_info)
	erc20_address := contract_info.EthERC20
	client := ethClient

	a, err := assetClient.GetAuction(auction_id)
	check(err)

	auction_addr := common.HexToAddress(a.EthAddr)
	// @todo: require platform either is 'quo' or 'eth'
	if platform != "eth" {
		auction_addr = common.HexToAddress(a.QuorumAddr)
		client = quoClient
		erc20_address = contract_info.QuoERC20
	}

	var auction_contract ecomm.AuctionContract
	//auction_contract, _ := english_auction.NewEnglishAuction(auction_addr, client)
	switch auc_type {
	case "english":
		auction_contract, _ = english_auction.NewEnglishAuction(auction_addr, client)

	case "dutch":
		auction_contract, _ = dutch_auction.NewDutchAuction(auction_addr, client)
	}

	check(err)

	bidT, err := cclib.NewTransactor(bid_key, "password")
	check(err)

	// Bid more than highest
	if amount.Cmp(big.NewInt(0)) == 0 {
		highest, err := auction_contract.HighestBid(&bind.CallOpts{}, big.NewInt(int64(auction_id)))
		highest.Div(highest, ecomm.DecimalB)
		check(err)
		amount.Add(highest, big.NewInt(1))
		//log.Printf("highest: %s bidAmount: %s", highest, amount)
	}

	//eventID := fmt.Sprintf("%s_%s_%s_%s", a.AssetID, platform, bidT.From.String()[36:], amount)
	keyWords := fmt.Sprintf("%s_%s_%s", platform, bidT.From.String()[36:], amount)
	ecomm.LogEvent(logInfoFile, a.AssetID, ecomm.BidEvent, keyWords, t, "", 0)

	// @todo: Make approve and bid in a single transaction
	// Approve amount of bid through ERC20 contract
	MDAI, _ := stable_coin.NewStableCoin(erc20_address, client)
	// valueB, err := MDAI.BalanceOf(&bind.CallOpts{}, auction_addr)
	// log.Printf("Auction contract orig balance: %s", valueB)

	tx1, _ := MDAI.Approve(bidT, auction_addr, big.NewInt(0).Mul(big.NewInt(amount.Int64()), ecomm.DecimalB))
	receipt1 := ecomm.WaitTx(client, tx1, "Approve Auction Contract's allowance")
	// allB, err := MDAI.Allowance(&bind.CallOpts{}, bidT.From, auction_addr)
	// log.Printf("Auction contract orig allowance: %s", allB)

	tx2, _ := auction_contract.Bid(bidT, big.NewInt(int64(auction_id)), big.NewInt(0).Mul(big.NewInt(amount.Int64()), ecomm.DecimalB))
	receipt2 := ecomm.WaitTx(client, tx2, fmt.Sprintf("Bid on Auction ID: %d through contract: %s", a.AuctionID, auction_addr))

	// valueB, err = MDAI.BalanceOf(&bind.CallOpts{}, auction_addr)
	// log.Printf("Auction contract after balance: %s", valueB)

	note := "Cost: " + strconv.FormatUint(receipt1.GasUsed, 10)
	note += " + " + strconv.FormatUint(receipt2.GasUsed, 10)
	note += " Bid: MDAI " + big.NewInt(0).Mul(big.NewInt(amount.Int64()), ecomm.DecimalB).String()

	total_cost := receipt1.GasUsed + receipt2.GasUsed
	ecomm.UpdateLog(logInfoFile, a.AssetID, ecomm.BidEvent, keyWords, total_cost, note)
}

func bidAuctionH(auction_id int, bidAmount *big.Int, bid_key string) {
	t := time.Now()

	var contract_info ecomm.ContractInfo
	ecomm.ReadJsonFile(contractInfoFile, &contract_info)
	//erc20_address := contract_info.EthERC20
	client := ethClient

	a, err := assetClient.GetAuction(auction_id)
	check(err)

	auction_addr := common.HexToAddress(a.EthAddr)
	// @todo: require platform either is 'quo' or 'eth'
	if platform != "eth" {
		auction_addr = common.HexToAddress(a.QuorumAddr)
		client = quoClient
		//erc20_address = contract_info.QuoERC20
	}

	var auction_contract ecomm.AuctionContractCloseBid
	//var auction_contract_close_bid ecomm.AuctionContractCloseBid

	switch auc_type {
	case "cb1p":
		auction_contract, err = cb1p_auction.NewCb1pAuction(auction_addr, client)
	case "cb2p":
		auction_contract, err = cb2p_auction.NewCb2pAuction(auction_addr, client)
		// case "cb1p":
		// 	auction_contract_cb, err = cb1p_auction.NewCb1pAuction(auction_addr, client)
	default:
		log.Printf("Incorrect Auction Type!")
		return
	}

	check(err)

	bidT, err := cclib.NewTransactor(bid_key, "password")
	check(err)

	keyWords := fmt.Sprintf("%s_%s_%s", platform, bidT.From.String()[36:], bidAmount)
	ecomm.LogEvent(logInfoFile, a.AssetID, ecomm.BidEvent, keyWords, t, "", 0)

	// Compute the hash of the bid amount
	// Solidity equivalent: keccak256(abi.encodePacked(bidAmount))
	bidHash := crypto.Keccak256Hash(bidAmount.Bytes())

	// Convert bidHash to [32]byte to match the Go binding's expectation
	var bidHashArray [32]byte
	copy(bidHashArray[:], bidHash.Bytes())

	tx, _ := auction_contract.Bid(bidT, big.NewInt(int64(auction_id)), bidHashArray)
	receipt := ecomm.WaitTx(client, tx, fmt.Sprintf("Bid on Auction ID: %d through contract: %s", a.AuctionID, auction_addr))

	fmt.Println("Transaction receipt:", receipt)

}

func withdraw(auction_id int, bid_key string) {
	t := time.Now()
	client := ethClient

	auction, err := assetClient.GetAuction(auction_id)
	check(err)

	auction_addr := common.HexToAddress(auction.EthAddr)
	if platform != "eth" {
		auction_addr = common.HexToAddress(auction.QuorumAddr)
		client = quoClient
	}

	bidT, err := cclib.NewTransactor(bid_key, "password")
	check(err)
	//eventID := auction.AssetID + "_" + platform + "_" + bidT.From.String()[36:]
	//@todo withdraw bidAmount
	keyWords := fmt.Sprintf("%s_%s", platform, bidT.From.String()[36:])

	ecomm.LogEvent(logInfoFile, auction.AssetID, ecomm.WithdrawEvent, keyWords, t, "", 0)

	// same interface for all 4 kinds auction contracts
	// @todo support all 4 auctions
	auction_contract, err := english_auction.NewEnglishAuction(auction_addr, client)
	check(err)

	tx, err := auction_contract.Withdraw(bidT, big.NewInt(int64(auction_id)))
	check(err)
	receipt := ecomm.WaitTx(client, tx, fmt.Sprintf("Withdraw bid on Auction ID: %d through contract: %s", auction.AuctionID, auction_addr))

	ecomm.UpdateLog(logInfoFile, auction.AssetID, ecomm.WithdrawEvent, keyWords, receipt.GasUsed, "")
	//debugTransaction(tx)
	// log
	/////////////
}

func check_winner(auction_id int, bid_key string) {
	client := ethClient
	bidT, err := cclib.NewTransactor(bid_key, password)
	check(err)

	// Get Auction Contract deployed on Eth/Quo
	//assetClient := ecomm.NewAssetClient() // return Fabric asset contract
	a, err := assetClient.GetAuction(auction_id)
	check(err)

	auction_addr := common.HexToAddress(a.EthAddr)
	if platform != "eth" {
		auction_addr = common.HexToAddress(a.QuorumAddr)
		client = quoClient
	}

	auction_contract, err := english_auction.NewEnglishAuction(auction_addr, client)
	check(err)

	// Check winner
	highestBidder, err := auction_contract.HighestBidder(&bind.CallOpts{}, big.NewInt(int64(auction_id)))
	check(err)

	if bidT.From == highestBidder {
		fmt.Println("Waiting your response (abt/prcd) for Auction", auction_id)
	} else {
		fmt.Println("highest bidder:", highestBidder.Hex())
	}

	highestBid, err := auction_contract.HighestBid(&bind.CallOpts{}, big.NewInt(int64(auction_id)))
	check(err)
	fmt.Println("highest bid:", highestBid)
}

func sign_auction_result(auction_id int, prcd bool, bid_key string) {
	// @reset timer
	t := time.Now()
	cclib.LastEventTimestamp.Set(t, timeInfoFile)

	client := ethClient

	//assetClient := ecomm.NewAssetClient() // return Fabric asset contract
	a, err := assetClient.GetAuction(auction_id)
	check(err)

	auction_addr := common.HexToAddress(a.EthAddr)
	if platform != "eth" {
		auction_addr = common.HexToAddress(a.QuorumAddr)
		client = quoClient
	}

	bidT, err := cclib.NewTransactor(bid_key, "password")
	check(err)

	// log
	payload, _ := json.Marshal(a)
	if prcd {
		cclib.LogEventToFile(logInfoFile, ecomm.CommitAuctionResultEvent, payload, t, timeInfoFile)
	} else {
		cclib.LogEventToFile(logInfoFile, ecomm.AbortAuctionResultEvent, payload, t, timeInfoFile)
	}

	auction_contract, err := english_auction.NewEnglishAuction(auction_addr, client)
	check(err)

	highestBidder, _ := auction_contract.HighestBidder(&bind.CallOpts{}, big.NewInt(int64(auction_id)))
	if highestBidder != bidT.From {
		log.Panicln("Not authorized to commit the auction")
	}

	highestBid, _ := auction_contract.HighestBid(&bind.CallOpts{}, big.NewInt(int64(auction_id)))

	auction_result := &ecomm.AuctionResult{
		Platform:    platform,
		AuctionID:   auction_id,
		AuctionAddr: auction_addr.Hex(),

		HighestBid:    *highestBid,
		HighestBidder: bidT.From.Hex(),
	}

	signer, _ := cclib.NewSigner(bid_key, password)
	sig, err := signer.Sign(auction_result.Hash())
	check(err)

	auction_result.Signature = sig
	jsonString, err := json.Marshal(auction_result)
	check(err)

	t = time.Now()

	var tx *types.Transaction
	var p_type string
	if prcd {
		tx, err = auction_contract.Commit(bidT, big.NewInt(int64(auction_id)), string(jsonString))
		p_type = "Commit"
		cclib.LogEventToFile(logInfoFile, ecomm.CommitAuctionResultEvent, jsonString, t, timeInfoFile)
	} else {
		tx, err = auction_contract.Abort(bidT, big.NewInt(int64(auction_id)), string(jsonString))
		p_type = "Abort"
		cclib.LogEventToFile(logInfoFile, ecomm.AbortAuctionResultEvent, jsonString, t, timeInfoFile)
	}
	check(err)

	receipt := ecomm.WaitTx(client, tx, fmt.Sprintf("Sign Auction Result on Auction ID: %d through contract: %s", a.AuctionID, auction_addr))
	//debugTransaction(tx)
	// log
	payload, _ = json.Marshal(&ecomm.Tx{
		Platform: platform,
		Type:     p_type,
		Hash:     receipt.TxHash,
	})

	t = time.Now()
	cclib.LogEventToFile(logInfoFile, ecomm.TransactionMinedEvent, payload, t, timeInfoFile)
}

func provide_feedback(auction_id int, feedback string, bid_key string) {
	// @reset timer
	t := time.Now()
	cclib.LastEventTimestamp.Set(t, timeInfoFile)

	parts := strings.SplitN(feedback, "@", 2)
	score, err := strconv.Atoi(parts[0])
	check(err)

	if len(parts) == 1 {
		feedback = ""
	} else {
		feedback = parts[1]
	}

	client := ethClient
	a, err := assetClient.GetAuction(auction_id)
	check(err)

	auction_addr := common.HexToAddress(a.EthAddr)
	if platform != "eth" {
		auction_addr = common.HexToAddress(a.QuorumAddr)
		client = quoClient
	}

	bidT, err := cclib.NewTransactor(bid_key, "password")
	check(err)

	// log
	payload, _ := json.Marshal(a)
	cclib.LogEventToFile(logInfoFile, ecomm.ProvideFeedbackEvent, payload, t, timeInfoFile)

	auction_contract, err := english_auction.NewEnglishAuction(auction_addr, client)
	check(err)

	tx, _ := auction_contract.ProvideFeedback(bidT, big.NewInt(int64(auction_id)), big.NewInt(int64(score)), feedback)
	receipt := ecomm.WaitTx(client, tx, fmt.Sprintf("Provide feedback on Auction ID: %d", a.AuctionID))
	t = time.Now()

	payload, _ = json.Marshal(&ecomm.Tx{
		Platform: platform,
		Type:     "Feedback",
		Hash:     receipt.TxHash,
	})

	cclib.LogEventToFile(logInfoFile, ecomm.TransactionMinedEvent, payload, t, timeInfoFile)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
