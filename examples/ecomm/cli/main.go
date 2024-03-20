package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/contracts/cb1p_auction"
	"github.com/Guy1m0/Blockchain-I-O/contracts/english_auction"
	"github.com/Guy1m0/Blockchain-I-O/contracts/eth_stable_coin"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/olekukonko/tablewriter"
)

const (
	key_path = "../../keys/"
	rootKey  = "../../keys/key0"

	auctionerKey = "../../keys/key1"
	bidder1Key   = "../../keys/key2"
	bidder2Key   = "../../keys/key3"
	password     = "password"

	contractInfoFile = "../contract_info.json"
	userInfoFile     = "../user_info.json"

	logCSVPath     = "../log.csv"
	defaultHeaders = "EventID,Event,AuctionType,StartTime,EndTime,KafkaReceived,GasCost,Note,TimeElapsed,KafkaTime\n"
)

var (
	rootT *bind.TransactOpts

	token_name = "MDai"

	ethClient *ethclient.Client
	quoClient *ethclient.Client

	plt = "eth"
	amt = "100"
)

func main() {
	rootT, _ = cclib.NewTransactor(rootKey, password)
	ethClient = ecomm.NewEthClient()
	quoClient = ecomm.NewQuorumClient()

	command := flag.String("c", "", "command")
	usr := flag.String("usr", "", "user name")
	flag.StringVar(&token_name, "t", token_name, "Stable coin token name")
	flag.StringVar(&amt, "amt", amt, "Set new user initial balance")
	flag.StringVar(&plt, "p", plt, "Platform for new user")
	flag.Parse()

	switch *command {
	case "init":
		initialize(token_name)
	case "setup":
		setup()
	case "display":
		display()
	case "add":
		add_user(*usr, plt, amt)
	default:
		fmt.Println("command not found")
	}

}

func reset_log() error {
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

// cleanFileContent opens the file at the given filePath and truncates it to zero length,
// effectively cleaning its content.
func cleanFileContent(filePath string) error {
	// Try to open the file in read mode
	_, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// If the file does not exist, use default headers
			return nil
		} else {
			return err
		}
	} else {
		err = os.Remove(filePath)
		check(err)
	}

	return nil
}

