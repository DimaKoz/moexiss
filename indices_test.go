package moexiss

import "testing"

func TestIndicesGetUrl(t *testing.T) {
	var income *IndicesRequestOptions = nil
	c := NewClient(nil)
	url, err := c.Indices.getUrl("sberp", income)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := url, `https://iss.moex.com/iss/securities/sberp/indices.json?iss.json=extended&iss.meta=off`; got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}

func TestIndicesGetUrlBadSecurity(t *testing.T) {
	var income *IndicesRequestOptions = nil
	c := NewClient(nil)
	_, err := c.Indices.getUrl("", income)
	if err != errBadSecurityParameter {
		t.Fatalf("Error: expecting error: %v \ngot %v \ninstead", errBadSecurityParameter, err)
	}
}
