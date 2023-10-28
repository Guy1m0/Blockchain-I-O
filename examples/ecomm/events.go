package ecomm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	AssetAddingEvent     = "auctioner.add_asset"
	AuctionStartingEvent = "relayer.start_auction"

	AuctionClosingEvent    = "auctioner.close_auction"
	AuctionCancelingEvent  = "auctioner.cancel_auction"
	AuctionFinalizingEvent = "relayer.fin_auction"

	// tx_mined on eth/quo
	// end AucCrEvt

	BidEvent                 = "bidder.bid_auc"
	WithdrawEvent            = "eth_quo.withdraw"
	CommitAuctionResultEvent = "bidder.cmt_result"
	AbortAuctionResultEvent  = "bidder.abt_result"

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
	AsseetID string
	Bidder   common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
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
