package ecomm

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strings"

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

func PrintFabricBalance(token *Chaincode, account string, label string) {
	b, err := token.EvaluateTransaction("BalanceOf", account)
	check(err)
	fmt.Printf("fabric ERC20 contract %s for account %s balance: %s\n", token.GetName(), label, string(b))
}

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
func WaitTx(client *ethclient.Client, tx *types.Transaction, label string) {
	fmt.Println(label + "...")
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	check(err)

	var status string
	if receipt.Status == 1 {
		status = "Success"
	} else {
		status = "Fail"
	}

	fmt.Printf("Transaction mined in block: %d with status: %s and cost: %d\n", receipt.BlockNumber, status, receipt.GasUsed)
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
