package moexiss

import "testing"

func TestStatGetUrl(t *testing.T) {
	c := NewClient(nil)
	gotURL, err := c.Stats.getUrl(EngineStock, "shares", nil)
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v \ninstead", err)
	}
	if got, expected := gotURL, `https://iss.moex.com/iss/engines/stock/markets/shares/secstats.json?iss.json=extended&iss.meta=off`; got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}

func TestStatGetUrlBadEngine(t *testing.T) {
	c := NewClient(nil)
	_, err := c.Stats.getUrl(EngineUndefined, "shares", nil)
	if got, expected := err, ErrBadEngineParameter; got != expected {
		t.Fatalf("Error: expecting %v error: \ngot %v \ninstead", expected, got)
	}
}

func TestStatGetUrlBadMarket(t *testing.T) {
	c := NewClient(nil)
	_, err := c.Stats.getUrl(EngineStock, "", nil)
	if got, expected := err, ErrBadMarketParameter; got != expected {
		t.Fatalf("Error: expecting %v error: \ngot %v \ninstead", expected, got)
	}
}
