package main

import (
	"flag"
	"strings"
	"sync"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	zkNodes = "localhost:2181"

	assetClient *ecomm.AssetClient
	//fabric_network *gateway.Network
	ccsvc *cclib.CCService

	auctionResults   map[int]*ecomm.FinalizeAuctionArgs
	auctionResultsMu sync.Mutex

	ethClient *ethclient.Client
	quoClient *ethclient.Client

	eth_ERC20    common.Address
	quo_ERC20    common.Address
	fabric_ERC20 string
)

const (
	platform = "eth"
	root_key = "../../keys/key0"
	password = "password"

	erc20InfoFile = "../erc20_info.json"
	logInfoFile   = "../log.csv"
)

func main() {
	// flag.StringVar(&platform, "p", platform, "Monitors wich platform")

	flag.StringVar(&zkNodes, "zk", zkNodes, "comma separated zoolkeeper nodes")
	flag.Parse()

	// Initialize
	auctionResults = make(map[int]*ecomm.FinalizeAuctionArgs)

	assetClient = ecomm.NewAssetClient()
	ethClient = ecomm.NewEthClient()
	quoClient = ecomm.NewQuorumClient()

	ccsvc, _ = cclib.NewEventService(strings.Split(zkNodes, ","), "relayer") //zookeeper node
	//check(err)

	var erc20_info ecomm.Erc20Info
	ecomm.ReadJsonFile(erc20InfoFile, &erc20_info)

	eth_ERC20 = erc20_info.EthERC20
	quo_ERC20 = erc20_info.QuoERC20
	fabric_ERC20 = erc20_info.FabricTokenName

	// @todo: separate ccsvc relayer in different platorms
	// and can not directly detects events on other platform if
	// and only if such events published by 'local' relayer

	// @todo: deploy 3 relayers monitor Fabric, Eth and Quo respectively
	ccsvc.Register(ecomm.AssetAddingEvent, logEvent)
	ccsvc.Register(ecomm.BidEvent, logEvent)
	ccsvc.Register(ecomm.AuctionStartingEvent, logEvent)
	ccsvc.Register(ecomm.AuctionClosingEvent, logEvent)
	ccsvc.Register(ecomm.AuctionCancelingEvent, logEvent)
	ccsvc.Register(ecomm.AuctionFinalizingEvent, logEvent)

	err := ccsvc.Start(true)
	check(err)

	startListeningForAssetEvents(assetClient)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
