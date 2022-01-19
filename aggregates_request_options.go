package moexiss

import (
	"net/url"
	"time"
)


//AggregatesRequestOptions contains options which can be used as arguments
//for building requests to get aggregated trading results for the date by market.
//MoEx ISS API docs: https://iss.moex.com/iss/reference/214
type AggregatesRequestOptions struct {
	lang Language //`lang` query parameter in url.URL
	date time.Time //`date` query parameter in url.URL
}

//AggregatesReqOptionsBuilder represents a builder of AggregatesRequestOptions struct
type AggregatesReqOptionsBuilder struct {
	options *AggregatesRequestOptions
}

//NewAggregatesReqOptionsBuilder is a constructor of AggregatesReqOptionsBuilder
func NewAggregatesReqOptionsBuilder() *AggregatesReqOptionsBuilder {
	return &AggregatesReqOptionsBuilder{options: &AggregatesRequestOptions{}}
}

//Build builds AggregatesRequestOptions from AggregatesReqOptionsBuilder
func (b *AggregatesReqOptionsBuilder) Build() *AggregatesRequestOptions {
	return b.options
}

//Lang sets 'lang' parameter to a request
func (b *AggregatesReqOptionsBuilder) Lang(lang Language) *AggregatesReqOptionsBuilder {
	b.options.lang = lang
	return b
}

//Date sets 'date' parameter to a request
//'date' is the date for which you want to display the data.
//By default(if none), for the last date in the trading results.
func (b *AggregatesReqOptionsBuilder) Date(date time.Time) *AggregatesReqOptionsBuilder {
	b.options.date = date
	return b
}

//addAggregatesRequestOptions sets parameters into *url.URL
//from AggregatesRequestOptions struct and returns it back
func addAggregatesRequestOptions(url *url.URL, options *AggregatesRequestOptions) *url.URL {
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
