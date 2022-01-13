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
		lang: LangRu,
		isTonightSession: true,
		date: date}
	bld := NewTurnoverReqOptionsBuilder()
	bld.Date(date).Lang(LangRu).IsTonightSession(true)
	if got, expected := bld.Build(), expectStruct; got == nil || *got != expected {
		t.Fatalf("Error: expecting `%v` \ngot `%v` \ninstead", expected, got)
	}

}