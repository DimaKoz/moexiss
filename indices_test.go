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
