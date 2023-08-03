package main

import (
	"flag"
	"strings"
	"sync"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	zkNodes = "localhost:2181"

	assetClient *ecomm.AssetClient
	ccsvc       *cclib.CCService

	auctionResults   map[int]*ecomm.FinalizeAuctionArgs
	auctionResultsMu sync.Mutex

	ethClient *ethclient.Client
	quoClient *ethclient.Client

	//eth_ERC20 eth_stable_coin
	//quo_ERC20 eth_stable_coin
)

const (
	root_key = "../../keys/key0"
	//auctionerKey = "../../keys/key1"
	//bidder1Key   = "../../keys/key2"
	//bidder2Key   = "../../keys/key3"
	password = "password"

	//fabricTokenName = "MDAI1"

	//setupInfoFile = "../setup_info.json"
	erc20InfoFile = "../erc20_info.json"
	logInfoFile   = "../log.json"
)

func main() {
	flag.StringVar(&zkNodes, "zk", zkNodes, "comma separated zoolkeeper nodes")
	flag.Parse()

	// Initialize
	auctionResults = make(map[int]*ecomm.FinalizeAuctionArgs)
	assetClient = ecomm.NewAssetClient()
	ethClient = ecomm.NewEthClient()
	quoClient = ecomm.NewQuorumClient()

	ccsvc, err := cclib.NewEventService(strings.Split(zkNodes, ","), "relayer") //zookeeper node
	check(err)

	// Register("event_", event_handler)
	// SignedAuctionResultEvent is the one when bidder accept such auction
	//ccsvc.Register(ecomm.SignedAuctionResultEvent, handleSignedAuctionResult)
	//ccsvc.Register(ecomm.AuctionCreatingEvent, handleAuctionCreating)
	ccsvc.Register(ecomm.AuctionEndingEvent, handleAuctionEnding)

	// Why not create a new event for new auction?
	err = ccsvc.Start(true)
	check(err)

	// check new auction posted on Asset contract on Fabric
	runAuctionListener()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// func load_setup_info() (*eth_stable_coin.EthStableCoin, *eth_stable_coin.EthStableCoin) {
// 	var erc20_info ecomm.Erc20Info
// 	ecomm.ReadJsonFile(erc20InfoFile, &erc20_info)

// 	eth_ERC20, err := eth_stable_coin.NewEthStableCoin(erc20_info.EthERC20, ethClient)
// 	check(err)

// 	quo_ERC20, err := eth_stable_coin.NewEthStableCoin(erc20_info.QuoERC20, quoClient)
// 	check(err)

// 	return eth_ERC20, quo_ERC20
// }
