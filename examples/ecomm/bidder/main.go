package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	zkNodes = "localhost:2181"
	//ethEndpoint = "localhost:8545"
	platform    = "ethereum"
	signerID    = "1"
	keyfile     = "../../keys/key1"
	keypassword = "password"

	ccsvc     *cclib.CCService
	ethClient *ethclient.Client
	quoClient *ethclient.Client
	//signer    *cclib.Signer
)

const (
	rootKey      = "../../keys/key0"
	auctionerKey = "../../keys/key1"
	bidder1Key   = "../../keys/key2"
	bidder2Key   = "../../keys/key3"
	password     = "password"

	fabricTokenName = "MDAI"

	setupInfoFile = "../setup_info.json"
)

func main() {
	createTopic := false

	var err error

	ethClient, err = ethclient.Dial(fmt.Sprintf("http://%s", "localhost:8545"))
	check(err)
	quoClient, err = ethclient.Dial(fmt.Sprintf("http://%s", "localhost:8546"))
	check(err)

	flag.StringVar(&zkNodes, "zk", zkNodes, "comma separated zoolkeeper nodes")
	flag.StringVar(&ethEndpoint, "eth", ethEndpoint, "eth endpoint")
	flag.StringVar(&platform, "p", platform, "platform")
	flag.StringVar(&signerID, "id", signerID, "signer id")
	flag.StringVar(&keyfile, "key", keyfile, "private key file")
	flag.BoolVar(&createTopic, "t", createTopic, "create kafka topic")
	flag.Parse()

	command := flag.String("c", "", "command")
	flag.Parse()

	switch *command {
	case "init":
		initialize()
	case "add":
		new_bidder()
	default:
		fmt.Println("command not found")
	}

}

func initialize() {
	_, err := cclib.NewSigner(bidder1Key, keypassword)
	check(err)

	ccsvc, err = cclib.NewEventService(
		strings.Split(zkNodes, ","),
		fmt.Sprintf("signer-%s-%s", platform, signerID),
	)
	check(err)

	ccsvc.Register(ecomm.AuctionEndingEvent, handleAuctionEnding)

	err = ccsvc.Start(true)
	check(err)

	select {}
}

func new_bidder() {

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
