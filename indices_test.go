package moexiss

import (
	"context"
	"github.com/buger/jsonparser"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestIndicesGetUrl(t *testing.T) {
	var income *IndicesRequestOptions = nil
	c := NewClient(nil)
	gotUrl, err := c.Indices.getUrl("sberp", income)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := gotUrl, `https://iss.moex.com/iss/securities/sberp/indices.json?iss.json=extended&iss.meta=off`; got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}

func TestIndicesGetUrlBadSecurity(t *testing.T) {
	var income *IndicesRequestOptions = nil
	c := NewClient(nil)
	_, err := c.Indices.getUrl("", income)
	if err != ErrBadSecurityParameter {
		t.Fatalf("Error: expecting error: %v \ngot %v \ninstead", ErrBadSecurityParameter, err)
	}
}

func TestParseIndicesItem(t *testing.T) {
	expectedStruct := Indices{
		IndexId:   "IMOEX",
		IndexName: "Индекс МосБиржи",
		From:      "2007-04-16",
		Till:      "2022-01-26",
	}
	var incomeJson = `
      {"SECID": "IMOEX", "SHORTNAME": "Индекс МосБиржи", "FROM": "2007-04-16", "TILL": "2022-01-26"}
`
	indicesItem := Indices{}
	err := parseIndicesItem([]byte(incomeJson), &indicesItem)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := indicesItem, expectedStruct; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseIndicesItemErrCases(t *testing.T) {
	type Case struct {
		incomeJson string
	}
	cases := []Case{
		// no SECID
		{`{"SECID1": "IMOEX", "SHORTNAME": "Индекс МосБиржи", "FROM": "2007-04-16", "TILL": "2022-01-26"}`},
		// no SHORTNAME
		{`{"SECID": "IMOEX", "SHORTNAME1": "Индекс МосБиржи", "FROM": "2007-04-16", "TILL": "2022-01-26"}`},
		// no FROM
		{`{"SECID": "IMOEX", "SHORTNAME": "Индекс МосБиржи", "FROM1": "2007-04-16", "TILL": "2022-01-26"}`},
		// no TILL
		{`{"SECID": "IMOEX", "SHORTNAME": "Индекс МосБиржи", "FROM": "2007-04-16", "TILL1": "2022-01-26"}`},
	}

	for i, c := range cases {
		indices := Indices{}
		if got, expected := parseIndicesItem([]byte(c.incomeJson), &indices), jsonparser.KeyPathNotFoundError; got != expected {
			t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead in %d case", expected, got, i)
		}

	}
}

func TestParseIndices(t *testing.T) {

	var incomeJson = `
[
      {"SECID": "IMOEX", "SHORTNAME": "Индекс МосБиржи", "FROM": "2007-04-16", "TILL": "2022-01-26"},
      {"SECID": "IMOEX2", "SHORTNAME": "Индекс МосБиржи (все сессии)", "FROM": "2020-06-22", "TILL": "2022-01-25"},
      {"SECID": "MCXSM", "SHORTNAME": "Индекс МосБиржи SMID", "FROM": "2014-01-06", "TILL": "2014-12-30"},
      {"SECID": "RUCGI", "SHORTNAME": "Нац. индекс корп. Управления", "FROM": "2021-06-18", "TILL": "2022-01-26"}
]
`
	indices := make([]Indices, 0)
	err := parseIndices([]byte(incomeJson), &indices)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := len(indices), 4; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseIndicesUnexpectedDataTypeError(t *testing.T) {

	var incomeJson = `
[
      []
]`
	indices := make([]Indices, 0)
	if got, expected := parseIndices([]byte(incomeJson), &indices), ErrUnexpectedDataType; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseIndicesError(t *testing.T) {

	var incomeJson = `
[
      {"SECID": "IMOEX", "SHORTNAME": "Индекс МосБиржи", "FROM": "2007-04-16", "TILL": "2022-01-26"},
      {"SECID": "RUCGI", "SHORTNAME1": "Нац. индекс корп. Управления", "FROM": "2021-06-18", "TILL": "2022-01-26"},
      {"SECID": "RUCGI", "SHORTNAME": "Нац. индекс корп. Управления", "FROM": "2021-06-18", "TILL": "2022-01-26"}
]`
	indices := make([]Indices, 0)
	if got, expected := parseIndices([]byte(incomeJson), &indices), jsonparser.KeyPathNotFoundError; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseIndicesResponse(t *testing.T) {
	var incomeJson = `
[
  {"charsetinfo": {"name": "utf-8"}},
  {
    "indices": [
      {"SECID": "IMOEX", "SHORTNAME": "Индекс МосБиржи", "FROM": "2007-04-16", "TILL": "2022-01-26"},
      {"SECID": "RTSI", "SHORTNAME": "Индекс РТС", "FROM": "2009-09-30", "TILL": "2022-01-26"}
]}
]
`
	expectedResponse := IndicesResponse{
		Indices: []Indices{
			{
				IndexId:   "IMOEX",
				IndexName: "Индекс МосБиржи",
				From:      "2007-04-16",
				Till:      "2022-01-26",
			},
			{
				IndexId:   "RTSI",
				IndexName: "Индекс РТС",
				From:      "2009-09-30",
				Till:      "2022-01-26"},
		},
	}
	indicesR := IndicesResponse{}
	var err error = nil
	if got, expected := parseIndicesResponse([]byte(incomeJson), &indicesR), err; got != expected {
		t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead", expected, got)
	}
	if got, expected := len(indicesR.Indices), len(expectedResponse.Indices); got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
	for i, gotItem := range indicesR.Indices {
		if got, expected := gotItem, expectedResponse.Indices[i]; got != expected {
			t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
		}
	}

}

func TestParseIndicesResponseNilError(t *testing.T) {
	var incomeJson = ``
	var indicesResponse *IndicesResponse = nil

	if got, expected := parseIndicesResponse([]byte(incomeJson), indicesResponse), ErrNilPointer; got != expected {
		t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseIndicesResponseError(t *testing.T) {
	var incomeJson = `
[
  {"charsetinfo": {"name": "utf-8"}},
  {
    "indices": [
      {"SECID": "IMOEX", "SHORTNAME": "Индекс МосБиржи", "FROM": "2007-04-16", "TILL": "2022-01-26"},
      {"SECID1": "RUCGI", "SHORTNAME": "Нац. индекс корп. Управления", "FROM": "2021-06-18", "TILL": "2022-01-26"}]}
]
`
	var indicesResponse = &IndicesResponse{}

	if got, expected := parseIndicesResponse([]byte(incomeJson), indicesResponse), jsonparser.KeyPathNotFoundError; got != expected {
		t.Fatalf("Error: expecting error: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestIndicesService_BadSecurityParam(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {}))
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)

	c.BaseURL, _ = url.Parse(srv.URL)
	_, err := c.Indices.Indices(context.Background(), "", nil)
	if got, expected := err, ErrBadSecurityParameter; got == nil || got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestIndicesService_BadUrl(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {}))
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)

	c.BaseURL, _ = url.Parse(srv.URL)
	_, err := c.Indices.Indices(context.Background(), "sber", nil)
	if got, expected := err, "BaseURL must have a trailing slash, but \""+srv.URL+"\" does not"; got == nil || got.Error() != expected {
		t.Fatalf("Error: expecting %v error \ngot %v  \ninstead", expected, got)
	}
}

func TestIndicesKeyPathNotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		str := `[{}]`
		_, _ = w.Write([]byte(str))
	}))
	defer srv.Close()

	httpClient := srv.Client()

	c := NewClient(httpClient)

	c.BaseURL, _ = url.Parse(srv.URL + "/")
	_, err := c.Indices.Indices(context.Background(), "jhgsd", nil)
	if got, expected := err, jsonparser.KeyPathNotFoundError; got == nil || got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v \ninstead", expected, got)
	}
}

func TestIndicesNilContextError(t *testing.T) {
	c := NewClient(nil)
	var ctx context.Context = nil
	_, err := c.Indices.Indices(ctx, "SBERP", nil)
	if got, expected := err, ErrNonNilContext; got == nil || got != expected {
		t.Fatalf("Error: expecting %v error \ngot %v \ninstead", expected, got)
	}
}

func TestIndicesService_Indices(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		str := `
[
  {
    "indices": [
      {"SECID": "IMOEX", "SHORTNAME": "Индекс МосБиржи", "FROM": "2007-04-16", "TILL": "2022-01-26"},
      {"SECID": "IMOEX2", "SHORTNAME": "Индекс МосБиржи (все сессии)", "FROM": "2020-06-22", "TILL": "2022-01-25"},
      {"SECID": "MICEXLC", "SHORTNAME": "MICEX LC", "FROM": "2006-07-03", "TILL": "2013-05-17"},
      {"SECID": "MICEXMC", "SHORTNAME": "MICEX MC", "FROM": "2006-09-04", "TILL": "2007-07-13"},
      {"SECID": "MOEXBC", "SHORTNAME": "Индекс голубых фишек", "FROM": "2019-09-20", "TILL": "2022-01-26"},
      {"SECID": "RUCGI", "SHORTNAME": "Нац. индекс корп. Управления", "FROM": "2021-06-18", "TILL": "2022-01-26"}]}
]

`
		_, _ = w.Write([]byte(str))
	}))
	defer srv.Close()

	httpClient := srv.Client()
	secId := "SBERP"
	c := NewClient(httpClient)
	c.BaseURL, _ = url.Parse(srv.URL + "/")
	result, err := c.Indices.Indices(context.Background(), secId, nil)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := len(result.Indices), 6; got != expected {
		t.Fatalf("Error: expecting: \n %v items\ngot:\n %v items\ninstead", expected, got)
	}
	if got, expected := result.SecurityId, secId; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}
