package main

import (
	"context"
	"fmt"

	"github.com/Guy1m0/Blockchain-I-O/contracts/eth_stable_coin"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

func load_bidder_key(name string) string {
	users, err := ecomm.ReadUsersFromFile(userInfoFile)
	check(err)

	for _, user := range users {
		//fmt.Println("Find ", name, "in", user.UserID)
		if name == user.UserID {
			fmt.Println("Find!", user.KeyFile)
			return user.KeyFile
		}
	}

	return "../../keys/key2"
}

// func load_ctcs() (*ethclient.Client, *bind.TransactOpts, *cclib.CCService, *cclib.Signer) {
// 	var bidderInfo ecomm.BidderInfo
// 	ecomm.ReadJsonFile(BidderInfoFile, &bidderInfo)

// 	client, err := ethclient.Dial(bidderInfo.Endpoint)
// 	check(err)

// 	bidT, err := cclib.NewTransactor(bidderInfo.Keyfile, "password")
// 	check(err)

// 	signer, _ := cclib.NewSigner(bid_key, "password")

// 	ccsvc, err := cclib.NewEventService(strings.Split(zkNodes, ","), "bidder")
// 	check(err)

// 	return client, bidT, ccsvc, signer

// }

// Call this when necessary
func debugTransaction(tx *types.Transaction) error {
	ctx := context.Background()
	txHash := tx.Hash()

	// get the underlying RPC client from the ethclient.Client
	rpcClient, err := rpc.Dial(fmt.Sprintf("http://%s", "localhost:8545"))
	check(err)

	var result interface{}
	err = rpcClient.CallContext(ctx, &result, "debug_traceTransaction", txHash, nil)

	if err != nil {
		return fmt.Errorf("failed to call client.CallContext: %v", err)
	}

	fmt.Printf("Debug info for transaction: %s\n", txHash.Hex())
	fmt.Printf("Result: %v\n", result)
	return nil
}

func load_ERC20() (*eth_stable_coin.EthStableCoin, *eth_stable_coin.EthStableCoin, *ecomm.Chaincode) {
	var erc20_info ecomm.Erc20Info
	ecomm.ReadJsonFile(erc20InfoFile, &erc20_info)

	eth_ERC20, err := eth_stable_coin.NewEthStableCoin(erc20_info.EthERC20, ethClient)
	check(err)

	quo_ERC20, err := eth_stable_coin.NewEthStableCoin(erc20_info.QuoERC20, quoClient)
	check(err)

	fabric_ERC20 := ecomm.NewChaincode(erc20_info.FabricTokenName)

	return eth_ERC20, quo_ERC20, fabric_ERC20
}
