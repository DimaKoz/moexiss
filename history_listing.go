package moexiss

import (
	"github.com/buger/jsonparser"
	"path"
)

//Listing struct represents listing of the security
type Listing struct {
	Ticker    string // "SECID"
	ShortName string // "SHORTNAME"
	FullName  string // "NAME"
	BoardId   string // "BOARDID"
	Decimals  int64  // "decimals"
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

func parseListing(data []byte, l *[]Listing) (err error) {

	var errInCb error
	_, err = jsonparser.ArrayEach(data, func(listingItemData []byte, dataType jsonparser.ValueType, offset int, errCb error) {
		if errInCb != nil {
			return
		}
		if dataType != jsonparser.Object {
			errInCb = ErrUnexpectedDataType
			return
		}

		listing := Listing{}
		errInCb = parseListingItem(listingItemData, &listing)
		if errInCb != nil {
			return
		}
		*l = append(*l, listing)

	})
	if err == nil && errInCb != nil {
		err = errInCb
	}
	return
}

func parseListingItem(data []byte, l *Listing) (err error) {

	ticker, err := parseStringWithDefaultValueByKey(data, listingKeyId, "")
	if err != nil {
		return
	}

	shortName, err := parseStringWithDefaultValueByKey(data, listingKeyShortName, "")
	if err != nil {
		return
	}

	name, err := parseStringWithDefaultValueByKey(data, listingKeyName, "")
	if err != nil {
		return
	}

	boardId, err := parseStringWithDefaultValueByKey(data, listingKeyBoardId, "")
	if err != nil {
		return
	}

	decimal, err := parseIntWithDefaultValue(data, listingKeyDecimals)
	if err != nil {
		return
	}

	from, err := parseStringWithDefaultValueByKey(data, listingKeyFrom, "")
	if err != nil {
		return
	}

	till, err := parseStringWithDefaultValueByKey(data, listingKeyTill, "")
	if err != nil {
		return
	}

	l.Ticker = ticker
	l.ShortName = shortName
	l.FullName = name
	l.BoardId = boardId
	l.Decimals = decimal
	l.From = from
	l.Till = till

	return
}
