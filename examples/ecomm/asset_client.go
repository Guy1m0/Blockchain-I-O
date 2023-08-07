package ecomm

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type AssetClient struct {
	contract *gateway.Contract
}

// load identity
func NewAssetClient() *AssetClient {
	// err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	// if err != nil {
	// 	log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environemnt variable: %v", err)
	// }

	credPath := filepath.Join(
		"../../../",
		"fabric-samples",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "User1@org1.example.com-cert.pem")
	if _, err := os.Stat(certPath); err != nil {
		certPath = filepath.Join(credPath, "signcerts", "cert.pem")
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		panic(err)
	}
	if len(files) != 1 {
		log.Fatalln("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())

	// Read the cert and key files
	cert, err := ioutil.ReadFile(certPath)
	if err != nil {
		panic(err)
	}

	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		panic(err)
	}

	configPath := filepath.Join(
		"../../../",
		"fabric-samples",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)

	// Create an in-memory wallet and add the X509 identity to it
	wallet := gateway.NewInMemoryWallet()
	x509Identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))
	err = wallet.Put("User1", x509Identity)
	if err != nil {
		log.Fatalf("Failed to put identity into wallet: %v\n", err)
	}

	// Connect to the gateway using the in-memory wallet
	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(configPath))),
		gateway.WithIdentity(wallet, "User1"),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gateway: %v\n", err)
	}
	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		log.Fatalf("Failed to get network: %v", err)
	}

	return &AssetClient{
		contract: network.GetContract("asset"),
	}
}

// func NewAssetClient_() *AssetClient {
// 	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
// 	if err != nil {
// 		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environemnt variable: %v", err)
// 	}

// 	ccpPath := filepath.Join(
// 		"../../../",
// 		"fabric-samples",
// 		"test-network",
// 		"organizations",
// 		"peerOrganizations",
// 		"org1.example.com",
// 		"connection-org1.yaml",
// 	)

// 	wallet, err := gateway.NewFileSystemWallet("wallet")
// 	if err != nil {
// 		log.Fatalf("Failed to create wallet: %v", err)
// 	}

// 	err = populateWallet(wallet)
// 	if err != nil {
// 		log.Fatalf("Failed to populate wallet contents: %v", err)
// 	}

// 	// error: Failed to apply identity option
// 	gw, err := gateway.Connect(
// 		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
// 		gateway.WithIdentity(wallet, "appUser"),
// 	)

// 	if err != nil {
// 		log.Fatalf("Failed to connect to gateway: %v", err)
// 	}
// 	defer gw.Close()

// 	network, err := gw.GetNetwork("mychannel")
// 	if err != nil {
// 		log.Fatalf("Failed to get network: %v", err)
// 	}

// 	return &AssetClient{
// 		contract: network.GetContract("asset"),
// 	}
// }

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
