package ecomm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	CommitAuctionResultEvent = "eth_quo.cmt_result"
	AbortAuctionResultEvent  = "eth_quo.abt_result"

	AssetAddingEvent       = "fabric.add_asset"
	AuctionStartingEvent   = "ccsvc.start_auction"
	AuctionFinalizingEvent = "ccsvc.fin_auction"
	AuctionClosingEvent    = "fabric.close_auction"
	AuctionCancelingEvent  = "fabric.cancel_auction"

	// tx_mined on eth/quo
	// end AucCrEvt

	BidEvent      = "eth_quo.bid_auc"
	WithdrawEvent = "eth_quo.withdraw"

	TransactionMinedEvent = "eth_quo.tx_mined"
	//SignAuctionResultEvent = "eth_quo.signed_result"

	ProvideFeedbackEvent = "eth_quo.provide_feedback"

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
