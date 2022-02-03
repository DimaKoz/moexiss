package moexiss

//Listing struct represents listing of the security
type Listing struct {
	Ticker    string // "SECID"
	ShortName string // "SHORTNAME"
	FullName  string // "NAME"
	BoardId   string // "BOARDID"
	Decimals  int32  // "decimals"
	From      string // "history_from"
	Till      string // "history_till"
}

const (
	listingKeyId        = "SECID"
	listingKeyShortName = "SHORTNAME"
	listingKeyName      = "NAME"
	listingKeyBoardId   = "BOARDID"
	listingKeyDecimals  = "decimals"
	listingKeyFrom      = "history_from"
	listingKeyTill      = "history_till"
)

// HistoryListingService gets a list of tradable/non-tradable securities
// with indication of tradability intervals by modes
// from the MoEx ISS API.
//
// MoEx ISS API docs:
// https://iss.moex.com/iss/reference/118
// https://iss.moex.com/iss/reference/119
// https://iss.moex.com/iss/reference/120
type HistoryListingService service
