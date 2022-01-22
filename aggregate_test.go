package moexiss

import (
	"github.com/buger/jsonparser"
	"testing"
)

func TestAggregatesGetUrl(t *testing.T) {
	var income *AggregateRequestOptions = nil
	c := NewClient(nil)

	if got, expected := c.Aggregates.getUrl("sberp", income), `https://iss.moex.com/iss/securities/sberp/aggregates.json?iss.json=extended&iss.meta=off`; got != expected {
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