// Deploy contracts and mint enough tokens
func initialize(token_name string) {
	fabricToken := ecomm.NewErc20Client(token_name)

	fmt.Println("Initialize Log and Info files")
	err := reset_log()
	check(err)

	err = cleanFileContent(userInfoFile)
	check(err)

	err = cleanFileContent(contractInfoFile)
	check(err)

	fmt.Println("Initialize Fabric Stable Coin: ", token_name)
	_, err = fabricToken.Initialize("Multi-Dai Stablecoin", "MDAI", "15")
	check(err)

	fmt.Printf("Mint %s on Frabic \n", token_name)
	_, err = fabricToken.Mint("10000000000")
	check(err)

	supply, _ := big.NewInt(0).SetString("1"+strings.Repeat("0", ecomm.Decimal+10), 10)

	//fmt.Println("Deploy ERC20 contracts on Eth and Quorum")
	eth_MDAI_addr, tx, eth_MDAI, _ := eth_stable_coin.DeployEthStableCoin(rootT, ethClient, big.NewInt(1))
	ecomm.WaitTx(ethClient, tx, "Deploy ERC20 Stable Coin on Ethereum")

	eth_english_addr, tx, _, _ := english_auction.DeployEnglishAuction(rootT, ethClient, eth_MDAI_addr)
	ecomm.WaitTx(ethClient, tx, "Deploy English Auction on Ethereum")

	// eth_dutch_addr, tx, _, _ := dutch_auction.DeployDutchAuction(rootT, ethClient, eth_MDAI_addr)
	// ecomm.WaitTx(ethClient, tx, "Deploy Dutch Auction on Ethereum")
	// //_ = debugTransaction(tx)

	eth_closed_bid_addr, tx, _, _ := cb1p_auction.DeployCb1pAuction(rootT, ethClient, eth_MDAI_addr)
	ecomm.WaitTx(ethClient, tx, "Deploy Closed Bid Auction on Ethereum")

	// eth_closed_bid_2nd_addr, tx, _, _ := cb2p_auction.DeployCb2pAuction(rootT, ethClient, eth_MDAI_addr)
	// ecomm.WaitTx(ethClient, tx, "Deploy Closed Bid Auction on Ethereum")

	tx, err = eth_MDAI.Mint(rootT, rootT.From, supply)
	check(err)
	ecomm.WaitTx(ethClient, tx, "Mint ERC20 Stable Coin on Ethereum")

	quo_MDAI_addr, tx, quo_MDAI, _ := eth_stable_coin.DeployEthStableCoin(rootT, quoClient, big.NewInt(1))
	ecomm.WaitTx(quoClient, tx, "Deploy ERC20 Stable Coin on Quorum")

	quo_english_addr, tx, _, _ := english_auction.DeployEnglishAuction(rootT, quoClient, quo_MDAI_addr)
	ecomm.WaitTx(quoClient, tx, "Deploy English Auction on Quorum")

	// quo_dutch_addr, tx, _, _ := dutch_auction.DeployDutchAuction(rootT, quoClient, quo_MDAI_addr)
	// ecomm.WaitTx(quoClient, tx, "Deploy Dutch Auction on Ethereum")

	quo_closed_bid_addr, tx, _, _ := cb1p_auction.DeployCb1pAuction(rootT, quoClient, quo_MDAI_addr)
	ecomm.WaitTx(quoClient, tx, "Deploy Closed Bid Auction on Quorum")

	// quo_closed_bid_2nd_addr, tx, _, _ := cb2p_auction.DeployCb2pAuction(rootT, quoClient, quo_MDAI_addr)
	// ecomm.WaitTx(quoClient, tx, "Deploy Closed Bid Auction on Ethereum")

	tx, err = quo_MDAI.Mint(rootT, rootT.From, supply)
	check(err)
	ecomm.WaitTx(quoClient, tx, "Mint ERC20 Stable Coin on Quorum")

	ecomm.WriteJsonFile(contractInfoFile, ecomm.ContractInfo{
		FabricTokenName: token_name,
		EthERC20:        eth_MDAI_addr,
		QuoERC20:        quo_MDAI_addr,
		EnglishAuction: ecomm.AuctionInfo{
			Owner:   rootT.From,
			QuoAddr: quo_english_addr,
			EthAddr: eth_english_addr,
		},
		// DutchAuction: ecomm.AuctionInfo{
		// 	Owner:   rootT.From,
		// 	QuoAddr: quo_dutch_addr,
		// 	EthAddr: eth_dutch_addr,
		// },
		Cb1pAuction: ecomm.AuctionInfo{
			Owner:   rootT.From,
			QuoAddr: quo_closed_bid_addr,
			EthAddr: eth_closed_bid_addr,
		},
		// Cb2pAuction: ecomm.AuctionInfo{
		// 	Owner:   rootT.From,
		// 	QuoAddr: quo_closed_bid_2nd_addr,
		// 	EthAddr: eth_closed_bid_2nd_addr,
		// },
	})
}

func setup() {
	aucT, err := cclib.NewTransactor(auctionerKey, password)
	check(err)

	bid1T, err := cclib.NewTransactor(bidder1Key, password)
	check(err)

	bid2T, err := cclib.NewTransactor(bidder2Key, password)
	check(err)
	//var user_info ecomm.UserInfo
	eth_ERC20, quo_ERC20, fabric_ERC20 := load_ERC20()

	fmt.Println("Setup account for 'Auctioner 1' on Fabirc")
	_, err = fabric_ERC20.Transfer(aucT.From.Hex(), "100")
	check(err)
	ecomm.AddUserToFile(userInfoFile, ecomm.UserInfo{
		UserID:  "Auctioner 1",
		Address: aucT.From,
		KeyFile: auctionerKey,
	})

	fmt.Println("Setup account for 'Bidder 1' on Ethereum")
	ecomm.TransferToken(ethClient, eth_ERC20, rootT, bid1T.From, 100)
	_, err = fabric_ERC20.Transfer(bid1T.From.Hex(), "0")
	check(err)
	ecomm.AddUserToFile(userInfoFile, ecomm.UserInfo{
		UserID:  "Bidder 1",
		Address: bid1T.From,
		KeyFile: bidder1Key,
	})

	fmt.Println("Setup account for 'Bidder 2' on Quorum")
	ecomm.TransferToken(quoClient, quo_ERC20, rootT, bid2T.From, 100)
	_, err = fabric_ERC20.Transfer(bid2T.From.Hex(), "0")
	check(err)
	ecomm.AddUserToFile(userInfoFile, ecomm.UserInfo{
		UserID:  "Bidder 2",
		Address: bid2T.From,
		KeyFile: bidder2Key,
	})
}

