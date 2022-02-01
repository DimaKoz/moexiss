package moexiss

import (
	"bufio"
	"bytes"
	"context"
	"github.com/buger/jsonparser"
	"path"
)

//Aggregate struct represents aggregated trading results
//for the date by the security by markets
type Aggregate struct {
	MarketName   string  // "market_name"
	MarketTitle  string  // "market_title"
	Engine       string  // "engine"
	TradeDate    string  // "tradedate"
	SecurityId   string  // "secid"
	Value        float64 // "value"
	Volume       int64   // "volume"
	NumberTrades int64   // "numtrades"
	UpdatedAt    string  // "updated_at"
}

//AggregatesResponse struct represents a response with aggregated trading results
type AggregatesResponse struct {
	SecurityId string
	Aggregates []Aggregate
	DatesFrom  string //
	DatesTill  string
}

const (
	aggregatesPartsUrl = "aggregates.json"

	aggKeyMarketName   = "market_name"
	aggKeyMarketTitle  = "market_title"
	aggKeyEngine       = "engine"
	aggKeyTradeDate    = "tradedate"
	aggKeySecurityId   = "secid"
	aggKeyValue        = "value"
	aggKeyVolume       = "volume"
	aggKeyNumberTrades = "numtrades"
	aggKeyUpdatedAt    = "updated_at"

	aggKeyAggregates = "aggregates"
	aggKeyDates      = "agregates.dates"
	aggKeyFrom       = "from"
	aggKeyTill       = "till"
)

// AggregateService gets aggregated trading results
// in the MoEx ISS API.
//
// MoEx ISS API docs: https://iss.moex.com/iss/reference/214
type AggregateService service

//Aggregates provides Aggregates of MoEx ISS
func (a *AggregateService) Aggregates(ctx context.Context, security string, opt *AggregateRequestOptions) (*AggregatesResponse, error) {

	url, err := a.getUrl(security, opt)
	if err != nil {
		return nil, err
	}
	req, err := a.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	_, err = a.client.Do(ctx, req, w)
	if err != nil {
		return nil, err
	}
	ar := AggregatesResponse{}
	err = parseAggregateResponse(b.Bytes(), &ar)
	if err != nil {
		return nil, err
	}
	ar.SecurityId = security
	return &ar, nil
}

//getUrl provides an url for a request of the aggregates with parameters from AggregateRequestOptions
//opt *AggregateRequestOptions can be nil, it is safe
func (a *AggregateService) getUrl(security string, opt *AggregateRequestOptions) (string, error) {
	if !isOkSecurityParam(security) {
		return "", errBadSecurityParameter
	}
	url, _ := a.client.BaseURL.Parse("securities")

	url.Path = path.Join(url.Path, security, aggregatesPartsUrl)
	gotUrl := addAggregateRequestOptions(url, opt)
	return gotUrl.String(), nil
}

func parseAggregateResponse(byteData []byte, aggregatesResponse *AggregatesResponse) error {
	var err error
	if aggregatesResponse == nil {
		err = errNilPointer
		return err
	}
	var errInCb error
	_, err = jsonparser.ArrayEach(byteData, func(aggregatesBytes []byte, _ jsonparser.ValueType, offset int, errCb error) {
		var data []byte
		var dataType jsonparser.ValueType
		data, dataType, _, errInCb = jsonparser.Get(aggregatesBytes, aggKeyAggregates)
		if errInCb == nil && data != nil && dataType == jsonparser.Array {
			errInCb = parseAggregates(data, &aggregatesResponse.Aggregates)
			if errInCb != nil {
				return
			}
		}
		data, dataType, _, errInCb = jsonparser.Get(aggregatesBytes, aggKeyDates)
		if errInCb == nil && data != nil && dataType == jsonparser.Array {
			errInCb = parseAggregatesDates(data, aggregatesResponse)
			if errInCb != nil {
				return
			}
		}
	})
	if err == nil && errInCb != nil {
		err = errInCb
	}
	return err
}

func parseAggregatesDates(dataDates []byte, ar *AggregatesResponse) (err error) {
	var errInCb error
	counter := 0
	_, err = jsonparser.ArrayEach(dataDates, func(aggregateDateData []byte, dataType jsonparser.ValueType, offset int, errCb error) {
		if counter > 0 { // getting the first item only
			return
		}
		counter++

		if dataType != jsonparser.Object {
			errInCb = errUnexpectedDataType
			return
		}

		errInCb = parseAggregatesDate(aggregateDateData, ar)
		if errInCb != nil {
			return
		}

	})
	if err == nil && errInCb != nil {
		err = errInCb
	}
	return

}

func parseAggregatesDate(data []byte, ar *AggregatesResponse) (err error) {

	from, err := parseStringWithDefaultValueByKey(data, aggKeyFrom, "")
	if err != nil {
		return
	}

	till, err := parseStringWithDefaultValueByKey(data, aggKeyTill, "")
	if err != nil {
		return
	}

	ar.DatesFrom = from
	ar.DatesTill = till
	return
}

func parseAggregates(byteData []byte, aggregates *[]Aggregate) (err error) {
	var errInCb error
	_, err = jsonparser.ArrayEach(byteData, func(aggregateItemData []byte, dataType jsonparser.ValueType, offset int, errCb error) {
		if errInCb != nil {
			return
		}
		if dataType != jsonparser.Object {
			errInCb = errUnexpectedDataType
			return
		}

		aggregate := Aggregate{}
		errInCb = parseAggregate(aggregateItemData, &aggregate)
		if errInCb != nil {
			return
		}
		*aggregates = append(*aggregates, aggregate)

	})
	if err == nil && errInCb != nil {
		err = errInCb
	}
	return
}

func parseAggregate(data []byte, a *Aggregate) (err error) {

	marketName, err := parseStringWithDefaultValueByKey(data, aggKeyMarketName, "")
	if err != nil {
		return
	}

	marketTitle, err := parseStringWithDefaultValueByKey(data, aggKeyMarketTitle, "")
	if err != nil {
		return
	}

	engine, err := parseStringWithDefaultValueByKey(data, aggKeyEngine, "")
	if err != nil {
		return
	}

	tradeDate, err := parseStringWithDefaultValueByKey(data, aggKeyTradeDate, "")
	if err != nil {
		return
	}

	secId, err := parseStringWithDefaultValueByKey(data, aggKeySecurityId, "")
	if err != nil {
		return
	}

	value, err := parseFloatWithDefaultValue(data, aggKeyValue)
	if err != nil {
		return
	}

	volume, err := parseIntWithDefaultValue(data, aggKeyVolume)
	if err != nil {
		return
	}

	numTrades, err := parseIntWithDefaultValue(data, aggKeyNumberTrades)
	if err != nil {
		return
	}

	updateAt, err := parseStringWithDefaultValueByKey(data, aggKeyUpdatedAt, "")
	if err != nil {
		return
	}

	a.MarketName = marketName
	a.MarketTitle = marketTitle
	a.Engine = engine
	a.TradeDate = tradeDate
	a.SecurityId = secId
	a.Value = value
	a.Volume = volume
	a.NumberTrades = numTrades
	a.UpdatedAt = updateAt

	return
}
