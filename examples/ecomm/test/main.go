package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"

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
	ethClient   *ethclient.Client
	quoClient   *ethclient.Client
	assetClient *ecomm.AssetClient

	// zkNodes = "localhost:2181"
	//platform = "eth"

	//aucT        *bind.TransactOpts
	asset_names []string

	usr_name          = "auctioneer 1"
	auc_type          = "english"
	support_auc_types = []string{"english", "dutch", "cb1p", "cb2p", "all"}

	logCSVPath = "../tmp/log.csv"

	defaultHeaders = "AssetID,Event,KeyWords,StartTime,EndTime,KafkaReceived,GasCost,Note,TimeElapsed,KafkaTime\n"

	//bid_key string
	//auc_key  = "../../keys/key1"
)

func main() {
	var err error
	ethClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8545", "localhost"))
	check(err)
	quoClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8546", "localhost"))
	check(err)
	assetClient = ecomm.NewAssetClient()
	asset_names, err = readNamesFromFile(assetNamesFile)
	check(err)
	// ccsvc, _ = cclib.NewEventService(strings.Split(zkNodes, ","), "relayer") //zookeeper node

	command := flag.String("c", "", "command")
	batch_size := flag.Int("s", 3, "Batch size for testing")

	// asset := flag.String("ast", "", "Asset name")
	//b := flag.String("id", "", "Auction ID")
	flag.StringVar(&usr_name, "usr", usr_name, "Load User/auctioneer Information")
	flag.StringVar(&auc_type, "t", auc_type, "Choose testing auction type")

	flag.Parse()
	// unique_names, _ := readUniqueNamesFromFile(assetNamesFile)
	// _ = writeNamesToFile(unique_names, assetNamesFile)

	switch *command {
	case "test":
		start := time.Now()
		log.Println("Initialize Log files")
		err := reset_log()
		check(err)

		auction_infos, _ := ecomm.ReadAuctionsFromFile(auctionInfoFile)
		check(err)

		s := len(auction_infos)
		//asset_name := asset_names[s]
		createTesting(s, *batch_size)
		var sleep_time int
		sleep_time = *batch_size * 10
		time.Sleep(time.Duration(sleep_time) * time.Second)

		auction_infos, _ = ecomm.ReadAuctionsFromFile(auctionInfoFile)
		s = len(auction_infos)
		bidTesting(auction_infos, s, *batch_size)

		last_id := auction_infos[s-1].AuctionID

		time.Sleep(15 * time.Second)
		closeTesting(last_id, *batch_size)

		time.Sleep(15 * time.Second)
		commitTesting(last_id, *batch_size)

		log.Printf("Testing execution took %s \n", time.Since(start))

	case "create":
		auction_infos, _ := ecomm.ReadAuctionsFromFile(auctionInfoFile)
		check(err)

		asset_name := asset_names[len(auction_infos)]

		create(asset_name, auc_type, usr_name)

	case "bidE":
		auction_infos, _ := ecomm.ReadAuctionsFromFile(auctionInfoFile)
		index := len(auction_infos) - 1
		if index == -1 {
			index = 0
		}
		auction_info := auction_infos[index]

		accounts, _ := ecomm.ReadUsersFromFile(userInfoFile)

		platform := "eth"
		userID := accounts[1].UserID
		bid_key := load_bidder_key(userID)
		log.Printf("Make bid on %s platforms with UserID: %s", platform, userID)
		bidAuction(auction_info.AuctionID, big.NewInt(0), bid_key, platform)

	case "bidQ":
		auction_infos, _ := ecomm.ReadAuctionsFromFile(auctionInfoFile)
		index := len(auction_infos) - 1
		if index == -1 {
			index = 0
		}
		auction_info := auction_infos[index]

		accounts, _ := ecomm.ReadUsersFromFile(userInfoFile)
		platform := "quo"
		userID := accounts[2].UserID
		bid_key := load_bidder_key(userID)
		log.Printf("Make bid on %s platforms with UserID: %s", platform, userID)
		bidAuction(auction_info.AuctionID, big.NewInt(0), bid_key, platform)

	case "bidH":
		auction_infos, _ := ecomm.ReadAuctionsFromFile(auctionInfoFile)
		index := len(auction_infos) - 1
		if index == -1 {
			index = 0
		}
		auction_info := auction_infos[index]

		accounts, _ := ecomm.ReadUsersFromFile(userInfoFile)
		platform := "eth"
		userID := accounts[2].UserID
		bid_key := load_bidder_key(userID)
		log.Printf("Make bidHash on %s platforms with UserID: %s", platform, userID)
		bidAuctionH(auction_info.AuctionID, big.NewInt(4), bid_key, platform)

		platform = "quo"
		log.Printf("Make bidHash on %s platforms with UserID: %s", platform, userID)
		bidAuctionH(auction_info.AuctionID, big.NewInt(5), bid_key, platform)

	case "revealA":
		return
	case "reveal":
		return
	case "close":
		auction_infos, _ := ecomm.ReadAuctionsFromFile(auctionInfoFile)
		index := len(auction_infos) - 1
		auction_info := auction_infos[index]
		//ccsvc.Register(ecomm.AuctionClosingEvent, autoCommit)
		//ccsvc.Start(true)

		close(auction_info.AuctionID)
	case "withE":
		auction_infos, _ := ecomm.ReadAuctionsFromFile(auctionInfoFile)
		index := len(auction_infos) - 1
		auction_info := auction_infos[index]

		accounts, _ := ecomm.ReadUsersFromFile(userInfoFile)

		platform := "eth"
		userID := accounts[1].UserID
		bid_key := load_bidder_key(userID)
		log.Printf("Withdraw bid on %s platforms with UserID: %s", platform, userID)
		withdraw(auction_info.AuctionID, bid_key, platform)
	case "withQ":
		auction_infos, _ := ecomm.ReadAuctionsFromFile(auctionInfoFile)
		index := len(auction_infos) - 1
		auction_info := auction_infos[index]

		accounts, _ := ecomm.ReadUsersFromFile(userInfoFile)

		platform := "quo"
		userID := accounts[2].UserID
		bid_key := load_bidder_key(userID)
		log.Printf("Withdraw bid on %s platforms with UserID: %s", platform, userID)
		withdraw(auction_info.AuctionID, bid_key, platform)
	case "commit":
		auction_infos, _ := ecomm.ReadAuctionsFromFile(auctionInfoFile)
		index := len(auction_infos) - 1
		auction_info := auction_infos[index]

		sign_auction_result(auction_info.AuctionID)
	default:
		fmt.Println("command not found")
	}
}

