package main

import (
	"flag"
	"strings"
	"sync"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/contracts/cb1p_auction"
	"github.com/Guy1m0/Blockchain-I-O/contracts/english_auction"
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

	english_auction_info ecomm.AuctionInfo
	cb1p_auction_info    ecomm.AuctionInfo

	eth_english_auction_contract english_auction.EnglishAuction
	eth_cb1p_contract            cb1p_auction.Cb1pAuction
)

const (
	platform = "eth"
	root_key = "../../keys/key0"
	password = "password"

	contractInfoFile = "../contract_info.json"
	logInfoFile      = "../log.csv"
)

func main() {
	// flag.StringVar(&platform, "p", platform, "Monitors wich platform")

	flag.StringVar(&zkNodes, "zk", zkNodes, "comma separated zoolkeeper nodes")
	flag.Parse()

	// Initialize
	auctionResults = make(map[int]*ecomm.FinalizeAuctionArgs)

	assetClient = ecomm.NewAssetClient()
	// ethClient = ecomm.NewEthClient()
	// quoClient = ecomm.NewQuorumClient()

	ccsvc, _ = cclib.NewEventService(strings.Split(zkNodes, ","), "relayer") //zookeeper node
	//check(err)

	var contract_info ecomm.ContractInfo
	ecomm.ReadJsonFile(contractInfoFile, &contract_info)

	eth_ERC20 = contract_info.EthERC20
	quo_ERC20 = contract_info.QuoERC20
	fabric_ERC20 = contract_info.FabricTokenName
	english_auction_info = contract_info.EnglishAuction

	// eth_english_auction_contract, err := english_auction.NewEnglishAuction(english_auction_info.EthAddr, ethClient)
	// quo_english_auction_contract, err := english_auction.NewEnglishAuction(english_auction_info.QuoAddr, quoClient)

	// eth_cb1p_contract, err  := cb1p_auction.NewCb1pAuction(cb1p_auction_info.EthAddr, ethClient)
	// quo_cb1p_contract, err  := cb1p_auction.NewCb1pAuction(cb1p_auction_info.QuoAddr, quoClient)
	// @todo: separate ccsvc relayer in different platorms
	// and can not directly detects events on other platform if
	// and only if such events published by 'local' relayer

	// @todo: deploy 3 relayers monitor Fabric, Eth and Quo respectively
	ccsvc.Register(ecomm.AssetAddingEvent, chainCodeEvent)
	ccsvc.Register(ecomm.AuctionStartingEvent, chainCodeEvent)
	ccsvc.Register(ecomm.AuctionClosingEvent, chainCodeEvent)
	ccsvc.Register(ecomm.AuctionCancelingEvent, chainCodeEvent)
	ccsvc.Register(ecomm.AuctionFinalizingEvent, chainCodeEvent)

	ccsvc.Register(ecomm.BidEvent, smartContractEvent)
	ccsvc.Register(ecomm.WithdrawEvent, smartContractEvent)

	err := ccsvc.Start(true)
	check(err)

	startFabricListener(assetClient)
	// go startAuctionListener("english_auction", english_auction.EthAddr.String(), "eth")
	// go startAuctionListener("english_auction", english_auction.QuoAddr.String(), "quo")
	// startListeningForAuctionEvents()
	// startListeningForAuctionEvents()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
