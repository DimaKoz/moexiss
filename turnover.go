package moexiss

import (
	"bufio"
	"bytes"
	"context"
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
	turnoverPartsUrl = "turnovers.json"

	turnoverKeyName        = "NAME"
	turnoverKeyId          = "ID"
	turnoverKeyValToday    = "VALTODAY"
	turnoverKeyValTodayUsd = "VALTODAY_USD"
	turnoverKeyNumTrades   = "NUMTRADES"
	turnoverKeyUpdateTime  = "UPDATETIME"
	turnoverKeyTitle       = "TITLE"
)

// TurnoverService gets turnovers on all the markets
// in the MoEx ISS API.
//
// MoEx ISS API docs: https://iss.moex.com/iss/reference/24
type TurnoverService service

//Turnovers provides a list of turnovers of markets of MoEx ISS
func (s *TurnoverService) Turnovers(ctx context.Context, opt *TurnoverRequestOptions) (*[]Turnover, error) {

	url := s.getUrl(opt, turnoversBlock)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	_, err = s.client.Do(ctx, req, w)
	if err != nil {
		return nil, err
	}
	t := make([]Turnover, 0)
	err = parseTurnoverResponse(b.Bytes(), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil

}

//getUrl provides an url for a request of the turnovers with parameters from TurnoverRequestOptions
//opt *TurnoverRequestOptions can be nil, it is safe
func (s *TurnoverService) getUrl(opt *TurnoverRequestOptions, onlyBlock turnoverBlock) string {
	url, _ := s.client.BaseURL.Parse(turnoverPartsUrl)
	gotURL := addTurnoverRequestOptions(url, opt, onlyBlock)
	return gotURL.String()
}

func parseTurnoverResponse(byteData []byte, turnovers *[]Turnover) error {
	var err error
	if turnovers == nil {
		err = ErrNilPointer
		return err
	}
	var errInCb error
	_, err = jsonparser.ArrayEach(byteData, func(turnoversBytes []byte, _ jsonparser.ValueType, offset int, errCb error) {
		var data []byte
		var dataType jsonparser.ValueType
		data, dataType, _, errInCb = jsonparser.Get(turnoversBytes, "turnovers")
		if errInCb == nil && data != nil && dataType == jsonparser.Array {
			errInCb = parseTurnovers(data, turnovers)
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
		if errInCb != nil {
			return
		}
		if dataType != jsonparser.Object {
			errInCb = ErrUnexpectedDataType
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
	name, err := parseStringWithDefaultValueByKey(data, turnoverKeyName, "")
	if err != nil {
		return
	}

	id, err := parseIntWithDefaultValue(data, turnoverKeyId)
	if err != nil {
		return
	}

	valToday, err := parseFloatWithDefaultValue(data, turnoverKeyValToday)
	if err != nil {
		return
	}

	valTodayUsd, err := parseFloatWithDefaultValue(data, turnoverKeyValTodayUsd)
	if err != nil {
		return
	}

	numTrades, err := parseIntWithDefaultValue(data, turnoverKeyNumTrades)
	if err != nil {
		return
	}

	updateTime, err := parseStringWithDefaultValueByKey(data, turnoverKeyUpdateTime, "")
	if err != nil {
		return
	}

	title, err := parseStringWithDefaultValueByKey(data, turnoverKeyTitle, "")
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