func createTesting(s, batch_size int) {
	var wg sync.WaitGroup // Use a WaitGroup to wait for all goroutines to finish
	//log.Println(len(asset_names), s, batch_size)
	for _, asset_name := range asset_names[s : s+batch_size] {
		wg.Add(1)                    // Increment the WaitGroup counter
		go func(asset_name string) { // Launch a goroutine for each create operation
			defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
			create(asset_name, auc_type, usr_name)
		}(asset_name) // Pass asset_name as an argument to the goroutine
	}

	wg.Wait() // Wait for all goroutines to finish
	log.Println("All assets have been added.")
}

func bidTesting(auction_infos []ecomm.AuctionInfo, s, batch_size int) {
	var wg sync.WaitGroup // Use a WaitGroup to wait for all goroutines to finish
	accounts, _ := ecomm.ReadUsersFromFile(userInfoFile)
	//log.Println(len(auction_infos), s, batch_size)

	auctions := auction_infos[s-batch_size : s]
	counter := len(auctions) * batch_size * (len(accounts) - 1)

	//log.Printf("counter: %d", counter)

	var auc_ind, acc_ind, auction_id, bidAmount int
	for index := 1; index <= counter; index++ {
		auc_ind = index % len(auctions)
		acc_ind = index%8 + 1
		auction_id = auctions[auc_ind].AuctionID

		platform := "quo"
		userID := accounts[acc_ind].UserID
		bid_key := load_bidder_key(userID)

		bidAmount = int(index/batch_size+(batch_size+1)%2) / 2
		bidAuction(auction_id, big.NewInt(int64(bidAmount)), bid_key, platform)

		time.Sleep(time.Duration(math.Floor(math.Log2(math.Float64frombits(uint64(batch_size))))) * time.Second)
		platform = "eth"
		userID = accounts[9-acc_ind].UserID
		bid_key = load_bidder_key(userID)

		bidAmount = int(index/batch_size+batch_size%2) / 2
		bidAuction(auction_id, big.NewInt(int64(bidAmount)), bid_key, platform)
		time.Sleep(time.Duration(math.Floor(math.Log2(math.Float64frombits(uint64(batch_size))))) * time.Second)

	}

	wg.Wait() // Wait for all goroutines to finish
	log.Println("All bids have been placed.")
}

func closeTesting(last_id, batch_size int) {
	for i := last_id - batch_size + 1; i <= last_id; i++ {
		close(i)
	}
	log.Println("All auctions have been closed.")
}

func commitTesting(last_id, batch_size int) {
	accounts, _ := ecomm.ReadUsersFromFile(userInfoFile)

	var platform string

	for i := last_id - batch_size + 1; i <= last_id; i++ {
		sign_auction_result(i)
	}
	log.Println("All auction results have been committed.")

	for i := 1; i < 9; i++ {
		bidKey := load_bidder_key(accounts[i].UserID)
		platform = "quo"
		withdraw(last_id, bidKey, platform)
		platform = "eth"
		withdraw(last_id, bidKey, platform)
	}
	log.Println("All bidders have withdrawed unsuccessfull bids.")

}

func load_auctioneer(name string) string {
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

func reset_log() error {
	// Ensure the directory exists
	dir := filepath.Dir(logCSVPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Open (or create) the file in write mode to reset it or create a new one
	file, err := os.Create(logCSVPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the column headers back to the file
	_, err = file.WriteString(defaultHeaders)
	if err != nil {
		return err
	}

	return nil
}
