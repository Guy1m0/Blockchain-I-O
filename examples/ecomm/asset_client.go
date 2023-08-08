package ecomm

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type AssetClient struct {
	contract *gateway.Contract
}

// load identity
func NewAssetClient() *AssetClient {
	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environment variable: %v", err)
	}

	walletPath := "wallet"
	// remove any existing wallet from prior runs
	os.RemoveAll(walletPath)
	wallet, err := gateway.NewFileSystemWallet(walletPath)
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}

	if !wallet.Exists("appUser") {
		err = populateWallet(wallet)
		if err != nil {
			log.Fatalf("Failed to populate wallet contents: %v", err)
		}
	}

	ccpPath := filepath.Join(
		"../../../",
		"fabric-samples",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)

	if err != nil {
		log.Fatalf("Failed to connect to gateway: %v", err)
	}
	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		log.Fatalf("Failed to get network: %v", err)
	}

	return &AssetClient{contract: network.GetContract("asset")}
}

func (cc *AssetClient) GetCCID() string {
	return "asset"
}

func (cc *AssetClient) AddAsset(id, owner string) ([]byte, error) {
	return cc.contract.SubmitTransaction("AddAsset", id, owner)
}

func (cc *AssetClient) StartAuction(args StartAuctionArgs) ([]byte, error) {
	b, _ := json.Marshal(args)
	return cc.contract.SubmitTransaction("StartAuction", string(b))
}

// Not use same assetID!!!!
// when auctioner tries to end one, which actually ends the latest Auction
// with same assetID
func (cc *AssetClient) EndAuction(assetID string) ([]byte, error) {
	return cc.contract.SubmitTransaction("EndAuction", assetID)
}

func (cc *AssetClient) FinalizeAuction(args FinalizeAuctionArgs) ([]byte, error) {
	b, _ := json.Marshal(args)
	return cc.contract.SubmitTransaction("FinalizeAuction", string(b))
}

func (cc *AssetClient) GetAsset(assetID string) (*Asset, error) {
	var asset Asset
	res, err := cc.contract.EvaluateTransaction("GetAsset", assetID)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(res, &asset)
	return &asset, err
}

func (cc *AssetClient) GetAuction(auctionID int) (*Auction, error) {
	var auction Auction
	res, err := cc.contract.EvaluateTransaction("GetAuction", strconv.Itoa(auctionID))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(res, &auction)
	return &auction, err
}

// better not use
func (cc *AssetClient) SetAuction(auction *Auction) ([]byte, error) {
	//var auction Auction
	b, _ := json.Marshal(auction)
	return cc.contract.EvaluateTransaction("setAuction", string(b))
	// if err != nil {
	// 	return nil, err
	// }
	// err = json.Unmarshal(res, &auction)
	// return &auction, err
}

func (cc *AssetClient) GetLastAuctionID() (int, error) {
	res, err := cc.contract.EvaluateTransaction("GetLastAuctionID")
	if err != nil {
		return 0, err
	}
	var id int
	err = json.Unmarshal(res, &id)
	return id, err
}

func (cc *AssetClient) Register(eventID string) (fab.Registration, <-chan *fab.CCEvent, error) {
	return cc.contract.RegisterEvent(eventID)
}

func (cc *AssetClient) Unregister(reg fab.Registration) {
	cc.contract.Unregister(reg)
}
