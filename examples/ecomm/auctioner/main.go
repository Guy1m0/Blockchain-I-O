package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/contracts/eth_auction"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	password = "password"

	zkNodes       = "localhost:2181"
	setupInfoFile = "../erc20_info.json"
)

type CreateAuctionRequest struct {
	AssetCC  []byte
	AssetID  []byte
	Platform string
}

var (
	ethClient *ethclient.Client
	quoClient *ethclient.Client

	assetClient *ecomm.AssetClient
	//ccsvc       *cclib.CCService

	asset     *ecomm.Asset
	myAuction *ecomm.Auction

	eth_ERC20 common.Address
	quo_ERC20 common.Address
)

func main() {
	// var erc20_info ecomm.Erc20Info
	// ecomm.ReadJsonFile(setupInfoFile, &erc20_info)

	// eth_ERC20 = erc20_info.EthERC20
	// quo_ERC20 = erc20_info.QuoERC20

	var err error
	ethClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8545", "localhost"))
	check(err)

	quoClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8546", "localhost"))
	check(err)

	assetClient = ecomm.NewAssetClient()

	command := flag.String("c", "", "command")
	//asset := flag.String("n", "", "Asset name")
	//id := flag.String("i", "", "Asset ID")
	flag.Parse()

	parts := strings.Split(*command, ":")
	cmd := parts[0]

	switch cmd {

	case "create":
		if len(parts) > 1 {
			args := strings.Split(parts[1], ",")
			if len(args) == 1 {
				create(args[0])
			}
		}
	case "end":
		if len(parts) > 1 {
			args := strings.Split(parts[1], ",")
			if len(args) == 1 {
				id, _ := strconv.Atoi(args[0])
				endAuction(id)
			}
		}
	case "init":
		initialize()
	default:
		fmt.Println("command not found")
	}

	//fmt.Println("[ethereum] Bidding auction")
	//bidAuction(ethClient, myAuction.EthAddr, "../../keys/key1", 500)

	//fmt.Println("[quorum] Bidding auction")
	//bidAuction(quorumClient, myAuction.QuorumAddr, "../../keys/key2", 1000)

	//fmt.Println("[fabric] Ending auction")
	//endAuction(myAuction)
}

// not need this i think
func initialize() {
	ccsvc, err := cclib.NewEventService(
		strings.Split(zkNodes, ","),
		fmt.Sprintf("auctioner"),
	)
	check(err)

	ccsvc.Register(ecomm.ProceedAuctionResultEvent, handleProceedingAuction)

	err = ccsvc.Start(true)
	check(err)
}

// Use key 1 as default auctioner
func create(asset_name string) {

	fmt.Println("[fabric] Adding asset")
	asset = addAsset(asset_name)

	fmt.Println("Starting auction")
	fmt.Println("[ethereum] Deploying auction")
	ethAddr := deployCrossChainAuction(ethClient, eth_ERC20)

	fmt.Println("[quorum] Deploying auction")
	quoAddr := deployCrossChainAuction(quoClient, quo_ERC20)

	fmt.Println("[fabric] Creating auction")
	myAuction = startAuction(asset.ID, ethAddr, quoAddr)

	// publish
	ccsvc, err := cclib.NewEventService(strings.Split(zkNodes, ","), "auctioner")
	check(err)

	payload, _ := json.Marshal(myAuction)
	ccsvc.Publish(ecomm.AuctionCreatingEvent, payload)
	// till this part, no relayer involved yet
}

// Default use ../../../keys/key1
func addAsset(id string) *ecomm.Asset {
	fmt.Println("Load Auctioner")
	aucT, err := cclib.NewTransactor("../../keys/key1", password)
	check(err)
	fmt.Println("Create New Auction")
	_, err = assetClient.AddAsset(id, aucT.From.Hex())
	check(err)
	time.Sleep(3 * time.Second)
	asset, err := assetClient.GetAsset(id)
	check(err)
	fmt.Println("Auction added on Fabric with owner: ", asset.Owner)
	return asset
}

