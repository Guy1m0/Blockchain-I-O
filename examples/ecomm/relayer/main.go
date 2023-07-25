package main

import (
	"flag"
	"strings"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
)

var (
	zkNodes = "localhost:2181"

	assetClient *ecomm.AssetClient
	ccsvc       *cclib.CCService

	auctionResults map[int]*ecomm.FinalizeAuctionArgs
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
	ccsvc.Register(ecomm.SignedAuctionResultEvent, handleSignedAuctionResult)
	err = ccsvc.Start(true)
	check(err)

	runAuctionListener()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
