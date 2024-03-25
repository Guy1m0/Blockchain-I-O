package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"sync"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	password = "password"

	userInfoFile     = "../tmp/user_info.json"
	contractInfoFile = "../tmp/contract_info.json"
	auctionInfoFile  = "../tmp/auction_info.json"
	assetNamesFile   = "../asset_names.txt"

	logInfoFile  = "../tmp/log.csv"
	timeInfoFile = "../tmp/timer"
)

var (
	ethClient *ethclient.Client
	quoClient *ethclient.Client

	assetClient *ecomm.AssetClient

	platform = "eth"

	aucT        *bind.TransactOpts
	asset_names []string

	usr_name          = "Auctioner 1"
	auc_type          = "english"
	support_auc_types = []string{"english", "dutch", "cb1p", "cb2p", "all"}

	bid_key string
	//auc_key  = "../../keys/key1"
)

func main() {
	var err error
	ethClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8545", "localhost"))
	check(err)
	quoClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8546", "localhost"))
	check(err)
	assetClient = ecomm.NewAssetClient()

	command := flag.String("c", "", "command")

	// asset := flag.String("ast", "", "Asset name")
	// id := flag.String("id", "", "Auction ID")
	flag.StringVar(&usr_name, "usr", usr_name, "Load User/Auctioner Information")
	flag.StringVar(&auc_type, "t", auc_type, "Choose testing auction type")

	flag.Parse()

	//fmt.Println("Load Auctioner: ", usr_name)
	auc_key := load_auctioner(usr_name)
	aucT, _ = cclib.NewTransactor(auc_key, password)

	asset_names, err := readNamesFromFile(assetNamesFile)

	check(err)
	// unique_names, _ := readUniqueNamesFromFile(assetNamesFile)
	// _ = writeNamesToFile(unique_names, assetNamesFile)

	switch *command {
	case "test":

	case "create":
		auction_infos, _ := ecomm.ReadAuctionsFromFile(auctionInfoFile)
		asset_name := asset_names[len(auction_infos)]

		create(asset_name, auc_type, usr_name)

	case "b-create":
		var wg sync.WaitGroup // Use a WaitGroup to wait for all goroutines to finish

		for _, asset_name := range asset_names[51:55] {
			wg.Add(1)                    // Increment the WaitGroup counter
			go func(asset_name string) { // Launch a goroutine for each create operation
				defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
				create(asset_name, auc_type, usr_name)
			}(asset_name) // Pass asset_name as an argument to the goroutine
		}

		wg.Wait() // Wait for all goroutines to finish
		log.Println("All assets have been added.")
	case "bid":
		auction_infos, _ := ecomm.ReadAuctionsFromFile(auctionInfoFile)
		index := len(auction_infos) - 1
		auction_info := auction_infos[index]

		accounts, _ := ecomm.ReadUsersFromFile(userInfoFile)

		platform = "eth"
		userID := accounts[1].UserID
		bid_key = load_bidder_key(userID)
		log.Printf("Make bid on %s platforms with UserID: %s", platform, userID)
		bidAuction(auction_info.AuctionID, big.NewInt(4))

		platform = "quo"
		userID = accounts[2].UserID
		bid_key = load_bidder_key(userID)
		log.Printf("Make bid on %s platforms with UserID: %s", platform, userID)
		bidAuction(auction_info.AuctionID, big.NewInt(5))

	case "bidH":
		return
	case "reveal":
		return
	// case "create":
	// 	create(*asset, *auc_type)
	// case "reveal":
	// 	id_, _ := strconv.Atoi(*id)
	// 	reveal(id_)
	// case "close":
	// 	id_, _ := strconv.Atoi(*id)
	// 	close(id_)
	// case "cancel":
	// 	id_, _ := strconv.Atoi(*id)
	// 	cancel(id_)
	// case "check":
	// 	id_, _ := strconv.Atoi(*id)
	// 	check_status(id_)
	default:
		fmt.Println("command not found")
	}
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

// stringInSlice checks if a string exists in a slice of strings.
func stringInSlice(slice []string, target string) bool {
	for _, item := range slice {
		if item == target {
			return true // Found the target string in the slice
		}
	}
	return false // Target string not found in the slice
}

// readNamesFromFile reads names from a file, one per line, and returns a slice of strings.
func readNamesFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err // Return an empty slice and the error
	}
	defer file.Close()

	var names []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		names = append(names, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return names, err
	}

	return names, nil
}
