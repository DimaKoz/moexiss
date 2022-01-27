package moexiss

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
