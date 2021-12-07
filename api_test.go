package moexiss

import (
	"errors"
	"net/http"
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

	if got, expected := CheckResponse(&resp), errors.New("status:[404] Not found"); got.Error() != expected.Error() {
		t.Fatalf("Error: expecting error with: %s \ngot %s  \ninstead", expected.Error(),  got.Error())
	}

}


func TestCheckResponse(t *testing.T) {
	resp := http.Response{StatusCode: 200, Status: "OK"}

	var nilErr error = nil
	if got, expected := CheckResponse(&resp), nilErr; got != expected {
		t.Fatalf("Error: expecting %v: \ngot %v  \ninstead", expected, got)
	}

}