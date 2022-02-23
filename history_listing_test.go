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

func TestHistoryListingGetUrl(t *testing.T) {
	c := NewClient(nil)
	gotURL, err := c.HistoryListing.getUrlListing(EngineStock, "shares", nil)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := gotURL, `https://iss.moex.com/iss/history/engines/stock/markets/shares/listing.json?iss.json=extended&iss.meta=off`; got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}

func TestHistoryListingGetUrlListingByBoard(t *testing.T) {
	c := NewClient(nil)
	gotURL, err := c.HistoryListing.getUrlListingByBoard(EngineStock, "shares", "TQTD", nil)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := gotURL, `https://iss.moex.com/iss/history/engines/stock/markets/shares/boards/TQTD/listing.json?iss.json=extended&iss.meta=off`; got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}

func TestHistoryListingGetUrlListingByBoardGroup(t *testing.T) {
	c := NewClient(nil)
	gotURL, err := c.HistoryListing.getUrlListingByBoardGroup(EngineStock, "shares", "6", nil)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := gotURL, `https://iss.moex.com/iss/history/engines/stock/markets/shares/boardgroups/6/listing.json?iss.json=extended&iss.meta=off`; got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}

func TestHistoryListingGetUrlListingByBoardGroupBadEngine(t *testing.T) {
	c := NewClient(nil)
	_, err := c.HistoryListing.getUrlListingByBoardGroup(EngineUndefined, "shares", "6", nil)
	if got, expected := err, ErrBadEngineParameter; got != expected {
		t.Fatalf("Error: expecting %v error: \ngot %v \ninstead", expected, got)
	}
}

func TestHistoryListingGetUrlListingByBoardGroupBadMarket(t *testing.T) {
	c := NewClient(nil)
	_, err := c.HistoryListing.getUrlListingByBoardGroup(EngineStock, "", "6", nil)
	if got, expected := err, ErrBadMarketParameter; got != expected {
		t.Fatalf("Error: expecting %v error: \ngot %v \ninstead", expected, got)
	}
}

func TestHistoryListingGetUrlListingByBoardBadBoardGroup(t *testing.T) {
	c := NewClient(nil)
	_, err := c.HistoryListing.getUrlListingByBoardGroup(EngineStock, "shares", "", nil)
	if got, expected := err, ErrBadBoardGroupParameter; got != expected {
		t.Fatalf("Error: expecting %v error: \ngot %v \ninstead", expected, got)
	}
}

func TestHistoryListingGetUrlListingByBoardBadBoard(t *testing.T) {
	c := NewClient(nil)
	_, err := c.HistoryListing.getUrlListingByBoard(EngineStock, "shares", "", nil)
	if got, expected := err, ErrBadBoardParameter; got != expected {
		t.Fatalf("Error: expecting %v error: \ngot %v \ninstead", expected, got)
	}
}

func TestHistoryListingGetUrlListingByBoardBadEngine(t *testing.T) {
	c := NewClient(nil)
	_, err := c.HistoryListing.getUrlListingByBoard(EngineUndefined, "shares", "TQTD", nil)
	if got, expected := err, ErrBadEngineParameter; got != expected {
		t.Fatalf("Error: expecting %v error: \ngot %v \ninstead", expected, got)
	}
}

