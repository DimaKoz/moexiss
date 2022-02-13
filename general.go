package moexiss

import (
	"errors"
	"github.com/buger/jsonparser"
	"unicode/utf8"
)

var (
	// General errors
	ErrNonNilContext          = errors.New("context must be non-nil")
	ErrUnexpectedDataType     = errors.New("unexpected data type")
	ErrNilPointer             = errors.New("nil pointer error")
	ErrBadSecurityParameter   = errors.New("bad 'security' parameter")
	ErrBadEngineParameter     = errors.New("bad 'engine' parameter")
	ErrBadMarketParameter     = errors.New("bad 'market' parameter")
	ErrBadBoardParameter      = errors.New("bad 'board' parameter")
	ErrBadBoardGroupParameter = errors.New("bad 'boardgroup' parameter")
	ErrEmptyServerResult      = errors.New("the empty answer")
)

func parseStringWithDefaultValue(fieldValue []byte) (string, error) {
	res, err := jsonparser.ParseString(fieldValue)
	if err != nil {
		return "", err
	}
	if res != "null" {
		return res, nil
	}
	return "", nil
}

func parseStringWithDefaultValueByKey(fieldValue []byte, key string, defaultValue string) (string, error) {
	valueData, _, _, err := jsonparser.Get(fieldValue, key)
	if string(valueData) == "null" {
		return defaultValue, nil
	}
	if err != nil {
		return defaultValue, err
	}
	value, err := jsonparser.ParseString(valueData)
	if err != nil {
		return defaultValue, err
	}
	return value, nil
}

func parseFloatWithDefaultValue(fieldValue []byte, key string) (float64, error) {
	valueData, _, _, err := jsonparser.Get(fieldValue, key)
	if string(valueData) == "null" {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	value, err := jsonparser.ParseFloat(valueData)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func parseIntWithDefaultValue(fieldValue []byte, key string) (int64, error) {
	valueData, _, _, err := jsonparser.Get(fieldValue, key)
	if string(valueData) == "null" {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	value, err := jsonparser.ParseInt(valueData)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func isOkSecurityParam(securityId string) bool {
	if securityId == "" {
		return false
	}
	minLen := 3
	return utf8.RuneCountInString(securityId) >= minLen
}
