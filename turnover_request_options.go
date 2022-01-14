package moexiss

import (
	"net/url"
	"time"
)

//TurnoverRequestOptions contains options which can be used as arguments
//for building requests to get current turnover on all the markets.
//MoEx ISS API docs: https://iss.moex.com/iss/reference/24
type TurnoverRequestOptions struct {
	lang Language //`lang` query parameter in url.URL

	isTonightSession bool //`is_tonight_session` query parameter in url.URL

	date time.Time //`date` query parameter in url.URL

}

//TurnoverReqOptionsBuilder represents a builder of TurnoverRequestOptions struct
type TurnoverReqOptionsBuilder struct {
	options *TurnoverRequestOptions
}

//NewTurnoverReqOptionsBuilder is a constructor for TurnoverReqOptionsBuilder
func NewTurnoverReqOptionsBuilder() *TurnoverReqOptionsBuilder {
	return &TurnoverReqOptionsBuilder{options: &TurnoverRequestOptions{}}
}

//Build builds TurnoverRequestOptions from TurnoverReqOptionsBuilder
func (b *TurnoverReqOptionsBuilder) Build() *TurnoverRequestOptions {
	return b.options
}

//Lang sets 'lang' parameter to a request of the current turnover
func (b *TurnoverReqOptionsBuilder) Lang(lang Language) *TurnoverReqOptionsBuilder {
	b.options.lang = lang
	return b
}

//IsTonightSession sets 'is_tonight_session' parameter to a request of the current turnover
func (b *TurnoverReqOptionsBuilder) IsTonightSession(isTonightSession bool) *TurnoverReqOptionsBuilder {
	b.options.isTonightSession = isTonightSession
	return b
}

//Date sets 'date' parameter to a request of the current turnover
func (b *TurnoverReqOptionsBuilder) Date(date time.Time) *TurnoverReqOptionsBuilder {
	b.options.date = date
	return b
}

//addTurnoverRequestOptions sets parameters into *url.URL
//from TurnoverRequestOptions struct and returns it back
func addTurnoverRequestOptions(url *url.URL, options *TurnoverRequestOptions, onlyBlock string) *url.URL {
	q := url.Query()
	q.Set("iss.meta", "off")
	q.Set("iss.only", onlyBlock)
	q.Set("iss.json", "extended")
	if options == nil {
		url.RawQuery = q.Encode()
		return url
	}

	if options.lang != LangUndefined {
		q.Set("lang", options.lang.String())
	}
	if options.isTonightSession {
		q.Set("is_tonight_session", "1")
	}

	if !options.date.IsZero() {
		q.Set("date", options.date.Format("2006-01-02"))
	}

	url.RawQuery = q.Encode()
	return url
}
