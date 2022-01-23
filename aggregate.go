package moexiss

import (
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
)

// AggregateService gets aggregated trading results
// in the MoEx ISS API.
//
// MoEx ISS API docs: https://iss.moex.com/iss/reference/214
type AggregateService service

//getUrl provides an url for a request of the aggregates with parameters from AggregateRequestOptions
//opt *AggregateRequestOptions can be nil, it is safe
func (s *AggregateService) getUrl(security string, opt *AggregateRequestOptions) string {
	url, _ := s.client.BaseURL.Parse("securities")

	url.Path = path.Join(url.Path, security, aggregatesPartsUrl)
	gotUrl := addAggregateRequestOptions(url, opt)
	return gotUrl.String()
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
