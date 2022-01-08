package moexiss

import (
	"bytes"
	"github.com/buger/jsonparser"
	"log"
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
var parseDurationOrigin = parseDuration
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
	parseDuration = funcCounter
	parseSecurityTypes = funcCounter
	parseSecurityGroups = funcCounter
	parseSecurityCollections = funcCounter
}

func restoreOverriddenFunctions() {
	parseEngines = parseEngOrigin
	parseMarkets = parseMarketsOrigin
	parseBoards = parseBoardsOrigin
	parseBoardGroups = parseBoardGroupsOrigin
	parseDuration = parseDurationOrigin
	parseSecurityTypes = parseSecurityTypesOrigin
	parseSecurityGroups = parseSecurityGroupsOrigin
	parseSecurityCollections = parseSecurityCollectionsOrigin
}

func TestParseIndexResponse(t *testing.T) {

	overrideParseFunctions()
	var incomeJson = `{
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
	err := parseIndexResponse([]byte(incomeJson), index)

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

	var incomeJson = `{
"engines": {},
}
`
	var index *Index = nil
	if got, expected := parseIndexResponse([]byte(incomeJson), index), errNilPointer; got != expected {
		t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseIndexResponseMalformedObjectError(t *testing.T) {

	var incomeJson = `{
"engines": {;,
;
`
	var index = &Index{}
	if got, expected := parseIndexResponse([]byte(incomeJson), index), jsonparser.MalformedObjectError; got != expected {
		t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseIndexResponseError(t *testing.T) {

	var incomeJson = `{
"engines": []}
`
	var index = &Index{}
	if got, expected := parseIndexResponse([]byte(incomeJson), index), errUnexpectedDataType; got != expected {
		t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseIndexResponseNoKeys(t *testing.T) {

	var incomeJson = `{
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
	_ = parseIndexResponse([]byte(incomeJson), index)
	restoreOverriddenFunctions()
	b := buf.Bytes()
	if got := string(b); got != expected {
		t.Fatalf("Error: expecting : \n%v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseIndexResponseUnknownKey(t *testing.T) {
	unknownKey := "unknownKey"
	indexKeys = append(indexKeys, unknownKey)
	var incomeJson = `{
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
	_ = parseIndexResponse([]byte(incomeJson), index)
	restoreOverriddenFunctions()
	indexKeys = indexKeys[0 : len(indexKeys)-1]
	b := buf.Bytes()
	if got := string(b); got != expected {
		t.Fatalf("Error: expecting : \n%v \ngot:\n%v\ninstead", expected, got)
	}
}

func TestIndexParseEngines(t *testing.T) {
	var incomeJson = `{
	"columns": ["id", "name", "title"], 
	"data": [
		[1, "stock", "Фондовый рынок и рынок депозитов"],
		[2, "state", "Рынок ГЦБ (размещение)"],
		[3, "currency", "Валютный рынок"],
		[4, "futures", "Срочный рынок"],
		[5, "commodity", "Товарный рынок"],
		[6, "interventions", "Товарные интервенции"],
		[7, "offboard", "ОТС-система"],
		[9, "agro", "Агро"]
	]
}
	`
	var index = NewIndex()
	err := parseEngines([]byte(incomeJson), index)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v  \ninstead", err)
	}
	if got, expected := len(index.Engines), 8; got != expected {
		t.Fatalf("Error: expecting items: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}
