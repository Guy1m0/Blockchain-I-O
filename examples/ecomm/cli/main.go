package main

import (
	"context"
	"flag"
	"fmt"
	"math/big"
	"strings"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/contracts/eth_arbitrage"
	"github.com/Guy1m0/Blockchain-I-O/contracts/eth_stable_coin"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// figure the users' action later
const (
	rootKey   = "../../keys/key0"
	excKey    = "../../keys/key1"
	lenderKey = "../../keys/key2"
	arbitKey  = "../../keys/key3"
	password  = "password"

	fabricTokenName = "MDAI"

	setupInfoFile = "../setup_info.json"
)

var (
	rootT   *bind.TransactOpts
	excT    *bind.TransactOpts
	lenderT *bind.TransactOpts
	arbT    *bind.TransactOpts

	ethClient *ethclient.Client
)

func main() {
	var err error
	// The one who deploys two ERC20 contracts in Eth and Fabric
	rootT, err = cclib.NewTransactor(rootKey, password)
	check(err)

	excT, err = cclib.NewTransactor(excKey, password)
	check(err)

	lenderT, err = cclib.NewTransactor(lenderKey, password)
	check(err)

	arbT, err = cclib.NewTransactor(arbitKey, password)
	check(err)

	ethClient = ecomm.NewEthClient()

	command := flag.String("c", "", "command")
	flag.Parse()

	switch *command {
	case "setup":
		setup()
	case "display":
		display()
	case "diffRate":
		diffRate()
	case "sameRate":
		sameRate()

	default:
		fmt.Println("command not found")
	}

}

func setup() {
	fmt.Println("Ethereum setup")
	supply, _ := big.NewInt(0).SetString("1"+strings.Repeat("0", ecomm.Decimal+10), 10)

	tokenAddr, tx, instance, err := eth_stable_coin.DeployEthStableCoin(rootT, ethClient, big.NewInt(1))
	//_, tx, _, err := eth_arbitrage.DeployERC20(rootT, ethClient, supply)

	ecomm.WaitTx(ethClient, tx, "deploy Eth Stable Coin")
	check(err)

	// err = debugTransaction(tx, ethClient)
	// if err != nil {
	// 	log.Fatalf("Failed to debug transaction: %v", err)
	// }

	fmt.Println("MDai address: ", tokenAddr)

	tx, err = instance.Mint(rootT, rootT.From, supply)
	check(err)
	ecomm.WaitTx(ethClient, tx, "Mint Enough Coins")

	ecomm.TransferToken(ethClient, instance, rootT, excT.From, 1000000)
	ecomm.TransferToken(ethClient, instance, rootT, lenderT.From, 100)

	ecomm.PrintTokenBalance(instance, excT.From, "MDai", "excT")
	ecomm.PrintTokenBalance(instance, lenderT.From, "MDai", "LenderT")

	// fmt.Println("Fabric setup")
	// fabricToken := ecomm.NewChaincode(fabricTokenName)

	// _, err = fabricToken.SubmitTransaction("SetBalance", excT.From.Hex(), "1000000")
	// check(err)
	// _, err = fabricToken.SubmitTransaction("SetBalance", lenderT.From.Hex(), "10000")
	// check(err)
	// fmt.Println(3 * time.Second)

	// ecomm.PrintFabricBalance(fabricToken, excT.From.Hex(), "exchange")
	// ecomm.PrintFabricBalance(fabricToken, lenderT.From.Hex(), "lender")

	// ecomm.WriteJsonFile(setupInfoFile, ecomm.SetupInfo{
	// 	Token1Address: token1Addr,
	// 	Token2Address: token2Addr,

	// 	Amm1Address: amm1Address,
	// 	Amm2Address: amm2Address,

	// 	FabricTokenName: fabricTokenName,

	// 	Exchange:    excT.From,
	// 	Lender:      lenderT.From,
	// 	Arbitrageur: arbT.From,
	// })
}

func display() {
	var setupInfo ecomm.SetupInfo
	ecomm.ReadJsonFile(setupInfoFile, &setupInfo)

	//token1, err := eth_arbitrage.NewERC20(setupInfo.Token1Address, ethClient)
	//check(err)

	//token2, err := eth_arbitrage.NewERC20(setupInfo.Token2Address, ethClient)
	//check(err)

	//ecomm.PrintTokenBalance(token1, excT.From, "eth token1", "exchange")
	//ecomm.PrintTokenBalance(token2, excT.From, "eth token2", "exchange")

	//ecomm.PrintTokenBalance(token1, arbT.From, "eth token1", "arbitrageur")
	//ecomm.PrintTokenBalance(token2, arbT.From, "eth token2", "arbitrageur")

	_, err := eth_arbitrage.NewAMM(setupInfo.Amm1Address, ethClient)
	check(err)
	_, err = eth_arbitrage.NewAMM(setupInfo.Amm2Address, ethClient)
	check(err)

	//ecomm.PrintAMMRate(amm1, "amm1")
	//ecomm.PrintAMMRate(amm2, "amm2")

	fabricToken := ecomm.NewChaincode(fabricTokenName)

	ecomm.PrintFabricBalance(fabricToken, excT.From.Hex(), "exchange")
	ecomm.PrintFabricBalance(fabricToken, lenderT.From.Hex(), "lender")
}

func diffRate() {
	var setupInfo ecomm.SetupInfo
	ecomm.ReadJsonFile(setupInfoFile, &setupInfo)

	amm1, err := eth_arbitrage.NewAMM(setupInfo.Amm1Address, ethClient)
	check(err)
	amm2, err := eth_arbitrage.NewAMM(setupInfo.Amm2Address, ethClient)
	check(err)

	tx, err := amm1.SetRate(rootT, big.NewInt(2), big.NewInt(3))
	check(err)
	ecomm.WaitTx(ethClient, tx, "set amm1 rate")

	tx, err = amm2.SetRate(rootT, big.NewInt(1), big.NewInt(1))
	check(err)
	ecomm.WaitTx(ethClient, tx, "set amm2 rate")

	fmt.Println("set different rate to make profit")
}

func sameRate() {
	var setupInfo ecomm.SetupInfo
	ecomm.ReadJsonFile(setupInfoFile, &setupInfo)

	amm1, err := eth_arbitrage.NewAMM(setupInfo.Amm1Address, ethClient)
	check(err)
	amm2, err := eth_arbitrage.NewAMM(setupInfo.Amm2Address, ethClient)
	check(err)

	tx, err := amm1.SetRate(rootT, big.NewInt(1), big.NewInt(1))
	check(err)
	ecomm.WaitTx(ethClient, tx, "set amm1 rate")

	tx, err = amm2.SetRate(rootT, big.NewInt(1), big.NewInt(1))
	check(err)
	ecomm.WaitTx(ethClient, tx, "set amm2 rate")

	fmt.Println("set same rate between two amm, arbitrageur should not make profit.")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func debugTransaction(tx *types.Transaction, client *ethclient.Client) error {
	ctx := context.Background()
	txHash := tx.Hash()

	// get the underlying RPC client from the ethclient.Client
	rpcClient := client.Client()

	var result interface{}
	err := rpcClient.CallContext(ctx, &result, "debug_traceTransaction", txHash, nil)

	if err != nil {
		return fmt.Errorf("failed to call client.CallContext: %v", err)
	}

	fmt.Printf("Debug info for transaction: %s\n", txHash.Hex())
	fmt.Printf("Result: %v\n", result)
	return nil
}
