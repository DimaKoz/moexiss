package moexiss

import (
	"net/url"
	"strconv"
)

// HistoryListingTradingStatus represents a type of trading status for listing of MoEx ISS API
type HistoryListingTradingStatus string

const (
	ListingTradingStatusUndefined HistoryListingTradingStatus = ""
	ListingTradingStatusAll       HistoryListingTradingStatus = "all"
	ListingTradingStatusNotTraded HistoryListingTradingStatus = "nottraded"
	ListingTradingStatusTraded    HistoryListingTradingStatus = "traded"
)

// String representations of HistoryListingTradingStatus values
func (hlt HistoryListingTradingStatus) String() string {
	return string(hlt)
}

// HistoryListingRequestOptions contains options which can be used as arguments
// for building requests to get listing information.
// MoEx ISS API docs:
//
// https://iss.moex.com/iss/reference/118
// https://iss.moex.com/iss/reference/119
// https://iss.moex.com/iss/reference/120
type HistoryListingRequestOptions struct {
	lang   Language                    // `lang` query parameter in url.URL
	start  uint64                      // `start` query parameter in url.URL
	status HistoryListingTradingStatus // `status` query parameter in url.URL
}

// HistoryListingRequestOptionsBuilder represents a builder of HistoryListingRequestOptions struct
type HistoryListingRequestOptionsBuilder struct {
	options *HistoryListingRequestOptions
}

// NewHistoryListingReqOptionsBuilder is a constructor of HistoryListingRequestOptionsBuilder
func NewHistoryListingReqOptionsBuilder() *HistoryListingRequestOptionsBuilder {
	return &HistoryListingRequestOptionsBuilder{options: &HistoryListingRequestOptions{}}
}

// Build builds HistoryListingRequestOptions from HistoryListingRequestOptionsBuilder
func (b *HistoryListingRequestOptionsBuilder) Build() *HistoryListingRequestOptions {
	return b.options
}

// Lang sets 'lang' parameter to a request
// Language of the result set: 'ru' or 'en'
// 'ru' by default
func (b *HistoryListingRequestOptionsBuilder) Lang(lang Language) *HistoryListingRequestOptionsBuilder {
	b.options.lang = lang
	return b
}

// Start sets 'start' parameter to a request
// Row number (the number of the first row is 0) to begin the result set with.
// If the result set contains no data then the specified number is greater than
// the total number of rows available.
// 0 by default
func (b *HistoryListingRequestOptionsBuilder) Start(start uint64) *HistoryListingRequestOptionsBuilder {
	b.options.start = start
	return b
}

// Status sets 'status' parameter to a request
// Trading status filter: ListingTradingStatusAll, ListingTradingStatusNotTraded or ListingTradingStatusTraded
// ListingTradingStatusAll status by default
func (b *HistoryListingRequestOptionsBuilder) Status(status HistoryListingTradingStatus) *HistoryListingRequestOptionsBuilder {
	b.options.status = status
	return b
}

// addHistoryListingRequestOptions sets parameters into *url.URL
// from HistoryListingRequestOptions struct and returns it back
func addHistoryListingRequestOptions(url *url.URL, options *HistoryListingRequestOptions) *url.URL {
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

	if options.start != 0 {
		q.Set("start", strconv.FormatUint(options.start, 10))
	}

	if options.status != ListingTradingStatusUndefined && options.status != ListingTradingStatusAll {
		q.Set("status", options.status.String())
	}

	url.RawQuery = q.Encode()
	return url
}
