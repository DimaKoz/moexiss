package moexiss

import "net/url"

// IndicesRequestOptions contains options which can be used as arguments
// for building requests to get a list of indices that include the security.
// MoEx ISS API docs: https://iss.moex.com/iss/reference/160
type IndicesRequestOptions struct {
	lang Language  // `lang` query parameter in url.URL
}

// IndicesRequestOptionsBuilder represents a builder of IndicesRequestOptions struct
type IndicesRequestOptionsBuilder struct {
	options *IndicesRequestOptions
}

// NewIndicesReqOptionsBuilder is a constructor of IndicesRequestOptionsBuilder
func NewIndicesReqOptionsBuilder() *IndicesRequestOptionsBuilder {
	return &IndicesRequestOptionsBuilder{options: &IndicesRequestOptions{}}
}

// Build builds IndicesRequestOptions from IndicesRequestOptionsBuilder
func (b *IndicesRequestOptionsBuilder) Build() *IndicesRequestOptions {
	return b.options
}

// Lang sets 'lang' parameter to a request
func (b *IndicesRequestOptionsBuilder) Lang(lang Language) *IndicesRequestOptionsBuilder {
	b.options.lang = lang
	return b
}

// addIndicesRequestOptions sets parameters into *url.URL
// from IndicesRequestOptions struct and returns it back
func addIndicesRequestOptions(url *url.URL, options *IndicesRequestOptions) *url.URL {
	q := url.Query()
	q.Set("iss.meta", "off")
	q.Set("iss.json", "extended")
	if options == nil {
		url.RawQuery = q.Encode()
		return url
	}

	if options.lang != LangUndefined {
		q.Set("lang", options.lang.String())
	}

	url.RawQuery = q.Encode()
	return url
}
