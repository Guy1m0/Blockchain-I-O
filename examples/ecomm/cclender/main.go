package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	zkNodes = "localhost:2181"

	ccsvc *cclib.CCService
)

// Fabric to Kafka
func main() {
	flag.StringVar(&zkNodes, "zk", zkNodes, "comma separated zoolkeeper nodes")
	flag.Parse()

	var err error

	ccsvc, err = cclib.NewEventService(strings.Split(zkNodes, ","), "cclender")
	check(err)
	ccsvc.Register(ecomm.OnLoanSuccessful, handleECommSuccessful)
	ccsvc.Register(ecomm.OnLoanFail, handleECommFail)
	err = ccsvc.Start(true)
	check(err)

	e := echo.New()
	e.Use(middleware.Recover())
	e.POST("/ecomm", createEComm)

	log.Fatal(e.Start(":9000"))
}

func createEComm(c echo.Context) error {
	var floan ecomm.EComm
	err := c.Bind(&floan)
	if err != nil {
		return err
	}

	fmt.Printf("register new ecomm, loan: %d, intrest: %d\n", floan.Loan, floan.Intrest)
	go listenLenderInitialize(&floan)
	return c.NoContent(http.StatusOK)
}

func listenLenderInitialize(floan *ecomm.EComm) {
	fmt.Println("listen lender initialize...")
	lenderCode := ecomm.NewChaincode(floan.LenderContract)
	for {
		resp, err := lenderCode.EvaluateTransaction("GetStatus")
		check(err)

		status, err := strconv.Atoi(string(resp))
		check(err)

		if status == 2 {
			break
		}
		time.Sleep(1 * time.Second)
	}

	payload, _ := json.Marshal(floan)
	err := ccsvc.Publish(ecomm.OnInitializedLending, payload)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
