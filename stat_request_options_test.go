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

func compareStatRequestOptions(a, b StatRequestOptions) bool {
	if a.TradingSessionType != b.TradingSessionType {
		return false
	}
	if len(a.TickerIds) != len(b.TickerIds) || len(a.BoardId) != len(b.BoardId) {
		return false
	}
	for i, v := range a.TickerIds {
		if b.TickerIds[i] != v {
			return false
		}
	}
	for k, v := range a.BoardId {
		if b.BoardId[k] != v {
			return false
		}
	}
	return true
}

func TestStatReqOptionsBuilder_Build(t *testing.T) {
	expectStruct := StatRequestOptions{}
	bld := NewStatReqOptionsBuilder()

	if got, expected := *bld.Build(), expectStruct; !compareStatRequestOptions(expected, got) {
		t.Fatalf("Error: expecting `%v` StatRequestOptions \ngot `%v` StatRequestOptions \ninstead", expected, got)
	}
}
