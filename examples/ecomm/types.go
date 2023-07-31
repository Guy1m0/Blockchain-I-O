package ecomm

import (
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
)

// update later
type EComm struct {
	LenderContract    string
	ArbitrageContract string

	Lender      string
	Arbitrageur string
	Exchange    string

	Loan    int64
	Intrest int64
}

func (fl *EComm) Hash() string {
	hash := fl.Hash32()
	return common.Bytes2Hex(hash[:])
}

func (fl *EComm) Hash32() [32]byte {
	h := sha3.New256()
	h.Write([]byte(fl.LenderContract))
	h.Write([]byte(fl.ArbitrageContract))
	h.Write([]byte(fl.Lender))
	h.Write([]byte(fl.Arbitrageur))
	h.Write([]byte(fl.Exchange))
	h.Write([]byte(fmt.Sprint(fl.Loan)))
	h.Write([]byte(fmt.Sprint(fl.Intrest)))
	hash := h.Sum(nil)
	ret := [32]byte{}
	copy(ret[:], hash)
	return ret
}

// update later
type CommitVote struct {
	LoanHash     string
	LenderSig    string
	ArbitrageSig string
}

type SetupInfo struct {
	Bidder1Address   common.Address
	Bidder2Address   common.Address
	AuctionerAddress common.Address

	FabricTokenName string

	EthERC20 common.Address
	QuoERC20 common.Address
}

type UserInfo struct {
	UserID  string
	Address common.Address
	KeyFile string
}

type Erc20Info struct {
	FabricTokenName string
	EthERC20        common.Address
	QuoERC20        common.Address
}

type BidderInfo struct {
	ZkNodes  string
	Endpoint string

	Keyfile  string
	Platform string

	Erc20 common.Address
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

	HighestBid         int
	HighestBidder      string
	HighestBidPlatform string
}

type StartAuctionArgs struct {
	AssetID    string
	EthAddr    string
	QuorumAddr string
	Signature  string
}

func (args *StartAuctionArgs) Hash() []byte {
	h := sha3.New256()
	h.Write([]byte(args.AssetID))
	h.Write([]byte(args.EthAddr))
	h.Write([]byte(args.QuorumAddr))
	return h.Sum(nil)
}

type AuctionResult struct {
	Platform      string
	HostAuctionID int
	AuctionAddr   string
	HighestBid    int
	HighestBidder string
}

func (args *AuctionResult) Hash() []byte {
	h := sha3.New256()
	h.Write([]byte(args.AuctionAddr))
	h.Write([]byte(args.Platform))
	h.Write([]byte(strconv.Itoa(args.HighestBid)))
	h.Write([]byte(args.HighestBidder))
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
