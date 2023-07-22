package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/contracts/eth_arbitrage"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	arbitKey = "../../keys/key3"
	password = "password"

	setupInfoFile  = "../setup_info.json"
	ecommFile      = "../flash_loan.json"
	commitVoteFile = "../commit_vote.json"

	loan    = 10000
	intrest = 100
)

var (
	lenderCodeName = "lender"

	setupInfo ecomm.SetupInfo

	arbT      *bind.TransactOpts
	ethClient *ethclient.Client
)

func main() {
	command := flag.String("c", "", "command")
	flag.StringVar(&lenderCodeName, "lc", lenderCodeName, "lender code name")

	flag.Parse()

	ecomm.ReadJsonFile(setupInfoFile, &setupInfo)

	var err error
	arbT, err = cclib.NewTransactor(arbitKey, password)
	check(err)

	ethClient = ecomm.NewEthClient()

	switch *command {
	case "setup":
		setup()
	case "register":
		register()
	case "execute":
		execute()
	case "display":
		display()

	default:
		fmt.Println("command not found")
	}
}

func setup() {
	arbAddr, tx, _, err := eth_arbitrage.DeployArbitrage(
		arbT, ethClient,
		setupInfo.Token1Address, setupInfo.Token2Address,
		setupInfo.Amm1Address, setupInfo.Amm2Address,
	)
	check(err)
	ecomm.WaitTx(ethClient, tx, "deploy arbitrage")

	floan := ecomm.EComm{
		LenderContract:    lenderCodeName,
		ArbitrageContract: arbAddr.Hex(),

		Lender:      setupInfo.Lender.Hex(),
		Arbitrageur: setupInfo.Arbitrageur.Hex(),
		Exchange:    setupInfo.Exchange.Hex(),

		Loan:    loan,
		Intrest: intrest,
	}

	arbitrage, err := eth_arbitrage.NewArbitrage(arbAddr, ethClient)
	check(err)

	loanB := big.NewInt(0).Mul(big.NewInt(loan), ecomm.DecimalB)
	intrestB := big.NewInt(0).Mul(big.NewInt(intrest), ecomm.DecimalB)

	tx, err = arbitrage.Setup(arbT,
		setupInfo.Lender, setupInfo.Arbitrageur, setupInfo.Exchange,
		loanB, intrestB, floan.Hash32(),
	)
	check(err)
	ecomm.WaitTx(ethClient, tx, "setup arbitrage contract")

	lenderCode := ecomm.NewChaincode(lenderCodeName)

	fljson, _ := json.Marshal(floan)
	_, err = lenderCode.SubmitTransaction("Setup", string(fljson), floan.Hash())
	check(err)
	fmt.Println("setup lender contract...")
	time.Sleep(3 * time.Second)

	resp, err := lenderCode.EvaluateTransaction("GetStatus")
	check(err)
	status, err := strconv.Atoi(string(resp))
	check(err)
	if status == 1 {
		fmt.Println("setup successful")
	}

	fmt.Println("writing commit vote")
	arbS, err := cclib.NewSigner(arbitKey, password)
	check(err)

	loanHash := floan.Hash32()
	sig, err := arbS.Sign(loanHash[:])
	check(err)

	commitVote := ecomm.CommitVote{
		LoanHash:     floan.Hash(),
		ArbitrageSig: hex.EncodeToString(sig),
	}

	ecomm.WriteJsonFile(ecommFile, floan)
	ecomm.WriteJsonFile(commitVoteFile, commitVote)
}

func register() {
	var floan ecomm.EComm
	ecomm.ReadJsonFile(ecommFile, &floan)

	b, _ := json.Marshal(floan)

	http.Post("http://localhost:9000/ecomm",
		"application/json",
		bytes.NewReader(b),
	)

}

func display() {
	var setupInfo ecomm.SetupInfo
	ecomm.ReadJsonFile(setupInfoFile, &setupInfo)

	var floan ecomm.EComm
	ecomm.ReadJsonFile(ecommFile, &floan)

	token1, err := eth_arbitrage.NewERC20(setupInfo.Token1Address, ethClient)
	check(err)

	ecomm.PrintTokenBalance(token1, common.HexToAddress(floan.ArbitrageContract), "token1", "arbitrage contract")
}

func execute() {
	var floan ecomm.EComm
	ecomm.ReadJsonFile(ecommFile, &floan)

	arbitrage, err := eth_arbitrage.NewArbitrage(
		common.HexToAddress(floan.ArbitrageContract), ethClient,
	)
	check(err)

	var commitVote ecomm.CommitVote
	ecomm.ReadJsonFile(commitVoteFile, &commitVote)

	rl, sl, vl := ecomm.SplitSignature(commitVote.LenderSig)
	ra, sa, va := ecomm.SplitSignature(commitVote.ArbitrageSig)
	tx, err := arbitrage.Execute(arbT,
		vl, rl, sl,
		va, ra, sa,
	)
	check(err)
	fmt.Println("excute arbitrage...")
	success, err := cclib.WaitTx(ethClient, tx.Hash())
	check(err)
	ecomm.PrintTxStatus(success)
	if success {
		return
	}
	// not success
	tx, err = arbitrage.Rollback(arbT)
	check(err)
	ecomm.WaitTx(ethClient, tx, "rollback")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
