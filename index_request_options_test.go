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
		enginesLang: EngLanguage,
		marketsLang: RusLanguage,
		boardsLang:  RusLanguage,
	}
	bld := NewIndexReqOptionsBuilder()
	bld.
		Market().Lang(RusLanguage).
		Engine().Lang(EngLanguage).
		Board().Lang(RusLanguage)
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

}
