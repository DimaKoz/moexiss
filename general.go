package moexiss

import "github.com/buger/jsonparser"

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
