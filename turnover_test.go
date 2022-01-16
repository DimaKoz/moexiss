package moexiss

import (
	"github.com/buger/jsonparser"
	"testing"
)

func TestParseTurnover(t *testing.T) {
	expectedStruct := Turnover{
		Name:        "stock",
		Id:          1,
		ValToday:    1988404.90786,
		ValTodayUsd: 26876.4019428,
		NumTrades:   2214956,
		UpdateTime:  "2021-02-24 23:50:29",
		Title:       "Securities Market",
	}
	var incomeJson = `
      {"NAME": "stock", "ID": 1, "VALTODAY": 1988404.90786, "VALTODAY_USD": 26876.4019428, "NUMTRADES": 2214956, "UPDATETIME": "2021-02-24 23:50:29", "TITLE": "Securities Market"}
`
	turnover := Turnover{}
	err := parseTurnover([]byte(incomeJson), &turnover)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := turnover, expectedStruct; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseTurnoverNilId(t *testing.T) {
	expectedStruct := Turnover{
		Name:        "TOTALS",
		Id:          0,
		ValToday:    1988404.90786,
		ValTodayUsd: 26876.4019428,
		NumTrades:   2214956,
		UpdateTime:  "2021-02-24 23:50:29",
		Title:       "Total on Moscow Exchange",
	}
	var incomeJson = `
      {"NAME": "TOTALS", "ID": null, "VALTODAY": 1988404.90786, "VALTODAY_USD": 26876.4019428, "NUMTRADES": 2214956, "UPDATETIME": "2021-02-24 23:50:29", "TITLE": "Total on Moscow Exchange"}
`
	turnover := Turnover{}
	err := parseTurnover([]byte(incomeJson), &turnover)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := turnover, expectedStruct; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseTurnoverCommodity(t *testing.T) {
	expectedStruct := Turnover{
		Name:        "commodity",
		Id:          5,
		ValToday:    0,
		ValTodayUsd: 0,
		NumTrades:   0,
		UpdateTime:  "2021-02-24 09:30:00",
		Title:       "Commodities Market",
	}
	var incomeJson = `
{"NAME": "commodity", "ID": 5, "VALTODAY": null, "VALTODAY_USD": null, "NUMTRADES": null, "UPDATETIME": "2021-02-24 09:30:00", "TITLE": "Commodities Market"}`
	turnover := Turnover{}
	err := parseTurnover([]byte(incomeJson), &turnover)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := turnover, expectedStruct; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseTurnoverErrCases(t *testing.T) {
	type Case struct {
		incomeJson string
		expected error
	}
	cases := []Case{
		// no NAME
		{`{"ID": "5", "VALTODAY": null, "VALTODAY_USD": null, "NUMTRADES": null, "UPDATETIME": "2021-02-24 09:30:00", "TITLE": "Commodities Market"}`,
			jsonparser.KeyPathNotFoundError,
		},
		// no ID
		{`{"NAME": "commodity", "VALTODAY": null, "VALTODAY_USD": null, "NUMTRADES": null, "UPDATETIME": "2021-02-24 09:30:00", "TITLE": "Commodities Market"}`,
			jsonparser.KeyPathNotFoundError,
		},
		// no VALTODAY
		{`{"NAME": "commodity","ID": "5", "VALTODAY_USD": null, "NUMTRADES": null, "UPDATETIME": "2021-02-24 09:30:00", "TITLE": "Commodities Market"}`,
			jsonparser.KeyPathNotFoundError,
		},
		// no VALTODAY_USD
		{`{"NAME": "commodity","ID": "5", "VALTODAY": null,  "NUMTRADES": null, "UPDATETIME": "2021-02-24 09:30:00", "TITLE": "Commodities Market"}`,
			jsonparser.KeyPathNotFoundError,
		},
		// no NUMTRADES
		{`{"NAME": "commodity","ID": "5", "VALTODAY": null, "VALTODAY_USD": null, "UPDATETIME": "2021-02-24 09:30:00", "TITLE": "Commodities Market"}`,
			jsonparser.KeyPathNotFoundError,
		},
		// no UPDATETIME
		{`{"NAME": "commodity","ID": "5", "VALTODAY": null, "VALTODAY_USD": null, "NUMTRADES": null, "TITLE": "Commodities Market"}`,
			jsonparser.KeyPathNotFoundError,
		},
		// no TITLE
		{`{"NAME": "commodity","ID": "5", "VALTODAY": null, "VALTODAY_USD": null, "NUMTRADES": null, "UPDATETIME": "2021-02-24 09:30:00"}`,
			jsonparser.KeyPathNotFoundError,
		},
	}

	for i, c := range cases{
		turnover := Turnover{}
		if got, expected := parseTurnover([]byte(c.incomeJson), &turnover), c.expected; got != expected {
			t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead in %d case", expected, got, i)
		}

	}
}
