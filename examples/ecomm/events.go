package ecomm

const (
	OnInitializedLending = "ecomm.initialize"
	OnLoanSuccessful     = "ecomm.successful"
	OnLoanFail           = "ecomm.fail"
)

var (
	SignedAuctionResultEvent = "signed_result@ccsvc"

	AuctionEndingEvent        = "auction_ending@ccsvc"
	ProceedAuctionResultEvent = "prcd_result@ccsvc"
	AbortAuctionResultEvent   = "abt_result@ccsvc"

	AuctionCreatingEvent = "auction_creating@ccsvc"
	//AddingAssetEvent     = "add_asset@fabric"
	// tx_mined on eth/quo
	// end AucCrEvt

	BiddingAuctionEvent   = "bid_auc@eth/quo"
	TransactionMinedEvent = "tx_mined@eth/quo"
	// end with BiddingAuctionEvent
)

type SignedAuctionResult struct {
	AuctionResult
	Signature []byte
}
