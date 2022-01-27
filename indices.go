package moexiss

import "path"

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

	indicesKeyId   = "SECID"
	indicesKeyName = "SHORTNAME"
	indicesKeyFrom = "FROM"
	indicesKeyTill = "TILL"
)

// IndicesService gets a list of the indices that include the security
// from the MoEx ISS API.
//
// MoEx ISS API docs: https://iss.moex.com/iss/reference/160
type IndicesService service

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
