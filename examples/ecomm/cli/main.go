package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// figure the users' action later
const (
	rootKey      = "../../keys/key0"
	auctionerKey = "../../keys/key1"
	bidder1Key   = "../../keys/key2"
	bidder2Key   = "../../keys/key3"
	password     = "password"

	fabricTokenName = "MDAI1"

	setupInfoFile = "../setup_info.json"
)

var (
	rootT *bind.TransactOpts
	aucT  *bind.TransactOpts
	bid1T *bind.TransactOpts
	bid2T *bind.TransactOpts

	ethClient *ethclient.Client
)

func main() {
	var err error
	// The one who deploys two ERC20 contracts in Eth and Fabric
	rootT, err = cclib.NewTransactor(rootKey, password)
	check(err)

	aucT, err = cclib.NewTransactor(auctionerKey, password)
	check(err)

	bid1T, err = cclib.NewTransactor(bidder1Key, password)
	check(err)

	bid2T, err = cclib.NewTransactor(bidder2Key, password)
	check(err)

	ethClient = ecomm.NewEthClient()

	command := flag.String("c", "", "command")
	flag.Parse()

	switch *command {
	case "setup":
		setup()
	case "display":
		display()
	// case "diffRate":
	// 	diffRate()
	// case "sameRate":
	// 	sameRate()

	default:
		fmt.Println("command not found")
	}

}

func setup() {
	//supply, _ := big.NewInt(0).SetString("1"+strings.Repeat("0", ecomm.Decimal+10), 10)

	fmt.Println("Ethereum setup")

	//tokenAddr, tx, instance, err := eth_stable_coin.DeployEthStableCoin(rootT, ethClient, big.NewInt(1))
	//_, tx, _, err := eth_arbitrage.DeployERC20(rootT, ethClient, supply)

	// ecomm.WaitTx(ethClient, tx, "deploy Eth Stable Coin")
	// check(err)

	// fmt.Println("MDai address: ", tokenAddr)

	// tx, err = instance.Mint(rootT, rootT.From, supply)
	// check(err)
	// ecomm.WaitTx(ethClient, tx, "Mint Enough Coins")

	// ecomm.TransferToken(ethClient, instance, rootT, bid1T.From, 1000000)
	// ecomm.TransferToken(ethClient, instance, rootT, bid2T.From, 1000000)
	// ecomm.TransferToken(ethClient, instance, rootT, aucT.From, 0)

	//ecomm.PrintTokenBalance(instance, excT.From, "MDai", "excT")
	//ecomm.PrintTokenBalance(instance, lenderT.From, "MDai", "LenderT")

	fmt.Println("Fabric setup")
	fabricToken := ecomm.NewChaincode(fabricTokenName)
	_, err := fabricToken.SubmitTransaction("Initialize", "Multi-Dai Stablecoin", "MDAI", "18")
	check(err)

	_, err = fabricToken.SubmitTransaction("Mint", "100000000")
	check(err)
	_, err = fabricToken.SubmitTransaction("Transfer", bid1T.From.Hex(), "1000000")
	check(err)

	// _, err = fabricToken.SubmitTransaction("SetBalance", bid1T.From.Hex(), "100")
	// check(err)
	// _, err = fabricToken.SubmitTransaction("SetBalance", bid2T.From.Hex(), "100")
	// check(err)
	// _, err = fabricToken.SubmitTransaction("SetBalance", aucT.From.Hex(), "0")
	// check(err)
	//fmt.Println(3 * time.Second)
	fmt.Println("Check Balance")
	ecomm.PrintFabricBalance(fabricToken, bid1T.From.Hex(), "Bidder 1")
	//ecomm.PrintFabricBalance(fabricToken, lenderT.From.Hex(), "lender")

	// ecomm.WriteJsonFile(setupInfoFile, ecomm.SetupInfo{
	// 	EthStableCoinAddress: tokenAddr,
	// 	FabricTokenName:      fabricTokenName,

	// 	Exchange:    excT.From,
	// 	Lender:      lenderT.From,
	// 	Arbitrageur: arbT.From,
	// })
}

func display() {
	// var setupInfo ecomm.SetupInfo
	// ecomm.ReadJsonFile(setupInfoFile, &setupInfo)

	//token1, err := eth_arbitrage.NewERC20(setupInfo.Token1Address, ethClient)
	//check(err)

	//token2, err := eth_arbitrage.NewERC20(setupInfo.Token2Address, ethClient)
	//check(err)

	//ecomm.PrintTokenBalance(token1, excT.From, "eth token1", "exchange")
	//ecomm.PrintTokenBalance(token2, excT.From, "eth token2", "exchange")

	//ecomm.PrintTokenBalance(token1, arbT.From, "eth token1", "arbitrageur")
	//ecomm.PrintTokenBalance(token2, arbT.From, "eth token2", "arbitrageur")

	// _, err := eth_arbitrage.NewAMM(setupInfo.Amm1Address, ethClient)
	// check(err)
	// _, err = eth_arbitrage.NewAMM(setupInfo.Amm2Address, ethClient)
	// check(err)

	//ecomm.PrintAMMRate(amm1, "amm1")
	//ecomm.PrintAMMRate(amm2, "amm2")

	fabricToken := ecomm.NewChaincode(fabricTokenName)

	ecomm.PrintFabricBalance(fabricToken, aucT.From.Hex(), "auctioner")
	ecomm.PrintFabricBalance(fabricToken, bid1T.From.Hex(), "bidder 1")
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
