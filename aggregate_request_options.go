package moexiss

import (
	"net/url"
	"time"
)

//AggregateRequestOptions contains options which can be used as arguments
//for building requests to get aggregated trading results for the date by market.
//MoEx ISS API docs: https://iss.moex.com/iss/reference/214
type AggregateRequestOptions struct {
	lang Language  // `lang` query parameter in url.URL
	date time.Time // `date` query parameter in url.URL
}

//AggregateReqOptionsBuilder represents a builder of AggregateRequestOptions struct
type AggregateReqOptionsBuilder struct {
	options *AggregateRequestOptions
}

//NewAggregateReqOptionsBuilder is a constructor of AggregateReqOptionsBuilder
func NewAggregateReqOptionsBuilder() *AggregateReqOptionsBuilder {
	return &AggregateReqOptionsBuilder{options: &AggregateRequestOptions{}}
}

//Build builds AggregateRequestOptions from AggregateReqOptionsBuilder
func (b *AggregateReqOptionsBuilder) Build() *AggregateRequestOptions {
	return b.options
}

//Lang sets 'lang' parameter to a request
func (b *AggregateReqOptionsBuilder) Lang(lang Language) *AggregateReqOptionsBuilder {
	b.options.lang = lang
	return b
}

//Date sets 'date' parameter to a request
//'date' is the date for which you want to display the data.
//By default(if none), for the last date in the trading results.
func (b *AggregateReqOptionsBuilder) Date(date time.Time) *AggregateReqOptionsBuilder {
	b.options.date = date
	return b
}

//addAggregateRequestOptions sets parameters into *url.URL
//from AggregateRequestOptions struct and returns it back
func addAggregateRequestOptions(url *url.URL, options *AggregateRequestOptions) *url.URL {
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
	if !options.date.IsZero() {
		q.Set("date", options.date.Format("2006-01-02"))
	}

	url.RawQuery = q.Encode()
	return url
}
