package moexiss

import (
	"strconv"
)

// TradingSession represents a type of trading sessions for intermediate day summary
// of MoEx ISS API
type TradingSession uint8

// These constants represent possible values of TradingSession
const (
	TradingSessionUndefined  TradingSession = 0
	TradingSessionMain       TradingSession = 1
	TradingSessionAdditional TradingSession = 2
	TradingSessionTotal      TradingSession = 3
)

// String representations of TradingSession values
func (ts TradingSession) String() string {
	return strconv.Itoa(int(ts))
}

// StatRequestOptions contains options which can be used as arguments
// for building requests to get intermediate day summary.
// MoEx ISS API docs: https://iss.moex.com/iss/reference/823
type StatRequestOptions struct {
	TradingSessionType TradingSession // `tradingsession` query parameter in url.URL
	TickerIds          []string       // `securities` query parameter in url.URL
	BoardId            []string       // `boardid` query parameter in url.URL
}

// StatReqOptionsBuilder represents a builder of StatRequestOptions struct
type StatReqOptionsBuilder struct {
	options *StatRequestOptions
}

// NewStatReqOptionsBuilder is a constructor of StatReqOptionsBuilder
func NewStatReqOptionsBuilder() *StatReqOptionsBuilder {
	return &StatReqOptionsBuilder{options: &StatRequestOptions{}}
}

// Build builds StatRequestOptions from StatReqOptionsBuilder
func (b *StatReqOptionsBuilder) Build() *StatRequestOptions {
	return b.options
}

// TypeTradingSession sets a type of trading session parameter to a request
// It allows to show data only for the required session.
func (b *StatReqOptionsBuilder) TypeTradingSession(ts TradingSession) *StatReqOptionsBuilder {
	b.options.TradingSessionType = ts
	return b
}

// AddTicker adds a ticker to a request
// It allows to show data only for the required tickers.
// No more than 10 tickers.
func (b *StatReqOptionsBuilder) AddTicker(ticker string) *StatReqOptionsBuilder {
	b.options.TickerIds = append(b.options.TickerIds, ticker)
	return b
}

// AddBoard adds a board to a request
// Filter the output by trading mode.
// No more than 10 boards.
func (b *StatReqOptionsBuilder) AddBoard(boardId string) *StatReqOptionsBuilder {
	b.options.BoardId = append(b.options.BoardId, boardId)
	return b
}
