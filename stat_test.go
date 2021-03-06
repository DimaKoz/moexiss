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

func TestStatGetUrl(t *testing.T) {
	c := NewClient(nil)
	gotURL, err := c.Stats.getUrl(EngineStock, "shares", nil)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := gotURL, `https://iss.moex.com/iss/engines/stock/markets/shares/secstats.json?iss.json=extended&iss.meta=off`; got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}

func TestStatGetUrlBadEngine(t *testing.T) {
	c := NewClient(nil)
	_, err := c.Stats.getUrl(EngineUndefined, "shares", nil)
	if got, expected := err, ErrBadEngineParameter; got != expected {
		t.Fatalf("Error: expecting %v error: \ngot %v \ninstead", expected, got)
	}
}

func TestStatGetUrlBadMarket(t *testing.T) {
	c := NewClient(nil)
	_, err := c.Stats.getUrl(EngineStock, "", nil)
	if got, expected := err, ErrBadMarketParameter; got != expected {
		t.Fatalf("Error: expecting %v error: \ngot %v \ninstead", expected, got)
	}
}

func TestParseSecStatItem(t *testing.T) {
	expectedStruct := SecStat{
		Ticker:           "GAZP",
		BoardId:          "TQBR",
		TrSession:        TradingSessionUndefined,
		Time:             "09:49:58",
		PriceMinusPrevPr: -22.98,
		VolToday:         47948300,
		ValToday:         12677905337,
		HighBid:          304.75,
		LowOffer:         250.92,
		LastOffer:        260.29,
		LastBid:          259.71,
		Open:             253.95,
		Low:              250.92,
		High:             273.99,
		Last:             260.29,
		LClosePrice:      0,
		NumTrades:        107517,
		WaPrice:          264.41,
		AdmittedQuote:    0,
		MarketPrice:      0,
		LCurrentPrice:    260.51,
		ClosingAucPrice:  0,
	}
	var incomeJSON = `
      {"SECID": "GAZP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:58", "PRICEMINUSPREVWAPRICE": -22.98, "VOLTODAY": 47948300, "VALTODAY": 12677905337, "HIGHBID": 304.75, "LOWOFFER": 250.92, "LASTOFFER": 260.29, "LASTBID": 259.71, "OPEN": 253.95, "LOW": 250.92, "HIGH": 273.99, "LAST": 260.29, "LCLOSEPRICE": null, "NUMTRADES": 107517, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 260.51, "CLOSINGAUCTIONPRICE": null}
`
	secStatItem := SecStat{}
	err := parseSecStatItem([]byte(incomeJSON), &secStatItem)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := secStatItem, expectedStruct; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseSecStatItemErrCases(t *testing.T) {
	type Case struct {
		incomeJSON string
	}
	cases := []Case{
		// no SECID
		{`{"SECID1": "GAZP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:58", "PRICEMINUSPREVWAPRICE": -22.98, "VOLTODAY": 47948300, "VALTODAY": 12677905337, "HIGHBID": 304.75, "LOWOFFER": 250.92, "LASTOFFER": 260.29, "LASTBID": 259.71, "OPEN": 253.95, "LOW": 250.92, "HIGH": 273.99, "LAST": 260.29, "LCLOSEPRICE": null, "NUMTRADES": 107517, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 260.51, "CLOSINGAUCTIONPRICE": null}`},
		// no BOARDID
		{`{"SECID": "GAZP", "BOARDID1": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:58", "PRICEMINUSPREVWAPRICE": -22.98, "VOLTODAY": 47948300, "VALTODAY": 12677905337, "HIGHBID": 304.75, "LOWOFFER": 250.92, "LASTOFFER": 260.29, "LASTBID": 259.71, "OPEN": 253.95, "LOW": 250.92, "HIGH": 273.99, "LAST": 260.29, "LCLOSEPRICE": null, "NUMTRADES": 107517, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 260.51, "CLOSINGAUCTIONPRICE": null}`},
		// no TRADINGSESSION
		{`{"SECID": "GAZP", "BOARDID": "TQBR", "TRADINGSESSION1": "0", "TIME": "09:49:58", "PRICEMINUSPREVWAPRICE": -22.98, "VOLTODAY": 47948300, "VALTODAY": 12677905337, "HIGHBID": 304.75, "LOWOFFER": 250.92, "LASTOFFER": 260.29, "LASTBID": 259.71, "OPEN": 253.95, "LOW": 250.92, "HIGH": 273.99, "LAST": 260.29, "LCLOSEPRICE": null, "NUMTRADES": 107517, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 260.51, "CLOSINGAUCTIONPRICE": null}`},
		// no TIME
		{`{"SECID": "GAZP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME1": "09:49:58", "PRICEMINUSPREVWAPRICE": -22.98, "VOLTODAY": 47948300, "VALTODAY": 12677905337, "HIGHBID": 304.75, "LOWOFFER": 250.92, "LASTOFFER": 260.29, "LASTBID": 259.71, "OPEN": 253.95, "LOW": 250.92, "HIGH": 273.99, "LAST": 260.29, "LCLOSEPRICE": null, "NUMTRADES": 107517, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 260.51, "CLOSINGAUCTIONPRICE": null}`},
		// no PRICEMINUSPREVWAPRICE
		{`{"SECID": "GAZP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:58", "PRICEMINUSPREVWAPRICE1": -22.98, "VOLTODAY": 47948300, "VALTODAY": 12677905337, "HIGHBID": 304.75, "LOWOFFER": 250.92, "LASTOFFER": 260.29, "LASTBID": 259.71, "OPEN": 253.95, "LOW": 250.92, "HIGH": 273.99, "LAST": 260.29, "LCLOSEPRICE": null, "NUMTRADES": 107517, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 260.51, "CLOSINGAUCTIONPRICE": null}`},
		// no VOLTODAY
		{`{"SECID": "GAZP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:58", "PRICEMINUSPREVWAPRICE": -22.98, "VOLTODAY1": 47948300, "VALTODAY": 12677905337, "HIGHBID": 304.75, "LOWOFFER": 250.92, "LASTOFFER": 260.29, "LASTBID": 259.71, "OPEN": 253.95, "LOW": 250.92, "HIGH": 273.99, "LAST": 260.29, "LCLOSEPRICE": null, "NUMTRADES": 107517, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 260.51, "CLOSINGAUCTIONPRICE": null}`},
		// no VALTODAY
		{`{"SECID": "GAZP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:58", "PRICEMINUSPREVWAPRICE": -22.98, "VOLTODAY": 47948300, "VALTODAY1": 12677905337, "HIGHBID": 304.75, "LOWOFFER": 250.92, "LASTOFFER": 260.29, "LASTBID": 259.71, "OPEN": 253.95, "LOW": 250.92, "HIGH": 273.99, "LAST": 260.29, "LCLOSEPRICE": null, "NUMTRADES": 107517, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 260.51, "CLOSINGAUCTIONPRICE": null}`},
		// no HIGHBID
		{`{"SECID": "GAZP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:58", "PRICEMINUSPREVWAPRICE": -22.98, "VOLTODAY": 47948300, "VALTODAY": 12677905337, "HIGHBID1": 304.75, "LOWOFFER": 250.92, "LASTOFFER": 260.29, "LASTBID": 259.71, "OPEN": 253.95, "LOW": 250.92, "HIGH": 273.99, "LAST": 260.29, "LCLOSEPRICE": null, "NUMTRADES": 107517, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 260.51, "CLOSINGAUCTIONPRICE": null}`},
		// no LOWOFFER
		{`{"SECID": "GAZP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:58", "PRICEMINUSPREVWAPRICE": -22.98, "VOLTODAY": 47948300, "VALTODAY": 12677905337, "HIGHBID": 304.75, "LOWOFFER1": 250.92, "LASTOFFER": 260.29, "LASTBID": 259.71, "OPEN": 253.95, "LOW": 250.92, "HIGH": 273.99, "LAST": 260.29, "LCLOSEPRICE": null, "NUMTRADES": 107517, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 260.51, "CLOSINGAUCTIONPRICE": null}`},
		// no LASTOFFER
		{`{"SECID": "GAZP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:58", "PRICEMINUSPREVWAPRICE": -22.98, "VOLTODAY": 47948300, "VALTODAY": 12677905337, "HIGHBID": 304.75, "LOWOFFER": 250.92, "LASTOFFER1": 260.29, "LASTBID": 259.71, "OPEN": 253.95, "LOW": 250.92, "HIGH": 273.99, "LAST": 260.29, "LCLOSEPRICE": null, "NUMTRADES": 107517, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 260.51, "CLOSINGAUCTIONPRICE": null}`},
		// no LASTBID
		{`{"SECID": "GAZP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:58", "PRICEMINUSPREVWAPRICE": -22.98, "VOLTODAY": 47948300, "VALTODAY": 12677905337, "HIGHBID": 304.75, "LOWOFFER": 250.92, "LASTOFFER": 260.29, "LASTBID1": 259.71, "OPEN": 253.95, "LOW": 250.92, "HIGH": 273.99, "LAST": 260.29, "LCLOSEPRICE": null, "NUMTRADES": 107517, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 260.51, "CLOSINGAUCTIONPRICE": null}`},
		// no OPEN
		{`{"SECID": "GAZP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:58", "PRICEMINUSPREVWAPRICE": -22.98, "VOLTODAY": 47948300, "VALTODAY": 12677905337, "HIGHBID": 304.75, "LOWOFFER": 250.92, "LASTOFFER": 260.29, "LASTBID": 259.71, "OPEN1": 253.95, "LOW": 250.92, "HIGH": 273.99, "LAST": 260.29, "LCLOSEPRICE": null, "NUMTRADES": 107517, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 260.51, "CLOSINGAUCTIONPRICE": null}`},
		// no LOW
		{`{"SECID": "GAZP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:58", "PRICEMINUSPREVWAPRICE": -22.98, "VOLTODAY": 47948300, "VALTODAY": 12677905337, "HIGHBID": 304.75, "LOWOFFER": 250.92, "LASTOFFER": 260.29, "LASTBID": 259.71, "OPEN": 253.95, "LOW1": 250.92, "HIGH": 273.99, "LAST": 260.29, "LCLOSEPRICE": null, "NUMTRADES": 107517, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 260.51, "CLOSINGAUCTIONPRICE": null}`},
		// no HIGH
		{`{"SECID": "GAZP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:58", "PRICEMINUSPREVWAPRICE": -22.98, "VOLTODAY": 47948300, "VALTODAY": 12677905337, "HIGHBID": 304.75, "LOWOFFER": 250.92, "LASTOFFER": 260.29, "LASTBID": 259.71, "OPEN": 253.95, "LOW": 250.92, "HIGH1": 273.99, "LAST": 260.29, "LCLOSEPRICE": null, "NUMTRADES": 107517, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 260.51, "CLOSINGAUCTIONPRICE": null}`},
		// no LAST
		{`{"SECID": "GAZP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:58", "PRICEMINUSPREVWAPRICE": -22.98, "VOLTODAY": 47948300, "VALTODAY": 12677905337, "HIGHBID": 304.75, "LOWOFFER": 250.92, "LASTOFFER": 260.29, "LASTBID": 259.71, "OPEN": 253.95, "LOW": 250.92, "HIGH": 273.99, "LAST1": 260.29, "LCLOSEPRICE": null, "NUMTRADES": 107517, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 260.51, "CLOSINGAUCTIONPRICE": null}`},
		// no LCLOSEPRICE
		{`{"SECID": "GAZP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:58", "PRICEMINUSPREVWAPRICE": -22.98, "VOLTODAY": 47948300, "VALTODAY": 12677905337, "HIGHBID": 304.75, "LOWOFFER": 250.92, "LASTOFFER": 260.29, "LASTBID": 259.71, "OPEN": 253.95, "LOW": 250.92, "HIGH": 273.99, "LAST": 260.29, "LCLOSEPRICE1": null, "NUMTRADES": 107517, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 260.51, "CLOSINGAUCTIONPRICE": null}`},
		// no NUMTRADES
		{`{"SECID": "GAZP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:58", "PRICEMINUSPREVWAPRICE": -22.98, "VOLTODAY": 47948300, "VALTODAY": 12677905337, "HIGHBID": 304.75, "LOWOFFER": 250.92, "LASTOFFER": 260.29, "LASTBID": 259.71, "OPEN": 253.95, "LOW": 250.92, "HIGH": 273.99, "LAST": 260.29, "LCLOSEPRICE": null, "NUMTRADES1": 107517, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 260.51, "CLOSINGAUCTIONPRICE": null}`},
		// no WAPRICE
		{`{"SECID": "GAZP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:58", "PRICEMINUSPREVWAPRICE": -22.98, "VOLTODAY": 47948300, "VALTODAY": 12677905337, "HIGHBID": 304.75, "LOWOFFER": 250.92, "LASTOFFER": 260.29, "LASTBID": 259.71, "OPEN": 253.95, "LOW": 250.92, "HIGH": 273.99, "LAST": 260.29, "LCLOSEPRICE": null, "NUMTRADES": 107517, "WAPRICE1": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 260.51, "CLOSINGAUCTIONPRICE": null}`},
		// no ADMITTEDQUOTE
		{`{"SECID": "GAZP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:58", "PRICEMINUSPREVWAPRICE": -22.98, "VOLTODAY": 47948300, "VALTODAY": 12677905337, "HIGHBID": 304.75, "LOWOFFER": 250.92, "LASTOFFER": 260.29, "LASTBID": 259.71, "OPEN": 253.95, "LOW": 250.92, "HIGH": 273.99, "LAST": 260.29, "LCLOSEPRICE": null, "NUMTRADES": 107517, "WAPRICE": 264.41, "ADMITTEDQUOTE1": null, "MARKETPRICE2": null, "LCURRENTPRICE": 260.51, "CLOSINGAUCTIONPRICE": null}`},
		// no MARKETPRICE2
		{`{"SECID": "GAZP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:58", "PRICEMINUSPREVWAPRICE": -22.98, "VOLTODAY": 47948300, "VALTODAY": 12677905337, "HIGHBID": 304.75, "LOWOFFER": 250.92, "LASTOFFER": 260.29, "LASTBID": 259.71, "OPEN": 253.95, "LOW": 250.92, "HIGH": 273.99, "LAST": 260.29, "LCLOSEPRICE": null, "NUMTRADES": 107517, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE21": null, "LCURRENTPRICE": 260.51, "CLOSINGAUCTIONPRICE": null}`},
		// no LCURRENTPRICE
		{`{"SECID": "GAZP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:58", "PRICEMINUSPREVWAPRICE": -22.98, "VOLTODAY": 47948300, "VALTODAY": 12677905337, "HIGHBID": 304.75, "LOWOFFER": 250.92, "LASTOFFER": 260.29, "LASTBID": 259.71, "OPEN": 253.95, "LOW": 250.92, "HIGH": 273.99, "LAST": 260.29, "LCLOSEPRICE": null, "NUMTRADES": 107517, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE1": 260.51, "CLOSINGAUCTIONPRICE": null}`},
		// no CLOSINGAUCTIONPRICE
		{`{"SECID": "GAZP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:58", "PRICEMINUSPREVWAPRICE": -22.98, "VOLTODAY": 47948300, "VALTODAY": 12677905337, "HIGHBID": 304.75, "LOWOFFER": 250.92, "LASTOFFER": 260.29, "LASTBID": 259.71, "OPEN": 253.95, "LOW": 250.92, "HIGH": 273.99, "LAST": 260.29, "LCLOSEPRICE": null, "NUMTRADES": 107517, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 260.51, "CLOSINGAUCTIONPRICE1": null}`},
	}

	for i, c := range cases {
		secStat := SecStat{}
		if got, expected := parseSecStatItem([]byte(c.incomeJSON), &secStat), jsonparser.KeyPathNotFoundError; got != expected {
			t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead in %d case", expected, got, i)
		}
	}
}

func TestParseSecStat(t *testing.T) {

	var incomeJSON = `
[
      {"SECID": "DSKY", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:55", "PRICEMINUSPREVWAPRICE": -7.12, "VOLTODAY": 1681450, "VALTODAY": 155748831, "HIGHBID": 114.32, "LOWOFFER": 85.88, "LASTOFFER": 92.58, "LASTBID": 92.52, "OPEN": 92, "LOW": 87.22, "HIGH": 96.16, "LAST": 92.54, "LCLOSEPRICE": null, "NUMTRADES": 10500, "WAPRICE": 92.62, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 92.8, "CLOSINGAUCTIONPRICE": null},
      {"SECID": "GAZP", "BOARDID": "SMAL", "TRADINGSESSION": "0", "TIME": "09:40:08", "PRICEMINUSPREVWAPRICE": -23.27, "VOLTODAY": 25, "VALTODAY": 6654, "HIGHBID": 270.42, "LOWOFFER": 258.12, "LASTOFFER": 271.29, "LASTBID": 261, "OPEN": 258.12, "LOW": 258.12, "HIGH": 287.99, "LAST": 260, "LCLOSEPRICE": null, "NUMTRADES": 16, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": null, "CLOSINGAUCTIONPRICE": null},
      {"SECID": "GAZP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:58", "PRICEMINUSPREVWAPRICE": -22.98, "VOLTODAY": 47948300, "VALTODAY": 12677905337, "HIGHBID": 304.75, "LOWOFFER": 250.92, "LASTOFFER": 260.29, "LASTBID": 259.71, "OPEN": 253.95, "LOW": 250.92, "HIGH": 273.99, "LAST": 260.29, "LCLOSEPRICE": null, "NUMTRADES": 107517, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 260.51, "CLOSINGAUCTIONPRICE": null},
      {"SECID": "SBERP", "BOARDID": "SMAL", "TRADINGSESSION": "0", "TIME": "09:33:40", "PRICEMINUSPREVWAPRICE": -17.24, "VOLTODAY": 38, "VALTODAY": 7321, "HIGHBID": 208.01, "LOWOFFER": 185, "LASTOFFER": 204.97, "LASTBID": 190.01, "OPEN": 190, "LOW": 185, "HIGH": 208.01, "LAST": 193, "LCLOSEPRICE": null, "NUMTRADES": 23, "WAPRICE": 193.01, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": null, "CLOSINGAUCTIONPRICE": null},
      {"SECID": "SBERP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:57", "PRICEMINUSPREVWAPRICE": -17.85, "VOLTODAY": 9160070, "VALTODAY": 1768007018, "HIGHBID": 221.66, "LOWOFFER": 175.23, "LASTOFFER": 192.47, "LASTBID": 192.27, "OPEN": 194.8, "LOW": 184, "HIGH": 199.87, "LAST": 192.39, "LCLOSEPRICE": null, "NUMTRADES": 38395, "WAPRICE": 193.01, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 190.91, "CLOSINGAUCTIONPRICE": null}]}
]
`
	secSt := make([]SecStat, 0)
	err := parseSecStat([]byte(incomeJSON), &secSt)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := len(secSt), 5; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseSecStatUnexpectedDataTypeError(t *testing.T) {

	var incomeJSON = `
[
      []
]`
	secSt := make([]SecStat, 0)
	if got, expected := parseSecStat([]byte(incomeJSON), &secSt), ErrUnexpectedDataType; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseSecStatError(t *testing.T) {

	var incomeJSON = `
[
      {"SECID": "DSKY", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:55", "PRICEMINUSPREVWAPRICE": -7.12, "VOLTODAY": 1681450, "VALTODAY": 155748831, "HIGHBID": 114.32, "LOWOFFER": 85.88, "LASTOFFER": 92.58, "LASTBID": 92.52, "OPEN": 92, "LOW": 87.22, "HIGH": 96.16, "LAST": 92.54, "LCLOSEPRICE": null, "NUMTRADES": 10500, "WAPRICE": 92.62, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 92.8, "CLOSINGAUCTIONPRICE": null},
      {"SECID": "GAZP", "BOARDID1": "SMAL", "TRADINGSESSION": "0", "TIME": "09:40:08", "PRICEMINUSPREVWAPRICE": -23.27, "VOLTODAY": 25, "VALTODAY": 6654, "HIGHBID": 270.42, "LOWOFFER": 258.12, "LASTOFFER": 271.29, "LASTBID": 261, "OPEN": 258.12, "LOW": 258.12, "HIGH": 287.99, "LAST": 260, "LCLOSEPRICE": null, "NUMTRADES": 16, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": null, "CLOSINGAUCTIONPRICE": null},
      {"SECID": "GAZP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:58", "PRICEMINUSPREVWAPRICE": -22.98, "VOLTODAY": 47948300, "VALTODAY": 12677905337, "HIGHBID": 304.75, "LOWOFFER": 250.92, "LASTOFFER": 260.29, "LASTBID": 259.71, "OPEN": 253.95, "LOW": 250.92, "HIGH": 273.99, "LAST": 260.29, "LCLOSEPRICE": null, "NUMTRADES": 107517, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 260.51, "CLOSINGAUCTIONPRICE": null},
      {"SECID": "SBERP", "BOARDID": "SMAL", "TRADINGSESSION": "0", "TIME": "09:33:40", "PRICEMINUSPREVWAPRICE": -17.24, "VOLTODAY": 38, "VALTODAY": 7321, "HIGHBID": 208.01, "LOWOFFER": 185, "LASTOFFER": 204.97, "LASTBID": 190.01, "OPEN": 190, "LOW": 185, "HIGH": 208.01, "LAST": 193, "LCLOSEPRICE": null, "NUMTRADES": 23, "WAPRICE": 193.01, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": null, "CLOSINGAUCTIONPRICE": null},
      {"SECID": "SBERP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:57", "PRICEMINUSPREVWAPRICE": -17.85, "VOLTODAY": 9160070, "VALTODAY": 1768007018, "HIGHBID": 221.66, "LOWOFFER": 175.23, "LASTOFFER": 192.47, "LASTBID": 192.27, "OPEN": 194.8, "LOW": 184, "HIGH": 199.87, "LAST": 192.39, "LCLOSEPRICE": null, "NUMTRADES": 38395, "WAPRICE": 193.01, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 190.91, "CLOSINGAUCTIONPRICE": null}]}
]`
	secSt := make([]SecStat, 0)
	if got, expected := parseSecStat([]byte(incomeJSON), &secSt), jsonparser.KeyPathNotFoundError; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseSecStatResponse(t *testing.T) {

	var incomeJSON = `
[
  {"charsetinfo": {"name": "utf-8"}},
  {
    "secstats": [
      {"SECID": "GAZP", "BOARDID": "SMAL", "TRADINGSESSION": "1", "TIME": "09:40:08", "PRICEMINUSPREVWAPRICE": -23.27, "VOLTODAY": 25, "VALTODAY": 6654, "HIGHBID": 270.42, "LOWOFFER": 258.12, "LASTOFFER": 271.29, "LASTBID": 261, "OPEN": 258.12, "LOW": 258.12, "HIGH": 287.99, "LAST": 260, "LCLOSEPRICE": null, "NUMTRADES": 16, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": null, "CLOSINGAUCTIONPRICE": null},
      {"SECID": "SBERP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:57", "PRICEMINUSPREVWAPRICE": -17.85, "VOLTODAY": 9160070, "VALTODAY": 1768007018, "HIGHBID": 221.66, "LOWOFFER": 175.23, "LASTOFFER": 192.47, "LASTBID": 192.27, "OPEN": 194.8, "LOW": 184, "HIGH": 199.87, "LAST": 192.39, "LCLOSEPRICE": null, "NUMTRADES": 38395, "WAPRICE": 193.01, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 190.91, "CLOSINGAUCTIONPRICE": null}]}
]
`
	expectedResponse := SecStatResponse{
		SecStats: []SecStat{
			{
				Ticker:           "GAZP",
				BoardId:          "SMAL",
				TrSession:        TradingSessionMain,
				Time:             "09:40:08",
				PriceMinusPrevPr: -23.27,
				VolToday:         25,
				ValToday:         6654,
				HighBid:          270.42,
				LowOffer:         258.12,
				LastOffer:        271.29,
				LastBid:          261,
				Open:             258.12,
				Low:              258.12,
				High:             287.99,
				Last:             260,
				LClosePrice:      0,
				NumTrades:        16,
				WaPrice:          264.41,
				AdmittedQuote:    0,
				MarketPrice:      0,
				LCurrentPrice:    0,
				ClosingAucPrice:  0,
			},
			{
				Ticker:           "SBERP",
				BoardId:          "TQBR",
				TrSession:        TradingSessionUndefined,
				Time:             "09:49:57",
				PriceMinusPrevPr: -17.85,
				VolToday:         9160070,
				ValToday:         1768007018,
				HighBid:          221.66,
				LowOffer:         175.23,
				LastOffer:        192.47,
				LastBid:          192.27,
				Open:             194.8,
				Low:              184,
				High:             199.87,
				Last:             192.39,
				LClosePrice:      0,
				NumTrades:        38395,
				WaPrice:          193.01,
				AdmittedQuote:    0,
				MarketPrice:      0,
				LCurrentPrice:    190.91,
				ClosingAucPrice:  0,
			},
		},
	}
	secStatR := SecStatResponse{}
	var err error = nil
	if got, expected := parseSecStatResponse([]byte(incomeJSON), &secStatR), err; got != expected {
		t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead", expected, got)
	}
	if got, expected := len(secStatR.SecStats), len(secStatR.SecStats); got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
	for i, gotItem := range secStatR.SecStats {
		if got, expected := gotItem, expectedResponse.SecStats[i]; got != expected {
			t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
		}
	}

}

func TestParseSecStatResponseResponseNilError(t *testing.T) {
	var incomeJSON = ``
	var secStatResponse *SecStatResponse = nil

	if got, expected := parseSecStatResponse([]byte(incomeJSON), secStatResponse), ErrNilPointer; got != expected {
		t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseSecStatResponseError(t *testing.T) {
	var incomeJSON = `
[
  {"charsetinfo": {"name": "utf-8"}},
  {
    "secstats": [
      {"SECID": "GAZP", "BOARDID1": "SMAL", "TRADINGSESSION": "1", "TIME": "09:40:08", "PRICEMINUSPREVWAPRICE": -23.27, "VOLTODAY": 25, "VALTODAY": 6654, "HIGHBID": 270.42, "LOWOFFER": 258.12, "LASTOFFER": 271.29, "LASTBID": 261, "OPEN": 258.12, "LOW": 258.12, "HIGH": 287.99, "LAST": 260, "LCLOSEPRICE": null, "NUMTRADES": 16, "WAPRICE": 264.41, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": null, "CLOSINGAUCTIONPRICE": null},
      {"SECID": "SBERP", "BOARDID": "TQBR", "TRADINGSESSION": "0", "TIME": "09:49:57", "PRICEMINUSPREVWAPRICE": -17.85, "VOLTODAY": 9160070, "VALTODAY": 1768007018, "HIGHBID": 221.66, "LOWOFFER": 175.23, "LASTOFFER": 192.47, "LASTBID": 192.27, "OPEN": 194.8, "LOW": 184, "HIGH": 199.87, "LAST": 192.39, "LCLOSEPRICE": null, "NUMTRADES": 38395, "WAPRICE": 193.01, "ADMITTEDQUOTE": null, "MARKETPRICE2": null, "LCURRENTPRICE": 190.91, "CLOSINGAUCTIONPRICE": null}]}
]
`
	secStatR := SecStatResponse{}

	if got, expected := parseSecStatResponse([]byte(incomeJSON), &secStatR), jsonparser.KeyPathNotFoundError; got != expected {
		t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestStatsService_GetSecStats(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		byteValueResult, err := getTestingData("secstats.json")
		if err != nil {
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(byteValueResult)
		if err != nil {
			fmt.Println(err)
		}
	}))
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)
	c.BaseURL, _ = url.Parse(srv.URL + "/")
	result, err := c.Stats.GetSecStats(context.Background(), EngineStock, "shares", nil)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := len(result.SecStats), 6; got != expected {
		t.Fatalf("Error: expecting: \n %v items\ngot:\n %v items\ninstead", expected, got)
	}
}

func TestStatsService_GetSecStats_KeyPathNotFound(t *testing.T) {
	srv := getEmptySrv()
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)

	c.BaseURL, _ = url.Parse(srv.URL + "/")
	_, err := c.Stats.GetSecStats(context.Background(), EngineStock, "shares", nil)
	if got, expected := err, jsonparser.KeyPathNotFoundError; got == nil || got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v \ninstead", expected, got)
	}
}

func TestStatsService_GetSecStats_BadUrl(t *testing.T) {
	srv := getEmptySrv()
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)

	c.BaseURL, _ = url.Parse(srv.URL)
	_, err := c.Stats.GetSecStats(context.Background(), EngineStock, "shares", nil)
	if got, expected := err, "BaseURL must have a trailing slash, but \""+srv.URL+"\" does not"; got == nil || got.Error() != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestStatsService_GetSecStats_EmptyEngine(t *testing.T) {
	srv := getEmptySrv()
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)

	c.BaseURL, _ = url.Parse(srv.URL)
	_, err := c.Stats.GetSecStats(context.Background(), "", "shares", nil)
	if got, expected := err, ErrBadEngineParameter; got == nil || got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestStatsServiceNilContextError(t *testing.T) {
	c := NewClient(nil)
	var ctx context.Context = nil
	_, err := c.Stats.GetSecStats(ctx, EngineStock, "shares", nil)
	if got, expected := err, ErrNonNilContext; got == nil || got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v \ninstead", expected, got)
	}
}
