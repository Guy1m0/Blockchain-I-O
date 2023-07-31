package main

import (
	"encoding/json"
	"fmt"

	"github.com/Guy1m0/Blockchain-I-O/contracts/eth_auction"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func handleAuctionEnding(payload []byte) {
	client, _, ccsvc, signer := load_ctcs()

	var a ecomm.Auction
	err := json.Unmarshal(payload, &a)
	check(err)

	var addr string
	if platform == "ethereum" {
		addr = a.EthAddr
	} else {
		addr = a.QuorumAddr
	}
	result := ecomm.AuctionResult{
		Platform:      platform,
		HostAuctionID: a.ID,
		AuctionAddr:   addr,
	}

	contract, err := eth_auction.NewAuction(common.HexToAddress(addr), client)
	check(err)
	opts := &bind.CallOpts{}

	highestBid, err := contract.HighestBid(opts)
	check(err)

	highestBidder, err := contract.HighestBidder(opts)
	check(err)

	result.HighestBid = int(highestBid.Int64())
	result.HighestBidder = highestBidder.Hex()

	signed := ecomm.SignedAuctionResult{
		AuctionResult: result,
	}
	signed.Signature, err = signer.Sign(result.Hash())
	check(err)

	b, _ := json.Marshal(signed)
	ccsvc.Publish(ecomm.SignedAuctionResultEvent, b)
}

func handleAuctionCreating(payload []byte) {
	fmt.Println("Handle Auction Creating")

	var a ecomm.Auction
	err := json.Unmarshal(payload, &a)
	check(err)

	var addr string
	if platform == "eth" {
		addr = a.EthAddr
	} else {
		addr = a.QuorumAddr
	}
	result := ecomm.AuctionResult{
		Platform:      platform,
		HostAuctionID: a.ID,
		AuctionAddr:   addr,
	}

	fmt.Println("Auction Addr: ", result.AuctionAddr)

}
