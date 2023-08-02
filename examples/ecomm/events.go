package ecomm

const (
	OnInitializedLending = "ecomm.initialize"
	OnLoanSuccessful     = "ecomm.successful"
	OnLoanFail           = "ecomm.fail"
)

var (
	SignedAuctionResultEvent = "signed_result@ccsvc"
	AuctionEndingEvent       = "auction_ending@ccsvc"
	AuctionCreatingEvent     = "auction_creating@ccsvc"

	ProceedAuctionResultEvent = "prcd_result@ccsvc"
	AbortAuctionResultEvent   = "abt_result@ccsvc"

	AddingAssetEvent      = "add_asset"
	BiddingAuctionEvent   = "bid_auc"
	TransactionMinedEvent = "tx_mined"
)

type SignedAuctionResult struct {
	AuctionResult
	Signature []byte
}
