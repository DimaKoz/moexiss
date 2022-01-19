package moexiss

import (
	"testing"
	"time"
)

func TestAggregatesReqOptionsBuilder_Build(t *testing.T) {
	expectStruct := AggregatesRequestOptions{}
	bld := NewAggregatesReqOptionsBuilder()

	if got, expected := *bld.Build(), expectStruct; got != expected {
		t.Fatalf("Error: expecting `%v` AggregatesRequestOptions \ngot `%v` AggregatesRequestOptions \ninstead", expected, got)
	}
}

func TestAggregatesReqOptionsBuilder_Lang(t *testing.T) {
	expectStruct := AggregatesRequestOptions{lang: LangEn}
	bld := NewAggregatesReqOptionsBuilder()
	bld.Lang(LangEn)
	if got, expected := *bld.Build(), expectStruct; got != expected {
		t.Fatalf("Error: expecting `%v` \ngot `%v` \ninstead", expected, got)
	}
}

func TestAggregatesReqOptionsBuilder_Date(t *testing.T) {
	date := time.Now()
	expectStruct := AggregatesRequestOptions{date: date}
	bld := NewAggregatesReqOptionsBuilder()
	bld.Date(date)
	if got, expected := *bld.Build(), expectStruct; got != expected {
		t.Fatalf("Error: expecting `%v` \ngot `%v` \ninstead", expected, got)
	}

}

func TestNewAggregatesReqOptionsBuilder(t *testing.T) {
	date := time.Now()
	expectStruct := AggregatesRequestOptions{
		lang:             LangRu,
		date:             date}
	bld := NewAggregatesReqOptionsBuilder()
	bld.Date(date).Lang(LangRu)
	if got, expected := *bld.Build(), expectStruct; got != expected {
		t.Fatalf("Error: expecting `%v` \ngot `%v` \ninstead", expected, got)
	}

}

func TestAddAggregatesRequestOptionsNilOptions(t *testing.T) {
	var income *AggregatesRequestOptions = nil
	c := NewClient(nil)
	url, _ := c.BaseURL.Parse("aggregates.json")
	gotUrl := addAggregatesRequestOptions(url, income)

	expected := `https://iss.moex.com/iss/aggregates.json?iss.json=extended&iss.meta=off`
	if got := gotUrl.String(); got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}

func TestAddAggregatesRequestOptions(t *testing.T) {
	var incomeOptions = NewAggregatesReqOptionsBuilder().
		Lang(LangEn).
		Date(time.Date(2021, 2, 24, 12, 0, 0, 0, time.UTC)).
		Build()
	c := NewClient(nil)
	url, _ := c.BaseURL.Parse("aggregates.json")
	gotUrl := addAggregatesRequestOptions(url, incomeOptions)

	expected := `https://iss.moex.com/iss/aggregates.json?date=2021-02-24&iss.json=extended&iss.meta=off&lang=en`
	if got := gotUrl.String(); got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}
