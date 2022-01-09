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
			usingFunc = parseDurations
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
	var errInCb error
	_, err = jsonparser.ArrayEach(byteData, func(engineItemBytes []byte, dataType jsonparser.ValueType, offset int, errCb error) {
		if errCb != nil {
			errInCb = errCb
			return
		}
		engineItem := &Engine{}
		errInCb = parseEngine(engineItem, engineItemBytes)
		if errInCb != nil {
			return
		}
		index.Engines = append(index.Engines, *engineItem)
	}, keyData)
	if err == nil && errInCb != nil {
		return errInCb
	}
	return
}

func parseEngine(engine *Engine, engineItemBytes []byte) (err error) {
	var errInCb error
	counter := 0
	var cb = func(fieldData []byte, dataType jsonparser.ValueType, offset int, errCb error) {
		if errCb != nil {
			errInCb = errCb
			return
		}

		switch counter {

		case 0:
			engine.Id, errInCb = jsonparser.ParseInt(fieldData)

		case 1:
			engine.Name, errInCb = parseStringWithDefaultValue(fieldData)

		case 2:
			engine.Title, errInCb = parseStringWithDefaultValue(fieldData)

		}
		if errInCb != nil {
			return
		}
		counter++
	}

	_, err = jsonparser.ArrayEach(engineItemBytes, cb)
	if err == nil && errInCb != nil {
		return errInCb
	}
	return

}

var parseMarkets = func(byteData []byte, index *Index) (err error) {
	var errInCb error
	_, err = jsonparser.ArrayEach(byteData, func(marketItemBytes []byte, dataType jsonparser.ValueType, offset int, errCb error) {
		if errCb != nil {
			errInCb = errCb
			return
		}
		marketItem := &Market{}
		errInCb = parseMarket(marketItem, marketItemBytes)
		if errInCb != nil {
			return
		}
		index.Markets = append(index.Markets, *marketItem)
	}, keyData)
	if err == nil && errInCb != nil {
		return errInCb
	}
	return
}

func parseMarket(market *Market, marketItemBytes []byte) (err error) {
	var errInCb error
	counter := 0
	var cb = func(fieldData []byte, dataType jsonparser.ValueType, offset int, errCb error) {
		if errCb != nil {
			errInCb = errCb
			return
		}

		switch counter {

		case 0:
			market.Id, errInCb = jsonparser.ParseInt(fieldData)

		case 1:
			market.Engine.Id, errInCb = jsonparser.ParseInt(fieldData)

		case 2:
			market.Engine.Name, errInCb = parseStringWithDefaultValue(fieldData)

		case 3:
			market.Engine.Title, errInCb = parseStringWithDefaultValue(fieldData)

		case 4:
			market.Name, errInCb = parseStringWithDefaultValue(fieldData)

		case 5:
			market.Title, errInCb = parseStringWithDefaultValue(fieldData)
		//case 6:
		//already presents in market.Id
		case 7:
			market.MarketPlace, errInCb = parseStringWithDefaultValue(fieldData)
		}
		if errInCb != nil {
			return
		}
		counter++
	}

	_, err = jsonparser.ArrayEach(marketItemBytes, cb)
	if err == nil && errInCb != nil {
		return errInCb
	}
	return

}

var parseBoards = func(byteData []byte, index *Index) (err error) {
	var errInCb error
	_, err = jsonparser.ArrayEach(byteData, func(boardItemBytes []byte, dataType jsonparser.ValueType, offset int, errCb error) {
		if errCb != nil {
			errInCb = errCb
			return
		}
		boardItem := &Board{}
		errInCb = parseBoard(boardItem, boardItemBytes)
		if errInCb != nil {
			return
		}
		index.Boards = append(index.Boards, *boardItem)
	}, keyData)
	if err == nil && errInCb != nil {
		err = errInCb
	}
	return
}

