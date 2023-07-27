package main

import (
	"flag"
	"fmt"
	"math/big"
	"strings"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/contracts/eth_auction"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	zkNodes     = "localhost:2181"
	Endpoint    = "localhost:8545"
	platform    = "eth"
	key_path    = "../../keys/"
	keyfile     = "key2"
	keypassword = "password"

	ccsvc  *cclib.CCService
	client *ethclient.Client
	//	quoClient *ethclient.Client
	signer *cclib.Signer

	erc20 common.Address
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
	var setupInfo ecomm.SetupInfo
	ecomm.ReadJsonFile(setupInfoFile, &setupInfo)

	// var err error
	// ethClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8545", "localhost"))
	// check(err)

	// quoClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8546", "localhost"))
	// check(err)

	//assetClient = ecomm.NewAssetClient() // return Fabric asset contract

	command := flag.String("c", "", "command")
	id := flag.String("a", "", "Asset ID")

	flag.StringVar(&platform, "p", platform, "platform")
	flag.StringVar(&keyfile, "key", keyfile, "private key file")
	//keyfile := flag.String("k", "", "private key file")

	flag.Parse()

	switch *command {
	case "init":
		initialize()
	case "load":
		load(platform, keyfile)
	case "bid":
		bidAuction(*id)
	case "sign":

	case "abort":

	}

	// fmt.Println("[ethereum] Bidding auction")
	// bidAuction(ethClient, myAuction.EthAddr, "../../keys/key1", 500)

	// fmt.Println("[quorum] Bidding auction")
	// bidAuction(quorumClient, myAuction.QuorumAddr, "../../keys/key2", 1000)

	// fmt.Println("[fabric] Ending auction")
	// endAuction(myAuction)
}

func initialize() {
	ccsvc, err := cclib.NewEventService(
		strings.Split(zkNodes, ","),
		fmt.Sprintf("bidder"),
	)
	check(err)

	ccsvc.Register(ecomm.AuctionEndingEvent, handleAuctionEnding)

	err = ccsvc.Start(true)
	check(err)
}

func load(platform string, key string) {
	//err new(error)
	if platform == "eth" {
		Endpoint = fmt.Sprintf("http://%s:8545", "localhost")
		//erc20 = setupInfoFile.
	} else {
		Endpoint = fmt.Sprintf("http://%s:8546", "localhost")
	}

	// not catch error
	client, _ = ethclient.Dial(Endpoint)
	signer, err := cclib.NewSigner(fmt.Sprintf("%s%s", key_path, keyfile), keypassword)
	check(err)

}

// also no relayer involved, 'locally' make bid
func bidAuction(addrHex, keyfile string, value int64) {

	ecomm.AssetClient

	addr := common.HexToAddress(addrHex)
	auctionSession := newAuctionSession(addr, client, keyfile)
	auctionSession.TransactOpts.Value = big.NewInt(value)
	tx, err := auctionSession.Bid(big.NewInt(value))
	check(err)
	success, err := cclib.WaitTx(client, tx.Hash())
	check(err)
	printTxStatus(success)
	if !success {
		panic("failed to bid auction")
	}
	auctionSession.TransactOpts.Value = big.NewInt(0)

	highestBidder, err := auctionSession.HighestBidder()
	check(err)
	fmt.Println("highest bidder:", highestBidder.Hex())

	highestBid, err := auctionSession.HighestBid()
	check(err)
	fmt.Println("highest bid:", highestBid)
}

func newAuctionSession(
	addr common.Address, eth *ethclient.Client, keyfile string,
) *eth_auction.EthAuctionSession {
	auth, err := cclib.NewTransactor(keyfile, "password")
	check(err)

	cc, err := eth_auction.NewEthAuction(addr, eth)
	check(err)

	return &eth_auction.EthAuctionSession{
		Contract:     cc,
		TransactOpts: *auth,
	}
}

func new_bidder() {

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
