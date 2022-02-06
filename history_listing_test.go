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

func TestParseListing(t *testing.T) {

	var incomeJson = `
[
      {"SECID": "CHMF", "SHORTNAME": "СевСт-ао", "NAME": "Северсталь (ПАО)ао", "BOARDID": "EQCC", "decimals": 1, "history_from": "2010-02-15", "history_till": "2011-05-27"},
      {"SECID": "CHMF", "SHORTNAME": "СевСт-ао", "NAME": "Северсталь (ПАО)ао", "BOARDID": "TQDP", "decimals": 1, "history_from": null, "history_till": null},
      {"SECID": "CHMK", "SHORTNAME": "ЧМК ао", "NAME": "\"ЧМК\" ПАО ао", "BOARDID": "TQBR", "decimals": 0, "history_from": "2014-06-09", "history_till": "2022-02-04"},
      {"SECID": "CHMK", "SHORTNAME": "ЧМК ао", "NAME": "\"ЧМК\" ПАО ао", "BOARDID": "TQNE", "decimals": 0, "history_from": "2013-09-02", "history_till": "2014-06-06"},
      {"SECID": "CHMK", "SHORTNAME": "ЧМК ао", "NAME": "\"ЧМК\" ПАО ао", "BOARDID": "EQNE", "decimals": 0, "history_from": "2008-12-12", "history_till": "2013-08-30"},
      {"SECID": "CHMZ", "SHORTNAME": "ЧМЗ ао", "NAME": "Чусовской мет.завод ОАО ао", "BOARDID": "EQNE", "decimals": 2, "history_from": "2011-12-08", "history_till": "2012-07-09"}
]
`
	listing := make([]Listing, 0)
	err := parseListing([]byte(incomeJson), &listing)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := len(listing), 6; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseListingUnexpectedDataTypeError(t *testing.T) {

	var incomeJson = `
[
      []
]`
	listing := make([]Listing, 0)
	if got, expected := parseListing([]byte(incomeJson), &listing), ErrUnexpectedDataType; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseListingError(t *testing.T) {

	var incomeJson = `
[
      {"SECID": "CHMF", "SHORTNAME": "СевСт-ао", "NAME": "Северсталь (ПАО)ао", "BOARDID": "EQCC", "decimals": 1, "history_from": "2010-02-15", "history_till": "2011-05-27"},
      {"SECID": "CHMF", "SHORTNAME1": "СевСт-ао", "NAME": "Северсталь (ПАО)ао", "BOARDID": "TQDP", "decimals": 1, "history_from": null, "history_till": null},
      {"SECID": "CHMK", "SHORTNAME": "ЧМК ао", "NAME": "\"ЧМК\" ПАО ао", "BOARDID": "TQBR", "decimals": 0, "history_from": "2014-06-09", "history_till": "2022-02-04"},
      {"SECID": "CHMK", "SHORTNAME": "ЧМК ао", "NAME": "\"ЧМК\" ПАО ао", "BOARDID": "TQNE", "decimals": 0, "history_from": "2013-09-02", "history_till": "2014-06-06"},
      {"SECID": "CHMK", "SHORTNAME": "ЧМК ао", "NAME": "\"ЧМК\" ПАО ао", "BOARDID": "EQNE", "decimals": 0, "history_from": "2008-12-12", "history_till": "2013-08-30"},
      {"SECID": "CHMZ", "SHORTNAME": "ЧМЗ ао", "NAME": "Чусовской мет.завод ОАО ао", "BOARDID": "EQNE", "decimals": 2, "history_from": "2011-12-08", "history_till": "2012-07-09"}
]`
	listing := make([]Listing, 0)
	if got, expected := parseListing([]byte(incomeJson), &listing), jsonparser.KeyPathNotFoundError; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}