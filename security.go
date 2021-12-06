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
	bytesSec, dataType, _, err := jsonparser.Get(b.Bytes(), "securities")
	if err != nil {
		return nil, err
	}
	if dataType != jsonparser.Object {
		return nil, fmt.Errorf("unknown type of 'securities'")
	}

	var errInArr error
	_, err = jsonparser.ArrayEach(bytesSec, func(secItemBytes []byte, dataType jsonparser.ValueType, offset int, err error) {
		if err != nil {
			errInArr = err
			return
		}

		secItem := &Security{}
		counter := 0

		var cb = func(ds []byte, dataType jsonparser.ValueType, offset int, err error) {
			switch counter {

			case 0:
				res, err := jsonparser.ParseInt(ds)
				if err == nil {
					secItem.Id = res
				} else {
					errInArr = err
					break
				}

			case 1:
				res, err := jsonparser.ParseString(ds)
				if err == nil {
					if res != "null" {
						secItem.SecId = res
					}
				} else {
					errInArr = err
					break
				}

			case 2:
				res, err := jsonparser.ParseString(ds)
				if err == nil {
					if res != "null" {
						secItem.ShortName = res
					}
				} else {
					errInArr = err
					break
				}

			case 3:
				res, err := jsonparser.ParseString(ds)
				if err == nil {
					if res != "null" {
						secItem.RegNumber = res
					}
				} else {
					errInArr = err
					break
				}

			case 4:
				res, err := jsonparser.ParseString(ds)
				if err == nil {
					if res != "null" {
						secItem.Name = res
					}
				} else {
					errInArr = err
					break
				}

			case 5:
				res, err := jsonparser.ParseString(ds)
				if err == nil {
					if res != "null" {
						secItem.Isin = res
					}
				} else {
					errInArr = err
					break
				}

			case 6:
				res, err := jsonparser.ParseInt(ds)
				if err == nil {
					if res == 1 {
						secItem.IsTraded = true
					}
				} else {
					errInArr = err
					break
				}

			case 7:
				res, err := jsonparser.ParseString(ds)
				if err == nil {
					if res != "null" {
						secItem.EmitentId = res
					}
				} else {
					errInArr = err
					break
				}

			case 8:
				res, err := jsonparser.ParseString(ds)
				if err == nil {
					if res != "null" {
						secItem.EmitentTitle = res
					}
				} else {
					errInArr = err
					break
				}

			case 9:
				res, err := jsonparser.ParseString(ds)
				if err == nil {
					if res != "null" {
						secItem.EmitentInn = res
					}
				} else {
					errInArr = err
					break
				}

			case 10:
				res, err := jsonparser.ParseString(ds)
				if err == nil {
					if res != "null" {
						secItem.EmitentOkpo = res
					}
				} else {
					errInArr = err
					break
				}

			case 11:
				res, err := jsonparser.ParseString(ds)
				if err == nil {
					if res != "null" {
						secItem.GosReg = res
					}
				} else {
					errInArr = err
					break
				}

			case 12:
				res, err := jsonparser.ParseString(ds)
				if err == nil {
					if res != "null" {
						secItem.Type = res
					}
				} else {
					errInArr = err
					break
				}

			case 13:
				res, err := jsonparser.ParseString(ds)
				if err == nil {
					if res != "null" {
						secItem.Group = res
					}
				} else {
					errInArr = err
					break
				}

			case 14:
				res, err := jsonparser.ParseString(ds)
				if err == nil {
					if res != "null" {
						secItem.PrimaryBoardId = res
					}
				} else {
					errInArr = err
					break
				}

			case 15:
				res, err := jsonparser.ParseString(ds)
				if err == nil {
					if res != "null" {
						secItem.MarketPriceBoardId = res
					}
				} else {
					errInArr = err
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

		securities = append(securities, *secItem)

	}, "data")
	if errInArr != nil {
		return nil, errInArr
	}
	return &securities, nil
}
