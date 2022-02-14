package moexiss

import (
	"testing"
)

func TestLanguage_String(t *testing.T) {
	if got, expected := LangEn.String(), "en"; got != expected {
		t.Fatalf("Error: expecting `%s` Lang \ngot `%s` Lang \ninstead", expected, got)
	}
	if got, expected := LangRu.String(), "ru"; got != expected {
		t.Fatalf("Error: expecting `%s` Lang \ngot `%s` Lang \ninstead", expected, got)
	}
	if got, expected := LangUndefined.String(), ""; got != expected {
		t.Fatalf("Error: expecting `%s` Lang \ngot `%s` Lang \ninstead", expected, got)
	}
}

func TestIndexReqOptionsBuilder_Build(t *testing.T) {
	expectStruct := &IndexRequestOptions{}
	bld := NewIndexReqOptionsBuilder()
	result := bld.Build()
	if result == nil {
		t.Fatalf("Error: expecting non-nil *IndexRequestOptions: got <nil> instead")
	}
	if got, expected := result.enginesLang, expectStruct.enginesLang; got != expected {
		t.Fatalf("Error: expecting `%s` Lang \ngot `%s` Lang \ninstead", expected, got)
	}
	if got, expected := result.marketsLang, expectStruct.marketsLang; got != expected {
		t.Fatalf("Error: expecting `%s` Lang \ngot `%s` Lang \ninstead", expected, got)
	}
	if got, expected := result.boardsLang, expectStruct.boardsLang; got != expected {
		t.Fatalf("Error: expecting `%s` Lang \ngot `%s` Lang \ninstead", expected, got)
	}

}

func TestIndexReqOptionsEngineBuilder_Lang(t *testing.T) {
	expectStruct := &IndexRequestOptions{enginesLang: LangEn}
	bld := NewIndexReqOptionsBuilder()
	bld.Engine().Lang(LangEn)
	result := bld.Build()
	if result == nil {
		t.Fatalf("Error: expecting non-nil *IndexRequestOptions: got <nil> instead")
	}
	if got, expected := result.enginesLang, expectStruct.enginesLang; got != expected {
		t.Fatalf("Error: expecting `%s` Lang \ngot `%s` Lang \ninstead", expected, got)
	}
}

func TestIndexReqOptionsDurationBuilder_Lang(t *testing.T) {
	expectStruct := &IndexRequestOptions{durationsLang: LangRu}
	bld := NewIndexReqOptionsBuilder()
	bld.Duration().Lang(LangRu)
	result := bld.Build()
	if result == nil {
		t.Fatalf("Error: expecting non-nil *IndexRequestOptions: got <nil> instead")
	}
	if got, expected := result.durationsLang, expectStruct.durationsLang; got != expected {
		t.Fatalf("Error: expecting `%s` Lang \ngot `%s` Lang \ninstead", expected, got)
	}
}

func TestIndexReqOptionsMarketBuilder_Lang(t *testing.T) {
	expectStruct := &IndexRequestOptions{marketsLang: LangRu}
	bld := NewIndexReqOptionsBuilder()
	bld.Market().Lang(LangRu)
	result := bld.Build()
	if result == nil {
		t.Fatalf("Error: expecting non-nil *IndexRequestOptions: got <nil> instead")
	}
	if got, expected := result.marketsLang, expectStruct.marketsLang; got != expected {
		t.Fatalf("Error: expecting `%s` Lang \ngot `%s` Lang \ninstead", expected, got)
	}
}

func TestIndexReqOptionsBuilder_LangRuEn(t *testing.T) {
	expectStruct := &IndexRequestOptions{
		enginesLang:             LangEn,
		marketsLang:             LangRu,
		boardsLang:              LangRu,
		boardGroupsLang:         LangEn,
		durationsLang:           LangEn,
		securityTypesLang:       LangRu,
		securityGroupsLang:      LangEn,
		securityCollectionsLang: LangRu,
	}
	bld := NewIndexReqOptionsBuilder()
	bld.
		Market().Lang(LangRu).
		Engine().Lang(LangEn).
		Board().Lang(LangRu).
		BoardGroup().Lang(LangEn).
		Duration().Lang(LangEn).
		SecurityType().Lang(LangRu).
		SecurityGroup().Lang(LangEn).
		SecurityCollection().Lang(LangRu)
	result := bld.Build()
	if result == nil {
		t.Fatalf("Error: expecting non-nil *IndexRequestOptions: got <nil> instead")
	}
	if got, expected := result.marketsLang, expectStruct.marketsLang; got != expected {
		t.Fatalf("Error: expecting `%s` Lang \ngot `%s` Lang \ninstead", expected, got)
	}
	if got, expected := result.enginesLang, expectStruct.enginesLang; got != expected {
		t.Fatalf("Error: expecting `%s` Lang \ngot `%s` Lang \ninstead", expected, got)
	}
	if got, expected := result.boardsLang, expectStruct.boardsLang; got != expected {
		t.Fatalf("Error: expecting `%s` Lang \ngot `%s` Lang \ninstead", expected, got)
	}
	if got, expected := result.boardGroupsLang, expectStruct.boardGroupsLang; got != expected {
		t.Fatalf("Error: expecting `%s` Lang \ngot `%s` Lang \ninstead", expected, got)
	}
	if got, expected := result.durationsLang, expectStruct.durationsLang; got != expected {
		t.Fatalf("Error: expecting `%s` Lang \ngot `%s` Lang \ninstead", expected, got)
	}
	if got, expected := result.securityTypesLang, expectStruct.securityTypesLang; got != expected {
		t.Fatalf("Error: expecting `%s` Lang \ngot `%s` Lang \ninstead", expected, got)
	}
	if got, expected := result.securityGroupsLang, expectStruct.securityGroupsLang; got != expected {
		t.Fatalf("Error: expecting `%s` Lang \ngot `%s` Lang \ninstead", expected, got)
	}
	if got, expected := result.securityCollectionsLang, expectStruct.securityCollectionsLang; got != expected {
		t.Fatalf("Error: expecting `%s` Lang \ngot `%s` Lang \ninstead", expected, got)
	}

}

