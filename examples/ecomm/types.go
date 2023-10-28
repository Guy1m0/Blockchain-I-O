package ecomm

import (
	"encoding/json"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
)

type UserInfo struct {
	UserID  string
	Address common.Address
	KeyFile string
}

// may not need
type AuctionInfo struct {
	AssetID    string
	ID         int
	EthAddress common.Address
	QuoAddress common.Address
}

type Erc20Info struct {
	FabricTokenName string
	EthERC20        common.Address
	QuoERC20        common.Address
}

type Asset struct {
	ID               string `json:"id"`
	Owner            string `json:"owner"`
	PendingAuctionID int    `json:"pendingAuctionId"`
}

type Auction struct {
	ID         int    `json:"id"`
	AssetID    string `json:"assetId"`
	EthAddr    string `json:"ethAddr"`
	QuorumAddr string `json:"quorumAddr"`
	Status     string `json:"status"`

	HighestBid         big.Int `json:"highestBid"`
	HighestBidder      string  `json:"highestBidder"`
	HighestBidPlatform string  `json:"highestBidPlatform"`
}

type Bid struct {
	Bidder      common.Address `json:"bidder"`
	BidAmount   big.Int        `json:"bidAmount"`
	AuctionAddr common.Address `json:"auctionAddr"`

	Platform string `json:"platform"`

	AuctionID int    `json:"auctionID"`
	AssetID   string `json:"assetID"`
}

type Tx struct {
	Platform string
	Type     string
	Hash     common.Hash
}

// Struct used as input for creating new Auction
type StartAuctionArgs struct {
	AssetID    string
	EthAddr    string
	QuorumAddr string

	Signature []byte
}

func (args *StartAuctionArgs) Hash() []byte {
	h := sha3.New256()

	h.Write([]byte(args.AssetID))
	h.Write([]byte(args.EthAddr))
	h.Write([]byte(args.QuorumAddr))

	return h.Sum(nil)
}

type AuctionResult struct {
	Platform    string
	AuctionID   int
	AuctionAddr string

	HighestBid    big.Int
	HighestBidder string

	Signature []byte
}

type SignedAuctionResult struct {
	AuctionResult
	Signature []byte
}

type EventWrapper struct {
	Type   string          `json:"type"`
	Result json.RawMessage `json:"result"`
}

func (ar *AuctionResult) Hash() []byte {
	h := sha3.New256()

	h.Write([]byte(ar.Platform))
	h.Write([]byte(strconv.Itoa(ar.AuctionID)))
	h.Write([]byte(ar.AuctionAddr))

	h.Write([]byte(ar.HighestBid.String()))
	h.Write([]byte(ar.HighestBidder))

	h.Write([]byte(""))

	return h.Sum(nil)
}

type CrossChainAuctionResult struct {
	AuctionResult
	Signatures [][]byte
}

type FinalizeAuctionArgs struct {
	AuctionID    int
	EthResult    CrossChainAuctionResult
	QuorumResult CrossChainAuctionResult
}
