package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/contracts/eth_auction"
	"github.com/Guy1m0/Blockchain-I-O/contracts/eth_stable_coin"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	zkNodes = "localhost:2181"
	//Endpoint = "localhost:8545"
	platform = "eth"

	bidder_name = "Bidder 1"
	bid_key     string

	//keypassword = "password"

	ccsvc     *cclib.CCService
	ethClient *ethclient.Client
	quoClient *ethclient.Client

	assetClient *ecomm.AssetClient
	// signer *cclib.Signer
)

const (
	// rootKey      = "../../keys/key0"
	// auctionerKey = "../../keys/key1"
	// bidder1Key   = "../../keys/key2"
	// bidder2Key   = "../../keys/key3"
	password = "password"

	// fabricTokenName = "MDAI"
	//default_name = "Bidder 1"
	//setupInfoFile  = "../setup_info.json"
	erc20InfoFile = "../erc20_info.json"
	userInfoFile  = "../user_info.json"
	logInfoFile   = "../log.json"
	timeInfoFile  = "../timer"
	//BidderInfoFile = "./bidder_info.json"
)

func main() {
	//var setupInfo ecomm.SetupInfo
	//ecomm.ReadJsonFile(setupInfoFile, &setupInfo)

	var err error
	ethClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8545", "localhost"))
	check(err)

	quoClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8546", "localhost"))
	check(err)

	//assetClient = ecomm.NewAssetClient() // return Fabric asset contract

	ccsvc, err = cclib.NewEventService(strings.Split(zkNodes, ","), "bidder")
	check(err)

	assetClient = ecomm.NewAssetClient()

	command := flag.String("c", "", "command")

	id_ := flag.String("id", "", "Auction ID")
	amount_ := flag.String("amt", "", "Bid amount")
	score_ := flag.String("s", "", "Auction rating")
	feedback := flag.String("fb", "", "Detail feedback")
	flag.StringVar(&platform, "p", platform, "platform")
	// flag.StringVar(&bid_key, "k", bid_key, "private key file")
	flag.StringVar(&bidder_name, "name", bidder_name, "Load Bidder Information")
	flag.Parse()

	bid_key = load_bidder_key(bidder_name)

	switch *command {
	// case "login":
	// 	login(platform, bidder_name)
	case "bid":
		amount := new(big.Int)
		amount.SetString(*amount_, 10)
		id, _ := strconv.Atoi(*id_)

		bidAuction(id, amount)
	case "check":
		id, _ := strconv.Atoi(*id_)
		check_winner(id)
	case "prcd":
		id, _ := strconv.Atoi(*id_)
		sign_auction_result(id, true)
	case "abt":
		id, _ := strconv.Atoi(*id_)
		sign_auction_result(id, false)
	case "with":
		id, _ := strconv.Atoi(*id_)
		withdraw(id)
	case "rate":
		id, _ := strconv.Atoi(*id_)
		score, _ := strconv.Atoi(*score_)
		provide_feedback(id, score, *feedback)
	}

	// fmt.Println("[ethereum] Bidding auction")
	// bidAuction(ethClient, myAuction.EthAddr, "../../keys/key1", 500)

	// fmt.Println("[quorum] Bidding auction")
	// bidAuction(quorumClient, myAuction.QuorumAddr, "../../keys/key2", 1000)

	// fmt.Println("[fabric] Ending auction")
	// endAuction(myAuction)
}

