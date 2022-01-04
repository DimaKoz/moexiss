package moexiss

import (
	"errors"
	"math"
	"net/http"
	"net/url"
	"testing"
)

func TestNewClientDefaultHttpClient(t *testing.T) {

	result := NewClient(nil)

	if result == nil {
		t.Fatalf("Error: expecting non-nil result: got nil instead")
	}

	if got, expected := result.BaseURL.String(), defaultBaseURL; got != expected {
		t.Fatalf("Error: expecting BaseURL: %s got %s instead", expected, got)
	}

	if got, expected := result.UserAgent, libraryUserAgent; got != expected {
		t.Fatalf("Error: expecting UserAgent: %s got %s instead", expected, got)
	}

}

func TestNewClientCustomHttpClient(t *testing.T) {

	result := NewClient(nil)

	if result == nil {
		t.Fatalf("Error: expecting non-nil result: got nil instead")
	}

	c := http.Client{}
	if got, expected := NewClient(&c).client, &c; got != expected {
		t.Fatalf("Error: expecting http.Client: \nAddress %p %v \ngot \nAddress %p %v  \ninstead", expected, expected, got, got)
	}

}

func TestCheckResponseError(t *testing.T) {
	resp := http.Response{StatusCode: 404, Status: "Not found"}

	if got, expected := CheckResponse(&resp), errors.New("status:[404] Not found"); got == nil || got.Error() != expected.Error() {
		t.Fatalf("Error: expecting error with: %s \ngot %v  \ninstead", expected.Error(), got)
	}

}

func TestCheckResponse(t *testing.T) {
	resp := http.Response{StatusCode: 200, Status: "OK"}

	var nilErr error = nil
	if got, expected := CheckResponse(&resp), nilErr; got != expected {
		t.Fatalf("Error: expecting %v: \ngot %v  \ninstead", expected, got)
	}

}

func TestNewRequestTrailingSlash(t *testing.T) {

	baseURL, _ := url.Parse("http://exmple.com")
	c := Client{BaseURL: baseURL}
	var nilReq *http.Request = nil
	var expectedErrStr = "BaseURL must have a trailing slash, but \"http://exmple.com\" does not"
	gotReq, gotErr := c.NewRequest("", "", nil)

	if got, expected := gotReq, nilReq; got != expected {
		t.Fatalf("Error: expecting %v: \ngot %v  \ninstead", expected, got)
		return
	}

	if gotErr == nil {
		t.Fatalf("Error: expecting error with %s: \ngot <nil> instead", expectedErrStr)
		return
	}

	if expectedErrStr != gotErr.Error() {
		t.Fatalf("Error: expecting error with %s: \ngot %s \ninstead", expectedErrStr, gotErr.Error())
	}

}

func TestNewRequestBadUrl(t *testing.T) {
	badUrl := ":badUrl"
	c := NewClient(nil)
	var nilReq *http.Request = nil
	var expectedErrStr = "parse \":badUrl\": missing protocol scheme"
	gotReq, gotErr := c.NewRequest("", badUrl, nil)

	if got, expected := gotReq, nilReq; got != expected {
		t.Fatalf("Error: expecting %v: \ngot %v  \ninstead", expected, got)
		return
	}

	if gotErr == nil {
		t.Fatalf("Error: expecting error with %s: \ngot <nil> instead", expectedErrStr)
		return
	}

	if expectedErrStr != gotErr.Error() {
		t.Fatalf("Error: expecting error with %s: \ngot %s \ninstead", expectedErrStr, gotErr.Error())
	}

}

func TestNewRequestBadBody(t *testing.T) {

	c := NewClient(nil)
	var nilReq *http.Request = nil
	var expectedErrStr = "json: unsupported value: +Inf"
	gotReq, gotErr := c.NewRequest("", "", math.Inf(1))

	if got, expected := gotReq, nilReq; got != expected {
		t.Fatalf("Error: expecting %v: \ngot %v  \ninstead", expected, got)
		return
	}

	if gotErr == nil {
		t.Fatalf("Error: expecting error with %s: \ngot <nil> instead", expectedErrStr)
		return
	}

	if expectedErrStr != gotErr.Error() {
		t.Fatalf("Error: expecting error with %s: \ngot %s \ninstead", expectedErrStr, gotErr.Error())
	}

}

func TestNewRequestBadMethod(t *testing.T) {

	c := NewClient(nil)
	var nilReq *http.Request = nil
	var expectedErrStr = "net/http: invalid method \"BadMethod/\""
	gotReq, gotErr := c.NewRequest("BadMethod/", "", nil)

	if got, expected := gotReq, nilReq; got != expected {
		t.Fatalf("Error: expecting %v: \ngot %v  \ninstead", expected, got)
		return
	}

	if gotErr == nil {
		t.Fatalf("Error: expecting error with %s: \ngot <nil> instead", expectedErrStr)
		return
	}

	if expectedErrStr != gotErr.Error() {
		t.Fatalf("Error: expecting error with %s: \ngot %s \ninstead", expectedErrStr, gotErr.Error())
	}

}

func TestNewRequest(t *testing.T) {

	c := NewClient(nil)

	gotReq, gotErr := c.NewRequest("GET", "", 2)

	if gotReq == nil {
		t.Fatalf("Error: expecting non-nil http.Request: got nil instead")
		return
	}

	if gotErr != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v instead", gotErr)
		return
	}

	if got, expected := gotReq.Header.Get("User-Agent"), libraryUserAgent; got != expected {
		t.Fatalf("Error: expected User-Agent %s: \ngot %s \ninstead", expected, got)
	}

	if got, expected := gotReq.Header.Get("Content-Type"), "application/json"; got != expected {
		t.Fatalf("Error: expected Content-Type %s: \ngot %s \ninstead", expected, got)
	}

}
