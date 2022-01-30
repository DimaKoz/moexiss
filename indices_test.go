package moexiss

import (
	"github.com/buger/jsonparser"
	"testing"
)

func TestIndicesGetUrl(t *testing.T) {
	var income *IndicesRequestOptions = nil
	c := NewClient(nil)
	url, err := c.Indices.getUrl("sberp", income)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := url, `https://iss.moex.com/iss/securities/sberp/indices.json?iss.json=extended&iss.meta=off`; got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}

func TestIndicesGetUrlBadSecurity(t *testing.T) {
	var income *IndicesRequestOptions = nil
	c := NewClient(nil)
	_, err := c.Indices.getUrl("", income)
	if err != errBadSecurityParameter {
		t.Fatalf("Error: expecting error: %v \ngot %v \ninstead", errBadSecurityParameter, err)
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
	if got, expected := parseIndices([]byte(incomeJson), &indices), errUnexpectedDataType; got != expected {
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

	if got, expected := parseIndicesResponse([]byte(incomeJson), indicesResponse), errNilPointer; got != expected {
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
