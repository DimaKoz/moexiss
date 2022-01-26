package moexiss

import (
	"github.com/buger/jsonparser"
	"testing"
)

func TestParseStringWithDefaultValueNull(t *testing.T) {
	var nullValue = []byte("null")
	expectedValue := ""
	var expectedErr error = nil
	got, err := parseStringWithDefaultValue(nullValue)
	if err != expectedErr {
		t.Fatalf("Error: expecting %v: \ngot %v  \ninstead", expectedErr, err)
	}
	if got != expectedValue {
		t.Fatalf("Error: expecting value: %s got %s instead", expectedValue, got)
	}
}

func TestParseStringWithDefaultValueErr(t *testing.T) {
	var errParseValue = []byte("\\")
	expectedValue := ""
	var expectedErr = jsonparser.MalformedValueError
	got, err := parseStringWithDefaultValue(errParseValue)
	if err != expectedErr {
		t.Fatalf("Error: expecting %v: \ngot %v  \ninstead", expectedErr, err)
	}
	if got != expectedValue {
		t.Fatalf("Error: expecting value: %s got %s instead", expectedValue, got)
	}
}

func TestParseStringWithDefaultValue(t *testing.T) {
	var ParseValue = "RU0009029557"
	expectedValue := ParseValue
	var expectedErr error = nil
	got, err := parseStringWithDefaultValue([]byte(ParseValue))
	if err != expectedErr {
		t.Fatalf("Error: expecting %v: \ngot %v  \ninstead", expectedErr, err)
	}
	if got != expectedValue {
		t.Fatalf("Error: expecting value: %s got %s instead", expectedValue, got)
	}
}

func TestParseStringWithDefaultValueByKeyBadString(t *testing.T) {
	var incomeJson = `
      {"market_name": null}
`
	expectedDefaultValue := "expectedDefaultValue"
	result, err := parseStringWithDefaultValueByKey([]byte(incomeJson), aggKeyMarketName, expectedDefaultValue)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}

	if got, expected := result, expectedDefaultValue; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseIntWithDefaultValueBadInt(t *testing.T) {
	var incomeJson = `
      {"market_name": "shares", "market_title": "Рынок акций", "engine": "stock", "tradedate": "2022-01-19", "secid": "SBERP", "value": 9833418828.24, "volume": 42115503.d, "numtrades": 144467, "updated_at": "2022-01-20 09:00:14"}
`
	aggregate := Aggregate{}
	if got, expected := parseAggregate([]byte(incomeJson), &aggregate), jsonparser.MalformedValueError; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestParseFloatWithDefaultValueBadFloat(t *testing.T) {
	var incomeJson = `
      {"market_name": "shares", "market_title": "Рынок акций", "engine": "stock", "tradedate": "2022-01-19", "secid": "SBERP", "value": 9833418828.24s, "volume": 42115503, "numtrades": 144467, "updated_at": "2022-01-20 09:00:14"}
`
	aggregate := Aggregate{}
	if got, expected := parseAggregate([]byte(incomeJson), &aggregate), jsonparser.MalformedValueError; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestGeneralIsOkSecurityParamEmptySecurity(t *testing.T) {
	if got, expected := isOkSecurityParam(""), false; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

func TestGeneralIsOkSecurityParamNonEmptySecurity(t *testing.T) {
	if got, expected := isOkSecurityParam("sber"), true; got != expected {
		t.Fatalf("Error: expecting: \n %v \ngot:\n %v \ninstead", expected, got)
	}
}

