package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
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
	ccsvc       *cclib.CCService

	zkNodes = "localhost:2181"
	//platform = "eth"

	//aucT        *bind.TransactOpts
	asset_names []string

	usr_name          = "auctioneer 1"
	auc_type          = "english"
	support_auc_types = []string{"english", "dutch", "cb1p", "cb2p", "all"}

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
	ccsvc, _ = cclib.NewEventService(strings.Split(zkNodes, ","), "relayer") //zookeeper node

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
		auction_infos, _ := ecomm.ReadAuctionsFromFile(auctionInfoFile)
		check(err)

		s := len(auction_infos)
		//asset_name := asset_names[s]
		createTesting(s, *batch_size)
		var sleep_time int
		sleep_time = *batch_size * 7
		time.Sleep(time.Duration(sleep_time) * time.Second)

		auction_infos, _ = ecomm.ReadAuctionsFromFile(auctionInfoFile)
		s = len(auction_infos)
		bidTesting(auction_infos, s, *batch_size)

		time.Sleep(time.Duration(sleep_time) * time.Second)
		closeTesting(s, *batch_size)

		time.Sleep(time.Duration(sleep_time*2) * time.Second)
		commitTesting(s, *batch_size)

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
		return
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
		log.Printf("Make bid on %s platforms with UserID: %s", platform, userID)
		withdraw(auction_info.AuctionID, bid_key, platform)
	case "withQ":
		auction_infos, _ := ecomm.ReadAuctionsFromFile(auctionInfoFile)
		index := len(auction_infos) - 1
		auction_info := auction_infos[index]

		accounts, _ := ecomm.ReadUsersFromFile(userInfoFile)

		platform := "quo"
		userID := accounts[2].UserID
		bid_key := load_bidder_key(userID)
		log.Printf("Make bid on %s platforms with UserID: %s", platform, userID)
		withdraw(auction_info.AuctionID, bid_key, platform)
	//	withdraw()
	// case "cancel":
	// 	id_, _ := strconv.Atoi(*id)
	// 	cancel(id_)
	// case "check":
	// 	id_, _ := strconv.Atoi(*id)
	// 	check_status(id_)
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

	auctions := auction_infos[s-batch_size-1 : s]
	counter := len(auctions) * batch_size

	var auc_ind, acc_ind, auction_id int
	for i := 0; i < counter; i++ {
		wg.Add(1)            // Increment the WaitGroup counter
		go func(index int) { // Launch a goroutine for each create operation
			defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
			auc_ind = index % len(auctions)
			acc_ind = index%8 + 1
			auction_id = auctions[auc_ind].AuctionID

			platform := "quo"
			userID := accounts[acc_ind].UserID
			bid_key := load_bidder_key(userID)
			// //log.Printf("User %s places bid %d MDAI for asset %d on %s platform", userID, i*5, size-i, platform)

			bidAuction(auction_id, big.NewInt(int64(index+1)), bid_key, platform)

			platform = "eth"
			// userID := accounts[acc_ind].UserID
			// bid_key := load_bidder_key(userID)
			//log.Printf("User %s places bid %d MDAI for asset %d on %s platform", userID, i*3, size-i, platform)
			bidAuction(auction_id, big.NewInt(int64(index+1)), bid_key, platform)

		}(i) // Pass asset_name as an argument to the goroutine
	}

	wg.Wait() // Wait for all goroutines to finish
	log.Println("All bids have been placed.")
}

func closeTesting(s, batch_size int) {
	var wg sync.WaitGroup // Use a WaitGroup to wait for all goroutines to finish
	//log.Println(len(asset_names), s, batch_size)
	for i := s - batch_size; i < s; i++ {
		wg.Add(1)                // Increment the WaitGroup counter
		go func(auctionID int) { // Launch a goroutine for each create operation
			defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
			close(auctionID)
		}(i) // Pass asset_name as an argument to the goroutine
	}

	wg.Wait() // Wait for all goroutines to finish
	log.Println("All auctions have been closed.")
}

func commitTesting(s, batch_size int) {
	var wg sync.WaitGroup // Use a WaitGroup to wait for all goroutines to finish
	accounts, _ := ecomm.ReadUsersFromFile(userInfoFile)
	var bidKeys []string
	//var platform string

	for i := s - batch_size; i < s; i++ {
		wg.Add(1)                // Increment the WaitGroup counter
		go func(auctionID int) { // Launch a goroutine for each create operation
			defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
			auction, _ := assetClient.GetAuction(auctionID)

			for j := 1; j < 9; j++ {
				if auction.HighestBidder != accounts[j].Address.String() {
					bidKey := load_bidder_key(accounts[j].UserID)
					bidKeys = append(bidKeys, bidKey)
				}
			}

			// for _, bidKey := range bidKeys {
			// 	platform = "quo"
			// 	withdraw(auctionID, bidKey, platform)
			// 	platform = "eth"
			// 	withdraw(auctionID, bidKey, platform)
			// }

			sign_auction_result(auctionID)
		}(i) // Pass asset_name as an argument to the goroutine
	}

	wg.Wait() // Wait for all goroutines to finish
	log.Println("All auction results have been committed.")
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
