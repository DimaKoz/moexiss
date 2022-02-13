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

func TestStatReqOptionsBuilder_TypeTradingSession(t *testing.T) {
	expectStruct := StatRequestOptions{TradingSessionType: TradingSessionMain}
	bld := NewStatReqOptionsBuilder()
	bld.TypeTradingSession(TradingSessionMain)

	if got, expected := *bld.Build(), expectStruct; !compareStatRequestOptions(expected, got) {
		t.Fatalf("Error: expecting `%v` StatRequestOptions \ngot `%v` StatRequestOptions \ninstead", expected, got)
	}
}

func TestStatReqOptionsBuilder_AddTicker(t *testing.T) {
	expectStruct := StatRequestOptions{
		TickerIds: []string{
			"SBERP",
			"GAZP",
			"DSKY",
		},
	}
	bld := NewStatReqOptionsBuilder()
	bld.AddTicker("SBERP")
	bld.AddTicker("GAZP")
	bld.AddTicker("DSKY")

	if got, expected := *bld.Build(), expectStruct; !compareStatRequestOptions(expected, got) {
		t.Fatalf("Error: expecting `%v` StatRequestOptions \ngot `%v` StatRequestOptions \ninstead", expected, got)
	}
}

func TestStatReqOptionsBuilder_AddBoard(t *testing.T) {
	expectStruct := StatRequestOptions{
		BoardId: []string{
			"TQBR",
			"SMAL",
		},
	}
	bld := NewStatReqOptionsBuilder()
	bld.AddBoard("TQBR")
	bld.AddBoard("SMAL")

	if got, expected := *bld.Build(), expectStruct; !compareStatRequestOptions(expected, got) {
		t.Fatalf("Error: expecting `%v` StatRequestOptions \ngot `%v` StatRequestOptions \ninstead", expected, got)
	}

}

func TestStatReqOptionsBuilder(t *testing.T) {
	expectStruct := StatRequestOptions{
		TradingSessionType: TradingSessionAdditional,
		BoardId: []string{
			"TQBR",
			"SMAL",
		},
		TickerIds: []string{
			"SBERP",
			"GAZP",
			"DSKY",
		},
	}
	bld := NewStatReqOptionsBuilder()
	bld.
		TypeTradingSession(TradingSessionAdditional).
		AddBoard("TQBR").
		AddBoard("SMAL").
		AddTicker("SBERP").
		AddTicker("GAZP").
		AddTicker("DSKY")

	if got, expected := *bld.Build(), expectStruct; !compareStatRequestOptions(expected, got) {
		t.Fatalf("Error: expecting `%v` StatRequestOptions \ngot `%v` StatRequestOptions \ninstead", expected, got)
	}

}
