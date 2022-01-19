package moexiss

import (
	"testing"
	"time"
)

func TestAggregateReqOptionsBuilder_Build(t *testing.T) {
	expectStruct := AggregateRequestOptions{}
	bld := NewAggregateReqOptionsBuilder()

	if got, expected := *bld.Build(), expectStruct; got != expected {
		t.Fatalf("Error: expecting `%v` AggregateRequestOptions \ngot `%v` AggregateRequestOptions \ninstead", expected, got)
	}
}

func TestAggregateReqOptionsBuilder_Lang(t *testing.T) {
	expectStruct := AggregateRequestOptions{lang: LangEn}
	bld := NewAggregateReqOptionsBuilder()
	bld.Lang(LangEn)
	if got, expected := *bld.Build(), expectStruct; got != expected {
		t.Fatalf("Error: expecting `%v` \ngot `%v` \ninstead", expected, got)
	}
}

func TestAggregateReqOptionsBuilder_Date(t *testing.T) {
	date := time.Now()
	expectStruct := AggregateRequestOptions{date: date}
	bld := NewAggregateReqOptionsBuilder()
	bld.Date(date)
	if got, expected := *bld.Build(), expectStruct; got != expected {
		t.Fatalf("Error: expecting `%v` \ngot `%v` \ninstead", expected, got)
	}

}

func TestNewAggregateReqOptionsBuilder(t *testing.T) {
	date := time.Now()
	expectStruct := AggregateRequestOptions{
		lang:             LangRu,
		date:             date}
	bld := NewAggregateReqOptionsBuilder()
	bld.Date(date).Lang(LangRu)
	if got, expected := *bld.Build(), expectStruct; got != expected {
		t.Fatalf("Error: expecting `%v` \ngot `%v` \ninstead", expected, got)
	}

}

func TestAddAggregateRequestOptionsNilOptions(t *testing.T) {
	var income *AggregateRequestOptions = nil
	c := NewClient(nil)
	url, _ := c.BaseURL.Parse("aggregates.json")
	gotUrl := addAggregateRequestOptions(url, income)

	expected := `https://iss.moex.com/iss/aggregates.json?iss.json=extended&iss.meta=off`
	if got := gotUrl.String(); got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}

func TestAddAggregateRequestOptions(t *testing.T) {
	var incomeOptions = NewAggregateReqOptionsBuilder().
		Lang(LangEn).
		Date(time.Date(2021, 2, 24, 12, 0, 0, 0, time.UTC)).
		Build()
	c := NewClient(nil)
	url, _ := c.BaseURL.Parse("aggregates.json")
	gotUrl := addAggregateRequestOptions(url, incomeOptions)

	expected := `https://iss.moex.com/iss/aggregates.json?date=2021-02-24&iss.json=extended&iss.meta=off&lang=en`
	if got := gotUrl.String(); got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}
