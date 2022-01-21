package moexiss

import "testing"

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
