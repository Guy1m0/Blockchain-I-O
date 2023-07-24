package main

import (
	"flag"
	"strings"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	zkNodes = "localhost:2181"

	excKey   = "../../keys/key1"
	password = "password"

	setupInfoFile = "../setup_info.json"

	ccsvc *cclib.CCService

	ethClient *ethclient.Client
	excT      *bind.TransactOpts
)

// Eth to Kafka
func main() {
	flag.StringVar(&zkNodes, "zk", zkNodes, "comma separated zoolkeeper nodes")
	flag.Parse()

	var err error

	ethClient = ecomm.NewEthClient()
	excT, err = cclib.NewTransactor(excKey, password)
	check(err)

	ccsvc, err = cclib.NewEventService(strings.Split(zkNodes, ","), "ccarbit")
	check(err)
	ccsvc.Register(ecomm.OnInitializedLending, handleECommInitialized)
	err = ccsvc.Start(true)
	check(err)

	select {}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
