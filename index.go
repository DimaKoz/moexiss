package moexiss

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"github.com/buger/jsonparser"
	"log"
)

const (
	indexPartsUrl = "index.json?iss.meta=off"

	keyEngines             = "engines"
	keyMarkets             = "markets"
	keyBoards              = "boards"
	keyBoardGroups         = "boardgroups"
	keyDurations           = "durations"
	keySecurityTypes       = "securitytypes"
	keySecurityGroups      = "securitygroups"
	keySecurityCollections = "securitycollections"
	keyData                = "data"
)

var indexKeys = []string{
	keyEngines,
	keyMarkets,
	keyBoards,
	keyBoardGroups,
	keyDurations,
	keySecurityTypes,
	keySecurityGroups,
	keySecurityCollections,
}

var errUnexpectedDataType = errors.New("unexpected data type")
var errNilPointer = errors.New("nil pointer error")

//GeneralFields it contains general fields of some other structures
type GeneralFields struct {
	Id    int64
	Name  string
	Title string
}

//Engine represents a description of the trading system
//in directories named in MoEx ISS API as 'engine'
type Engine struct {
	GeneralFields
}

//Market represents a description of the market and its attributes
type Market struct {
	Engine Engine
	GeneralFields
	MarketPlace string
}

//Board represent a description of the board and its attributes
type Board struct {
	Id           int64
	BoardGroupId int64
	EngineId     int64
	MarketId     int64
	BoardId      string
	BoardTitle   string
	IsTraded     bool
	HasCandles   bool
	IsPrimary    bool
}

//BoardGroup represent a description of the board group and its attributes
type BoardGroup struct {
	Engine     Engine
	MarketId   int64
	MarketName string
	GeneralFields
	IsDefault bool
	IsTraded  bool
}

//Duration represent a description of the duration and its attributes
type Duration struct {
	Interval int64
	Duration int64
	Title    string
	Hint     string
}

//SecurityType represent a description of the security type and its attributes
type SecurityType struct {
	Engine Engine
	GeneralFields
	SecurityGroupName string
}

//SecurityGroup represent a description of the security group and its attributes
type SecurityGroup struct {
	GeneralFields
	IsHidden bool
}

//SecurityCollection represent a description of the security collection and its attributes
type SecurityCollection struct {
	GeneralFields
	SecurityGroupId int64
}

//Index represents a result of the request of the index
type Index struct {
	Engines             []Engine
	Markets             []Market
	Boards              []Board
	BoardGroups         []BoardGroup
	Durations           []Duration
	SecurityTypes       []SecurityType
	SecurityGroups      []SecurityGroup
	SecurityCollections []SecurityCollection
}

func NewIndex() *Index {
	result := &Index{
		Engines:             make([]Engine, 0),
		Markets:             make([]Market, 0),
		Boards:              make([]Board, 0),
		BoardGroups:         make([]BoardGroup, 0),
		Durations:           make([]Duration, 0),
		SecurityTypes:       make([]SecurityType, 0),
		SecurityGroups:      make([]SecurityGroup, 0),
		SecurityCollections: make([]SecurityCollection, 0),
	}
	return result
}

// IndexService provides access to the index related functions
// in the MoEx ISS API.
//
// MoEx ISS API docs: https://iss.moex.com/iss/reference/28
type IndexService service

//TODO
func (s *IndexService) List(ctx context.Context, opt *IndexRequestOptions) (*Index, error) {
	//TODO
	url := s.getUrl(opt)
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

	return nil, nil
}

//getUrl provides an url for a request of the index with parameters from IndexRequestOptions
//opt *IndexRequestOptions can be nil, it is safe
func (s *IndexService) getUrl(opt *IndexRequestOptions) string {
	url, _ := s.client.BaseURL.Parse(indexPartsUrl)
	gotUrl := addIndexRequestOptions(url, opt)
	return gotUrl.String()
}

