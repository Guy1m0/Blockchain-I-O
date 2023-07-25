package main

import (
	"encoding/hex"
	"flag"
	"fmt"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	bidderKey = "../../keys/key2"
	password  = "password"

	setupInfoFile = "../setup_info.json"
	//ecommFile      = "../flash_loan.json"
	commitVoteFile = "../commit_vote.json"
)

var (
	bidderS *cclib.Signer
)

func main() {
	var err error

	bidderS, err = cclib.NewSigner(bidderKey, password)
	check(err)

	command := flag.String("c", "", "command")
	flag.Parse()

	switch *command {
	case "bid":
		bid()
	case "sign":
		sign()
	case "abort":
		abort()
	case "verify":
		verify()
	case "initialize":
		initialize()

	default:
		fmt.Println("command not found")
	}
}

func bid() {
	return
}

func sign() {
	var commitVote ecomm.CommitVote
	ecomm.ReadJsonFile(commitVoteFile, &commitVote)

	hash, err := hex.DecodeString(commitVote.LoanHash)
	check(err)
	sig, err := bidderS.Sign(hash)
	check(err)

	commitVote.LenderSig = hex.EncodeToString(sig)

	ecomm.WriteJsonFile(commitVoteFile, commitVote)
}

func abort() {
	return
}

func initialize() {
	var setupInfo ecomm.SetupInfo
	ecomm.ReadJsonFile(setupInfoFile, &setupInfo)

	//var floan ecomm.EComm
	//ecomm.ReadJsonFile(ecommFile, &floan)

	var commitVote ecomm.CommitVote
	ecomm.ReadJsonFile(commitVoteFile, &commitVote)

	// lenderCode := ecomm.NewChaincode(floan.LenderContract)

	// _, err := lenderCode.SubmitTransaction("Initialize", commitVote.LenderSig, commitVote.ArbitrageSig)
	// check(err)
	// fmt.Println("initialize lender contract...")
	// time.Sleep(3 * time.Second)

	// resp, err := lenderCode.EvaluateTransaction("GetStatus")
	// check(err)
	// status, err := strconv.Atoi(string(resp))
	// check(err)
	// if status == 2 {
	// 	fmt.Println("initialize successful")
	// }
	// fabricToken := ecomm.NewChaincode(setupInfo.FabricTokenName)
	// ecomm.PrintFabricBalance(fabricToken, lenderS.Address().Hex(), "lender")
}

func verify() {

	var floan ecomm.EComm
	ecomm.ReadJsonFile(ecommFile, &floan)

	var commitVote ecomm.CommitVote
	ecomm.ReadJsonFile(commitVoteFile, &commitVote)

	err := verifySignature(commitVote.LoanHash, commitVote.LenderSig, floan.Lender)
	check(err)

	err = verifySignature(commitVote.LoanHash, commitVote.ArbitrageSig, floan.Arbitrageur)
	check(err)

	fmt.Println("Comit vote is valid")
}

func verifySignature(hash_, signature_, address_ string) error {
	hash, err := hex.DecodeString(hash_)
	if err != nil {
		return nil
	}
	signature, err := hex.DecodeString(signature_)
	if err != nil {
		return nil
	}
	address := common.HexToAddress(address_)
	pubkey, err := crypto.SigToPub(hash, signature)
	if err != nil {
		return err
	}
	recover := crypto.PubkeyToAddress(*pubkey)
	if address.Hex() != recover.Hex() {
		return fmt.Errorf("invalid signer address")
	}
	return nil
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
