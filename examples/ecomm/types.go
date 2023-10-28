package ecomm

import (
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
	ID               string
	Owner            string
	PendingAuctionID int
}

type Auction struct {
	ID         int
	AssetID    string
	EthAddr    string
	QuorumAddr string

	Status string

	HighestBid         big.Int
	HighestBidder      string
	HighestBidPlatform string
}

type Bid struct {
	Bidder common.Address

	BidAmount   big.Int
	AuctionAddr common.Address
	Platform    string

	AuctionID int
	AssetID   string
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

	HighestBid    int
	HighestBidder string

	Signature []byte
}

type SignedAuctionResult struct {
	AuctionResult
	Signature []byte
}

func (ar *AuctionResult) Hash() []byte {
	h := sha3.New256()

	h.Write([]byte(ar.Platform))
	h.Write([]byte(strconv.Itoa(ar.AuctionID)))
	h.Write([]byte(ar.AuctionAddr))

	h.Write([]byte(strconv.Itoa(ar.HighestBid)))
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
