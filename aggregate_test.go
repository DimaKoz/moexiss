package moexiss

import (
	"context"
	"fmt"
	"github.com/buger/jsonparser"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestParseAggregateResponse(t *testing.T) {
	var incomeJson = `
[
{"charsetinfo": {"name": "utf-8"}},
{
"aggregates": [
{"market_name": "shares", "market_title": "Рынок акций", "engine": "stock", "tradedate": "2022-01-19", "secid": "SBERP", "value": 9833418828.24, "volume": 42115503, "numtrades": 144467, "updated_at": "2022-01-21 09:00:15"},
{"market_name": "moexboard", "market_title": "MOEX Board", "engine": "stock", "tradedate": "2022-01-19", "secid": "SBERP", "value": null, "volume": null, "numtrades": 0, "updated_at": "2022-01-21 09:00:15"}],
"agregates.dates": [
{"from": "2011-11-21", "till": "2022-01-21"}, {"from": "0000-00-00", "till": "0000-00-00"}]}
]
`
	expectedResponse := AggregatesResponse{
		DatesFrom: "2011-11-21",
		DatesTill: "2022-01-21",
		Aggregates: []Aggregate{
			{
				MarketName:   "shares",
				MarketTitle:  "Рынок акций",
				Engine:       "stock",
				TradeDate:    "2022-01-19",
				SecurityId:   "SBERP",
				Value:        9833418828.24,
				Volume:       42115503,
				NumberTrades: 144467,
				UpdatedAt:    "2022-01-21 09:00:15"},
			{
				MarketName:   "moexboard",
				MarketTitle:  "MOEX Board",
				Engine:       "stock",
				TradeDate:    "2022-01-19",
				SecurityId:   "SBERP",
				Value:        0,
				Volume:       0,
				NumberTrades: 0,
				UpdatedAt:    "2022-01-21 09:00:15"},
		},
	}
	aggregatesR := AggregatesResponse{}
	var err error = nil
	if got, expected := parseAggregateResponse([]byte(incomeJson), &aggregatesR), err; got != expected {
		t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead", expected, got)
	}
	if got, expected := len(aggregatesR.Aggregates), len(expectedResponse.Aggregates); got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
	for i, gotItem := range aggregatesR.Aggregates {
		if got, expected := gotItem, expectedResponse.Aggregates[i]; got != expected {
			t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
		}
	}
	if got, expected := aggregatesR.DatesFrom, expectedResponse.DatesFrom; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
	if got, expected := aggregatesR.DatesTill, expectedResponse.DatesTill; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}

}

