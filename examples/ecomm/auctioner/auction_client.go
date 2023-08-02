package main

import (
	"fmt"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/contracts/eth_auction"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func load_auctioner(name string) {
	users, err := ecomm.ReadUsersFromFile(userInfoFile)
	check(err)

	for _, user := range users {
		if name == user.UserID {
			auc_key = user.KeyFile
			return
		}
	}
}

// Default use ../../../keys/key1
func addAsset(id string) *ecomm.Asset {
	fmt.Println("Create New Auction")
	_, err := assetClient.AddAsset(id, aucT.From.Hex())
	check(err)
	// @wait
	time.Sleep(3 * time.Second)
	asset, err := assetClient.GetAsset(id)
	check(err)

	fmt.Println("Auction added on Fabric with owner: ", asset.Owner)
	//fmt.Println("Pending Auction ID?", asset.PendingAuctionID)
	//fmt.Println("Asset ID: ", asset.ID)
	return asset
}

// Auction Contract is already deployed in Fabric Network
// Just create a asset/auction obj in one global variable stored in this
func startAuction(assetID, ethAddr, quorumAddr string) *ecomm.Auction {
	args := ecomm.StartAuctionArgs{
		AssetID:    assetID,
		EthAddr:    ethAddr,
		QuorumAddr: quorumAddr,
	}
	_, err := assetClient.StartAuction(args)
	check(err)
	// @wait
	time.Sleep(3 * time.Second)
	fmt.Println("Started auction for asset")

	auctionID, err := assetClient.GetLastAuctionID()
	check(err)
	fmt.Println("AuctionID: ", auctionID)

	a, err := assetClient.GetAuction(auctionID)
	check(err)
	return a
}

// this is only used for recording bid
// Use Auctioner 1's key1 to deploy contract
func deployCrossChainAuction(client *ethclient.Client, erc20 common.Address) string {
	auth, err := cclib.NewTransactor(root_key, password)
	check(err)

	addr, tx, _, err := eth_auction.DeployEthAuction(auth, client, erc20)
	check(err)

	success, err := cclib.WaitTx(client, tx.Hash())
	check(err)

	printTxStatus(success)
	fmt.Printf("Auction contract address: %s\n", addr.Hex())

	return addr.Hex()
}

func newAuctionSession(
	addr common.Address, eth *ethclient.Client, keyfile string,
) *eth_auction.EthAuctionSession {
	auth, err := cclib.NewTransactor(keyfile, password)
	check(err)

	cc, err := eth_auction.NewEthAuction(addr, eth)
	check(err)

	return &eth_auction.EthAuctionSession{
		Contract:     cc,
		TransactOpts: *auth,
	}
}

func printTxStatus(success bool) {
	if success {
		fmt.Println("Transaction successful")
	} else {
		fmt.Println("Transaction failed")
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
