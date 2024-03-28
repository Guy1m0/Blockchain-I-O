package ecomm

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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	platform = "eth"
	auc_type = "english"

	usr_name = "Bidder 1"
	bid_key  string

	ethClient *ethclient.Client
	quoClient *ethclient.Client

	assetClient *AssetClient
)

const (
	contractInfoFile = "../tmp/contract_info.json"
	userInfoFile     = "../tmp/user_info.json"
	logInfoFile      = "../tmp/log.csv"
	timeInfoFile     = "../tmp/timer"
)

func load_bidder_key(name string) string {
	users, err := ReadUsersFromFile(userInfoFile)
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

// func load_ctcs() (*ethclient.Client, *bind.TransactOpts, *cclib.CCService, *cclib.Signer) {
// 	var bidderInfo BidderInfo
// 	ReadJsonFile(BidderInfoFile, &bidderInfo)

// 	client, err := ethclient.Dial(bidderInfo.Endpoint)
// 	check(err)

// 	bidT, err := cclib.NewTransactor(bidderInfo.Keyfile, "password")
// 	check(err)

// 	signer, _ := cclib.NewSigner(bid_key, "password")

// 	ccsvc, err := cclib.NewEventService(strings.Split(zkNodes, ","), "bidder")
// 	check(err)

// 	return client, bidT, ccsvc, signer

// }

func load_ERC20() (*stable_coin.StableCoin, *stable_coin.StableCoin, *Erc20Client) {
	var contract_info ContractInfo
	ReadJsonFile(contractInfoFile, &contract_info)

	eth_ERC20, err := stable_coin.NewStableCoin(contract_info.EthERC20, ethClient)
	check(err)

	quo_ERC20, err := stable_coin.NewStableCoin(contract_info.QuoERC20, quoClient)
	check(err)

	fabric_ERC20 := NewErc20Client(contract_info.FabricTokenName)

	return eth_ERC20, quo_ERC20, fabric_ERC20
}

// also no relayer involved, 'locally' make bid
func bidAuction(auction_id int, amount *big.Int) {
	t := time.Now()

	var contract_info ContractInfo
	ReadJsonFile(contractInfoFile, &contract_info)
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

	var auction_contract AuctionContract
	//var auction_contract_close_bid AuctionContractCloseBid

	switch auc_type {
	case "english":
		auction_contract, err = english_auction.NewEnglishAuction(auction_addr, client)
	case "dutch":
		auction_contract, err = dutch_auction.NewDutchAuction(auction_addr, client)
	}

	check(err)

	bidT, err := cclib.NewTransactor(bid_key, "password")
	check(err)

	//eventID := a.AssetID + "_" + platform + "_" + bidT.From.String()[36:]
	keyWords := fmt.Sprintf("%s_%s_%s", platform, bidT.From.String()[36:], amount)
	LogEvent(logInfoFile, a.AssetID, BidEvent, keyWords, t, "", 0)

	// @todo: Make approve and bid in a single transaction
	// Approve amount of bid through ERC20 contract
	MDAI, _ := stable_coin.NewStableCoin(erc20_address, client)
	tx1, _ := MDAI.Approve(bidT, auction_addr, big.NewInt(0).Mul(big.NewInt(amount.Int64()), DecimalB))
	receipt1 := WaitTx(client, tx1, "Approve Auction Contract's allowance")

	tx2, _ := auction_contract.Bid(bidT, big.NewInt(int64(auction_id)), big.NewInt(0).Mul(big.NewInt(amount.Int64()), DecimalB))
	receipt2 := WaitTx(client, tx2, fmt.Sprintf("Bid on Auction ID: %d through contract: %s", a.AuctionID, auction_addr))

	note := "Cost: " + strconv.FormatUint(receipt1.GasUsed, 10)
	note += " + " + strconv.FormatUint(receipt2.GasUsed, 10)
	note += " Bid: MDAI " + big.NewInt(0).Mul(big.NewInt(amount.Int64()), DecimalB).String()

	total_cost := receipt1.GasUsed + receipt2.GasUsed
	UpdateLog(logInfoFile, a.AssetID, BidEvent, keyWords, total_cost, note)
}

func bidAuctionH(auction_id int, bidAmount *big.Int) {
	t := time.Now()

	var contract_info ContractInfo
	ReadJsonFile(contractInfoFile, &contract_info)
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

	var auction_contract AuctionContractCloseBid
	//var auction_contract_close_bid AuctionContractCloseBid

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

	eventID := a.AssetID + "_" + platform + "_" + bidT.From.String()[36:]
	LogEvent(logInfoFile, BidEvent, eventID, auc_type, t, "", 0)

	// Compute the hash of the bid amount
	// Solidity equivalent: keccak256(abi.encodePacked(bidAmount))
	bidHash := crypto.Keccak256Hash(bidAmount.Bytes())

	// Convert bidHash to [32]byte to match the Go binding's expectation
	var bidHashArray [32]byte
	copy(bidHashArray[:], bidHash.Bytes())

	tx, _ := auction_contract.Bid(bidT, big.NewInt(int64(auction_id)), bidHashArray)
	receipt := WaitTx(client, tx, fmt.Sprintf("Bid on Auction ID: %d through contract: %s", a.AuctionID, auction_addr))

	fmt.Println("Transaction receipt:", receipt)

}

func withdraw(auction_id int) {
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
	eventID := auction.AssetID + "_" + platform + "_" + bidT.From.String()[36:]

	LogEvent(logInfoFile, WithdrawEvent, eventID, auc_type, t, "", 0)

	// same interface for all 4 kinds auction contracts
	auction_contract, err := english_auction.NewEnglishAuction(auction_addr, client)
	check(err)

	tx, err := auction_contract.Withdraw(bidT, big.NewInt(int64(auction_id)))
	check(err)
	receipt := WaitTx(client, tx, fmt.Sprintf("Withdraw bid on Auction ID: %d through contract: %s", auction.AuctionID, auction_addr))

	UpdateLog(logInfoFile, WithdrawEvent, eventID, auc_type, receipt.GasUsed, "")
	//debugTransaction(tx)
	// log
	/////////////
}

func check_winner(auction_id int) {
	client := ethClient
	bidT, err := cclib.NewTransactor(bid_key, password)
	check(err)

	// Get Auction Contract deployed on Eth/Quo
	//assetClient := NewAssetClient() // return Fabric asset contract
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

func sign_auction_result(auction_id int, prcd bool) {
	// @reset timer
	t := time.Now()
	cclib.LastEventTimestamp.Set(t, timeInfoFile)

	client := ethClient

	//assetClient := NewAssetClient() // return Fabric asset contract
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
		cclib.LogEventToFile(logInfoFile, CommitAuctionResultEvent, payload, t, timeInfoFile)
	} else {
		cclib.LogEventToFile(logInfoFile, AbortAuctionResultEvent, payload, t, timeInfoFile)
	}

	auction_contract, err := english_auction.NewEnglishAuction(auction_addr, client)
	check(err)

	highestBidder, _ := auction_contract.HighestBidder(&bind.CallOpts{}, big.NewInt(int64(auction_id)))
	if highestBidder != bidT.From {
		log.Panicln("Not authorized to commit the auction")
	}

	highestBid, _ := auction_contract.HighestBid(&bind.CallOpts{}, big.NewInt(int64(auction_id)))

	auction_result := &AuctionResult{
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
		cclib.LogEventToFile(logInfoFile, CommitAuctionResultEvent, jsonString, t, timeInfoFile)
	} else {
		tx, err = auction_contract.Abort(bidT, big.NewInt(int64(auction_id)), string(jsonString))
		p_type = "Abort"
		cclib.LogEventToFile(logInfoFile, AbortAuctionResultEvent, jsonString, t, timeInfoFile)
	}
	check(err)

	receipt := WaitTx(client, tx, fmt.Sprintf("Sign Auction Result on Auction ID: %d through contract: %s", a.AuctionID, auction_addr))
	//debugTransaction(tx)
	// log
	payload, _ = json.Marshal(&Tx{
		Platform: platform,
		Type:     p_type,
		Hash:     receipt.TxHash,
	})

	t = time.Now()
	cclib.LogEventToFile(logInfoFile, TransactionMinedEvent, payload, t, timeInfoFile)
}

func provide_feedback(auction_id int, feedback string) {
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
	cclib.LogEventToFile(logInfoFile, ProvideFeedbackEvent, payload, t, timeInfoFile)

	auction_contract, err := english_auction.NewEnglishAuction(auction_addr, client)
	check(err)

	tx, _ := auction_contract.ProvideFeedback(bidT, big.NewInt(int64(auction_id)), big.NewInt(int64(score)), feedback)
	receipt := WaitTx(client, tx, fmt.Sprintf("Provide feedback on Auction ID: %d", a.AuctionID))
	t = time.Now()

	payload, _ = json.Marshal(&Tx{
		Platform: platform,
		Type:     "Feedback",
		Hash:     receipt.TxHash,
	})

	cclib.LogEventToFile(logInfoFile, TransactionMinedEvent, payload, t, timeInfoFile)
}
