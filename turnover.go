package moexiss

//Turnover struct represents market turnovers
type Turnover struct {
	Name        string  // "NAME" Market text identifier
	Id          int64   // "ID" Market ID
	ValToday    float64 // "VALTODAY" Value of Concluded Transactions, million RUB
	ValTodayUsd float64 // "VALTODAY_USD" Value of Concluded Transactions, million USD
	NumTrades   int64   // "NUMTRADES" Quantity of Trades per Day, units
	UpdateTime  string  // "UPDATETIME" Time of Last Updating
	Title       string  // "TITLE" Market title
}

const (
	turnoverKeyName        = "NAME"
	turnoverKeyId          = "ID"
	turnoverKeyValToday    = "VALTODAY"
	turnoverKeyValTodayUsd = "VALTODAY_USD"
	turnoverKeyNumTrades   = "NUMTRADES"
	turnoverKeyUpdateTime  = "UPDATETIME"
	turnoverKeyTitle       = "TITLE"
)
