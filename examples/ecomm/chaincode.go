package ecomm

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type Chaincode struct {
	contract *gateway.Contract
	name     string
}

// const (
// 	walletInfoDir = "/wallet"
// )

func NewChaincode(name string) *Chaincode {
	log.Println("============ application-golang starts ============")

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

	channelName := "mychannel"
	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}

	log.Println("--> Connecting to channel", channelName)
	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		log.Fatalf("Failed to get network: %v", err)
	}

	return &Chaincode{
		contract: network.GetContract(name),
		name:     name,
	}
}

// not use populate wallet
// func NewChaincode_(name string) *Chaincode {
// 	credPath := filepath.Join(
// 		"../../../",
// 		"fabric-samples",
// 		"test-network",
// 		"organizations",
// 		"peerOrganizations",
// 		"org1.example.com",
// 		"users",
// 		"User1@org1.example.com",
// 		"msp",
// 	)

// 	certPath := filepath.Join(credPath, "signcerts", "User1@org1.example.com-cert.pem")
// 	if _, err := os.Stat(certPath); err != nil {
// 		certPath = filepath.Join(credPath, "signcerts", "cert.pem")
// 	}

// 	keyDir := filepath.Join(credPath, "keystore")
// 	// there's a single file in this dir containing the private key
// 	files, err := ioutil.ReadDir(keyDir)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if len(files) != 1 {
// 		log.Fatalln("keystore folder should have contain one file")
// 	}
// 	keyPath := filepath.Join(keyDir, files[0].Name())

// 	// Read the cert and key files
// 	// cert := fs.readFileSync(certPath).toString()
// 	// cert, err := ioutil.ReadFile(filepath.Clean(certPath))
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	// key, err := ioutil.ReadFile(filepath.Clean(keyPath))
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	configPath := filepath.Join(
// 		"../../../",
// 		"fabric-samples",
// 		"test-network",
// 		"organizations",
// 		"peerOrganizations",
// 		"org1.example.com",
// 		"connection-org1.yaml",
// 	)

// 	// fmt.Println("cert:", string(cert))

// 	// Create an in-memory wallet and add the X509 identity to it
// 	//wallet := gateway.NewInMemoryWallet()
// 	wallet, _ := gateway.NewFileSystemWallet("wallet")

// 	// identity: X509Identity := {
// 	// 	credentials: {
// 	// 		certificate: certificatePEM,
// 	// 		privateKey: privateKeyPEM,
// 	// 	},
// 	// 	mspId: 'Org1MSP',
// 	// 	type: 'X.509',
// 	// };

// 	x509Identity := gateway.NewX509Identity("Org1MSP", certPath, keyPath)

// 	// need marshal?

// 	err = wallet.Put("appUser", x509Identity)
// 	if err != nil {
// 		log.Fatalf("Failed to put identity into wallet: %v\n", err)
// 	}

// 	// Connect to the gateway using the in-memory wallet
// 	gw, err := gateway.Connect(
// 		gateway.WithConfig(config.FromFile(filepath.Clean(configPath))),
// 		gateway.WithIdentity(wallet, "appUser"),
// 	)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to gateway: %v\n", err)
// 	}
// 	defer gw.Close()

// 	network, err := gw.GetNetwork("mychannel")
// 	if err != nil {
// 		log.Fatalf("Failed to get network: %v", err)
// 	}

// 	return &Chaincode{
// 		contract: network.GetContract(name),
// 		name:     name,
// 	}
// }

func (cc *Chaincode) GetName() string {
	return cc.name
}

func (cc *Chaincode) SubmitTransaction(name string, args ...string) ([]byte, error) {
	return cc.contract.SubmitTransaction(name, args...)
}

func (cc *Chaincode) EvaluateTransaction(name string, args ...string) ([]byte, error) {
	return cc.contract.EvaluateTransaction(name, args...)
}

func populateWallet(wallet *gateway.Wallet) error {
	log.Println("============ Populating wallet ============")
	credPath := filepath.Join(
		"../../..",
		"fabric-samples",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	// read the certificate pem
	// cert, err := os.ReadFile(filepath.Clean(certPath))
	// if err != nil {
	// 	return err
	// }

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := os.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	// key, err := os.ReadFile(filepath.Clean(keyPath))
	// if err != nil {
	// 	return err
	// }

	identity := gateway.NewX509Identity("Org1MSP", certPath, keyPath)

	return wallet.Put("appUser", identity)
}
