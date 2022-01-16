package moexiss

import (
	"testing"
	"time"
)

func TestTurnoverReqOptionsBuilder_Build(t *testing.T) {
	expectStruct := &TurnoverRequestOptions{}
	bld := NewTurnoverReqOptionsBuilder()

	if got, expected := bld.Build(), expectStruct; *got != *expected {
		t.Fatalf("Error: expecting `%v` TurnoverRequestOptions \ngot `%v` TurnoverRequestOptions \ninstead", expected, got)
	}
}

func TestTurnoverReqOptionsBuilder_Lang(t *testing.T) {
	expectStruct := &TurnoverRequestOptions{lang: LangEn}
	bld := NewTurnoverReqOptionsBuilder()
	bld.Lang(LangEn)
	if got, expected := bld.Build(), expectStruct; got == nil || *got != *expected {
		t.Fatalf("Error: expecting `%v` \ngot `%v` \ninstead", expected, got)
	}
}

func TestTurnoverReqOptionsBuilder_IsTonightSession(t *testing.T) {
	expectStruct := TurnoverRequestOptions{isTonightSession: true}
	bld := NewTurnoverReqOptionsBuilder()
	bld.IsTonightSession(true)
	if got, expected := bld.Build(), expectStruct; got == nil || *got != expected {
		t.Fatalf("Error: expecting `%v` \ngot `%v` \ninstead", expected, got)
	}
}

func TestTurnoverReqOptionsBuilder_Date(t *testing.T) {
	date := time.Now()
	expectStruct := TurnoverRequestOptions{date: date}
	bld := NewTurnoverReqOptionsBuilder()
	bld.Date(date)
	if got, expected := bld.Build(), expectStruct; got == nil || *got != expected {
		t.Fatalf("Error: expecting `%v` \ngot `%v` \ninstead", expected, got)
	}

}

func TestNewTurnoverReqOptionsBuilder(t *testing.T) {
	date := time.Now()
	expectStruct := TurnoverRequestOptions{
		lang:             LangRu,
		isTonightSession: true,
		date:             date}
	bld := NewTurnoverReqOptionsBuilder()
	bld.Date(date).Lang(LangRu).IsTonightSession(true)
	if got, expected := bld.Build(), expectStruct; got == nil || *got != expected {
		t.Fatalf("Error: expecting `%v` \ngot `%v` \ninstead", expected, got)
	}

}

func TestAddTurnoverRequestOptionsNilOptions(t *testing.T) {
	var income *TurnoverRequestOptions = nil
	c := NewClient(nil)
	url, _ := c.BaseURL.Parse("turnovers.json")
	gotUrl := addTurnoverRequestOptions(url, income, turnoversPrevDateBlock)

	expected := `https://iss.moex.com/iss/turnovers.json?iss.json=extended&iss.meta=off&iss.only=turnoversprevdate`
	if got := gotUrl.String(); got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}

func TestAddTurnoverRequestOptions(t *testing.T) {
	var incomeOptions = NewTurnoverReqOptionsBuilder().
		Lang(LangEn).
		Date(time.Date(2021, 2, 24, 12, 0, 0, 0, time.UTC)).
		IsTonightSession(true).Build()
	c := NewClient(nil)
	url, _ := c.BaseURL.Parse("turnovers.json")
	gotUrl := addTurnoverRequestOptions(url, incomeOptions, turnoversBlock)

	expected := `https://iss.moex.com/iss/turnovers.json?date=2021-02-24&is_tonight_session=1&iss.json=extended&iss.meta=off&iss.only=turnovers&lang=en`
	if got := gotUrl.String(); got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}

func TestTurnoverBlock_String(t *testing.T) {
	type Case struct {
		income   turnoverBlock
		expected string
	}

	cases := []Case{
		{
			income:   turnoversBlockUndefined,
			expected: ""},
		{
			income:   turnoversBlock,
			expected: "turnovers"},
		{
			income:   turnoversPrevDateBlock,
			expected: "turnoversprevdate"},
	}

	for _, c := range cases {
		if got := c.income.String(); got != c.expected {
			t.Fatalf("Error: expecting :\n`%s` \ngot \n`%s` \ninstead", c.expected, got)
		}
	}
}
