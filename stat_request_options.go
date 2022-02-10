package moexiss

import (
	"strconv"
)

// TradingSession represents a type of trading sessions for intermediate day summary
// of MoEx ISS API
type TradingSession uint8

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
