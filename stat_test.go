package moexiss

import "testing"

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
