package moexiss

import "testing"

func TestHistoryListingTradingStatus_String(t *testing.T) {
	if got, expected := ListingTradingStatusNotTraded.String(), "nottraded"; got != expected {
		t.Fatalf("Error: expecting `%s` \ngot `%s` \ninstead", expected, got)
	}
	if got, expected := ListingTradingStatusAll.String(), "all"; got != expected {
		t.Fatalf("Error: expecting `%s` \ngot `%s` \ninstead", expected, got)
	}
	if got, expected := ListingTradingStatusTraded.String(), "traded"; got != expected {
		t.Fatalf("Error: expecting `%s` \ngot `%s` \ninstead", expected, got)
	}
	if got, expected := ListingTradingStatusUndefined.String(), ""; got != expected {
		t.Fatalf("Error: expecting `%s` \ngot `%s` \ninstead", expected, got)
	}

}

func TestHistoryListingRequestOptionsBuilder_Build(t *testing.T) {
	expectStruct := HistoryListingRequestOptions{}
	bld := NewHistoryListingReqOptionsBuilder()

	if got, expected := *bld.Build(), expectStruct; got != expected {
		t.Fatalf("Error: expecting `%v` HistoryListingRequestOptions \ngot `%v` HistoryListingRequestOptions \ninstead", expected, got)
	}
}

func TestHistoryListingRequestOptionsBuilder_Lang(t *testing.T) {
	expectStruct := HistoryListingRequestOptions{lang: LangEn}
	bld := NewHistoryListingReqOptionsBuilder()
	bld.Lang(LangEn)
	if got, expected := *bld.Build(), expectStruct; got != expected {
		t.Fatalf("Error: expecting `%v` \ngot `%v` \ninstead", expected, got)
	}
}

func TestHistoryListingRequestOptionsBuilder_Start(t *testing.T) {
	expectStruct := HistoryListingRequestOptions{start: 42}
	bld := NewHistoryListingReqOptionsBuilder()
	bld.Start(42)
	if got, expected := *bld.Build(), expectStruct; got != expected {
		t.Fatalf("Error: expecting `%v` \ngot `%v` \ninstead", expected, got)
	}
}

func TestHistoryListingRequestOptionsBuilder_Status(t *testing.T) {
	expectStruct := HistoryListingRequestOptions{status: ListingTradingStatusTraded}
	bld := NewHistoryListingReqOptionsBuilder()
	bld.Status(ListingTradingStatusTraded)
	if got, expected := *bld.Build(), expectStruct; got != expected {
		t.Fatalf("Error: expecting `%v` \ngot `%v` \ninstead", expected, got)
	}
}

func TestAddHistoryListingRequestOptionsNil(t *testing.T) {
	var income *HistoryListingRequestOptions = nil
	c := NewClient(nil)
	url, _ := c.BaseURL.Parse("test.json")
	gotUrl := addHistoryListingRequestOptions(url, income)

	expected := `https://iss.moex.com/iss/test.json?iss.json=extended&iss.meta=off`
	if got := gotUrl.String(); got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}

func TestAddHistoryListingRequestOptions(t *testing.T) {
	var income = NewHistoryListingReqOptionsBuilder().
		Status(ListingTradingStatusNotTraded).
		Start(42).
		Lang(LangEn).
		Build()

	c := NewClient(nil)
	url, _ := c.BaseURL.Parse("test.json")
	gotUrl := addHistoryListingRequestOptions(url, income)

	expected := `https://iss.moex.com/iss/test.json?iss.json=extended&iss.meta=off&lang=en&start=42&status=nottraded`
	if got := gotUrl.String(); got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}
