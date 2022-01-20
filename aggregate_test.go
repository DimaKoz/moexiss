package moexiss

import "testing"

func TestAggregatesGetUrl(t *testing.T) {
	var income *AggregateRequestOptions = nil
	c := NewClient(nil)

	if got, expected := c.Aggregates.getUrl("sberp", income), `https://iss.moex.com/iss/securities/sberp/aggregates.json?iss.json=extended&iss.meta=off`; got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}

