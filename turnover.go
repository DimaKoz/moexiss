package moexiss

import (
	"github.com/buger/jsonparser"
)

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

func parseTurnoverResponse(byteData []byte, turnovers *[]Turnover) error {
	var err error
	if turnovers == nil {
		err = errNilPointer
		return err
	}
	var errInCb error
	_, err = jsonparser.ArrayEach(byteData, func(turnoversBytes []byte, _ jsonparser.ValueType, offset int, errCb error) {
		var bytes []byte
		var dataType jsonparser.ValueType
		bytes, dataType, _, errInCb = jsonparser.Get(turnoversBytes, "turnovers")
		if errInCb == nil && bytes != nil && dataType == jsonparser.Array {
			errInCb = parseTurnovers(bytes, turnovers)
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

func parseTurnovers(byteData []byte, turnovers *[]Turnover) (err error) {
	var errInCb error
	_, err = jsonparser.ArrayEach(byteData, func(turnoverItemData []byte, dataType jsonparser.ValueType, offset int, errCb error) {
		if dataType != jsonparser.Object {
			errInCb = errUnexpectedDataType
			return
		}

		turnover := Turnover{}
		errInCb = parseTurnover(turnoverItemData, &turnover)
		if errInCb != nil {
			return
		}
		*turnovers = append(*turnovers, turnover)

	})
	if err == nil && errInCb != nil {
		err = errInCb
	}
	return
}

func parseTurnover(data []byte, t *Turnover) (err error) {
	nullValueData := "null"
	nameData, _, _, err := jsonparser.Get(data, turnoverKeyName)
	if err != nil {
		return
	}
	name, err := parseStringWithDefaultValue(nameData)
	if err != nil {
		return
	}
	idData, _, _, err := jsonparser.Get(data, turnoverKeyId)

	if string(idData) == nullValueData {
		idData = []byte("0")
	}
	if err != nil {
		return
	}
	id, err := jsonparser.ParseInt(idData)
	if err != nil {
		return
	}
	valTodayData, _, _, err := jsonparser.Get(data, turnoverKeyValToday)
	if string(valTodayData) == nullValueData {
		valTodayData = []byte("0")
	}
	if err != nil {
		return
	}
	valToday, err := jsonparser.ParseFloat(valTodayData)
	if err != nil {
		return
	}
	valTodayUsdData, _, _, err := jsonparser.Get(data, turnoverKeyValTodayUsd)
	if string(valTodayUsdData) == nullValueData {
		valTodayUsdData = []byte("0")
	}
	if err != nil {
		return
	}
	valTodayUsd, err := jsonparser.ParseFloat(valTodayUsdData)
	if err != nil {
		return
	}
	numTradesData, _, _, err := jsonparser.Get(data, turnoverKeyNumTrades)
	if string(numTradesData) == nullValueData {
		numTradesData = []byte("0")
	}
	if err != nil {
		return
	}
	numTrades, err := jsonparser.ParseInt(numTradesData)
	if err != nil {
		return
	}
	updateTimeData, _, _, err := jsonparser.Get(data, turnoverKeyUpdateTime)
	if err != nil {
		return
	}
	updateTime, err := parseStringWithDefaultValue(updateTimeData)
	if err != nil {
		return
	}
	titleData, _, _, err := jsonparser.Get(data, turnoverKeyTitle)
	if err != nil {
		return
	}
	title, err := parseStringWithDefaultValue(titleData)
	if err != nil {
		return
	}

	t.Name = name
	t.Id = id
	t.ValToday = valToday
	t.ValTodayUsd = valTodayUsd
	t.NumTrades = numTrades
	t.UpdateTime = updateTime
	t.Title = title

	return
}