package main

import (
	"flag"
	"fmt"
	"math/big"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/contracts/eth_auction"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"

	//"github.com/Guy1m0/Blockchain-I-O/examples/ecomm/auction"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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

type CreateAuctionRequest struct {
	AssetCC  []byte
	AssetID  []byte
	Platform string
}

var (
	ethClient    *ethclient.Client
	quorumClient *ethclient.Client
	assetClient  *ecomm.AssetClient

	asset     *ecomm.Asset
	myAuction *ecomm.Auction
)

func main() {
	var setupInfo ecomm.SetupInfo
	ecomm.ReadJsonFile(setupInfoFile, &setupInfo)

	var err error
	ethClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8545", "localhost"))
	check(err)

	quorumClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8546", "localhost"))
	check(err)

	assetClient = ecomm.NewAssetClient()

	command := flag.String("c", "", "command")
	asset := flag.String("a", "", "Asset name")
	flag.Parse()

	switch *command {
	case "create":
		create(*asset)
	case "end":

	}

}

func create(asset_name string) {

	fmt.Println("[fabric] Adding asset")
	asset = addAsset(asset_name)

	fmt.Println("Starting auction")
	fmt.Println("[ethereum] Deploying auction")
	ethAddr := deployCrossChainAuction(ethClient)

	fmt.Println("[quorum] Deploying auction")
	quorumAddr := deployCrossChainAuction(quorumClient)

	fmt.Println("[fabric] Creating auction")
	myAuction = startAuction(asset.ID, ethAddr, quorumAddr)

	fmt.Println("[ethereum] Bidding auction")
	bidAuction(ethClient, myAuction.EthAddr, "../../keys/key1", 500)

	fmt.Println("[quorum] Bidding auction")
	bidAuction(quorumClient, myAuction.QuorumAddr, "../../keys/key2", 1000)

	fmt.Println("[fabric] Ending auction")
	endAuction(myAuction)
}

func addAsset(id string) *ecomm.Asset {
	auth, err := cclib.NewTransactor("../../../keys/key1", "password")
	check(err)
	_, err = assetClient.AddAsset(id, auth.From.Hex())
	check(err)
	time.Sleep(3 * time.Second)
	asset, err := assetClient.GetAsset(id)
	check(err)
	fmt.Println("Asset added, owner: ", asset.Owner)
	return asset
}

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

func endAuction(a *ecomm.Auction) {
	_, err := assetClient.EndAuction(a.AssetID)
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
}

func bidAuction(client *ethclient.Client, addrHex, keyfile string, value int64) {
	addr := common.HexToAddress(addrHex)
	auctionSession := newAuctionSession(addr, client, keyfile)
	auctionSession.TransactOpts.Value = big.NewInt(value)
	tx, err := auctionSession.Bid()
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

func deployCrossChainAuction(client *ethclient.Client) string {
	auth, err := cclib.NewTransactor("../../keys/key0", "password")
	check(err)

	addr, tx, _, err := eth_auction.DeployAuction(auth, client)
	check(err)

	success, err := cclib.WaitTx(client, tx.Hash())
	check(err)

	printTxStatus(success)
	fmt.Printf("Auction contract address: %s\n", addr.Hex())

	return addr.Hex()
}

func newAuctionSession(
	addr common.Address, eth *ethclient.Client, keyfile string,
) *eth_auction.AuctionSession {
	auth, err := cclib.NewTransactor(keyfile, "password")
	check(err)

	cc, err := eth_auction.NewAuction(addr, eth)
	check(err)

	return &eth_auction.AuctionSession{
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
