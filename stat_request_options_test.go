package moexiss

import "testing"

func TestTradingSession_String(t *testing.T) {
	if got, expected := TradingSessionUndefined.String(), "0"; got != expected {
		t.Fatalf("Error: expecting `%s` TradingSession \ngot `%s` TradingSession \ninstead", expected, got)
	}
	if got, expected := TradingSessionMain.String(), "1"; got != expected {
		t.Fatalf("Error: expecting `%s` TradingSession \ngot `%s` TradingSession \ninstead", expected, got)
	}
	if got, expected := TradingSessionAdditional.String(), "2"; got != expected {
		t.Fatalf("Error: expecting `%s` TradingSession \ngot `%s` TradingSession \ninstead", expected, got)
	}
	if got, expected := TradingSessionTotal.String(), "3"; got != expected {
		t.Fatalf("Error: expecting `%s` TradingSession \ngot `%s` TradingSession \ninstead", expected, got)
	}
}
