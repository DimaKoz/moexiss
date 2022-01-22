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