func TestIndexReqOptionsBoardGroupBuilderAllOptions(t *testing.T) {
	expectStruct := &IndexRequestOptions{
		boardGroupsLang:     LangRu,
		boardGroupsEngine:   EngineStock,
		boardGroupsIsTraded: true}
	bld := NewIndexReqOptionsBuilder()
	bld.BoardGroup().Lang(LangRu).WithEngine(EngineStock).IsTraded(true)
	result := bld.Build()
	if result == nil {
		t.Fatalf("Error: expecting non-nil *IndexRequestOptions: got <nil> instead")
	}
	if got, expected := result.boardGroupsEngine, expectStruct.boardGroupsEngine; got != expected {
		t.Fatalf("Error: expecting boardGroupsEngine :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
	if got, expected := result.boardGroupsIsTraded, expectStruct.boardGroupsIsTraded; got != expected {
		t.Fatalf("Error: expecting boardGroupsIsTraded :\n`%t`  \ngot \n`%t` \ninstead", expected, got)
	}
	if got, expected := result.boardGroupsLang, expectStruct.boardGroupsLang; got != expected {
		t.Fatalf("Error: expecting boardGroupsLang `%s` \ngot \n`%s` \ninstead", expected, got)
	}

}

func TestIndexReqOptionsSecurityTypesBuilderAllOptions(t *testing.T) {
	expectStruct := &IndexRequestOptions{
		securityTypesLang:   LangEn,
		securityTypesEngine: EngineStock}
	bld := NewIndexReqOptionsBuilder()
	bld.SecurityType().Lang(LangEn).WithEngine(EngineStock)
	result := bld.Build()
	if result == nil {
		t.Fatalf("Error: expecting non-nil *IndexRequestOptions: got <nil> instead")
	}
	if got, expected := result.boardGroupsEngine, expectStruct.boardGroupsEngine; got != expected {
		t.Fatalf("Error: expecting boardGroupsEngine :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
	if got, expected := result.boardGroupsLang, expectStruct.boardGroupsLang; got != expected {
		t.Fatalf("Error: expecting boardGroupsLang `%s` \ngot \n`%s` \ninstead", expected, got)
	}

}

func TestIndexReqOptionsSecurityGroupsBuilderAllOptions(t *testing.T) {
	expectStruct := &IndexRequestOptions{
		securityGroupsLang:         LangRu,
		securityGroupsEngine:       EngineStock,
		securityGroupsHideInactive: true,
	}
	bld := NewIndexReqOptionsBuilder()
	bld.SecurityGroup().Lang(LangRu).WithEngine(EngineStock).HideInactive(true)
	result := bld.Build()
	if result == nil {
		t.Fatalf("Error: expecting non-nil *IndexRequestOptions: got <nil> instead")
	}
	if got, expected := result.securityGroupsEngine, expectStruct.securityGroupsEngine; got != expected {
		t.Fatalf("Error: expecting securityGroupsEngine :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
	if got, expected := result.securityGroupsLang, expectStruct.securityGroupsLang; got != expected {
		t.Fatalf("Error: expecting securityGroupsLang `%s` \ngot \n`%s` \ninstead", expected, got)
	}
	if got, expected := result.securityGroupsHideInactive, expectStruct.securityGroupsHideInactive; got != expected {
		t.Fatalf("Error: expecting securityGroupsHideInactive :\n`%t`  \ngot \n`%t` \ninstead", expected, got)
	}

}

func TestAddIndexRequestOptions(t *testing.T) {
	income := NewIndexReqOptionsBuilder().
		Engine().Lang(LangEn).
		Market().Lang(LangEn).
		Board().Lang(LangEn).
		BoardGroup().Lang(LangEn).WithEngine(EngineCurrency).IsTraded(true).
		Duration().Lang(LangEn).
		SecurityType().Lang(LangEn).WithEngine(EngineFutures).
		SecurityGroup().Lang(LangRu).WithEngine(EngineStock).HideInactive(true).
		SecurityCollection().Lang(LangRu).
		Build()
	c := NewClient(nil)
	url, _ := c.BaseURL.Parse(indexPartsUrl)
	gotURL := addIndexRequestOptions(url, income)

	expected := `https://iss.moex.com/iss/index.json?` +
		`boardgroups.engine=currency&boardgroups.is_traded=1&boardgroups.lang=en&` +
		`boards.lang=en&` +
		`durations.lang=en&` +
		`engines.lang=en&` +
		`iss.meta=off&` +
		`markets.lang=en&` +
		`securitycollections.lang=ru&` +
		`securitygroups.hide_inactive=1&securitygroups.lang=ru&securitygroups.trade_engine=stock&` +
		`securitytypes.engine=futures&securitytypes.lang=en`

	if got := gotURL.String(); got != expected {
		t.Fatalf("Error: expecting url :\n`%s`  \ngot \n`%s` \ninstead", expected, got)
	}
}

func TestAddIndexRequestOptionsNilOptions(t *testing.T) {
	var income *IndexRequestOptions = nil
	c := NewClient(nil)
	url, _ := c.BaseURL.Parse("index.json")
	gotURL := addIndexRequestOptions(url, income)

	expected := `https://iss.moex.com/iss/index.json`
	if got := gotURL.String(); got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}