// also no relayer involved, 'locally' make bid
func bidAuction(auction_id int, amount *big.Int) {

	// if auction addr is already known
	// Get Auction Contract deployed on Eth/Quo
	var erc20_info ecomm.Erc20Info
	ecomm.ReadJsonFile(erc20InfoFile, &erc20_info)
	erc20_address := erc20_info.EthERC20

	client := ethClient

	// return Fabric asset contract
	a, err := assetClient.GetAuction(auction_id)
	check(err)

	auction_addr := common.HexToAddress(a.EthAddr)
	if platform != "eth" {
		auction_addr = common.HexToAddress(a.QuorumAddr)
		client = quoClient
		erc20_address = erc20_info.QuoERC20
	}

	auction_contract, err := eth_auction.NewEthAuction(auction_addr, client)
	check(err)

	bidT, err := cclib.NewTransactor(bid_key, "password")
	check(err)

	// log
	payload, _ := json.Marshal(&ecomm.Bid{
		Bidder:      bidT.From,
		BidAmount:   *amount,
		AuctionAddr: auction_addr,
		Platform:    platform,
		AuctionID:   auction_id,
		AssetID:     a.AssetID,
	})

	// @reset timer
	t := time.Now()
	cclib.LastEventTimestamp.Set(t, timeInfoFile)
	cclib.LogEventToFile(logInfoFile, ecomm.BiddingAuctionEvent, payload, t, timeInfoFile)

	// Approve amount of bid through ERC20 contract
	MDAI, _ := eth_stable_coin.NewEthStableCoin(erc20_address, client)
	tx, _ := MDAI.Approve(bidT, auction_addr, big.NewInt(0).Mul(big.NewInt(amount.Int64()), ecomm.DecimalB))
	ecomm.WaitTx(client, tx, "Approve Auction Contract's allowance")

	// alw, err := MDAI.Allowance(&bind.CallOpts{}, bidT.From, auction_addr)
	tx, err = auction_contract.Bid(bidT, big.NewInt(0).Mul(big.NewInt(amount.Int64()), ecomm.DecimalB))
	check(err)
	receipt := ecomm.WaitTx(client, tx, fmt.Sprintf("Bid on Auction ID: %d through contract: %s", a.ID, auction_addr))
	//debugTransaction(tx)

	// log
	t = time.Now()
	payload, _ = json.Marshal(&ecomm.Tx{
		Platform: platform,
		Type:     "Bid",
		Hash:     receipt.TxHash,
	})
	cclib.LogEventToFile(logInfoFile, ecomm.TransactionMinedEvent, payload, t, timeInfoFile)
	cclib.LastEventTimestamp.Set(t, timeInfoFile)
}

func check_winner(auction_id int) {
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

	auction_contract, err := eth_auction.NewEthAuction(auction_addr, client)
	check(err)

	// Check winner
	highestBidder, err := auction_contract.HighestBidder(&bind.CallOpts{})
	check(err)

	if bidT.From == highestBidder {
		fmt.Println("Waiting your response (abt/prcd) for Auction", auction_id)
	} else {
		fmt.Println("highest bidder:", highestBidder.Hex())
	}

	highestBid, err := auction_contract.HighestBid(&bind.CallOpts{})
	check(err)
	fmt.Println("highest bid:", highestBid)
}

// func procd_auction(auction_id int) {
// 	// Check Status of Auctio

// 	client := ethClient

// 	//assetClient := ecomm.NewAssetClient() // return Fabric asset contract
// 	a, err := assetClient.GetAuction(auction_id)
// 	check(err)

// 	auction_addr := common.HexToAddress(a.EthAddr)
// 	if platform != "eth" {
// 		auction_addr = common.HexToAddress(a.QuorumAddr)
// 		client = quoClient
// 	}

// 	// @reset timer
// 	t := time.Now()
// 	cclib.LastEventTimestamp.Set(t, timeInfoFile)

// 	bidT, err := cclib.NewTransactor(bid_key, "password")
// 	check(err)

// 	// log
// 	payload, _ := json.Marshal(a)
// 	cclib.LogEventToFile(logInfoFile, ecomm.ProceedAuctionResultEvent, payload, t, timeInfoFile)

// 	auction_contract, err := eth_auction.NewEthAuction(auction_addr, client)
// 	check(err)

// 	highestBidder, _ := auction_contract.HighestBidder(&bind.CallOpts{})
// 	if highestBidder != bidT.From {
// 		log.Panicln("Not authorized to proceed the auction")
// 	}

// 	highestBid, _ := auction_contract.HighestBid(&bind.CallOpts{})

// 	auction_result := &ecomm.AuctionResult{
// 		Platform:    platform,
// 		AuctionID:   auction_id,
// 		AuctionAddr: auction_addr.Hex(),

// 		HighestBid:    int(highestBid.Int64()),
// 		HighestBidder: bidT.From.Hex(),
// 	}

// 	signer, _ := cclib.NewSigner(bid_key, password)
// 	sig, err := signer.Sign(auction_result.Hash())
// 	check(err)

// 	auction_result.Signature = sig
// 	jsonString, err := json.Marshal(auction_result)
// 	check(err)

// 	tx, err := auction_contract.Proceed(bidT, string(jsonString))
// 	check(err)
// 	receipt := ecomm.WaitTx(client, tx, fmt.Sprintf("Commit to vote on Auction ID: %d through contract: %s", a.ID, auction_addr))
// 	//debugTransaction(tx)
// 	// log
// 	payload, _ = json.Marshal(&ecomm.Tx{
// 		Platform: platform,
// 		Type:     "Proceed",
// 		Hash:     receipt.TxHash,
// 	})

