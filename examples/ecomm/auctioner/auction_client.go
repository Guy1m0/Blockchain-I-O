package main

import (
	"fmt"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
