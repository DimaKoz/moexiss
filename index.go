package moexiss

import (
	"context"
)

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

// IndexService provides access to the index related functions
// in the MoEx ISS API.
//
// MoEx ISS API docs: https://iss.moex.com/iss/reference/28
type IndexService service

//TODO
func (s *IndexService) List(ctx context.Context) (*Index, error) {
	//TODO
	return nil, nil
}
