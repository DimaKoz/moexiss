package moexiss

import (
	"github.com/buger/jsonparser"
	"testing"
)

func TestHistoryListingGetUrl(t *testing.T) {
	c := NewClient(nil)
	gotUrl, err := c.HistoryListing.getUrlListing(EngineStock, "shares", nil)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := gotUrl, `https://iss.moex.com/iss/history/engines/stock/markets/shares/listing.json?iss.json=extended&iss.meta=off`; got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}

func TestParseListingItem(t *testing.T) {
	expectedStruct := Listing{
		Ticker:    "BSPB",
		ShortName: "БСП ао",
		FullName:  "ПАО \"Банк \"Санкт-Петербург\" ао",
		BoardId:   "TQBR",
		Decimals:  2,
		From:      "2014-06-09",
		Till:      "2022-02-04",
	}
	var incomeJson = `
      {"SECID": "BSPB", "SHORTNAME": "БСП ао", "NAME": "ПАО \"Банк \"Санкт-Петербург\" ао", "BOARDID": "TQBR", "decimals": 2, "history_from": "2014-06-09", "history_till": "2022-02-04"}
`
	listingItem := Listing{}
	err := parseListingItem([]byte(incomeJson), &listingItem)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := listingItem, expectedStruct; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseListingItemErrCases(t *testing.T) {
	type Case struct {
		incomeJson string
	}
	cases := []Case{
		// no SECID
		{`{"SECID1": "BSPB", "SHORTNAME": "БСП ао", "NAME": "ПАО \"Банк \"Санкт-Петербург\" ао", "BOARDID": "TQBR", "decimals": 2, "history_from": "2014-06-09", "history_till": "2022-02-04"}`},
		// no SHORTNAME
		{`{"SECID": "BSPB", "SHORTNAME1": "БСП ао", "NAME": "ПАО \"Банк \"Санкт-Петербург\" ао", "BOARDID": "TQBR", "decimals": 2, "history_from": "2014-06-09", "history_till": "2022-02-04"}`},
		// no NAME
		{`{"SECID": "BSPB", "SHORTNAME": "БСП ао", "NAME1": "ПАО \"Банк \"Санкт-Петербург\" ао", "BOARDID": "TQBR", "decimals": 2, "history_from": "2014-06-09", "history_till": "2022-02-04"}`},
		// no BOARDID
		{`{"SECID": "BSPB", "SHORTNAME": "БСП ао", "NAME": "ПАО \"Банк \"Санкт-Петербург\" ао", "BOARDID1": "TQBR", "decimals": 2, "history_from": "2014-06-09", "history_till": "2022-02-04"}`},
		// no decimals
		{`{"SECID": "BSPB", "SHORTNAME": "БСП ао", "NAME": "ПАО \"Банк \"Санкт-Петербург\" ао", "BOARDID": "TQBR", "decimals1": 2, "history_from": "2014-06-09", "history_till": "2022-02-04"}`},
		// no history_from
		{`{"SECID": "BSPB", "SHORTNAME": "БСП ао", "NAME": "ПАО \"Банк \"Санкт-Петербург\" ао", "BOARDID": "TQBR", "decimals": 2, "history_from1": "2014-06-09", "history_till": "2022-02-04"}`},
		// no history_till
		{`{"SECID": "BSPB", "SHORTNAME": "БСП ао", "NAME": "ПАО \"Банк \"Санкт-Петербург\" ао", "BOARDID": "TQBR", "decimals": 2, "history_from": "2014-06-09", "history_till1": "2022-02-04"}`},
	}

	for i, c := range cases {
		listing := Listing{}
		if got, expected := parseListingItem([]byte(c.incomeJson), &listing), jsonparser.KeyPathNotFoundError; got != expected {
			t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead in %d case", expected, got, i)
		}

	}
}
