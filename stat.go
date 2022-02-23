package moexiss

import (
	"path"
	"unicode/utf8"
)

const (
	statsPartsUrl = "secstats.json"

	secStatKeyTicker           = "SECID"
	secStatKeyBoardId          = "BOARDID"
	secStatKeyTrSession        = "TRADINGSESSION"
	secStatKeyTime             = "TIME"
	secStatKeyPriceMinusPrevPr = "PRICEMINUSPREVWAPRICE"
	secStatKeyVolToday         = "VOLTODAY"
	secStatKeyValToday         = "VALTODAY"
	secStatKeyHighBid          = "HIGHBID"
	secStatKeyLowOffer         = "LOWOFFER"
	secStatKeyLastOffer        = "LASTOFFER"
	secStatKeyLastBid          = "LASTBID"
	secStatKeyOpen             = "OPEN"
	secStatKeyLow              = "LOW"
	secStatKeyHigh             = "HIGH"
	secStatKeyLast             = "LAST"
	secStatKeyLClosePrice      = "LCLOSEPRICE"
	secStatKeyNumTrades        = "NUMTRADES"
	secStatKeyWaPrice          = "WAPRICE"
	secStatKeyAdmittedQuote    = "ADMITTEDQUOTE"
	secStatKeyMarketPrice      = "MARKETPRICE2"
	secStatKeyLCurrentPrice    = "LCURRENTPRICE"
	secStatKeyClosingAucPrice  = "CLOSINGAUCTIONPRICE"
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

	url, _ := s.client.BaseURL.Parse(enginePartOfPath)

	url.Path = path.Join(url.Path, engine.String(), marketsPartOfPath, market, statsPartsUrl)
	gotURL := addStatRequestOptions(url, opt)
	return gotURL.String(), nil
}

func parseSecStatItem(data []byte, ss *SecStat) (err error) {

	ticker, err := parseStringWithDefaultValueByKey(data, secStatKeyTicker, "")
	if err != nil {
		return
	}

	boardId, err := parseStringWithDefaultValueByKey(data, secStatKeyBoardId, "")
	if err != nil {
		return
	}

	trSessionStr, err := parseStringWithDefaultValueByKey(data, secStatKeyTrSession, "0")
	if err != nil {
		return
	}
	trSession := getTradingSession(trSessionStr)

	time, err := parseStringWithDefaultValueByKey(data, secStatKeyTime, "")
	if err != nil {
		return
	}

	priceMinus, err := parseFloatWithDefaultValue(data, secStatKeyPriceMinusPrevPr)
	if err != nil {
		return
	}

	volToday, err := parseIntWithDefaultValue(data, secStatKeyVolToday)
	if err != nil {
		return
	}

	valToday, err := parseIntWithDefaultValue(data, secStatKeyValToday)
	if err != nil {
		return
	}

	highBid, err := parseFloatWithDefaultValue(data, secStatKeyHighBid)
	if err != nil {
		return
	}

	lowOffer, err := parseFloatWithDefaultValue(data, secStatKeyLowOffer)
	if err != nil {
		return
	}

	lastOffer, err := parseFloatWithDefaultValue(data, secStatKeyLastOffer)
	if err != nil {
		return
	}

	lastBid, err := parseFloatWithDefaultValue(data, secStatKeyLastBid)
	if err != nil {
		return
	}

	open, err := parseFloatWithDefaultValue(data, secStatKeyOpen)
	if err != nil {
		return
	}

	low, err := parseFloatWithDefaultValue(data, secStatKeyLow)
	if err != nil {
		return
	}

	high, err := parseFloatWithDefaultValue(data, secStatKeyHigh)
	if err != nil {
		return
	}

	last, err := parseFloatWithDefaultValue(data, secStatKeyLast)
	if err != nil {
		return
	}

	lClosePrice, err := parseFloatWithDefaultValue(data, secStatKeyLClosePrice)
	if err != nil {
		return
	}

	numTrades, err := parseIntWithDefaultValue(data, secStatKeyNumTrades)
	if err != nil {
		return
	}

	waPrice, err := parseFloatWithDefaultValue(data, secStatKeyWaPrice)
	if err != nil {
		return
	}

	admittedQuote, err := parseFloatWithDefaultValue(data, secStatKeyAdmittedQuote)
	if err != nil {
		return
	}

	marketPrice, err := parseFloatWithDefaultValue(data, secStatKeyMarketPrice)
	if err != nil {
		return
	}

	lCurrentPrice, err := parseFloatWithDefaultValue(data, secStatKeyLCurrentPrice)
	if err != nil {
		return
	}

	closingAucPrice, err := parseFloatWithDefaultValue(data, secStatKeyClosingAucPrice)
	if err != nil {
		return
	}

	ss.Ticker = ticker
	ss.BoardId = boardId
	ss.TrSession = trSession
	ss.Time = time
	ss.PriceMinusPrevPr = priceMinus
	ss.VolToday = volToday
	ss.ValToday = valToday
	ss.HighBid = highBid
	ss.LowOffer = lowOffer
	ss.LastOffer = lastOffer
	ss.LastBid = lastBid
	ss.Open = open
	ss.Low = low
	ss.High = high
	ss.Last = last
	ss.LClosePrice = lClosePrice
	ss.NumTrades = numTrades
	ss.WaPrice = waPrice
	ss.AdmittedQuote = admittedQuote
	ss.MarketPrice = marketPrice
	ss.LCurrentPrice = lCurrentPrice
	ss.ClosingAucPrice = closingAucPrice

	return
}
