package moexiss

import (
	"testing"
)

func TestIndicesReqOptionsBuilder_Build(t *testing.T) {
	expectStruct := IndicesRequestOptions{}
	bld := NewIndicesReqOptionsBuilder()

	if got, expected := *bld.Build(), expectStruct; got != expected {
		t.Fatalf("Error: expecting `%v` IndicesRequestOptions \ngot `%v` IndicesRequestOptions \ninstead", expected, got)
	}
}

func TestIndicesReqOptionsBuilder_Lang(t *testing.T) {
	expectStruct := IndicesRequestOptions{lang: LangEn}
	bld := NewIndicesReqOptionsBuilder()
	bld.Lang(LangEn)
	if got, expected := *bld.Build(), expectStruct; got != expected {
		t.Fatalf("Error: expecting `%v` \ngot `%v` \ninstead", expected, got)
	}
}

func TestNewIndicesRequestOptions(t *testing.T) {
	expectStruct := IndicesRequestOptions{
		lang: LangRu,
	}
	bld := NewIndicesReqOptionsBuilder()
	bld.Lang(LangRu)
	if got, expected := *bld.Build(), expectStruct; got != expected {
		t.Fatalf("Error: expecting `%v` \ngot `%v` \ninstead", expected, got)
	}

}

func TestAddIndicesRequestOptionsNilOptions(t *testing.T) {
	var income *IndicesRequestOptions = nil
	c := NewClient(nil)
	url, _ := c.BaseURL.Parse("test.json")
	gotURL := addIndicesRequestOptions(url, income)

	expected := `https://iss.moex.com/iss/test.json?iss.json=extended&iss.meta=off`
	if got := gotURL.String(); got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}

func TestAddIndicesRequestOptions(t *testing.T) {
	var incomeOptions = NewIndicesReqOptionsBuilder().
		Lang(LangEn).
		Build()
	c := NewClient(nil)
	url, _ := c.BaseURL.Parse("test.json")
	gotURL := addIndicesRequestOptions(url, incomeOptions)

	expected := `https://iss.moex.com/iss/test.json?iss.json=extended&iss.meta=off&lang=en`
	if got := gotURL.String(); got != expected {
		t.Fatalf("Error: expecting url :\n`%s` \ngot \n`%s` \ninstead", expected, got)
	}
}
