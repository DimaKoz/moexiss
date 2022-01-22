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

