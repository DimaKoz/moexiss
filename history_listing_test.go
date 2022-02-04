package moexiss

import "testing"

func TestHistoryListingGetUrl(t *testing.T) {
	c := NewClient(nil)
	gotUrl, err := c.HistoryListing.getUrlListing(EngineStock, "shares", nil)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := gotUrl, `https://iss.moex.com/iss/history/engines/stock/markets/shares/listing.json?iss.json=extended&iss.meta=off`; got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}
