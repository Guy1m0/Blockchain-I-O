package main

import (
	"flag"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/contracts/eth_auction"
	"github.com/Guy1m0/Blockchain-I-O/contracts/eth_stable_coin"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	zkNodes  = "localhost:2181"
	Endpoint = "localhost:8545"
	platform = "eth"
	key_path = "../../keys/"
	keyfile  = "key2"
	//keypassword = "password"

	// ccsvc  *cclib.CCService
	// client *ethclient.Client
	// signer *cclib.Signer
)

const (
	// rootKey      = "../../keys/key0"
	// auctionerKey = "../../keys/key1"
	// bidder1Key   = "../../keys/key2"
	// bidder2Key   = "../../keys/key3"
	//password = "password"

	// fabricTokenName = "MDAI"

	setupInfoFile  = "../setup_info.json"
	erc20InfoFile  = "../erc20_info.json"
	BidderInfoFile = "./bidder_info.json"
)

func main() {
	//var setupInfo ecomm.SetupInfo
	//ecomm.ReadJsonFile(setupInfoFile, &setupInfo)

	// var err error
	// ethClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8545", "localhost"))
	// check(err)

	// quoClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8546", "localhost"))
	// check(err)

	//assetClient = ecomm.NewAssetClient() // return Fabric asset contract

	command := flag.String("c", "", "command")
	id_ := flag.String("id", "", "Asset ID")
	value_ := flag.String("v", "", "Bid value")

	flag.StringVar(&platform, "p", platform, "platform")
	flag.StringVar(&keyfile, "k", keyfile, "private key file")
	//keyfile := flag.String("k", "", "private key file")

	flag.Parse()

	switch *command {
	case "load":
		load(platform, keyfile)
	case "bid":
		value := new(big.Int)
		value.SetString(*value_, 10)
		id, _ := strconv.Atoi(*id_)
		bidAuction(id, value)
		// case "sign":

		// case "abort":

	}

	// fmt.Println("[ethereum] Bidding auction")
	// bidAuction(ethClient, myAuction.EthAddr, "../../keys/key1", 500)

	// fmt.Println("[quorum] Bidding auction")
	// bidAuction(quorumClient, myAuction.QuorumAddr, "../../keys/key2", 1000)

	// fmt.Println("[fabric] Ending auction")
	// endAuction(myAuction)
}

// func initialize() {
// 	ccsvc, err := cclib.NewEventService(
// 		strings.Split(zkNodes, ","),
// 		fmt.Sprintf("bidder"),
// 	)
// 	check(err)

// 	ccsvc.Register(ecomm.AuctionCreatingEvent, handleAuctionCreating)
// 	ccsvc.Register(ecomm.AuctionEndingEvent, handleAuctionEnding)

// 	err = ccsvc.Start(true)
// 	check(err)
// }

func load(platform string, key string) {
	var setupInfo ecomm.SetupInfo
	ecomm.ReadJsonFile(setupInfoFile, &setupInfo)

	keyfile := fmt.Sprintf("%s%s", key_path, key)
	erc20 := setupInfo.EthERC20
	//err new(error)
	if platform == "eth" {
		Endpoint = fmt.Sprintf("http://%s:8545", "localhost")
	} else {
		Endpoint = fmt.Sprintf("http://%s:8546", "localhost")
		erc20 = setupInfo.QuoERC20
	}

	ecomm.WriteJsonFile(BidderInfoFile, ecomm.BidderInfo{
		ZkNodes:  zkNodes,
		Endpoint: Endpoint,
		Keyfile:  keyfile,
		Platform: platform,
		Erc20:    erc20,
	})

	// not catch error
	//client, _ = ethclient.Dial(Endpoint)
	//signer, _ = cclib.NewSigner(fmt.Sprintf("%s%s", key_path, keyfile), keypassword)
	//check(err)

}

// also no relayer involved, 'locally' make bid
func bidAuction(id int, value *big.Int) {
	var bidderInfo ecomm.BidderInfo
	ecomm.ReadJsonFile(BidderInfoFile, &bidderInfo)

	client, err := ethclient.Dial(bidderInfo.Endpoint)
	check(err)

	bidT, err := cclib.NewTransactor(bidderInfo.Keyfile, "password")
	check(err)

	// Get Auction Contract deployed on Eth/Quo
	assetClient := ecomm.NewAssetClient() // return Fabric asset contract
	a, err := assetClient.GetAuction(id)
	check(err)

	Auction_addr_ := a.EthAddr
	if bidderInfo.Platform != "eth" {
		Auction_addr_ = a.QuorumAddr
	}
	Auction_addr := common.HexToAddress(Auction_addr_)

	// Approve amount of bid through ERC20 contract
	MDAI, _ := eth_stable_coin.NewEthStableCoin(bidderInfo.Erc20, client)
	tx, _ := MDAI.Approve(bidT, Auction_addr, value)
	ecomm.WaitTx(client, tx, "Approve Auction Contract's allowance")

	Auction, err := eth_auction.NewEthAuction(Auction_addr, client)
	check(err)

	tx, err = Auction.Bid(bidT, value)
	check(err)
	ecomm.WaitTx(client, tx, "Bid on Asset/Auction Contract")

	highestBidder, err := Auction.HighestBidder(&bind.CallOpts{})
	check(err)
	fmt.Println("highest bidder:", highestBidder.Hex())

	highestBid, err := Auction.HighestBid(&bind.CallOpts{})
	check(err)
	fmt.Println("highest bid:", highestBid)
}

func load_ctcs() (*ethclient.Client, *bind.TransactOpts, *cclib.CCService, *cclib.Signer) {
	var bidderInfo ecomm.BidderInfo
	ecomm.ReadJsonFile(BidderInfoFile, &bidderInfo)

	client, err := ethclient.Dial(bidderInfo.Endpoint)
	check(err)

	bidT, err := cclib.NewTransactor(bidderInfo.Keyfile, "password")
	check(err)

	signer, _ := cclib.NewSigner(fmt.Sprintf("%s%s", key_path, keyfile), "password")

	ccsvc, err := cclib.NewEventService(
		strings.Split(zkNodes, ","),
		fmt.Sprintf("bidder"),
	)
	check(err)

	return client, bidT, ccsvc, signer

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
