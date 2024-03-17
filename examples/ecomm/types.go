package ecomm

import (
	"encoding/json"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

type UserInfo struct {
	UserID  string
	Address common.Address
	KeyFile string
}

// may not need
type AuctionInfo struct {
	Owner   common.Address
	EthAddr common.Address
	QuoAddr common.Address
}

type ContractInfo struct {
	FabricTokenName string
	EthERC20        common.Address
	QuoERC20        common.Address
	EnglishAuction  AuctionInfo
	DutchAuction    AuctionInfo
	Cb1pAuction     AuctionInfo
	Cb2pAuction     AuctionInfo
}

// type EnglishAuctionInfo struct {
// 	Owner   common.Address
// 	EthAddr common.Address
// 	QuoAddr common.Address
// }

// type ClosedBidAuctionInfo struct {
// 	Owner   common.Address
// 	EthAddr common.Address
// 	QuoAddr common.Address
// }

type Asset struct {
	ID               string `json:"id"`
	Owner            string `json:"owner"`
	PendingAuctionID int    `json:"pendingAuctionId"`
}

type Auction struct {
	ID         int    `json:"id"`
	AssetID    string `json:"assetId"`
	AucType    string `json:"aucType"`
	EthAddr    string `json:"ethAddr"`
	QuorumAddr string `json:"quorumAddr"`

	Status string `json:"status"`

	HighestBid         string `json:"highestBid"`
	HighestBidder      string `json:"highestBidder"`
	HighestBidPlatform string `json:"highestBidPlatform"`
}

type Bid struct {
	Bidder      common.Address `json:"bidder"`
	BidAmount   string         `json:"bidAmount"`
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
	AucType    string
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

type AssetAddingEventPayload struct {
	AssetID string `json:"assetId"`
	AucType string `json:"aucType"`
}

type StartAuctionEventPayload struct {
	ID      int    `json:"id"`
	AucType string `json:"aucType"`
	Owner   string `json:"owner"`
}

func VerifySignature(hash, signature []byte, addr string) bool {
	pubkey, err := crypto.SigToPub(hash, signature)
	if err != nil {
		return false
	}

	if addr == crypto.PubkeyToAddress(*pubkey).Hex() {
		return true
	}
	return false
}
