package moexiss

import (
	"bytes"
	"context"
	"fmt"
	"github.com/buger/jsonparser"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync/atomic"
	"testing"
)

func TestIndexGetUrl(t *testing.T) {
	var income *IndexRequestOptions = nil
	c := NewClient(nil)

	if got, expected := c.Index.getUrl(income), `https://iss.moex.com/iss/index.json?iss.meta=off`; got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}

var funcCallingCounter uint32

var parseEngOrigin = parseEngines
var parseMarketsOrigin = parseMarkets
var parseBoardsOrigin = parseBoards
var parseBoardGroupsOrigin = parseBoardGroups
var parseDurationOrigin = parseDurations
var parseSecurityTypesOrigin = parseSecurityTypes
var parseSecurityGroupsOrigin = parseSecurityGroups
var parseSecurityCollectionsOrigin = parseSecurityCollections

var funcCounter = func(byteData []byte, index *Index) (err error) {
	atomic.AddUint32(&funcCallingCounter, 1)
	return nil
}

func overrideParseFunctions() {
	parseEngines = funcCounter
	parseMarkets = funcCounter
	parseBoards = funcCounter
	parseBoardGroups = funcCounter
	parseDurations = funcCounter
	parseSecurityTypes = funcCounter
	parseSecurityGroups = funcCounter
	parseSecurityCollections = funcCounter
}

func restoreOverriddenFunctions() {
	parseEngines = parseEngOrigin
	parseMarkets = parseMarketsOrigin
	parseBoards = parseBoardsOrigin
	parseBoardGroups = parseBoardGroupsOrigin
	parseDurations = parseDurationOrigin
	parseSecurityTypes = parseSecurityTypesOrigin
	parseSecurityGroups = parseSecurityGroupsOrigin
	parseSecurityCollections = parseSecurityCollectionsOrigin
}

func TestParseIndexResponse(t *testing.T) {

	overrideParseFunctions()
	var incomeJSON = `{
"engines": {},
"markets": {},
"boards": {},
"boardgroups": {},
"durations": {},
"securitytypes": {},
"securitygroups": {},
"securitycollections": {}}
`
	index := &Index{}
	err := parseIndexResponse([]byte(incomeJSON), index)

	//restore functions
	restoreOverriddenFunctions()

	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v  \ninstead", err)
	}
	if got, expected := int(funcCallingCounter), 8; got != expected {
		t.Fatalf("Error: wrong func calls :\n`%d` got `%d` instead", expected, got)
	}
}

