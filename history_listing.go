package moexiss

import (
	"bufio"
	"bytes"
	"context"
	"github.com/buger/jsonparser"
	"path"
	"unicode/utf8"
)

// Listing struct represents listing of the security
type Listing struct {
	Ticker    string // "SECID"
	ShortName string // "SHORTNAME"
	FullName  string // "NAME"
	BoardId   string // "BOARDID"
	Decimals  int64  // "decimals"
	From      string // "history_from"
	Till      string // "history_till"
}

// ListingResponse struct represents a response with listing of the security
type ListingResponse struct {
	Engine       EngineName
	Market       string
	BoardGroupId string
	Listing      []Listing
}

const (
	listingKeyId         = "SECID"
	listingKeyShortName  = "SHORTNAME"
	listingKeyName       = "NAME"
	listingKeyBoardId    = "BOARDID"
	listingKeyDecimals   = "decimals"
	listingKeyFrom       = "history_from"
	listingKeyTill       = "history_till"
	listingKeySecurities = "securities"

	historyListingFilePartsUrl = "listing.json"
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

// GetListing provides a list of tradable/non-tradable securities
func (hl *HistoryListingService) GetListing(ctx context.Context, engine EngineName, market string, opt *HistoryListingRequestOptions) (*ListingResponse, error) {
	url, err := hl.getUrlListing(engine, market, opt)
	if err != nil {
		return nil, err
	}
	req, err := hl.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	_, err = hl.client.Do(ctx, req, w)
	if err != nil {
		return nil, err
	}
	lr := ListingResponse{}
	err = parseListingResponse(b.Bytes(), &lr)
	if err != nil {
		return nil, err
	}
	lr.Engine = engine
	lr.Market = market
	return &lr, nil
}

// GetListingByBoardGroup provides security listing information for a given boardgroup
func (hl *HistoryListingService) GetListingByBoardGroup(ctx context.Context, engine EngineName, market string, boardGroupId string, opt *HistoryListingRequestOptions) (*ListingResponse, error) {
	url, err := hl.getUrlListingByBoardGroup(engine, market, boardGroupId, opt)
	if err != nil {
		return nil, err
	}
	req, err := hl.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	_, err = hl.client.Do(ctx, req, w)
	if err != nil {
		return nil, err
	}
	lr := ListingResponse{}
	err = parseListingResponse(b.Bytes(), &lr)
	if err != nil {
		return nil, err
	}
	lr.Engine = engine
	lr.Market = market
	lr.BoardGroupId = boardGroupId
	return &lr, nil
}

// GetListingByBoard provides security listing information for a given board
func (hl *HistoryListingService) GetListingByBoard(ctx context.Context, engine EngineName, market string, boardId string, opt *HistoryListingRequestOptions) (*ListingResponse, error) {
	url, err := hl.getUrlListingByBoard(engine, market, boardId, opt)
	if err != nil {
		return nil, err
	}
	req, err := hl.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	_, err = hl.client.Do(ctx, req, w)
	if err != nil {
		return nil, err
	}
	lr := ListingResponse{}
	err = parseListingResponse(b.Bytes(), &lr)
	if err != nil {
		return nil, err
	}
	lr.Engine = engine
	lr.Market = market
	return &lr, nil
}

// getUrlListing provides an url to get information on when securities were traded on which boards
func (hl *HistoryListingService) getUrlListing(engine EngineName, market string, opt *HistoryListingRequestOptions) (string, error) {
	if engine == EngineUndefined {
		return "", ErrBadEngineParameter
	}
	marketMinLen := 3
	if market == "" || utf8.RuneCountInString(market) < marketMinLen {
		return "", ErrBadMarketParameter
	}

	url, _ := hl.client.BaseURL.Parse(historyPartOfPath)

	url.Path = path.Join(url.Path, enginePartOfPath, engine.String(), marketsPartOfPath, market, historyListingFilePartsUrl)
	gotURL := addHistoryListingRequestOptions(url, opt)
	return gotURL.String(), nil
}

// getUrlListingByBoard provides an url to get security listing information for a given board
func (hl *HistoryListingService) getUrlListingByBoard(engine EngineName, market string, boardId string, opt *HistoryListingRequestOptions) (string, error) {
	if engine == EngineUndefined {
		return "", ErrBadEngineParameter
	}
	marketMinLen := 3
	if market == "" || utf8.RuneCountInString(market) < marketMinLen {
		return "", ErrBadMarketParameter
	}
	boardMinLen := 4
	if boardId == "" || utf8.RuneCountInString(boardId) < boardMinLen {
		return "", ErrBadBoardParameter
	}
	url, _ := hl.client.BaseURL.Parse(historyPartOfPath)

	url.Path = path.Join(url.Path, enginePartOfPath, engine.String(), marketsPartOfPath, market, "boards", boardId, historyListingFilePartsUrl)
	gotURL := addHistoryListingRequestOptions(url, opt)
	return gotURL.String(), nil
}

// getUrlListingByBoardGroup provides an url to get security listing information for a given boardgroup
func (hl *HistoryListingService) getUrlListingByBoardGroup(engine EngineName, market string, boardGroupId string, opt *HistoryListingRequestOptions) (string, error) {
	if engine == EngineUndefined {
		return "", ErrBadEngineParameter
	}
	marketMinLen := 3
	if market == "" || utf8.RuneCountInString(market) < marketMinLen {
		return "", ErrBadMarketParameter
	}

	if boardGroupId == "" {
		return "", ErrBadBoardGroupParameter
	}
	url, _ := hl.client.BaseURL.Parse(historyPartOfPath)

	url.Path = path.Join(url.Path, enginePartOfPath, engine.String(), marketsPartOfPath, market, "boardgroups", boardGroupId, historyListingFilePartsUrl)
	gotURL := addHistoryListingRequestOptions(url, opt)
	return gotURL.String(), nil
}

func parseListingResponse(byteData []byte, listingResponse *ListingResponse) error {
	var err error
	if listingResponse == nil {
		err = ErrNilPointer
		return err
	}
	var errInCb error
	_, err = jsonparser.ArrayEach(byteData, func(listingBytes []byte, _ jsonparser.ValueType, offset int, errCb error) {
		var data []byte
		var dataType jsonparser.ValueType
		data, dataType, _, errInCb = jsonparser.Get(listingBytes, listingKeySecurities)
		if errInCb == nil && data != nil && dataType == jsonparser.Array {
			errInCb = parseListing(data, &listingResponse.Listing)
			if errInCb != nil {
				return
			}
		}
	})
	if err == nil && errInCb != nil {
		err = errInCb
	}
	if err == nil && len(listingResponse.Listing) == 0 {
		return ErrEmptyServerResult
	}
	return err
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
