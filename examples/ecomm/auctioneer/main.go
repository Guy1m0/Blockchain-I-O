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
	logInfoFile   = "../log.csv"
	timeInfoFile  = "../timer"
)

var (
	assetClient *ecomm.AssetClient

	aucT     *bind.TransactOpts
	usr_name = "Auctioner 1"
	//auc_key  = "../../keys/key1"
)

func main() {
	assetClient = ecomm.NewAssetClient()

	command := flag.String("c", "", "command")
	asset := flag.String("ast", "", "Asset name")
	id := flag.String("id", "", "Auction ID")
	flag.StringVar(&usr_name, "usr", usr_name, "Load User/Auctioner Information")
	flag.Parse()

	fmt.Println("Load Auctioner: ", usr_name)
	auc_key := load_auctioner(usr_name)
	aucT, _ = cclib.NewTransactor(auc_key, password)

	switch *command {
	case "create":
		create(*asset)
	case "close":
		id_, _ := strconv.Atoi(*id)
		close(id_)
	case "cancel":
		id_, _ := strconv.Atoi(*id)
		cancel(id_)
	case "check":
		id_, _ := strconv.Atoi(*id)
		check_status(id_)
	default:
		fmt.Println("command not found")
	}
}

// Use key 1 as default auctioner
func create(asset_name string) {
	t := time.Now()

	log.Println("[fabric] Adding asset")
	_, err := assetClient.AddAsset(asset_name, aucT.From.Hex())
	check(err)
	ecomm.LogEvent(logInfoFile, ecomm.AssetAddingEvent, asset_name, t, "", 0)
}

func cancel(auctionID int) {
	t := time.Now()

	a, err := assetClient.GetAuction(auctionID)
	check(err)

	if a.Status != "open" {
		err = fmt.Errorf("auction status error")
		check(err)
	}

	log.Println("[fabric] Cancel auction")
	_, err = assetClient.CancelAuction(auctionID)
	check(err)

	ecomm.LogEvent(logInfoFile, ecomm.AuctionCancelingEvent, a.AssetID, t, "", 0)
}

func close(auctionID int) {
	t := time.Now()
	cclib.LastEventTimestamp.Set(t, timeInfoFile)

	a, err := assetClient.GetAuction(auctionID)
	check(err)

	if a.Status != "open" {
		err = fmt.Errorf("auction status error")
		check(err)
	}

	log.Println("[fabric] Conclude auction")
	_, err = assetClient.CloseAuction(auctionID)
	check(err)

	payload, _ := json.Marshal(a)
	t = time.Now()
	cclib.LogEventToFile(logInfoFile, ecomm.AuctionClosingEvent, payload, t, timeInfoFile)

	//@reset
	cclib.LastEventTimestamp.Set(t, timeInfoFile)
}

func check_status(auctionID int) {
	a, err := assetClient.GetAuction(auctionID)
	check(err)

	fmt.Println("auction ID:", auctionID, "AssetID:", a.AssetID, "Status: ", a.Status)

}

func load_auctioner(name string) string {
	users, err := ecomm.ReadUsersFromFile(userInfoFile)
	check(err)

	for _, user := range users {
		if name == user.UserID {
			return user.KeyFile
		}
	}

	return "../../keys/key1"
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
