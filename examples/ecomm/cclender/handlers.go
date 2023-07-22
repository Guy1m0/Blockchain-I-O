package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
)

func handleECommSuccessful(payload []byte) {
	var floan ecomm.EComm
	err := json.Unmarshal(payload, &floan)
	check(err)

	lenderCode := ecomm.NewChaincode(floan.LenderContract)
	_, err = lenderCode.SubmitTransaction("EndLoan", "true")
	check(err)
	fmt.Println("Ending flash loan. Refund loan + intrest: ", floan.Loan+floan.Intrest)
	time.Sleep(3 * time.Second)
	confirmEndLoan(lenderCode)
}

func handleECommFail(payload []byte) {
	var floan ecomm.EComm
	err := json.Unmarshal(payload, &floan)
	check(err)

	lenderCode := ecomm.NewChaincode(floan.LenderContract)
	_, err = lenderCode.SubmitTransaction("EndLoan", "false")
	check(err)
	fmt.Println("Ending flash loan. Refund loan: ", floan.Loan)
	time.Sleep(3 * time.Second)
	confirmEndLoan(lenderCode)
}

func confirmEndLoan(lenderCode *ecomm.Chaincode) {
	resp, err := lenderCode.EvaluateTransaction("GetStatus")
	check(err)

	status, err := strconv.Atoi(string(resp))
	check(err)

	if status == 3 {
		fmt.Println("ended loan")
	}
}
