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
)

const (
	password = "password"

	userInfoFile  = "../user_info.json"
	erc20InfoFile = "../erc20_info.json"
	logInfoFile   = "../log.json"
	timeInfoFile  = "../timer"
)

var (
	assetClient *ecomm.AssetClient

	aucT     *bind.TransactOpts
	usr_name = "Auctioner 1"
	auc_key  = "../../keys/key1"
)

func main() {
	assetClient = ecomm.NewAssetClient()

	command := flag.String("c", "", "command")
	asset := flag.String("ast", "", "Asset name")
	id := flag.String("id", "", "Auction ID")
	flag.StringVar(&usr_name, "usr", usr_name, "Load User/Auctioner Information")
	flag.Parse()

	fmt.Println("Load Auctioner: ", usr_name)
	load_auctioner(usr_name)
	aucT, _ = cclib.NewTransactor(auc_key, password)

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
	// @reset timer
	t := time.Now()
	cclib.LastEventTimestamp.Set(t, timeInfoFile)

	asset := &ecomm.Asset{
		ID:    asset_name,
		Owner: aucT.From.Hex(),
	}

	log.Println("[fabric] Adding asset")
	_, err := assetClient.AddAsset(asset_name, aucT.From.Hex())
	check(err)

	payload, _ := json.Marshal(asset)
	cclib.LogEventToFile(logInfoFile, ecomm.AddingAssetEvent, payload, t, timeInfoFile)
}

func close(auctionID int) {
	// @reset timer
	t := time.Now()
	cclib.LastEventTimestamp.Set(t, timeInfoFile)

	a, err := assetClient.GetAuction(auctionID)
	check(err)

	if a.Status != "open" {
		err = fmt.Errorf("auction status error")
		check(err)
	}

	log.Println("[fabric] Closing auction")
	_, err = assetClient.CloseAuction(auctionID)
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
