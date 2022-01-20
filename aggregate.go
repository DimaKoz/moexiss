package moexiss

import "path"


const (
	aggregatesPartsUrl = "aggregates.json"
)

// AggregateService gets aggregated trading results
// in the MoEx ISS API.
//
// MoEx ISS API docs: https://iss.moex.com/iss/reference/214
type AggregateService service


//getUrl provides an url for a request of the aggregates with parameters from AggregateRequestOptions
//opt *AggregateRequestOptions can be nil, it is safe
func (s *AggregateService) getUrl(security string, opt *AggregateRequestOptions) string {
	url, _ := s.client.BaseURL.Parse("securities")

	url.Path = path.Join(url.Path, security, aggregatesPartsUrl)
	gotUrl := addAggregateRequestOptions(url, opt)
	return gotUrl.String()
}
