package main

import (
	"encoding/json"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/contracts/eth_arbitrage"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func handleECommInitialized(payload []byte) {
	var floan ecomm.EComm
	err := json.Unmarshal(payload, &floan)
	check(err)

	// transfer token to arbitrage contract
	var setupInfo ecomm.SetupInfo
	ecomm.ReadJsonFile(setupInfoFile, &setupInfo)

	token1, err := eth_arbitrage.NewERC20(setupInfo.Token1Address, ethClient)
	check(err)

	ecomm.TransferToken(
		ethClient, token1, excT, common.HexToAddress(floan.ArbitrageContract), floan.Loan,
	)

	status := listenArbitrageEnd(&floan)
	if status == 2 {
		ccsvc.Publish(ecomm.OnLoanSuccessful, payload)
	} else {
		ccsvc.Publish(ecomm.OnLoanFail, payload)
	}
}

func listenArbitrageEnd(floan *ecomm.EComm) uint8 {
	addr := common.HexToAddress(floan.ArbitrageContract)
	arbitrage, err := eth_arbitrage.NewArbitrage(addr, ethClient)
	check(err)

	for {
		status, err := arbitrage.Status(&bind.CallOpts{})
		check(err)
		if status >= 2 {
			return status
		}
		time.Sleep(1 * time.Second)
	}
}