// Auction Contract is already deployed in Fabric Network
// Just create a asset/auction obj in one global variable stored in this
func startAuction(assetID, ethAddr, quorumAddr string) *ecomm.Auction {
	args := ecomm.StartAuctionArgs{
		AssetID:    assetID,
		EthAddr:    ethAddr,
		QuorumAddr: quorumAddr,
	}
	_, err := assetClient.StartAuction(args)
	check(err)
	time.Sleep(3 * time.Second)
	fmt.Println("Started auction for asset")

	auctionID, err := assetClient.GetLastAuctionID()
	check(err)
	fmt.Println("AuctionID: ", auctionID)

	a, err := assetClient.GetAuction(auctionID)
	check(err)
	return a
}

func endAuction(auctionID int) {
	a, err := assetClient.GetAuction(auctionID)
	check(err)

	_, err = assetClient.EndAuction(a.AssetID)
	check(err)

	for {
		time.Sleep(1 * time.Second)
		a, err = assetClient.GetAuction(a.ID)
		check(err)
		if a.Status == "Ended" {

			fmt.Println("Auction Ended")
			fmt.Println("Highest Bidder: ", a.HighestBidder)
			fmt.Println("Highest Bid: ", a.HighestBid)
			fmt.Println("Highest Bid Platform: ", a.HighestBidPlatform)

			asset, err := assetClient.GetAsset(a.AssetID)
			check(err)
			fmt.Println("New Asset Owner: ", asset.Owner)

			break
		}
	}

	// publish
	ccsvc, err := cclib.NewEventService(strings.Split(zkNodes, ","), "auctioner")
	check(err)

	payload, _ := json.Marshal(a)
	ccsvc.Publish(ecomm.AuctionEndingEvent, payload)
}

// this is only used for recording bid
// Use Auctioner 1's key1 to deploy contract
func deployCrossChainAuction(client *ethclient.Client, erc20 common.Address) string {
	auth, err := cclib.NewTransactor("../../keys/key1", password)
	check(err)

	addr, tx, _, err := eth_auction.DeployEthAuction(auth, client, erc20)
	check(err)

	success, err := cclib.WaitTx(client, tx.Hash())
	check(err)

	printTxStatus(success)
	fmt.Printf("Auction contract address: %s\n", addr.Hex())

	return addr.Hex()
}

func newAuctionSession(
	addr common.Address, eth *ethclient.Client, keyfile string,
) *eth_auction.EthAuctionSession {
	auth, err := cclib.NewTransactor(keyfile, password)
	check(err)

	cc, err := eth_auction.NewEthAuction(addr, eth)
	check(err)

	return &eth_auction.EthAuctionSession{
		Contract:     cc,
		TransactOpts: *auth,
	}
}

func printTxStatus(success bool) {
	if success {
		fmt.Println("Transaction successful")
	} else {
		fmt.Println("Transaction failed")
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// also no relayer involved, 'locally' make bid
// func bidAuction(addrHex, keyfile string, value int64) {
// 	addr := common.HexToAddress(addrHex)
// 	auctionSession := newAuctionSession(addr, client, keyfile)
// 	auctionSession.TransactOpts.Value = big.NewInt(value)
// 	tx, err := auctionSession.Bid(big.NewInt(value))
// 	check(err)
// 	success, err := cclib.WaitTx(client, tx.Hash())
// 	check(err)
// 	printTxStatus(success)
// 	if !success {
// 		panic("failed to bid auction")
// 	}
// 	auctionSession.TransactOpts.Value = big.NewInt(0)

// 	highestBidder, err := auctionSession.HighestBidder()
// 	check(err)
// 	fmt.Println("highest bidder:", highestBidder.Hex())

// 	highestBid, err := auctionSession.HighestBid()
// 	check(err)
// 	fmt.Println("highest bid:", highestBid)
// }

func handleProceedingAuction(payload []byte) {
	var a ecomm.Auction
	err := json.Unmarshal(payload, &a)
	check(err)
}
