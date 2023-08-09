package ecomm

const (
	OnInitializedLending = "ecomm.initialize"
	OnLoanSuccessful     = "ecomm.successful"
	OnLoanFail           = "ecomm.fail"
)

var (
	SignedAuctionResultEvent = "auction.signed_result"

	AuctionEndingEvent        = "ccsvc.auction_ending"
	ProceedAuctionResultEvent = "ccsvc.prcd_result"
	AbortAuctionResultEvent   = "ccsvc.abt_result"

	AuctionCreatingEvent = "ccsvc.auction_creating"
	//AddingAssetEvent     = "add_asset@fabric"
	// tx_mined on eth/quo
	// end AucCrEvt

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
