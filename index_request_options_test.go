package moexiss

import "testing"

func TestLanguage_String(t *testing.T) {
	if got, expected := EngLanguage.String(), "en"; got != expected {
		t.Fatalf("Error: expecting `%s` Lang \ngot `%s` Lang \ninstead", expected, got)
	}
	if got, expected := RusLanguage.String(), "ru"; got != expected {
		t.Fatalf("Error: expecting `%s` Lang \ngot `%s` Lang \ninstead", expected, got)
	}
	if got, expected := UndefinedLanguage.String(), ""; got != expected {
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
	expectStruct := &IndexRequestOptions{enginesLang: EngLanguage}
	bld := NewIndexReqOptionsBuilder()
	bld.Engine().Lang(EngLanguage)
	result := bld.Build()
	if result == nil {
		t.Fatalf("Error: expecting non-nil *IndexRequestOptions: got <nil> instead")
	}
	if got, expected := result.enginesLang, expectStruct.enginesLang; got != expected {
		t.Fatalf("Error: expecting `%s` Lang \ngot `%s` Lang \ninstead", expected, got)
	}
}

func TestIndexReqOptionsDurationBuilder_Lang(t *testing.T) {
	expectStruct := &IndexRequestOptions{durationsLang: RusLanguage}
	bld := NewIndexReqOptionsBuilder()
	bld.Duration().Lang(RusLanguage)
	result := bld.Build()
	if result == nil {
		t.Fatalf("Error: expecting non-nil *IndexRequestOptions: got <nil> instead")
	}
	if got, expected := result.durationsLang, expectStruct.durationsLang; got != expected {
		t.Fatalf("Error: expecting `%s` Lang \ngot `%s` Lang \ninstead", expected, got)
	}
}

func TestIndexReqOptionsMarketBuilder_Lang(t *testing.T) {
	expectStruct := &IndexRequestOptions{marketsLang: RusLanguage}
	bld := NewIndexReqOptionsBuilder()
	bld.Market().Lang(RusLanguage)
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
		enginesLang:        EngLanguage,
		marketsLang:        RusLanguage,
		boardsLang:         RusLanguage,
		boardGroupsLang:    EngLanguage,
		durationsLang:      EngLanguage,
		securityTypesLang:  RusLanguage,
		securityGroupsLang: EngLanguage,
	}
	bld := NewIndexReqOptionsBuilder()
	bld.
		Market().Lang(RusLanguage).
		Engine().Lang(EngLanguage).
		Board().Lang(RusLanguage).
		BoardGroup().Lang(EngLanguage).
		Duration().Lang(EngLanguage).
		SecurityType().Lang(RusLanguage).
		SecurityGroup().Lang(EngLanguage)
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

}

func TestIndexReqOptionsBoardGroupBuilderAllOptions(t *testing.T) {
	expectStruct := &IndexRequestOptions{
		boardGroupsLang:     RusLanguage,
		boardGroupsEngine:   "stock",
		boardGroupsIsTraded: true}
	bld := NewIndexReqOptionsBuilder()
	bld.BoardGroup().Lang(RusLanguage).WithEngine("stock").IsTraded(true)
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
		securityTypesLang:   EngLanguage,
		securityTypesEngine: "stock"}
	bld := NewIndexReqOptionsBuilder()
	bld.SecurityType().Lang(EngLanguage).WithEngine("stock")
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
		securityGroupsLang:         RusLanguage,
		securityGroupsEngine:       "stock",
		securityGroupsHideInactive: true,
	}
	bld := NewIndexReqOptionsBuilder()
	bld.SecurityGroup().Lang(RusLanguage).WithEngine("stock").HideInactive(true)
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