func parseIndexResponse(byteData []byte, index *Index) (err error) {
	if index == nil {
		err = errNilPointer
		return
	}
	for _, key := range indexKeys {
		bytes, dataType, _, err := jsonparser.Get(byteData, key)
		if err != nil {
			if err != jsonparser.KeyPathNotFoundError {
				return err
			} else {
				log.Println(err.Error(), "for the key:", key)
				continue
			}

		}
		if dataType != jsonparser.Object {
			return errUnexpectedDataType
		}
		var usingFunc = func(byteData []byte, index *Index) (err error) { return }
		switch key {
		case keyEngines:
			usingFunc = parseEngines
		case keyMarkets:
			usingFunc = parseMarkets
		case keyBoards:
			usingFunc = parseBoards
		case keyBoardGroups:
			usingFunc = parseBoardGroups
		case keyDurations:
			usingFunc = parseDuration
		case keySecurityTypes:
			usingFunc = parseSecurityTypes
		case keySecurityGroups:
			usingFunc = parseSecurityGroups
		case keySecurityCollections:
			usingFunc = parseSecurityCollections
		default:
			log.Println("unknown key:", key)
		}
		err = usingFunc(bytes, index)
		if err != nil {
			return err
		}
	}

	return
}

var parseEngines = func(byteData []byte, index *Index) (err error) {

	_, err = jsonparser.ArrayEach(byteData, func(engineItemBytes []byte, dataType jsonparser.ValueType, offset int, errCb error) {
		if errCb != nil {
			err = errCb
			return
		}
		engineItem := &Engine{}
		err = parseEngine(engineItem, engineItemBytes)
		if err != nil {
			return
		}
		index.Engines = append(index.Engines, *engineItem)
	}, keyData)

	return
}

func parseEngine(engine *Engine, engineItemBytes []byte) (err error) {

	counter := 0
	var cb = func(fieldData []byte, dataType jsonparser.ValueType, offset int, errCb error) {
		if errCb != nil {
			err = errCb
			return
		}

		switch counter {

		case 0:
			engine.Id, err = jsonparser.ParseInt(fieldData)

		case 1:
			engine.Name, err = parseStringWithDefaultValue(fieldData)

		case 2:
			engine.Title, err = parseStringWithDefaultValue(fieldData)

		}
		if err != nil {
			return
		}
		counter++
	}

	_, err = jsonparser.ArrayEach(engineItemBytes, cb)

	return

}

var parseMarkets = func(byteData []byte, index *Index) (err error) {
	_, err = jsonparser.ArrayEach(byteData, func(marketItemBytes []byte, dataType jsonparser.ValueType, offset int, errCb error)  {
		if errCb != nil {
			err = errCb
			return
		}
		marketItem := &Market{}
		err = parseMarket(marketItem, marketItemBytes)
		if err != nil {
			return
		}
		index.Markets = append(index.Markets, *marketItem)
	}, keyData)

	return
}

func parseMarket(market *Market, marketItemBytes []byte) (err error) {

	counter := 0
	var cb = func(fieldData []byte, dataType jsonparser.ValueType, offset int, errCb error) {
		if errCb != nil {
			err = errCb
			return
		}

		switch counter {

		case 0:
			market.Id, err = jsonparser.ParseInt(fieldData)

		case 1:
			market.Engine.Id, err = jsonparser.ParseInt(fieldData)

		case 2:
			market.Engine.Name, err = parseStringWithDefaultValue(fieldData)

		case 3:
			market.Engine.Title, err = parseStringWithDefaultValue(fieldData)

		case 4:
			market.Name, err = parseStringWithDefaultValue(fieldData)

		case 5:
			market.Title, err = parseStringWithDefaultValue(fieldData)
		//case 6:
		//already presents in market.Id
		case 7:
			market.MarketPlace, err = parseStringWithDefaultValue(fieldData)
		}
		if err != nil {
			return
		}
		counter++
	}

	_, err = jsonparser.ArrayEach(marketItemBytes, cb)

	return

}

var parseBoards = func(byteData []byte, index *Index) (err error) {
	return nil
}

var parseBoardGroups = func(byteData []byte, index *Index) (err error) {
	return nil
}

var parseDuration = func(byteData []byte, index *Index) (err error) {
	return nil
}

var parseSecurityTypes = func(byteData []byte, index *Index) (err error) {
	return nil
}

var parseSecurityGroups = func(byteData []byte, index *Index) (err error) {
	return nil
}

var parseSecurityCollections = func(byteData []byte, index *Index) (err error) {
	return nil
}
