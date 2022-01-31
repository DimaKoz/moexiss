package moexiss

import (
	"bufio"
	"bytes"
	"context"
	"github.com/buger/jsonparser"
	"path"
)

//Indices struct represents a list of the indices that include the security
type Indices struct {
	IndexId   string // "SECID"
	IndexName string // "SHORTNAME"
	From      string // "FROM"
	Till      string // "TILL"
}

//IndicesResponse struct represents a response with the list of the indices
type IndicesResponse struct {
	SecurityId string
	Indices    []Indices
}

const (
	indicesPartsUrl = "indices.json"

	indicesKeyId      = "SECID"
	indicesKeyName    = "SHORTNAME"
	indicesKeyFrom    = "FROM"
	indicesKeyTill    = "TILL"
	indicesKeyIndices = "indices"
)

// IndicesService gets a list of the indices that include the security
// from the MoEx ISS API.
//
// MoEx ISS API docs: https://iss.moex.com/iss/reference/160
type IndicesService service

//Aggregates provides Aggregates of MoEx ISS
func (i *IndicesService) Indices(ctx context.Context, security string, opt *IndicesRequestOptions) (*IndicesResponse, error) {

	url, err := i.getUrl(security, opt)
	if err != nil {
		return nil, err
	}
	req, err := i.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	_, err = i.client.Do(ctx, req, w)
	if err != nil {
		return nil, err
	}
	ir := IndicesResponse{}
	err = parseIndicesResponse(b.Bytes(), &ir)
	if err != nil {
		return nil, err
	}
	ir.SecurityId = security
	return &ir, nil
}

// getUrl provides an url for a request of indices with parameters from IndicesRequestOptions
// opt *IndicesRequestOptions can be nil, it is safe
// 'security' parameter must not be empty otherwise getUrl returns errBadSecurityParameter
func (i *IndicesService) getUrl(security string, opt *IndicesRequestOptions) (string, error) {
	if !isOkSecurityParam(security) {
		return "", errBadSecurityParameter
	}
	url, _ := i.client.BaseURL.Parse("securities")

	url.Path = path.Join(url.Path, security, indicesPartsUrl)
	gotUrl := addIndicesRequestOptions(url, opt)
	return gotUrl.String(), nil
}

func parseIndicesResponse(byteData []byte, indicesResponse *IndicesResponse) error {
	var err error
	if indicesResponse == nil {
		err = errNilPointer
		return err
	}
	var errInCb error
	_, err = jsonparser.ArrayEach(byteData, func(indicesBytes []byte, _ jsonparser.ValueType, offset int, errCb error) {
		var data []byte
		var dataType jsonparser.ValueType
		data, dataType, _, errInCb = jsonparser.Get(indicesBytes, indicesKeyIndices)
		if errInCb == nil && data != nil && dataType == jsonparser.Array {
			errInCb = parseIndices(data, &indicesResponse.Indices)
			if errInCb != nil {
				return
			}
		}
	})
	if err == nil && errInCb != nil {
		err = errInCb
	}
	return err
}

func parseIndices(data []byte, i *[]Indices) (err error) {

	var errInCb error
	_, err = jsonparser.ArrayEach(data, func(aggregateItemData []byte, dataType jsonparser.ValueType, offset int, errCb error) {
		if errInCb != nil {
			return
		}
		if dataType != jsonparser.Object {
			errInCb = errUnexpectedDataType
			return
		}

		indices := Indices{}
		errInCb = parseIndicesItem(aggregateItemData, &indices)
		if errInCb != nil {
			return
		}
		*i = append(*i, indices)

	})
	if err == nil && errInCb != nil {
		err = errInCb
	}
	return
}

func parseIndicesItem(data []byte, i *Indices) (err error) {

	id, err := parseStringWithDefaultValueByKey(data, indicesKeyId, "")
	if err != nil {
		return
	}

	name, err := parseStringWithDefaultValueByKey(data, indicesKeyName, "")
	if err != nil {
		return
	}

	from, err := parseStringWithDefaultValueByKey(data, indicesKeyFrom, "")
	if err != nil {
		return
	}

	till, err := parseStringWithDefaultValueByKey(data, indicesKeyTill, "")
	if err != nil {
		return
	}

	i.IndexId = id
	i.IndexName = name
	i.From = from
	i.Till = till

	return
}
