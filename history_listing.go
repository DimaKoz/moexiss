package moexiss

import "path"

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

// getUrlListing provides an url to get information on when securities were traded on which boards
func (i *HistoryListingService) getUrlListing(engine EngineName, market string, opt *HistoryListingRequestOptions) (string, error) {
	url, _ := i.client.BaseURL.Parse("history/engines")

	url.Path = path.Join(url.Path, engine.String(), "markets", market, "listing.json")
	gotUrl := addHistoryListingRequestOptions(url, opt)
	return gotUrl.String(), nil
}