func TestHistoryListingGetUrlListingByBoardBadMarket(t *testing.T) {
	c := NewClient(nil)
	_, err := c.HistoryListing.getUrlListingByBoard(EngineStock, "", "TQTD", nil)
	if got, expected := err, ErrBadMarketParameter; got != expected {
		t.Fatalf("Error: expecting %v error: \ngot %v \ninstead", expected, got)
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
	var incomeJSON = `
      {"SECID": "BSPB", "SHORTNAME": "БСП ао", "NAME": "ПАО \"Банк \"Санкт-Петербург\" ао", "BOARDID": "TQBR", "decimals": 2, "history_from": "2014-06-09", "history_till": "2022-02-04"}
`
	listingItem := Listing{}
	err := parseListingItem([]byte(incomeJSON), &listingItem)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := listingItem, expectedStruct; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseListingItemErrCases(t *testing.T) {
	type Case struct {
		incomeJSON string
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
		if got, expected := parseListingItem([]byte(c.incomeJSON), &listing), jsonparser.KeyPathNotFoundError; got != expected {
			t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead in %d case", expected, got, i)
		}

	}
}

func TestParseListing(t *testing.T) {

	var incomeJSON = `
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
	err := parseListing([]byte(incomeJSON), &listing)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := len(listing), 6; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseListingUnexpectedDataTypeError(t *testing.T) {

	var incomeJSON = `
[
      []
]`
	listing := make([]Listing, 0)
	if got, expected := parseListing([]byte(incomeJSON), &listing), ErrUnexpectedDataType; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseListingError(t *testing.T) {

	var incomeJSON = `
[
      {"SECID": "CHMF", "SHORTNAME": "СевСт-ао", "NAME": "Северсталь (ПАО)ао", "BOARDID": "EQCC", "decimals": 1, "history_from": "2010-02-15", "history_till": "2011-05-27"},
      {"SECID": "CHMF", "SHORTNAME1": "СевСт-ао", "NAME": "Северсталь (ПАО)ао", "BOARDID": "TQDP", "decimals": 1, "history_from": null, "history_till": null},
      {"SECID": "CHMK", "SHORTNAME": "ЧМК ао", "NAME": "\"ЧМК\" ПАО ао", "BOARDID": "TQBR", "decimals": 0, "history_from": "2014-06-09", "history_till": "2022-02-04"},
      {"SECID": "CHMK", "SHORTNAME": "ЧМК ао", "NAME": "\"ЧМК\" ПАО ао", "BOARDID": "TQNE", "decimals": 0, "history_from": "2013-09-02", "history_till": "2014-06-06"},
      {"SECID": "CHMK", "SHORTNAME": "ЧМК ао", "NAME": "\"ЧМК\" ПАО ао", "BOARDID": "EQNE", "decimals": 0, "history_from": "2008-12-12", "history_till": "2013-08-30"},
      {"SECID": "CHMZ", "SHORTNAME": "ЧМЗ ао", "NAME": "Чусовской мет.завод ОАО ао", "BOARDID": "EQNE", "decimals": 2, "history_from": "2011-12-08", "history_till": "2012-07-09"}
]`
	listing := make([]Listing, 0)
	if got, expected := parseListing([]byte(incomeJSON), &listing), jsonparser.KeyPathNotFoundError; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseListingResponse(t *testing.T) {

	var incomeJSON = `
[
  {"charsetinfo": {"name": "utf-8"}},
  {
    "securities": [
      {"SECID": "CHMF", "SHORTNAME": "СевСт-ао", "NAME": "Северсталь (ПАО)ао", "BOARDID": "EQCC", "decimals": 1, "history_from": "2010-02-15", "history_till": "2011-05-27"},
      {"SECID": "CHMF", "SHORTNAME": "СевСт-ао", "NAME": "Северсталь (ПАО)ао", "BOARDID": "TQDP", "decimals": 1, "history_from": null, "history_till": null}
]}
]
`
	expectedResponse := ListingResponse{
		Listing: []Listing{
			{
				Ticker:    "CHMF",
				ShortName: "СевСт-ао",
				FullName:  "Северсталь (ПАО)ао",
				BoardId:   "EQCC",
				Decimals:  1,
				From:      "2010-02-15",
				Till:      "2011-05-27",
			},
			{
				Ticker:    "CHMF",
				ShortName: "СевСт-ао",
				FullName:  "Северсталь (ПАО)ао",
				BoardId:   "TQDP",
				Decimals:  1,
				From:      "",
				Till:      "",
			},
		},
	}
	listingR := ListingResponse{}
	var err error = nil
	if got, expected := parseListingResponse([]byte(incomeJSON), &listingR), err; got != expected {
		t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead", expected, got)
	}
	if got, expected := len(listingR.Listing), len(expectedResponse.Listing); got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
	for i, gotItem := range listingR.Listing {
		if got, expected := gotItem, expectedResponse.Listing[i]; got != expected {
			t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
		}
	}

}

func TestParseListingResponseNilError(t *testing.T) {
	var incomeJSON = ``
	var listingR *ListingResponse = nil

	if got, expected := parseListingResponse([]byte(incomeJSON), listingR), ErrNilPointer; got != expected {
		t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseListingResponseError(t *testing.T) {
	var incomeJSON = `
[
  {"charsetinfo": {"name": "utf-8"}},
  {
    "securities": [
      {"SECID": "CHMF", "SHORTNAME1": "СевСт-ао", "NAME": "Северсталь (ПАО)ао", "BOARDID": "EQCC", "decimals": 1, "history_from": "2010-02-15", "history_till": "2011-05-27"},
      {"SECID": "CHMF", "SHORTNAME": "СевСт-ао", "NAME": "Северсталь (ПАО)ао", "BOARDID": "TQDP", "decimals": 1, "history_from": null, "history_till": null}
]}
]
`
	var listingR = &ListingResponse{}

	if got, expected := parseListingResponse([]byte(incomeJSON), listingR), jsonparser.KeyPathNotFoundError; got != expected {
		t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseListingResponseEmpty(t *testing.T) {
	var incomeJSON = `
[
{"charsetinfo": {"name": "utf-8"}},
{
"securities": [
]}
]
`
	var listingR = &ListingResponse{}

	if got, expected := parseListingResponse([]byte(incomeJSON), listingR), ErrEmptyServerResult; got != expected {
		t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

//A handler to return expected results
//TestingHistoryListingHandler emulates an external server
func TestingHistoryListingHandler(w http.ResponseWriter, _ *http.Request) {

	byteValueResult, err := getTestingData("history_listing.json")
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

func TestHistoryListingService_Listing(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(TestingHistoryListingHandler))
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)
	c.BaseURL, _ = url.Parse(srv.URL + "/")
	result, err := c.HistoryListing.GetListing(context.Background(), EngineStock, "shares", nil)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := len(result.Listing), 100; got != expected {
		t.Fatalf("Error: expecting: \n %v items\ngot:\n %v items\ninstead", expected, got)
	}
}

func TestHistoryListingService_ListingBadEngineParam(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {}))
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)

	c.BaseURL, _ = url.Parse(srv.URL)
	_, err := c.HistoryListing.GetListing(context.Background(), "", "shares", nil)
	if got, expected := err, ErrBadEngineParameter; got == nil || got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestHistoryListingService_ListingBadMarketParam(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {}))
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)

	c.BaseURL, _ = url.Parse(srv.URL)
	_, err := c.HistoryListing.GetListing(context.Background(), EngineStock, "", nil)
	if got, expected := err, ErrBadMarketParameter; got == nil || got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestHistoryListingService_BadUrl(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {}))
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)

	c.BaseURL, _ = url.Parse(srv.URL)
	_, err := c.HistoryListing.GetListing(context.Background(), EngineStock, "shares", nil)
	if got, expected := err, "BaseURL must have a trailing slash, but \""+srv.URL+"\" does not"; got == nil || got.Error() != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestHistoryListingKeyPathNotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		str := `[{}]`
		_, _ = w.Write([]byte(str))
	}))
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)

	c.BaseURL, _ = url.Parse(srv.URL + "/")
	_, err := c.HistoryListing.GetListing(context.Background(), EngineStock, "shares", nil)
	if got, expected := err, jsonparser.KeyPathNotFoundError; got == nil || got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v \ninstead", expected, got)
	}
}

func TestHistoryListingNilContextError(t *testing.T) {
	c := NewClient(nil)
	var ctx context.Context = nil
	_, err := c.HistoryListing.GetListing(ctx, EngineStock, "shares", nil)
	if got, expected := err, ErrNonNilContext; got == nil || got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v \ninstead", expected, got)
	}
}

func TestHistoryListingService_ListingByBoard(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		byteValueResult, err := getTestingData("history_listing_board.json")
		if err != nil {
			return
		}

		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		_, err = writer.Write(byteValueResult)
		if err != nil {
			fmt.Println(err)
		}

	}))
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)
	c.BaseURL, _ = url.Parse(srv.URL + "/")
	result, err := c.HistoryListing.GetListingByBoard(context.Background(), EngineStock, "shares", "TQTD", nil)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := len(result.Listing), 80; got != expected {
		t.Fatalf("Error: expecting: \n %v items\ngot:\n %v items\ninstead", expected, got)
	}
}

func TestHistoryListingByBoardNilContextError(t *testing.T) {
	c := NewClient(nil)
	var ctx context.Context = nil
	_, err := c.HistoryListing.GetListingByBoard(ctx, EngineStock, "shares", "TQTD", nil)
	if got, expected := err, ErrNonNilContext; got == nil || got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v \ninstead", expected, got)
	}
}

func TestHistoryListingByBoardKeyPathNotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		str := `[{}]`
		_, _ = w.Write([]byte(str))
	}))
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)

	c.BaseURL, _ = url.Parse(srv.URL + "/")
	_, err := c.HistoryListing.GetListingByBoard(context.Background(), EngineStock, "shares", "TQTD", nil)
	if got, expected := err, jsonparser.KeyPathNotFoundError; got == nil || got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v \ninstead", expected, got)
	}
}

func TestHistoryListingByBoardService_BadUrl(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {}))
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)

	c.BaseURL, _ = url.Parse(srv.URL)
	_, err := c.HistoryListing.GetListingByBoard(context.Background(), EngineStock, "shares", "TQTD", nil)
	if got, expected := err, "BaseURL must have a trailing slash, but \""+srv.URL+"\" does not"; got == nil || got.Error() != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestHistoryListingService_ListingBadBoardParam(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {}))
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)

	c.BaseURL, _ = url.Parse(srv.URL)
	_, err := c.HistoryListing.GetListingByBoard(context.Background(), EngineStock, "shares", "", nil)
	if got, expected := err, ErrBadBoardParameter; got == nil || got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestHistoryListingService_ListingBadBoardGroupParam(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {}))
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)

	c.BaseURL, _ = url.Parse(srv.URL)
	_, err := c.HistoryListing.GetListingByBoardGroup(context.Background(), EngineStock, "shares", "", nil)
	if got, expected := err, ErrBadBoardGroupParameter; got == nil || got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestHistoryListingByBoardGroupService_BadUrl(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {}))
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)

	c.BaseURL, _ = url.Parse(srv.URL)
	_, err := c.HistoryListing.GetListingByBoardGroup(context.Background(), EngineStock, "shares", "TQTD", nil)
	if got, expected := err, "BaseURL must have a trailing slash, but \""+srv.URL+"\" does not"; got == nil || got.Error() != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestHistoryListingByBoardGroupKeyPathNotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		str := `[{}]`
		_, _ = w.Write([]byte(str))
	}))
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)

	c.BaseURL, _ = url.Parse(srv.URL + "/")
	_, err := c.HistoryListing.GetListingByBoardGroup(context.Background(), EngineStock, "shares", "6", nil)
	if got, expected := err, jsonparser.KeyPathNotFoundError; got == nil || got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v \ninstead", expected, got)
	}
}

func TestHistoryListingByBoardGroupNilContextError(t *testing.T) {
	c := NewClient(nil)
	var ctx context.Context = nil
	_, err := c.HistoryListing.GetListingByBoardGroup(ctx, EngineStock, "shares", "6", nil)
	if got, expected := err, ErrNonNilContext; got == nil || got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v \ninstead", expected, got)
	}
}

func TestHistoryListingService_ListingByBoardGroup(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		byteValueResult, err := getTestingData("history_listing_boardgroups.json")
		if err != nil {
			return
		}

		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		_, err = writer.Write(byteValueResult)
		if err != nil {
			fmt.Println(err)
		}

	}))
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)
	c.BaseURL, _ = url.Parse(srv.URL + "/")
	result, err := c.HistoryListing.GetListingByBoardGroup(context.Background(), EngineStock, "shares", "6", nil)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := len(result.Listing), 100; got != expected {
		t.Fatalf("Error: expecting: \n %v items\ngot:\n %v items\ninstead", expected, got)
	}
}
