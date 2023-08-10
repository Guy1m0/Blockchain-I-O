package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	password = "password"

	zkNodes       = "localhost:2181"
	userInfoFile  = "../user_info.json"
	erc20InfoFile = "../erc20_info.json"
	logInfoFile   = "../log.json"
	timeInfoFile  = "../timer"

	root_key = "../../keys/key0"
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

	aucT     *bind.TransactOpts
	auc_name = "Auctioner 1"
	auc_key  = "../../keys/key1"

	//ccsvc       *cclib.CCService

	asset *ecomm.Asset
	// myAuction *ecomm.Auction

	eth_ERC20 common.Address
	quo_ERC20 common.Address
)

func main() {
	var erc20_info ecomm.Erc20Info
	ecomm.ReadJsonFile(erc20InfoFile, &erc20_info)

	eth_ERC20 = erc20_info.EthERC20
	quo_ERC20 = erc20_info.QuoERC20

	var err error
	ethClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8545", "localhost"))
	check(err)

	quoClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8546", "localhost"))
	check(err)

	assetClient = ecomm.NewAssetClient()

	command := flag.String("c", "", "command")
	asset := flag.String("ast", "", "Asset name")
	id := flag.String("id", "", "Auction ID")
	flag.StringVar(&auc_name, "name", auc_name, "Load Auctioner Information")
	flag.Parse()

	fmt.Println("Load Auctioner: ", auc_name)
	load_auctioner(auc_name)

	aucT, _ = cclib.NewTransactor(auc_key, password)
	// check(err)
	// parts := strings.Split(*command, ":")
	// cmd := parts[0]

	switch *command {
	case "create":
		create(*asset)
	case "close":
		id_, _ := strconv.Atoi(*id)
		close(id_)
	case "check":
		id_, _ := strconv.Atoi(*id)
		check_status(id_)
	default:
		fmt.Println("command not found")
	}
}

// Use key 1 as default auctioner
func create(asset_name string) {
	asset := &ecomm.Asset{
		ID:    asset_name,
		Owner: aucT.From.Hex(),
	}

	// @reset timer
	t := time.Now()
	cclib.LastEventTimestamp.Set(t, timeInfoFile)

	log.Println("[fabric] Adding asset")
	_, err := assetClient.AddAsset(asset_name, aucT.From.Hex())
	check(err)

	payload, _ := json.Marshal(asset)
	cclib.LogEventToFile(logInfoFile, ecomm.AddingAssetEvent, payload, t, timeInfoFile)
}

func close(auctionID int) {
	a, err := assetClient.GetAuction(auctionID)
	check(err)

	if a.Status != "open" {
		//cclib.LastEventTimestamp.Set(time.Time{})
		err = fmt.Errorf("auction status error")
		check(err)
	}

	// @reset timer
	t := time.Now()
	cclib.LastEventTimestamp.Set(t, timeInfoFile)

	log.Println("[fabric] Closing auction")

	_, err = assetClient.CloseAuction(a.AssetID)
	check(err)

	payload, _ := json.Marshal(a)
	cclib.LogEventToFile(logInfoFile, ecomm.AuctionClosingEvent, payload, t, timeInfoFile)
}

func check_status(auctionID int) {
	a, err := assetClient.GetAuction(auctionID)
	check(err)

	fmt.Println("auction ID:", auctionID, "AssetID:", a.AssetID, "Status: ", a.Status)

}

func load_auctioner(name string) {
	users, err := ecomm.ReadUsersFromFile(userInfoFile)
	check(err)

	for _, user := range users {
		if name == user.UserID {
			auc_key = user.KeyFile
			return
		}
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