func TestParseAggregateResponseNilError(t *testing.T) {
	var incomeJson = ``
	var aggregatesResponse *AggregatesResponse = nil

	if got, expected := parseAggregateResponse([]byte(incomeJson), aggregatesResponse), ErrNilPointer; got != expected {
		t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseAggregateResponseError(t *testing.T) {
	var incomeJson = `
[
{"charsetinfo": {"name": "utf-8"}},
{
"aggregates": [
{"market_name1": "shares", "market_title": "Рынок акций", "engine": "stock", "tradedate": "2022-01-19", "secid": "SBERP", "value": 9833418828.24, "volume": 42115503, "numtrades": 144467, "updated_at": "2022-01-21 09:00:15"},
{"market_name": "moexboard", "market_title": "MOEX Board", "engine": "stock", "tradedate": "2022-01-19", "secid": "SBERP", "value": null, "volume": null, "numtrades": 0, "updated_at": "2022-01-21 09:00:15"}],
"agregates.dates": [
{"from": "2011-11-21", "till": "2022-01-21"}]}
]
`
	var aggregatesResponse = &AggregatesResponse{}

	if got, expected := parseAggregateResponse([]byte(incomeJson), aggregatesResponse), jsonparser.KeyPathNotFoundError; got != expected {
		t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}
func TestParseAggregateResponseErrorBadDatesFrom(t *testing.T) {
	var incomeJson = `
[
{"charsetinfo": {"name": "utf-8"}},
{
"aggregates": [
{"market_name": "moexboard", "market_title": "MOEX Board", "engine": "stock", "tradedate": "2022-01-19", "secid": "SBERP", "value": null, "volume": null, "numtrades": 0, "updated_at": "2022-01-21 09:00:15"}],
"agregates.dates": [
{"from1": "2011-11-21", "till": "2022-01-21"}]}
]
`
	var aggregatesResponse = &AggregatesResponse{}

	if got, expected := parseAggregateResponse([]byte(incomeJson), aggregatesResponse), jsonparser.KeyPathNotFoundError; got != expected {
		t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseAggregateResponseDatesError(t *testing.T) {
	var incomeJson = `
[
{"charsetinfo": {"name": "utf-8"}},
{
"aggregates": [
{"market_name": "shares", "market_title": "Рынок акций", "engine": "stock", "tradedate": "2022-01-19", "secid": "SBERP", "value": 9833418828.24, "volume": 42115503, "numtrades": 144467, "updated_at": "2022-01-21 09:00:15"},
{"market_name": "moexboard", "market_title": "MOEX Board", "engine": "stock", "tradedate": "2022-01-19", "secid": "SBERP", "value": null, "volume": null, "numtrades": 0, "updated_at": "2022-01-21 09:00:15"}],
"agregates.dates": [
{"from": "2011-11-21", "till1": "2022-01-21"}]}
]
`
	var aggregatesResponse = &AggregatesResponse{}

	if got, expected := parseAggregateResponse([]byte(incomeJson), aggregatesResponse), jsonparser.KeyPathNotFoundError; got != expected {
		t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseAggregateResponseDatesWrongJsonObject(t *testing.T) {
	var incomeJson = `
[
{"charsetinfo": {"name": "utf-8"}},
{
"aggregates": [
{"market_name": "moexboard", "market_title": "MOEX Board", "engine": "stock", "tradedate": "2022-01-19", "secid": "SBERP", "value": null, "volume": null, "numtrades": 0, "updated_at": "2022-01-21 09:00:15"}],
"agregates.dates": [[
{"from": "2011-11-21", "till1": "2022-01-21"},{"from": "0", "till1": "0"}]]}
]
`
	var aggregatesResponse = &AggregatesResponse{}

	if got, expected := parseAggregateResponse([]byte(incomeJson), aggregatesResponse), ErrUnexpectedDataType; got != expected {
		t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestAggregatesGetUrl(t *testing.T) {
	var income *AggregateRequestOptions = nil
	c := NewClient(nil)
	url, err := c.Aggregates.getUrl("sberp", income)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := url, `https://iss.moex.com/iss/securities/sberp/aggregates.json?iss.json=extended&iss.meta=off`; got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}

func TestParseAggregate(t *testing.T) {
	expectedStruct := Aggregate{
		MarketName:   "shares",
		MarketTitle:  "Рынок акций",
		Engine:       "stock",
		TradeDate:    "2022-01-19",
		SecurityId:   "SBERP",
		Value:        9833418828.24,
		Volume:       42115503,
		NumberTrades: 144467,
		UpdatedAt:    "2022-01-20 09:00:14",
	}
	var incomeJson = `
      {"market_name": "shares", "market_title": "Рынок акций", "engine": "stock", "tradedate": "2022-01-19", "secid": "SBERP", "value": 9833418828.24, "volume": 42115503, "numtrades": 144467, "updated_at": "2022-01-20 09:00:14"}
`
	aggregate := Aggregate{}
	err := parseAggregate([]byte(incomeJson), &aggregate)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := aggregate, expectedStruct; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseAggregateErrCases(t *testing.T) {
	type Case struct {
		incomeJson string
		expected   error
	}
	cases := []Case{
		// no market_name
		{`{"1": "shares", "market_title": "Рынок акций", "engine": "stock", "tradedate": "2022-01-19", "secid": "SBERP", "value": 9833418828.24, "volume": 42115503, "numtrades": 144467, "updated_at": "2022-01-20 09:00:14"}`,
			jsonparser.KeyPathNotFoundError,
		},
		// no market_title
		{`{"market_name": "shares", "1": "Рынок акций", "engine": "stock", "tradedate": "2022-01-19", "secid": "SBERP", "value": 9833418828.24, "volume": 42115503, "numtrades": 144467, "updated_at": "2022-01-20 09:00:14"}`,
			jsonparser.KeyPathNotFoundError,
		},
		// no engine
		{`{"market_name": "shares", "market_title": "Рынок акций", "engine1": "stock", "tradedate": "2022-01-19", "secid": "SBERP", "value": 9833418828.24, "volume": 42115503, "numtrades": 144467, "updated_at": "2022-01-20 09:00:14"}`,
			jsonparser.KeyPathNotFoundError,
		},
		// no tradedate
		{`{"market_name": "shares", "market_title": "Рынок акций", "engine": "stock", "tradedate1": "2022-01-19", "secid": "SBERP", "value": 9833418828.24, "volume": 42115503, "numtrades": 144467, "updated_at": "2022-01-20 09:00:14"}`,
			jsonparser.KeyPathNotFoundError,
		},
		// no secid
		{`{"market_name": "shares", "market_title": "Рынок акций", "engine": "stock", "tradedate": "2022-01-19", "secid1": "SBERP", "value": 9833418828.24, "volume": 42115503, "numtrades": 144467, "updated_at": "2022-01-20 09:00:14"}`,
			jsonparser.KeyPathNotFoundError,
		},
		// no value
		{`{"market_name": "shares", "market_title": "Рынок акций", "engine": "stock", "tradedate": "2022-01-19", "secid": "SBERP", "value1": 9833418828.24, "volume": 42115503, "numtrades": 144467, "updated_at": "2022-01-20 09:00:14"}`,
			jsonparser.KeyPathNotFoundError,
		},
		// no volume
		{`{"market_name": "shares", "market_title": "Рынок акций", "engine": "stock", "tradedate": "2022-01-19", "secid": "SBERP", "value": 9833418828.24, "volume1": 42115503, "numtrades": 144467, "updated_at": "2022-01-20 09:00:14"}`,
			jsonparser.KeyPathNotFoundError,
		},
		// no numtrades
		{`{"market_name": "shares", "market_title": "Рынок акций", "engine": "stock", "tradedate": "2022-01-19", "secid": "SBERP", "value": 9833418828.24, "volume": 42115503, "numtrades1": 144467, "updated_at": "2022-01-20 09:00:14"}`,
			jsonparser.KeyPathNotFoundError,
		},
		// no updated_at
		{`{"market_name": "shares", "market_title": "Рынок акций", "engine": "stock", "tradedate": "2022-01-19", "secid": "SBERP", "value": 9833418828.24, "volume": 42115503, "numtrades": 144467, "updated_at1": "2022-01-20 09:00:14"}`,
			jsonparser.KeyPathNotFoundError,
		},
	}

	for i, c := range cases {
		aggregate := Aggregate{}
		if got, expected := parseAggregate([]byte(c.incomeJson), &aggregate), c.expected; got != expected {
			t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead in %d case", expected, got, i)
		}

	}
}

func TestParseAggregates(t *testing.T) {

	var incomeJson = `
[
      {"market_name": "shares", "market_title": "Рынок акций", "engine": "stock", "tradedate": "2022-01-19", "secid": "SBERP", "value": 9833418828.24, "volume": 42115503, "numtrades": 144467, "updated_at": "2022-01-20 09:00:14"},
      {"market_name": "ndm", "market_title": "Режим переговорных сделок", "engine": "stock", "tradedate": "2022-01-19", "secid": "SBERP", "value": 179995527.30, "volume": 751890, "numtrades": 3, "updated_at": "2022-01-20 09:00:14"},
      {"market_name": "otc", "market_title": "ОТС", "engine": "stock", "tradedate": "2022-01-19", "secid": "SBERP", "value": 20456116.30, "volume": 87020, "numtrades": 2, "updated_at": "2022-01-20 09:00:14"},
      {"market_name": "repo", "market_title": "Рынок сделок РЕПО", "engine": "stock", "tradedate": "2022-01-19", "secid": "SBERP", "value": 9397389429.44, "volume": 46971852, "numtrades": 3320, "updated_at": "2022-01-20 09:00:14"},
      {"market_name": "moexboard", "market_title": "MOEX Board", "engine": "stock", "tradedate": "2022-01-19", "secid": "SBERP", "value": null, "volume": null, "numtrades": 0, "updated_at": "2022-01-20 09:00:14"}
]
`
	aggregates := make([]Aggregate, 0)
	err := parseAggregates([]byte(incomeJson), &aggregates)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := len(aggregates), 5; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseAggregatesUnexpectedDataTypeError(t *testing.T) {

	var incomeJson = `
[
      []
]`
	aggregates := make([]Aggregate, 0)
	if got, expected := parseAggregates([]byte(incomeJson), &aggregates), ErrUnexpectedDataType; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseAggregatesError(t *testing.T) {

	var incomeJson = `
[
      {"market_name1": "repo", "market_title": "Рынок сделок РЕПО", "engine": "stock", "tradedate": "2022-01-19", "secid": "SBERP", "value": 9397389429.44, "volume": 46971852, "numtrades": 3320, "updated_at": "2022-01-20 09:00:14"}
]`
	aggregates := make([]Aggregate, 0)
	if got, expected := parseAggregates([]byte(incomeJson), &aggregates), jsonparser.KeyPathNotFoundError; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestAggregatesNilContextError(t *testing.T) {
	c := NewClient(nil)
	var ctx context.Context = nil
	_, err := c.Aggregates.Aggregates(ctx, "SBERP", nil)
	if got, expected := err, ErrNonNilContext; got == nil || got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v \ninstead", expected, got)
	}
}

func TestAggregatesKeyPathNotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		str := `[
  {}]
`
		_, _ = w.Write([]byte(str))
	}))
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)

	c.BaseURL, _ = url.Parse(srv.URL + "/")
	_, err := c.Aggregates.Aggregates(context.Background(), "jhgsd", nil)
	if got, expected := err, jsonparser.KeyPathNotFoundError; got == nil || got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v \ninstead", expected, got)
	}
}

func TestAggregateService_AggregatesBadUrl(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {}))
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)

	c.BaseURL, _ = url.Parse(srv.URL)
	_, err := c.Aggregates.Aggregates(context.Background(), "sber", nil)
	if got, expected := err, "BaseURL must have a trailing slash, but \""+srv.URL+"\" does not"; got == nil || got.Error() != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestAggregateService_AggregatesBadSecurityParam(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {}))
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)

	c.BaseURL, _ = url.Parse(srv.URL)
	_, err := c.Aggregates.Aggregates(context.Background(), "", nil)
	if got, expected := err, ErrBadSecurityParameter; got == nil || got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

//A handler to return expected results
//TestingAggregatesHandler emulates an external server
func TestingAggregatesHandler(w http.ResponseWriter, _ *http.Request) {

	byteValueResult, err := getTestingData("aggregates.json")
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

func TestAggregateService_Aggregates(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(TestingAggregatesHandler))
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)
	c.BaseURL, _ = url.Parse(srv.URL + "/")
	result, err := c.Aggregates.Aggregates(context.Background(), "SBERP", nil)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := len(result.Aggregates), 5; got != expected {
		t.Fatalf("Error: expecting: \n %v items\ngot:\n %v items\ninstead", expected, got)
	}
}