func parseBoard(board *Board, boardItemBytes []byte) (err error) {
	var errInCb error
	counter := 0
	var cb = func(fieldData []byte, dataType jsonparser.ValueType, offset int, errCb error) {
		if errCb != nil {
			errInCb = errCb
			return
		}

		switch counter {

		case 0:
			board.Id, errInCb = jsonparser.ParseInt(fieldData)

		case 1:
			board.BoardGroupId, errInCb = jsonparser.ParseInt(fieldData)

		case 2:
			board.EngineId, errInCb = jsonparser.ParseInt(fieldData)

		case 3:
			board.MarketId, errInCb = jsonparser.ParseInt(fieldData)

		case 4:
			board.BoardId, errInCb = parseStringWithDefaultValue(fieldData)

		case 5:
			board.BoardTitle, errInCb = parseStringWithDefaultValue(fieldData)

		case 6:
			board.IsTraded = string(fieldData) == "1"

		case 7:
			board.HasCandles = string(fieldData) == "1"

		case 8:
			board.IsPrimary = string(fieldData) == "1"

		}
		if errInCb != nil {
			return
		}
		counter++
	}

	_, err = jsonparser.ArrayEach(boardItemBytes, cb)
	if err == nil && errInCb != nil {
		return errInCb
	}
	return

}

var parseBoardGroups = func(byteData []byte, index *Index) (err error) {
	var errInCb error
	_, err = jsonparser.ArrayEach(byteData, func(bgItemData []byte, dataType jsonparser.ValueType, offset int, errCb error) {
		if errCb != nil {
			errInCb = errCb
			return
		}
		boardGroupItem := &BoardGroup{}
		errInCb = parseBoardGroup(boardGroupItem, bgItemData)
		if errInCb != nil {
			return
		}
		index.BoardGroups = append(index.BoardGroups, *boardGroupItem)
	}, keyData)
	if err == nil && errInCb != nil {
		err = errInCb
	}
	return

}

func parseBoardGroup(bg *BoardGroup, boardGroupItemData []byte) (err error) {
	var errInCb error
	counter := 0
	var cb = func(fieldData []byte, dataType jsonparser.ValueType, offset int, errCb error) {
		if errCb != nil {
			errInCb = errCb
			return
		}

		switch counter {

		case 0:
			bg.Id, errInCb = jsonparser.ParseInt(fieldData)

		case 1:
			bg.Engine.Id, errInCb = jsonparser.ParseInt(fieldData)

		case 2:
			bg.Engine.Name, errInCb = parseStringWithDefaultValue(fieldData)

		case 3:
			bg.Engine.Title, errInCb = parseStringWithDefaultValue(fieldData)

		case 4:
			bg.MarketId, errInCb = jsonparser.ParseInt(fieldData)

		case 5:
			bg.MarketName, errInCb = parseStringWithDefaultValue(fieldData)

		case 6:
			bg.Name, errInCb = parseStringWithDefaultValue(fieldData)

		case 7:
			bg.Title, errInCb = parseStringWithDefaultValue(fieldData)

		case 8:
			bg.IsDefault = string(fieldData) == "1"

		//case 9: Nothing to do
		case 10:
			bg.IsTraded = string(fieldData) == "1"

		}
		if errInCb != nil {
			return
		}
		counter++
	}

	_, err = jsonparser.ArrayEach(boardGroupItemData, cb)
	if err == nil && errInCb != nil {
		return errInCb
	}
	return

}

var parseDurations = func(byteData []byte, index *Index) (err error) {
	var errInCb error
	_, err = jsonparser.ArrayEach(byteData, func(durationData []byte, dataType jsonparser.ValueType, offset int, errCb error) {
		if errCb != nil {
			errInCb = errCb
			return
		}
		durationItem := &Duration{}
		errInCb = parseDuration(durationItem, durationData)
		if errInCb != nil {
			return
		}
		index.Durations = append(index.Durations, *durationItem)
	}, keyData)
	if err == nil && errInCb != nil {
		err = errInCb
	}
	return

}

