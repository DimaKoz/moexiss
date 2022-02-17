package moexiss

import (
	"path"
	"unicode/utf8"
)

const (
	statsPartsUrl = "secstats.json"
)

// SecStat struct represents intermediate day summary
// by the security by markets
type SecStat struct {
	Ticker           string         // "SECID"
	BoardId          string         // "BOARDID"
	TrSession        TradingSession // "TRADINGSESSION"
	Time             string         // "TIME"
	PriceMinusPrevPr float64        // "PRICEMINUSPREVWAPRICE"
	VolToday         int64          // "VOLTODAY"
	ValToday         int64          // "VALTODAY"
	HighBid          float64        // "HIGHBID"
	LowOffer         float64        // "LOWOFFER"
	LastOffer        float64        // "LASTOFFER"
	LastBid          float64        // "LASTBID"
	Open             float64        // "OPEN"
	Low              float64        // "LOW"
	High             float64        // "HIGH"
	Last             float64        // "LAST"
	LClosePrice      float64        // "LCLOSEPRICE"
	NumTrades        int64          // "NUMTRADES"
	WaPrice          float64        // "WAPRICE"
	AdmittedQuote    float64        // "ADMITTEDQUOTE"
	MarketPrice      float64        // "MARKETPRICE2"
	LCurrentPrice    float64        // "LCURRENTPRICE"
	ClosingAucPrice  float64        // "CLOSINGAUCTIONPRICE"
}

// StatsService gets intermediate day summary
// from the MoEx ISS API.
//
// MoEx ISS API docs: https://iss.moex.com/iss/reference/823
type StatsService service

// getUrl provides an url to get intermediate day summary
// opt *StatRequestOptions can be nil, it is safe
func (s *StatsService) getUrl(engine EngineName, market string, opt *StatRequestOptions) (string, error) {
	if engine == EngineUndefined {
		return "", ErrBadEngineParameter
	}
	marketMinLen := 3
	if market == "" || utf8.RuneCountInString(market) < marketMinLen {
		return "", ErrBadMarketParameter
	}

	url, _ := s.client.BaseURL.Parse("engines")

	url.Path = path.Join(url.Path, engine.String(), "markets", market, statsPartsUrl)
	gotURL := addStatRequestOptions(url, opt)
	return gotURL.String(), nil
}