func display() {
	DecimalB, _ := big.NewInt(0).SetString("1"+strings.Repeat("0", 15), 10)
	// var contract_info ecomm.ConractInfo
	// ecomm.ReadJsonFile(contractInfoFile, &contract_info)

	users, err := ecomm.ReadUsersFromFile(userInfoFile)
	check(err)

	eth_ERC20, quo_ERC20, fabric_ERC20 := load_ERC20()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"User ID", "Ethereum", "Quorum", "Fabric"})

	for _, user := range users {
		// Eth balance
		valueB, _ := eth_ERC20.BalanceOf(&bind.CallOpts{}, user.Address)
		eth_balance := big.NewInt(0).Div(valueB, DecimalB).String()
		// Quo balance
		valueB, _ = quo_ERC20.BalanceOf(&bind.CallOpts{}, user.Address)
		quo_balance := big.NewInt(0).Div(valueB, DecimalB).String()
		// Fabric balance
		b, _ := fabric_ERC20.BalanceOf(user.Address.Hex())

		row := []string{
			user.UserID,
			eth_balance,
			quo_balance,
			string(b),
		}
		table.Append(row)
		//fmt.Println(user.UserID, user.Address, user.KeyFile)
	}
	table.Render()
}

func add_user(user_id string, platform string, amount string) {
	// Get contract address and corresponding client
	users, err := ecomm.ReadUsersFromFile(userInfoFile)
	check(err)

	user_key := fmt.Sprintf("%s%s%d", key_path, "key", len(users)+1)
	userT, err := cclib.NewTransactor(user_key, password)
	check(err)

	amout_, _ := strconv.ParseInt(amount, 10, 64)
	eth_ERC20, quo_ERC20, fabric_ERC20 := load_ERC20()

	if platform == "eth" {
		ecomm.TransferToken(ethClient, eth_ERC20, rootT, userT.From, amout_)
		_, err = fabric_ERC20.Transfer(userT.From.Hex(), "0")
		check(err)
	} else if platform == "quo" {
		ecomm.TransferToken(quoClient, quo_ERC20, rootT, userT.From, amout_)
		_, err = fabric_ERC20.Transfer(userT.From.Hex(), "0")
		check(err)
	} else {
		_, err = fabric_ERC20.Transfer(userT.From.Hex(), amount)
		check(err)
	}

	ecomm.AddUserToFile(userInfoFile, ecomm.UserInfo{
		UserID:  user_id,
		Address: userT.From,
		KeyFile: user_key,
	})
}

func load_ERC20() (*eth_stable_coin.EthStableCoin, *eth_stable_coin.EthStableCoin, *ecomm.Erc20Client) {
	var contract_info ecomm.ContractInfo
	ecomm.ReadJsonFile(contractInfoFile, &contract_info)

	eth_ERC20, err := eth_stable_coin.NewEthStableCoin(contract_info.EthERC20, ethClient)
	check(err)

	quo_ERC20, err := eth_stable_coin.NewEthStableCoin(contract_info.QuoERC20, quoClient)
	check(err)

	fabric_ERC20 := ecomm.NewErc20Client(contract_info.FabricTokenName)

	return eth_ERC20, quo_ERC20, fabric_ERC20
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