func parseDuration(d *Duration, durationData []byte) (err error) {
	var errInCb error
	counter := 0
	var cb = func(fieldData []byte, dataType jsonparser.ValueType, offset int, errCb error) {
		if errCb != nil {
			errInCb = errCb
			return
		}

		switch counter {

		case 0:
			d.Interval, errInCb = jsonparser.ParseInt(fieldData)

		case 1:
			d.Duration, errInCb = jsonparser.ParseInt(fieldData)

		//case 2: Do nothing for 2 value
		case 3:
			d.Title, errInCb = parseStringWithDefaultValue(fieldData)

		case 4:
			d.Hint, errInCb = parseStringWithDefaultValue(fieldData)

		}
		if errInCb != nil {
			return
		}
		counter++
	}

	_, err = jsonparser.ArrayEach(durationData, cb)
	if err == nil && errInCb != nil {
		return errInCb
	}
	return
}

var parseSecurityTypes = func(byteData []byte, index *Index) (err error) {
	var errInCb error
	_, err = jsonparser.ArrayEach(byteData, func(stData []byte, dataType jsonparser.ValueType, offset int, errCb error) {
		if errCb != nil {
			errInCb = errCb
			return
		}
		securityTypeItem := &SecurityType{}
		errInCb = parseSecurityType(securityTypeItem, stData)
		if errInCb != nil {
			return
		}
		index.SecurityTypes = append(index.SecurityTypes, *securityTypeItem)
	}, keyData)
	if err == nil && errInCb != nil {
		err = errInCb
	}
	return
}

func parseSecurityType(st *SecurityType, stData []byte) (err error) {
	var errInCb error
	counter := 0
	var cb = func(fieldData []byte, dataType jsonparser.ValueType, offset int, errCb error) {
		if errCb != nil {
			errInCb = errCb
			return
		}

		switch counter {

		case 0:
			st.Id, errInCb = jsonparser.ParseInt(fieldData)

		case 1:
			st.Engine.Id, errInCb = jsonparser.ParseInt(fieldData)

		case 2:
			st.Engine.Name, errInCb = parseStringWithDefaultValue(fieldData)

		case 3:
			st.Engine.Title, errInCb = parseStringWithDefaultValue(fieldData)

		case 4:
			st.Name, errInCb = parseStringWithDefaultValue(fieldData)

		case 5:
			st.Title, errInCb = parseStringWithDefaultValue(fieldData)

		case 6:
			st.SecurityGroupName, errInCb = parseStringWithDefaultValue(fieldData)

		}
		if errInCb != nil {
			return
		}
		counter++
	}

	_, err = jsonparser.ArrayEach(stData, cb)
	if err == nil && errInCb != nil {
		return errInCb
	}
	return
}

var parseSecurityGroups = func(byteData []byte, index *Index) (err error) {
	var errInCb error
	_, err = jsonparser.ArrayEach(byteData, func(sgData []byte, dataType jsonparser.ValueType, offset int, errCb error) {
		if errCb != nil {
			errInCb = errCb
			return
		}
		securityGroupItem := &SecurityGroup{}
		errInCb = parseSecurityGroup(securityGroupItem, sgData)
		if errInCb != nil {
			return
		}
		index.SecurityGroups = append(index.SecurityGroups, *securityGroupItem)
	}, keyData)
	if err == nil && errInCb != nil {
		err = errInCb
	}
	return
}

func parseSecurityGroup(sg *SecurityGroup, sgData []byte) (err error) {
	var errInCb error
	counter := 0
	var cb = func(fieldData []byte, dataType jsonparser.ValueType, offset int, errCb error) {
		if errCb != nil {
			errInCb = errCb
			return
		}

		switch counter {

		case 0:
			sg.Id, errInCb = jsonparser.ParseInt(fieldData)

		case 1:
			sg.Name, errInCb = parseStringWithDefaultValue(fieldData)

		case 2:
			sg.Title, errInCb = parseStringWithDefaultValue(fieldData)

		case 3:
			sg.IsHidden = string(fieldData) == "1"

		}
		if errInCb != nil {
			return
		}
		counter++
	}

	_, err = jsonparser.ArrayEach(sgData, cb)
	if err == nil && errInCb != nil {
		return errInCb
	}
	return
}

var parseSecurityCollections = func(byteData []byte, index *Index) (err error) {
	return nil
}
