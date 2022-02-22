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

func TestGetTradingSession(t *testing.T) {
	type Case struct {
		income   string
		expected TradingSession
	}
	cases := []Case{
		{
			income:   "",
			expected: TradingSessionUndefined},
		{
			income:   "dhe",
			expected: TradingSessionUndefined},
		{
			income:   "0",
			expected: TradingSessionUndefined},
		{
			income:   "1",
			expected: TradingSessionMain},
		{
			income:   "2",
			expected: TradingSessionAdditional},
		{
			income:   "3",
			expected: TradingSessionTotal},
	}
	for _, c := range cases {
		if got := getTradingSession(c.income); got != c.expected {
			t.Fatalf("Error: expecting :\n`%s` \ngot \n`%s` \ninstead", c.expected, got)
		}
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

func TestAddStatRequestOptions(t *testing.T) {
	var incomeOptions = NewStatReqOptionsBuilder().
		TypeTradingSession(TradingSessionMain).
		AddBoard("TQBR").
		AddBoard("SMAL").
		AddTicker("SBERP").
		AddTicker("GAZP").
		AddTicker("DSKY").
		Build()
	c := NewClient(nil)
	url, _ := c.BaseURL.Parse("test.json")
	gotURL := addStatRequestOptions(url, incomeOptions)

	expected := `https://iss.moex.com/iss/test.json?boardid=TQBR%2CSMAL&iss.json=extended&iss.meta=off&securities=SBERP%2CGAZP%2CDSKY&tradingsession=1`
	if got := gotURL.String(); got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}

func TestAddStatRequestOptionsMore10(t *testing.T) {
	var incomeOptions = NewStatReqOptionsBuilder().
		TypeTradingSession(TradingSessionMain).
		AddTicker("").
		AddBoard("0000").
		AddBoard("0001").
		AddBoard("0002").
		AddBoard("0003").
		AddBoard("0004").
		AddBoard("0005").
		AddBoard("0006").
		AddBoard("0007").
		AddBoard("0008").
		AddBoard("0009").
		AddBoard("0010").
		Build()
	c := NewClient(nil)
	url, _ := c.BaseURL.Parse("test.json")
	gotURL := addStatRequestOptions(url, incomeOptions)

	expected := `https://iss.moex.com/iss/test.json?boardid=0000%2C0001%2C0002%2C0003%2C0004%2C0005%2C0006%2C0007%2C0008%2C0009&iss.json=extended&iss.meta=off&tradingsession=1`
	if got := gotURL.String(); got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}

func TestAddStatRequestOptionsNil(t *testing.T) {
	c := NewClient(nil)
	url, _ := c.BaseURL.Parse("test.json")
	gotURL := addStatRequestOptions(url, nil)

	expected := `https://iss.moex.com/iss/test.json?iss.json=extended&iss.meta=off`
	if got := gotURL.String(); got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}
