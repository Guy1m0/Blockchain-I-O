package ecomm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	OnInitializedLending = "ecomm.initialize"
	OnLoanSuccessful     = "ecomm.successful"
	OnLoanFail           = "ecomm.fail"
)

var (
	SignedAuctionResultEvent = "auction.signed_result"

	ProceedAuctionResultEvent = "ccsvc.prcd_result"
	AbortAuctionResultEvent   = "ccsvc.abt_result"
	AuctionStartingEvent      = "ccsvc.start_auction"

	// tx_mined on eth/quo
	// end AucCrEvt
	AddingAssetEvent   = "fabric.add_asset"
	AuctionEndingEvent = "fabric.end_auction"

	BiddingAuctionEvent   = "eth_quo.bid_auc"
	WithdrawEvent         = "eth_quo.withdraw"
	TransactionMinedEvent = "eth_quo.tx_mined"

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
	Winner common.Address
	Amount *big.Int
	Id     string
	Raw    types.Log // Blockchain specific contextual infos
}

type WaitResponseEvent struct {
	Winner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}
