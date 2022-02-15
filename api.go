package moexiss

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"strings"
)

// Global constants.
const (
	libraryName    = "MoEx ISS"
	libraryVersion = "v0.0.1"
	defaultBaseURL = "https://iss.moex.com/iss/"
)

// User Agent should always following the below style.
//
//       MoExIss (OS; ARCH) LIB/VER APP/VER
const (
	libraryUserAgentPrefix = "MoExIss (" + runtime.GOOS + "; " + runtime.GOARCH + ") "
	libraryUserAgent       = libraryUserAgentPrefix + libraryName + "/" + libraryVersion
)

// Response represent a response of BareDo an Do functions
type Response struct {
	*http.Response
}

type service struct {
	client *Client
}

// Client structure represents a client of MoEx ISS API
type Client struct {
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// User agent used when communicating with the MoEx Iss API.
	UserAgent string

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	Securities     *SecuritiesService
	Index          *IndexService
	Turnovers      *TurnoverService
	Aggregates     *AggregateService
	Indices        *IndicesService
	HistoryListing *HistoryListingService
	Stats          *StatsService
}

// NewClient creates an instance of Client
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: libraryUserAgent}
	c.common.client = c

	c.Securities = (*SecuritiesService)(&c.common)
	c.Index = (*IndexService)(&c.common)
	c.Turnovers = (*TurnoverService)(&c.common)
	c.Aggregates = (*AggregateService)(&c.common)
	c.Indices = (*IndicesService)(&c.common)
	c.HistoryListing = (*HistoryListingService)(&c.common)
	c.Stats = (*StatsService)(&c.common)
	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer interface,
// the raw response body will be written to v, without attempting to first
// decode it. If v is nil, and no error happens, the response is returned as is.
//
// The provided ctx must be non-nil, if it is nil an error is returned. If it
// is canceled or times out, ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.BareDo(ctx, req)
	if err != nil {
		return resp, err
	}

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		var b []byte
		b, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		decErr := json.Unmarshal(b, &v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}
	if err != nil {
		return nil, err
	}
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return resp, err
}

// BareDo sends an API request and lets you handle the api response. If an error
// or API Error occurs, the error will contain more information. Otherwise you
// are supposed to read and close the response's Body.
//
// The provided ctx must be non-nil, if it is nil an error is returned. If it is
// canceled or times out, ctx.Err() will be returned.
func (c *Client) BareDo(ctx context.Context, req *http.Request) (*Response, error) {
	if ctx == nil {
		return nil, ErrNonNilContext
	}

	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// returning *url.Error.
		if e, ok := err.(*url.Error); ok {
			return nil, e
		}

		return nil, err

	}

	response := &Response{resp}

	err = CheckResponse(resp)
	if err != nil {
		clErr := resp.Body.Close()
		if clErr != nil {
			return nil, fmt.Errorf("got some errors: \n%s \nand \n%s", err.Error(), clErr.Error())
		}
		return nil, err
	}
	return response, err
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	return fmt.Errorf("status:[%d] %s", r.StatusCode, r.Status)
}