func TestParseIndexResponseNilIndex(t *testing.T) {

	var incomeJSON = `{
"engines": {},
}
`
	var index *Index = nil
	if got, expected := parseIndexResponse([]byte(incomeJSON), index), ErrNilPointer; got != expected {
		t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseIndexResponseMalformedObjectError(t *testing.T) {

	var incomeJSON = `{
"engines": {;,
;
`
	var index = &Index{}
	if got, expected := parseIndexResponse([]byte(incomeJSON), index), jsonparser.MalformedObjectError; got != expected {
		t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseIndexResponseError(t *testing.T) {

	var incomeJSON = `{
"engines": []}
`
	var index = &Index{}
	if got, expected := parseIndexResponse([]byte(incomeJSON), index), ErrUnexpectedDataType; got != expected {
		t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseIndexResponseNoKeys(t *testing.T) {

	var incomeJSON = `{
"engines": {}}
`
	expected := `Key path not found for the key: markets
Key path not found for the key: boards
Key path not found for the key: boardgroups
Key path not found for the key: durations
Key path not found for the key: securitytypes
Key path not found for the key: securitygroups
Key path not found for the key: securitycollections
`
	buf := new(bytes.Buffer)
	log.SetFlags(0)
	log.SetOutput(buf)
	var index = &Index{}
	overrideParseFunctions()
	_ = parseIndexResponse([]byte(incomeJSON), index)
	restoreOverriddenFunctions()
	b := buf.Bytes()
	if got := string(b); got != expected {
		t.Fatalf("Error: expecting : \n%v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseIndexResponseUnknownKey(t *testing.T) {
	unknownKey := "unknownKey"
	indexKeys = append(indexKeys, unknownKey)
	var incomeJSON = `{
		"engines": {},
		"markets": {},
		"boards": {},
		"boardgroups": {},
		"durations": {},
		"unknownKey": {},
		"securitytypes": {},
		"securitygroups": {},
		"securitycollections": {}}
	`
	expected := `unknown key: unknownKey
`
	buf := new(bytes.Buffer)
	log.SetFlags(0)
	log.SetOutput(buf)
	var index = &Index{}
	overrideParseFunctions()
	_ = parseIndexResponse([]byte(incomeJSON), index)
	restoreOverriddenFunctions()
	indexKeys = indexKeys[0 : len(indexKeys)-1]
	b := buf.Bytes()
	if got := string(b); got != expected {
		t.Fatalf("Error: expecting : \n%v \ngot:\n%v\ninstead", expected, got)
	}
}

func TestIndexParseEngines(t *testing.T) {
	var incomeJSON = `{
	"columns": ["id", "name", "title"], 
	"data": [
		[1, "stock", "???????????????? ?????????? ?? ?????????? ??????????????????"],
		[2, "state", "?????????? ?????? (????????????????????)"],
		[3, "currency", "???????????????? ??????????"],
		[4, "futures", "?????????????? ??????????"],
		[5, "commodity", "???????????????? ??????????"],
		[6, "interventions", "???????????????? ??????????????????????"],
		[7, "offboard", "??????-??????????????"],
		[9, "agro", "????????"]
	]
}
	`
	var index = newIndex()
	err := parseEngines([]byte(incomeJSON), index)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v  \ninstead", err)
	}
	if got, expected := len(index.Engines), 8; got != expected {
		t.Fatalf("Error: expecting items: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestIndexParseMarkets(t *testing.T) {
	var incomeJSON = `{
	"columns": ["id", "trade_engine_id", "trade_engine_name", "trade_engine_title", "market_name", "market_title", "market_id", "marketplace"], 
	"data": [
		[51, 9, "agro", "????????", "sugar", "?????????? ??????????????", 51, null],
		[5, 1, "stock", "???????????????? ?????????? ?? ?????????? ??????????????????", "index", "?????????????? ?????????????????? ??????????", 5, "INDICES"],
		[21, 3, "currency", "???????????????? ??????????", "basket", "???????????????????? ??????????????", 21, null],
		[12, 4, "futures", "?????????????? ??????????", "main", "?????????????? ??????????????????????", 12, null],
		[23, 1, "stock", "???????????????? ?????????? ?? ?????????? ??????????????????", "standard", "Standard", 23, null],
		[25, 1, "stock", "???????????????? ?????????? ?? ?????????? ??????????????????", "classica", "Classica", 25, null]
	]
}
	`
	var index = newIndex()
	err := parseMarkets([]byte(incomeJSON), index)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v  \ninstead", err)
	}
	if got, expected := len(index.Markets), 6; got != expected {
		t.Fatalf("Error: expecting items: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestIndexParseBoards(t *testing.T) {
	var incomeJSON = `{
	"columns": ["id", "board_group_id", "engine_id", "market_id", "boardid", "board_title", "is_traded", "has_candles", "is_primary"], 
	"data": [
		[177, 57, 1, 1, "TQIF", "??+: ?????? - ????????????????.", 1, 1, 1],
		[244, 72, 1, 33, "MXBD", "MOEX Board", 1, 0, 1]
	]
}
	`
	var index = newIndex()
	err := parseBoards([]byte(incomeJSON), index)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v  \ninstead", err)
	}
	if got, expected := len(index.Boards), 2; got != expected {
		t.Fatalf("Error: expecting items: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestIndexParseBoardsError(t *testing.T) {
	var incomeJSON = `{
	"data": [
		[177;
	]
}
	`
	var index = newIndex()
	if got, expected := parseBoards([]byte(incomeJSON), index), jsonparser.MalformedArrayError; got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestIndexParseEnginesError(t *testing.T) {
	var incomeJSON = `{
		"data": [
		[1;
	]
}
	`
	var index = newIndex()
	if got, expected := parseEngines([]byte(incomeJSON), index), jsonparser.MalformedArrayError; got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestIndexParseMarketsError(t *testing.T) {
	var incomeJSON = `{
	"data": [
		[51;
	]
}
	`
	var index = newIndex()
	if got, expected := parseMarkets([]byte(incomeJSON), index), jsonparser.MalformedArrayError; got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestIndexParseMarketsUnknownValueTypeError(t *testing.T) {
	var incomeJSON = `{
	"data": [
		{}
	]
}
	`
	var index = newIndex()
	if got, expected := parseMarkets([]byte(incomeJSON), index), jsonparser.UnknownValueTypeError; got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestIndexParseEnginesUnknownValueTypeError(t *testing.T) {
	var incomeJSON = `{
	"data": [
		{}
	]
}
	`
	var index = newIndex()
	if got, expected := parseEngines([]byte(incomeJSON), index), jsonparser.UnknownValueTypeError; got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestIndexParseBoardsUnknownValueTypeError(t *testing.T) {
	var incomeJSON = `{
	"data": [
		{}
	]
}
	`
	var index = newIndex()
	if got, expected := parseBoards([]byte(incomeJSON), index), jsonparser.UnknownValueTypeError; got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestIndexParseBoardGroups(t *testing.T) {
	var incomeJSON = `{
	"columns": ["id", "trade_engine_id", "trade_engine_name", "trade_engine_title", "market_id", "market_name", "name", "title", "is_default", "board_group_id", "is_traded"], 
	"data": [
		[9, 1, "stock", "???????????????? ?????????? ?? ?????????? ??????????????????", 5, "index", "stock_index", "??????????????", 1, 9, 1],
		[104, 1, "stock", "???????????????? ?????????? ?? ?????????? ??????????????????", 5, "index", "stock_index_inav", "INAV", 0, 104, 1],
		[15, 4, "futures", "?????????????? ??????????", 12, "main", "futures", "?????????????? ??????????????????????", 1, 15, 0]
	]
}
	`
	var index = newIndex()
	err := parseBoardGroups([]byte(incomeJSON), index)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v  \ninstead", err)
	}
	if got, expected := len(index.BoardGroups), 3; got != expected {
		t.Fatalf("Error: expecting items: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestIndexParseBoardGroupsMalformedArrayError(t *testing.T) {
	var incomeJSON = `{
	"data": [
		[51;
	]
}
	`
	var index = newIndex()
	if got, expected := parseBoardGroups([]byte(incomeJSON), index), jsonparser.MalformedArrayError; got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestIndexParseBoardGroupsUnknownValueTypeError(t *testing.T) {
	var incomeJSON = `{
	"data": [
		{}
	]
}
	`
	var index = newIndex()
	if got, expected := parseBoardGroups([]byte(incomeJSON), index), jsonparser.UnknownValueTypeError; got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestIndexParseDurations(t *testing.T) {
	var incomeJSON = `{
	"columns": ["interval", "duration", "days", "title", "hint"], 
	"data": [
		[1, 60, null, "????????????", "1??"],
		[10, 600, null, "10 ??????????", "10??"],
		[60, 3600, null, "??????", "1??"],
		[24, 86400, null, "????????", "1??"],
		[7, 604800, null, "????????????", "1??"],
		[31, 2678400, null, "??????????", "1??"],
		[4, 8035200, null, "??????????????", "1??"]
	]
}
	`
	var index = newIndex()
	err := parseDurations([]byte(incomeJSON), index)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v  \ninstead", err)
	}
	if got, expected := len(index.Durations), 7; got != expected {
		t.Fatalf("Error: expecting items: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestIndexParseDurationsMalformedArrayError(t *testing.T) {
	var incomeJSON = `{
	"data": [
		[51;
	]
}
	`
	var index = newIndex()
	if got, expected := parseDurations([]byte(incomeJSON), index), jsonparser.MalformedArrayError; got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestIndexParseDurationsUnknownValueTypeError(t *testing.T) {
	var incomeJSON = `{
	"data": [
		{}
	]
}
	`
	var index = newIndex()
	if got, expected := parseDurations([]byte(incomeJSON), index), jsonparser.UnknownValueTypeError; got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestIndexParseSecurityTypes(t *testing.T) {
	var incomeJSON = `{
	"columns": ["id", "trade_engine_id", "trade_engine_name", "trade_engine_title", "security_type_name", "security_type_title", "security_group_name"], 
	"data": [
		[3, 1, "stock", "???????????????? ?????????? ?? ?????????? ??????????????????", "common_share", "?????????? ????????????????????????", "stock_shares"],
		[1, 1, "stock", "???????????????? ?????????? ?? ?????????? ??????????????????", "preferred_share", "?????????? ?????????????????????????????????? ", "stock_shares"]
	]
}
	`
	expectedSt := SecurityType{
		GeneralFields:     GeneralFields{Id: 3, Name: "common_share", Title: "?????????? ????????????????????????"},
		Engine:            Engine{GeneralFields{Id: 1, Name: "stock", Title: "???????????????? ?????????? ?? ?????????? ??????????????????"}},
		SecurityGroupName: "stock_shares",
	}
	var index = newIndex()
	err := parseSecurityTypes([]byte(incomeJSON), index)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v  \ninstead", err)
	}
	if got, expected := len(index.SecurityTypes), 2; got != expected {
		t.Fatalf("Error: expecting items: \n %v \ngot:\n %v \ninstead", expected, got)
	}
	if got, expected := index.SecurityTypes[0], expectedSt; expected != got {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestIndexParseSecurityTypesMalformedArrayError(t *testing.T) {
	var incomeJSON = `{
	"data": [
		[51;
	]
}
	`
	var index = newIndex()
	if got, expected := parseSecurityTypes([]byte(incomeJSON), index), jsonparser.MalformedArrayError; got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestIndexParseSecurityTypesUnknownValueTypeError(t *testing.T) {
	var incomeJSON = `{
	"data": [
		{}
	]
}
	`
	var index = newIndex()
	if got, expected := parseSecurityTypes([]byte(incomeJSON), index), jsonparser.UnknownValueTypeError; got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestIndexParseSecurityGroups(t *testing.T) {
	var incomeJSON = `{
	"columns": ["id", "name", "title", "is_hidden"], 
	"data": [
		[12, "stock_index", "??????????????", 0],
		[4, "stock_shares", "??????????", 0],
		[22, "stock_mortgage", "?????????????????? ????????????????????", 1]
	]
}
	`
	expectedSg := SecurityGroup{
		GeneralFields: GeneralFields{Id: 22, Name: "stock_mortgage", Title: "?????????????????? ????????????????????"},
		IsHidden:      true,
	}
	var index = newIndex()
	err := parseSecurityGroups([]byte(incomeJSON), index)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v  \ninstead", err)
	}
	if got, expected := len(index.SecurityGroups), 3; got != expected {
		t.Fatalf("Error: expecting items: \n %v \ngot:\n %v \ninstead", expected, got)
	}
	if got, expected := index.SecurityGroups[2], expectedSg; expected != got {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestIndexParseSecurityGroupsMalformedArrayError(t *testing.T) {
	var incomeJSON = `{
	"data": [
		[51;
	]
}
	`
	var index = newIndex()
	if got, expected := parseSecurityGroups([]byte(incomeJSON), index), jsonparser.MalformedArrayError; got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestIndexParseSecurityGroupsUnknownValueTypeError(t *testing.T) {
	var incomeJSON = `{
	"data": [
		{}
	]
}
	`
	var index = newIndex()
	if got, expected := parseSecurityGroups([]byte(incomeJSON), index), jsonparser.UnknownValueTypeError; got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestIndexParseSecurityCollections(t *testing.T) {
	var incomeJSON = `{
	"columns": ["id", "name", "title", "security_group_id"], 
	"data": [
		[72, "stock_index_all", "?????? ??????????????", 12],
		[255, "currency_futures_delivery_eur", "EUR\/RUB ?????????????????????? ??????????????", 28]
	]
}
	`
	expectedSc := SecurityCollection{
		GeneralFields:   GeneralFields{Id: 72, Name: "stock_index_all", Title: "?????? ??????????????"},
		SecurityGroupId: 12,
	}
	var index = newIndex()
	err := parseSecurityCollections([]byte(incomeJSON), index)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v  \ninstead", err)
	}
	if got, expected := len(index.SecurityCollections), 2; got != expected {
		t.Fatalf("Error: expecting items: \n %v \ngot:\n %v \ninstead", expected, got)
	}
	if got, expected := index.SecurityCollections[0], expectedSc; expected != got {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestIndexParseSecurityCollectionsMalformedArrayError(t *testing.T) {
	var incomeJSON = `{
	"data": [
		[51;
	]
}
	`
	var index = newIndex()
	if got, expected := parseSecurityCollections([]byte(incomeJSON), index), jsonparser.MalformedArrayError; got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestIndexParseSecurityCollectionsUnknownValueTypeError(t *testing.T) {
	var incomeJSON = `{
	"data": [
		{}
	]
}
	`
	var index = newIndex()
	if got, expected := parseSecurityCollections([]byte(incomeJSON), index), jsonparser.UnknownValueTypeError; got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

//A handler to return expected results
//TestingIndexHandler emulates an external server
func TestingIndexHandler(w http.ResponseWriter, _ *http.Request) {
	log.Println("IndexHandler")
	byteValueResult, err := getTestingData("index.json")
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(byteValueResult)
	if err != nil {
		fmt.Println(err)
	}

}

func TestGetIndexList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(TestingIndexHandler))
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)
	c.BaseURL, _ = url.Parse(srv.URL + "/")
	result, err := c.Index.List(context.Background(), nil)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := len(result.SecurityCollections), 103; got != expected {
		t.Fatalf("Error: expecting: \n %v items\ngot:\n %v items\ninstead", expected, got)
	}
}

func TestGetIndexListBaseUrl(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(TestingIndexHandler))
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)

	c.BaseURL, _ = url.Parse(srv.URL)
	_, err := c.Index.List(context.Background(), nil)
	if got, expected := err, "BaseURL must have a trailing slash, but \""+srv.URL+"\" does not"; got == nil || got.Error() != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestGetIndexListKeyPathNotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("{\"engines\":{}"))
	}))
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)

	c.BaseURL, _ = url.Parse(srv.URL + "/")
	_, err := c.Index.List(context.Background(), nil)
	if got, expected := err, jsonparser.KeyPathNotFoundError; got == nil || got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v \ninstead", expected, got)
	}
}

func TestGetIndexListNilContextError(t *testing.T) {
	c := NewClient(nil)
	var ctx context.Context = nil
	_, err := c.Index.List(ctx, nil)
	if got, expected := err, ErrNonNilContext; got == nil || got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v \ninstead", expected, got)
	}
}
