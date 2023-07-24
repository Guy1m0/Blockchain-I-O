package main

import (
	"context"
	"flag"
	"fmt"
	"math/big"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/contracts/eth_stable_coin"
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
	quoClient *ethclient.Client
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
	quoClient = ecomm.NewQuorumClient()

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
	eth_MDAI_addr, tx, eth_MDAI, _ := eth_stable_coin.DeployEthStableCoin(rootT, ethClient, big.NewInt(1))
	//fmt.Println("Eth address:", etheth_MDAI_addr.Hex())
	ecomm.WaitTx(ethClient, tx, "deploy Eth Stable Coin")

	tx, err := eth_MDAI.Mint(rootT, rootT.From, big.NewInt(10000))
	check(err)
	ecomm.WaitTx(ethClient, tx, "Mint Enough Coins")

	ecomm.TransferToken(ethClient, eth_MDAI, rootT, bid1T.From, *big.NewInt(10000))

	fmt.Println("Quorum setup")
	quo_MDAI_addr, tx, quo_MDAI, _ := eth_stable_coin.DeployEthStableCoin(rootT, quoClient, big.NewInt(1))
	ecomm.WaitTx(quoClient, tx, "deploy Quorum Stable Coin")
	//_, err = bind.WaitMined(context.Background(), quoClient, tx)
	//check(err)
	//ecomm.WaitTx(ethClient, tx, "deploy Quorum Stable Coin")

	tx, err = quo_MDAI.Mint(rootT, rootT.From, big.NewInt(10000))
	check(err)
	ecomm.WaitTx(quoClient, tx, "deploy Quorum Stable Coin")

	ecomm.TransferToken(quoClient, quo_MDAI, rootT, bid2T.From, *big.NewInt(10000))

	// ecomm.TransferToken(ethClient, instance, rootT, bid2T.From, 1000000)
	// ecomm.TransferToken(ethClient, instance, rootT, aucT.From, 0)

	//ecomm.PrintTokenBalance(instance, excT.From, "MDai", "excT")
	//ecomm.PrintTokenBalance(instance, lenderT.From, "MDai", "LenderT")

	fmt.Println("Fabric setup")
	fabricToken := ecomm.NewChaincode(fabricTokenName)
	fmt.Println("Initialize MDai ERC20 contract")
	_, err = fabricToken.SubmitTransaction("Initialize", "Multi-Dai Stablecoin", "MDAI", "18")
	check(err)

	//fmt.Println("Mint token")
	//_, err = fabricToken.SubmitTransaction("Mint", "0")
	//check(err)
	fmt.Println("Transfer tokens")
	_, err = fabricToken.SubmitTransaction("Transfer", bid1T.From.Hex(), "10")
	check(err)
	_, err = fabricToken.SubmitTransaction("Transfer", bid2T.From.Hex(), "10")
	check(err)
	_, err = fabricToken.SubmitTransaction("Transfer", aucT.From.Hex(), "10")
	check(err)

	// _, err = fabricToken.SubmitTransaction("SetBalance", bid1T.From.Hex(), "100")
	// check(err)
	// _, err = fabricToken.SubmitTransaction("SetBalance", bid2T.From.Hex(), "100")
	// check(err)
	// _, err = fabricToken.SubmitTransaction("SetBalance", aucT.From.Hex(), "0")
	// check(err)
	//fmt.Println(3 * time.Second)
	//fmt.Println("Check Balance")
	//ecomm.PrintFabricBalance(fabricToken, bid1T.From.Hex(), "Bidder 1")
	//ecomm.PrintFabricBalance(fabricToken, lenderT.From.Hex(), "lender")

	ecomm.WriteJsonFile(setupInfoFile, ecomm.SetupInfo{
		Bidder1Address:   bid1T.From,
		Bidder2Address:   bid2T.From,
		AuctionerAddress: aucT.From,

		FabricTokenName: fabricTokenName,

		EthERC20: eth_MDAI_addr,
		QuoERC20: quo_MDAI_addr,
	})
}

func display() {
	var setupInfo ecomm.SetupInfo
	ecomm.ReadJsonFile(setupInfoFile, &setupInfo)

	eth_ERC20, err := eth_stable_coin.NewEthStableCoin(setupInfo.EthERC20, ethClient)
	check(err)

	quo_ERC20, err := eth_stable_coin.NewEthStableCoin(setupInfo.QuoERC20, quoClient)
	check(err)

	fmt.Println("Check Balance in Eth")
	ecomm.PrintTokenBalance(eth_ERC20, setupInfo.Bidder1Address, "MDai", "Bidder_1")
	ecomm.PrintTokenBalance(eth_ERC20, setupInfo.Bidder2Address, "MDai", "Bidder_2")
	ecomm.PrintTokenBalance(eth_ERC20, setupInfo.AuctionerAddress, "MDai", "Auctioner")

	fmt.Println("Check Balance in Quorum")
	ecomm.PrintTokenBalance(quo_ERC20, setupInfo.Bidder1Address, "MDai", "Bidder_1")
	ecomm.PrintTokenBalance(quo_ERC20, setupInfo.Bidder2Address, "MDai", "Bidder_2")
	ecomm.PrintTokenBalance(quo_ERC20, setupInfo.AuctionerAddress, "MDai", "Auctioner")

	fmt.Println("Check Balance in Fabric")
	fabricToken := ecomm.NewChaincode(fabricTokenName)

	ecomm.PrintFabricBalance(fabricToken, setupInfo.Bidder1Address.Hex(), "Bidder_1")
	ecomm.PrintFabricBalance(fabricToken, setupInfo.Bidder2Address.Hex(), "Bidder_2")
	ecomm.PrintFabricBalance(fabricToken, setupInfo.AuctionerAddress.Hex(), "Auctioner")

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