// 	t = time.Now()
// 	cclib.LogEventToFile(logInfoFile, ecomm.ProceedAuctionResultEvent, payload, t, timeInfoFile)

// }

func sign_auction_result(auction_id int, prcd bool) {
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
		cclib.LogEventToFile(logInfoFile, ecomm.ProceedAuctionResultEvent, payload, t, timeInfoFile)
	} else {
		cclib.LogEventToFile(logInfoFile, ecomm.AbortAuctionResultEvent, payload, t, timeInfoFile)
	}

	auction_contract, err := eth_auction.NewEthAuction(auction_addr, client)
	check(err)

	highestBidder, _ := auction_contract.HighestBidder(&bind.CallOpts{})
	if highestBidder != bidT.From {
		log.Panicln("Not authorized to proceed the auction")
	}

	highestBid, _ := auction_contract.HighestBid(&bind.CallOpts{})

	auction_result := &ecomm.AuctionResult{
		Platform:    platform,
		AuctionID:   auction_id,
		AuctionAddr: auction_addr.Hex(),

		HighestBid:    int(highestBid.Int64()),
		HighestBidder: bidT.From.Hex(),
	}

	signer, _ := cclib.NewSigner(bid_key, password)
	sig, err := signer.Sign(auction_result.Hash())
	check(err)

	auction_result.Signature = sig
	jsonString, err := json.Marshal(auction_result)
	check(err)

	t = time.Now()
	cclib.LogEventToFile(logInfoFile, ecomm.SignedAuctionResultEvent, jsonString, t, timeInfoFile)

	var tx *types.Transaction
	var p_type string
	if prcd {
		tx, err = auction_contract.Proceed(bidT, string(jsonString))
		p_type = "Proceed"
	} else {
		tx, err = auction_contract.Abort(bidT, string(jsonString))
		p_type = "Abort"
	}
	check(err)

	receipt := ecomm.WaitTx(client, tx, fmt.Sprintf("Sign Auction Result on Auction ID: %d through contract: %s", a.ID, auction_addr))
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

func provide_feedback(auction_id int, score int, feedback string) {
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

	cclib.LogEventToFile(logInfoFile, ecomm.FeedBackEvent, payload, t, timeInfoFile)

	auction_contract, err := eth_auction.NewEthAuction(auction_addr, client)
	check(err)

	tx, _ := auction_contract.ProvideFeedback(bidT, big.NewInt(int64(score)), feedback)
	receipt := ecomm.WaitTx(client, tx, fmt.Sprintf("Provide feedback on Auction ID: %d", a.ID))
	//debugTransaction(tx)
	// log

	payload, _ = json.Marshal(&ecomm.Tx{
		Platform: platform,
		Type:     "Feedback",
		Hash:     receipt.TxHash,
	})

	t = time.Now()
	cclib.LogEventToFile(logInfoFile, ecomm.TransactionMinedEvent, payload, t, timeInfoFile)
}

func withdraw(auction_id int) {
	client := ethClient

	//assetClient := ecomm.NewAssetClient() // return Fabric asset contract
	a, err := assetClient.GetAuction(auction_id)
	check(err)

	auction_addr := common.HexToAddress(a.EthAddr)
	if platform != "eth" {
		auction_addr = common.HexToAddress(a.QuorumAddr)
		client = quoClient
	}

	// @reset timer
	t := time.Now()
	cclib.LastEventTimestamp.Set(t, timeInfoFile)

	bidT, err := cclib.NewTransactor(bid_key, "password")
	check(err)

	// log
	payload, _ := json.Marshal(a)
	cclib.LogEventToFile(logInfoFile, ecomm.WithdrawEvent, payload, t, timeInfoFile)

	auction_contract, err := eth_auction.NewEthAuction(auction_addr, client)
	check(err)

	tx, err := auction_contract.Withdraw(bidT)
	check(err)
	receipt := ecomm.WaitTx(client, tx, fmt.Sprintf("Commit to vote on Auction ID: %d through contract: %s", a.ID, auction_addr))
	//debugTransaction(tx)
	// log
	payload, _ = json.Marshal(&ecomm.Tx{
		Platform: platform,
		Type:     "Withdraw",
		Hash:     receipt.TxHash,
	})

	t = time.Now()
	cclib.LogEventToFile(logInfoFile, ecomm.WithdrawEvent, payload, t, timeInfoFile)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
