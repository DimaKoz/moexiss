package moexiss

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/buger/jsonparser"
)

type Security struct {
	Id                 int64  //"0"
	SecId              string //"1"
	ShortName          string //"2"
	RegNumber          string //"3"
	Name               string //"4"
	Isin               string //"5"
	IsTraded           bool   //"6"
	EmitentId          string //"7"
	EmitentTitle       string //"8"
	EmitentInn         string //"9"
	EmitentOkpo        string //"10"
	GosReg             string //"11"
	Type               string //"12"
	Group              string //"13"
	PrimaryBoardId     string //"14"
	MarketPriceBoardId string //"15"
}

type SecuritiesRequest struct {
	Query string //Argument 'q', minimum 3 symbols
}

// SecuritiesService provides access to the security related functions
// in the MoEx ISS API.
//
// MoEx ISS API docs: https://iss.moex.com/iss/reference/5
type SecuritiesService service

func (s *SecuritiesService) List(ctx context.Context) (*[]Security, error) {
	var u string = "securities.json"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	securities := make([]Security, 0, 100)
	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	_, err = s.client.Do(ctx, req, w)
	if err != nil {
		return nil, err
	}
	err = parseSecurities(&securities, b.Bytes())
	if err != nil {
		return nil, err
	}
	return &securities, nil
}

func parseSecurities(securities *[]Security, byteData []byte) error {
	bytesSec, dataType, _, err := jsonparser.Get(byteData, "securities")
	if err != nil {
		return err
	}
	if dataType != jsonparser.Object {
		return fmt.Errorf("unknown type of 'securities'")
	}

	var errInArr error
	_, err = jsonparser.ArrayEach(bytesSec, func(secItemBytes []byte, dataType jsonparser.ValueType, offset int, err error) {
		if err != nil {
			errInArr = err
			return
		}

		secItem := &Security{}

		errInArr = parseSecurityItem(secItem, secItemBytes)
		if errInArr != nil {
			return
		}

		*securities = append(*securities, *secItem)

	}, "data")
	if errInArr != nil {
		return errInArr
	}
	return nil
}

func parseSecurityItem(s *Security, secItemBytes []byte) (err error) {
	counter := 0
	var errInArr error
	var cb = func(fieldData []byte, dataType jsonparser.ValueType, offset int, err error) {
		switch counter {

		case 0:
			res, err := jsonparser.ParseInt(fieldData)
			if err == nil {
				s.Id = res
			} else {
				errInArr = err
				break
			}

		case 1:
			s.SecId, errInArr = parseStringWithDefaultValue(fieldData)
			if errInArr != nil {
				break
			}

		case 2:
			s.ShortName, errInArr = parseStringWithDefaultValue(fieldData)
			if errInArr != nil {
				break
			}

		case 3:
			s.RegNumber, errInArr = parseStringWithDefaultValue(fieldData)
			if errInArr != nil {
				break
			}

		case 4:
			s.Name, errInArr = parseStringWithDefaultValue(fieldData)
			if errInArr != nil {
				break
			}

		case 5:
			s.Isin, errInArr = parseStringWithDefaultValue(fieldData)
			if errInArr != nil {
				break
			}

		case 6:
			res, err := jsonparser.ParseInt(fieldData)
			if err == nil {
				s.IsTraded = res == 1
			} else {
				errInArr = err
				break
			}

		case 7:
			s.EmitentId, errInArr = parseStringWithDefaultValue(fieldData)
			if errInArr != nil {
				break
			}

		case 8:
			s.EmitentTitle, errInArr = parseStringWithDefaultValue(fieldData)
			if errInArr != nil {
				break
			}

		case 9:
			s.EmitentInn, errInArr = parseStringWithDefaultValue(fieldData)
			if errInArr != nil {
				break
			}

		case 10:
			s.EmitentOkpo, errInArr = parseStringWithDefaultValue(fieldData)
			if errInArr != nil {
				break
			}

		case 11:
			s.GosReg, errInArr = parseStringWithDefaultValue(fieldData)
			if errInArr != nil {
				break
			}

		case 12:
			s.Type, errInArr = parseStringWithDefaultValue(fieldData)
			if errInArr != nil {
				break
			}

		case 13:
			s.Group, errInArr = parseStringWithDefaultValue(fieldData)
			if errInArr != nil {
				break
			}

		case 14:
			s.PrimaryBoardId, errInArr = parseStringWithDefaultValue(fieldData)
			if errInArr != nil {
				break
			}

		case 15:
			s.MarketPriceBoardId, errInArr = parseStringWithDefaultValue(fieldData)
			if errInArr != nil {
				break
			}

		}
		counter++
	}

	_, err = jsonparser.ArrayEach(secItemBytes, cb)
	if errInArr != nil {
		return
	}
	if err != nil {
		errInArr = err
		return
	}

	return nil
}

func parseStringWithDefaultValue(fieldValue []byte) (string, error) {
	res, err := jsonparser.ParseString(fieldValue)
	if err != nil {
		return "", err
	}
	if res != "null" {
		return res, nil
	}
	return "", nil
}
