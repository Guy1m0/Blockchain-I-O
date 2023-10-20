package ecomm

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/contracts/eth_auction"
	"github.com/Guy1m0/Blockchain-I-O/contracts/eth_stable_coin"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	Decimal     = 15
	DecimalB, _ = big.NewInt(0).SetString("1"+strings.Repeat("0", Decimal), 10)
)

const (
	//root_key = "../keys/key0"
	password = "password"
)

func NewEthClient() *ethclient.Client {
	client, err := ethclient.Dial(fmt.Sprintf("http://%s:8545", "localhost"))
	check(err)
	return client
}

func NewQuorumClient() *ethclient.Client {
	client, err := ethclient.Dial(fmt.Sprintf("http://%s:8546", "localhost"))
	check(err)
	return client
}

// func PrintFabricBalance(token *Chaincode, account string, label string) {
// 	b, err := token.EvaluateTransaction("BalanceOf", account)
// 	check(err)
// 	fmt.Printf("fabric ERC20 contract %s for account %s balance: %s\n", token.GetName(), label, string(b))
// }

func TransferToken(client *ethclient.Client, token *eth_stable_coin.EthStableCoin, auth *bind.TransactOpts, to common.Address, amount int64) {
	tx, err := token.Transfer(auth, to, big.NewInt(0).Mul(big.NewInt(amount), DecimalB))
	check(err)
	WaitTx(client, tx, "transfer token")
}

func PrintTokenBalance(token *eth_stable_coin.EthStableCoin, address common.Address, tokenName, accountName string) {
	valueB, err := token.BalanceOf(&bind.CallOpts{}, address)
	check(err)
	fmt.Printf("%s %s balance: %s\n",
		tokenName, accountName,
		big.NewInt(0).Div(valueB, DecimalB).String(),
	)
}
func WaitTx(client *ethclient.Client, tx *types.Transaction, label string) *types.Receipt {
	fmt.Println(label + "...")
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	check(err)

	// @todo: debug transaction if status == "Fail"

	// transferEventSignature := []byte("Transfer(address,address,uint256)")
	// hash := crypto.Keccak256Hash(transferEventSignature)

	// for _, log := range receipt.Logs {
	// 	if log.Topics[0].Hex() == hash.Hex() {
	// 		src := common.BytesToAddress(log.Topics[1].Bytes())
	// 		dst := common.BytesToAddress(log.Topics[2].Bytes())

	// 		data := new(big.Int)
	// 		data.SetBytes(log.Data)
	// 		amount := data.Uint64()

	// 		fmt.Printf("Transfer event: src: %s, dst: %s, amount: %d\n", src.Hex(), dst.Hex(), amount)
	// 	}
	// }

	var status string
	if receipt.Status == 1 {
		status = "Success"
	} else {
		status = "Fail"
	}

	fmt.Printf("Transaction mined in block: %d with status: %s and cost: %d\n", receipt.BlockNumber, status, receipt.GasUsed)
	return receipt
}

// Call this when necessary
func debugTransaction(tx *types.Transaction, endpoint string) error {
	ctx := context.Background()
	txHash := tx.Hash()

	// get the underlying RPC client from the ethclient.Client
	rpcClient, err := rpc.Dial(endpoint)

	var result interface{}
	err = rpcClient.CallContext(ctx, &result, "debug_traceTransaction", txHash, nil)

	if err != nil {
		return fmt.Errorf("failed to call client.CallContext: %v", err)
	}

	fmt.Printf("Debug info for transaction: %s\n", txHash.Hex())
	fmt.Printf("Result: %v\n", result)
	return nil
}

func PrintTxStatus(success bool) {
	if success {
		fmt.Println("Transaction successful")
	} else {
		fmt.Println("Transaction failed")
	}
}

func ReadJsonFile(filepath string, val interface{}) {
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()

	d := json.NewDecoder(f)
	err = d.Decode(val)
	check(err)
}

func WriteJsonFile(filepath string, val interface{}) {
	f, err := os.Create(filepath)
	check(err)
	defer f.Close()

	e := json.NewEncoder(f)
	e.SetIndent("", "  ")
	err = e.Encode(val)
	check(err)
}

func AddUserToFile(filepath string, newUser UserInfo) {
	// Read existing users
	var users []UserInfo
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			// File does not exist yet, will be created later
			users = []UserInfo{}
		} else {
			// Other error
			panic(err)
		}
	} else {
		err = json.Unmarshal(file, &users)
		if err != nil {
			panic(err)
		}
	}

	// Add new user
	users = append(users, newUser)

	// Write back to file
	file, err = json.MarshalIndent(users, "", "  ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filepath, file, 0644)
	if err != nil {
		panic(err)
	}
}

func ReadUsersFromFile(filepath string) ([]UserInfo, error) {
	// Read file
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON
	var users []UserInfo
	err = json.Unmarshal(file, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func SplitSignature(sig string) (r [32]byte, s [32]byte, v uint8) {
	b := common.Hex2Bytes(sig)
	copy(r[:], b[:32])
	copy(s[:], b[32:64])
	v = b[64]
	return
}

// this is only used for recording bid
// Use Auctioner 1's key1 to deploy contract
func DeployCrossChainAuction(client *ethclient.Client, erc20 common.Address, asset_id string, root_key string) (string, *types.Receipt) {
	auth, err := cclib.NewTransactor(root_key, password)
	check(err)

	addr, tx, _, err := eth_auction.DeployEthAuction(auth, client, erc20, asset_id)
	check(err)

	receipt := WaitTx(client, tx, fmt.Sprintf("Deploy Auction contract with address: %s", addr.Hex()))

	return addr.Hex(), receipt
}

// Auction Contract is already deployed in Fabric Network
// Just create a asset/auction obj in one global variable stored in this
// sleep 3s
func StartAuction(assetClient *AssetClient, assetID, ethAddr, quorumAddr string) *Auction {
	args := StartAuctionArgs{
		AssetID:    assetID,
		EthAddr:    ethAddr,
		QuorumAddr: quorumAddr,
	}
	_, err := assetClient.StartAuction(args)
	check(err)
	// @wait
	time.Sleep(10 * time.Microsecond) // 0.01s = 100 ms
	fmt.Println("Started auction for asset")

	auctionID, err := assetClient.GetLastAuctionID()
	check(err)
	fmt.Println("AuctionID: ", auctionID)

	a, err := assetClient.GetAuction(auctionID)
	check(err)
	return a
}

var mu sync.Mutex

func UpdateCSVLog(filePath, event, eventID, cost string, eventTime map[string]time.Time) error {
	mu.Lock()
	defer mu.Unlock()

	// Open CSV for appending
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Create a new log entry
	log := EventLog{
		Event:   event,
		EventID: eventID,
		Cost:    cost,
	}

	if t, ok := eventTime["Start"]; ok {
		log.StartTime = t
	}
	if t, ok := eventTime["End"]; ok {
		log.EndTime = t
	}
	if t, ok := eventTime["KafkaReceived"]; ok {
		log.KafkaReceived = t
	}

	// Write the new log entry to the CSV
	err = writer.Write([]string{
		log.Event,
		log.EventID,
		log.StartTime.String(),
		log.EndTime.String(),
		log.KafkaReceived.String(),
		log.Cost,
	})

	return err
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
