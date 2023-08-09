package ecomm

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
