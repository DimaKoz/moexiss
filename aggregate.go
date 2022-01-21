package moexiss

import (
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
	nullValueData := "null"

	marketNameData, _, _, err := jsonparser.Get(data, aggKeyMarketName)
	if err != nil {
		return
	}
	marketName, err := parseStringWithDefaultValue(marketNameData)
	if err != nil {
		return
	}

	marketTitleData, _, _, err := jsonparser.Get(data, aggKeyMarketTitle)
	if err != nil {
		return
	}
	marketTitle, err := parseStringWithDefaultValue(marketTitleData)
	if err != nil {
		return
	}

	engineData, _, _, err := jsonparser.Get(data, aggKeyEngine)
	if err != nil {
		return
	}
	engine, err := parseStringWithDefaultValue(engineData)
	if err != nil {
		return
	}

	tradeDateData, _, _, err := jsonparser.Get(data, aggKeyTradeDate)
	if err != nil {
		return
	}
	tradeDate, err := parseStringWithDefaultValue(tradeDateData)
	if err != nil {
		return
	}

	secIdData, _, _, err := jsonparser.Get(data, aggKeySecurityId)
	if err != nil {
		return
	}
	secId, err := parseStringWithDefaultValue(secIdData)
	if err != nil {
		return
	}

	valueData, _, _, err := jsonparser.Get(data, aggKeyValue)
	if string(valueData) == nullValueData {
		valueData = []byte("0")
	}
	if err != nil {
		return
	}
	value, err := jsonparser.ParseFloat(valueData)
	if err != nil {
		return
	}

	volumeData, _, _, err := jsonparser.Get(data, aggKeyVolume)
	if err != nil {
		return
	}
	if string(volumeData) == nullValueData {
		volumeData = []byte("0")
	}
	volume, err := jsonparser.ParseInt(volumeData)
	if err != nil {
		return
	}

	numTradesData, _, _, err := jsonparser.Get(data, aggKeyNumberTrades)
	if err != nil {
		return
	}
	if string(numTradesData) == nullValueData {
		volumeData = []byte("0")
	}
	numTrades, err := jsonparser.ParseInt(numTradesData)
	if err != nil {
		return
	}

	updateAtData, _, _, err := jsonparser.Get(data, aggKeyUpdatedAt)
	if err != nil {
		return
	}
	updateAt, err := parseStringWithDefaultValue(updateAtData)
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
