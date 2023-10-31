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
	AuctionFinalizingEvent = "auctioner.fin_auction"

	AuctionStateUpdatingEvent = "relayer.update_auction_state"

	BidEvent                 = "bidder.bid_auc"
	WithdrawEvent            = "bidder.withdraw"
	CommitAuctionResultEvent = "bidder.cmt_result"
	AbortAuctionResultEvent  = "bidder.abt_result"

	TransactionMinedEvent = "eth_quo.tx_mined"
	//SignAuctionResultEvent = "eth_quo.signed_result"

	ProvideFeedbackEvent = "eth_quo.provide_feedback"

	KafkaReceivedEvent   = "kafka.received"
	RelayerDetectedEvent = "relayer.detected"
	// end with BiddingAuctionEvent
)

type HighestBidIncreased struct {
	Id     string
	Bidder common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

type WithdrawBid struct {
	Id     string
	Bidder common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

type DecisionMade struct {
	Winner     common.Address
	Amount     *big.Int
	Id         string
	Prcd       bool
	JsonString string
	Raw        types.Log // Blockchain specific contextual infos
}

type WaitResponse struct {
	Winner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}
