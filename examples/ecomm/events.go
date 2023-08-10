package ecomm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	ProceedAuctionResultEvent = "eth_quo.prcd_result"
	AbortAuctionResultEvent   = "eth_quo.abt_result"
	AuctionStartingEvent      = "ccsvc.start_auction"
	AuctionClosedEvent        = "ccsvc.auction_closed"

	// tx_mined on eth/quo
	// end AucCrEvt
	AddingAssetEvent    = "fabric.add_asset"
	AuctionClosingEvent = "fabric.close_auction"

	BiddingAuctionEvent      = "eth_quo.bid_auc"
	WithdrawEvent            = "eth_quo.withdraw"
	TransactionMinedEvent    = "eth_quo.tx_mined"
	SignedAuctionResultEvent = "eth_quo.signed_result"
	FeedBackEvent            = "eth_quo.provide_feedback"

	KafkaReceivedEvent   = "kafka.received"
	RelayerDetectedEvent = "relayer.detected"
	// end with BiddingAuctionEvent
)

type SignedAuctionResult struct {
	AuctionResult
	Signature []byte
}

type HighestBidIncreasedEvent struct {
	Bidder common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

type DecisionMadeEvent struct {
	Winner     common.Address
	Amount     *big.Int
	Id         string
	Prcd       bool
	JsonString string
	Raw        types.Log // Blockchain specific contextual infos
}

type WaitResponseEvent struct {
	Winner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}
