package ecomm

const (
	OnInitializedLending = "ecomm.initialize"
	OnLoanSuccessful     = "ecomm.successful"
	OnLoanFail           = "ecomm.fail"
)

var (
	SignedAuctionResultEvent = "signed_result"
	AuctionEndingEvent       = "auction_ending"
	AuctionCreatingEvent     = "auction_creating"

	ProceedAuctionResultEvent = "prcd_result"
	AbortAuctionResultEvent   = "abt_result"
)

type SignedAuctionResult struct {
	AuctionResult
	Signature []byte
}
