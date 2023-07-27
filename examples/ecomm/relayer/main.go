package main

import (
	"flag"
	"strings"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/contracts/eth_stable_coin"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	zkNodes = "localhost:2181"

	assetClient *ecomm.AssetClient
	ccsvc       *cclib.CCService

	auctionResults map[int]*ecomm.FinalizeAuctionArgs

	ethClient *ethclient.Client
	quoClient *ethclient.Client

	//eth_ERC20 eth_stable_coin
	//quo_ERC20 eth_stable_coin
)

const (
	rootKey      = "../../keys/key0"
	auctionerKey = "../../keys/key1"
	bidder1Key   = "../../keys/key2"
	bidder2Key   = "../../keys/key3"
	password     = "password"

	fabricTokenName = "MDAI1"

	setupInfoFile = "../setup_info.json"
)

func main() {
	flag.StringVar(&zkNodes, "zk", zkNodes, "comma separated zoolkeeper nodes")
	flag.Parse()

	auctionResults = make(map[int]*ecomm.FinalizeAuctionArgs)
	assetClient = ecomm.NewAssetClient()

	var err error

	ccsvc, err = cclib.NewEventService(strings.Split(zkNodes, ","), "relayer") //zookeeper node
	check(err)

	// Register("event_", event_handler)
	// SignedAuctionResultEvent is the one when bidder accept such auction
	ccsvc.Register(ecomm.SignedAuctionResultEvent, handleSignedAuctionResult)

	// Why not create a new event for new auction?
	err = ccsvc.Start(true)
	check(err)

	runAuctionListener()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func load_setup_info() {
	var setupInfo ecomm.SetupInfo
	ecomm.ReadJsonFile(setupInfoFile, &setupInfo)

	eth_ERC20, err := eth_stable_coin.NewEthStableCoin(setupInfo.EthERC20, ethClient)
	check(err)

	quo_ERC20, err := eth_stable_coin.NewEthStableCoin(setupInfo.QuoERC20, quoClient)
	check(err)
}
