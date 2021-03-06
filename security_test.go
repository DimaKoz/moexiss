package moexiss

import (
	"github.com/buger/jsonparser"
	"testing"
)

var secArray = []string{
	`[5444, "SBERP", "Сбербанк-п", "20301481B", "Сбербанк России ПАО ап", "RU0009029557", 1, 1199, "Публичное акционерное общество \"Сбербанк России\"", "7707083893", "00032537", "20301481B", "preferred_share", "stock_shares", "TQBR", "TQBR"]`,
	`[414220919, "AA-RM", "Alcoa", null, "Alcoa Corporation", "US0138721065", 1, 1375753, "Alcoa Corporation (Алкоа Корпорэйшн)", "0000005270", null, null, "common_share", "stock_shares", "FQBR", "FQBR"]`,
}

var dataField = `{ "data": [ ` +
	secArray[0] +
	`,` +
	secArray[1] +
	`]}`

var omittedSecResp = `{ "securities":` +
	dataField +
	`}`

func TestParseSecurityItem(t *testing.T) {
	expected := Security{
		Id:                 5444,
		SecId:              "SBERP",
		ShortName:          "Сбербанк-п",
		RegNumber:          "20301481B",
		Name:               "Сбербанк России ПАО ап",
		Isin:               "RU0009029557",
		IsTraded:           true,
		EmitentId:          "1199",
		EmitentTitle:       "Публичное акционерное общество \"Сбербанк России\"",
		EmitentInn:         "7707083893",
		EmitentOkpo:        "00032537",
		GosReg:             "20301481B",
		Type:               "preferred_share",
		Group:              "stock_shares",
		PrimaryBoardId:     "TQBR",
		MarketPriceBoardId: "TQBR",
	}
	got := Security{}
	err := parseSecurityItem(&got, []byte(secArray[0]))
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v  \ninstead", err)
	}

	if got != expected {
		t.Fatalf("Error: expected \n%v : \ngot \n%v  \ninstead", expected, got)
	}

}

func TestParseSecurityItemBadBytes(t *testing.T) {
	expected := jsonparser.UnknownValueTypeError
	got := Security{}
	err := parseSecurityItem(&got, []byte("[s]"))
	if err != expected {
		t.Fatalf("Error: expecting %v error: \ngot %v  \ninstead", expected, err)
	}
}

func TestParseSecurityItemBadInt(t *testing.T) {
	expected := jsonparser.OverflowIntegerError
	got := Security{}
	err := parseSecurityItem(&got, []byte("[-9923372036854775809]"))
	if err != expected {
		t.Fatalf("Error: expecting %v error: \ngot %v  \ninstead", expected, err)
	}
}

func TestParseSecurityItemNil(t *testing.T) {
	expectedText := "<nil> pointer passed instead of *Security"
	var nilSecurity *Security = nil
	err := parseSecurityItem(nilSecurity, []byte(""))
	if err == nil {
		t.Fatalf("Error: expecting error: \n%v \ngot %v  \ninstead", expectedText, err)
	}
	if err.Error() != expectedText {
		t.Fatalf("Error: expecting %v error: \ngot %v  \ninstead", expectedText, err)
	}
}

func TestParseSecurities(t *testing.T) {
	income := dataField
	securities := make([]Security, 0, 2)
	err := parseSecurities(&securities, []byte(income))
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v  \ninstead", err)
	}
	if got, expected := len(securities), len(secArray); got != expected {
		t.Fatalf("Error: expecting %d items \ngot %d items \ninstead", expected, got)
	}
}

func TestParseSecuritiesResponse(t *testing.T) {
	income := omittedSecResp
	securities := make([]Security, 0, 2)
	err := parseSecuritiesResponse(&securities, []byte(income))
	if err != nil {
		t.Fatalf("Error: expecting <nil> error: \ngot %v  \ninstead", err)
	}
	if got, expected := len(securities), len(secArray); got != expected {
		t.Fatalf("Error: expecting %d items \ngot %d items \ninstead", expected, got)
	}
}

func TestParseSecuritiesResponseKeyPathNotFoundError(t *testing.T) {
	income := `{ "securities; }`
	securities := make([]Security, 0, 2)
	if got, expected := parseSecuritiesResponse(&securities, []byte(income)), jsonparser.KeyPathNotFoundError; got != expected {
		t.Fatalf("Error: expecting %v error: got %v instead", got, expected)
	}
}

func TestParseSecuritiesWrongTypeJson(t *testing.T) {
	income := `{ "securities": [] }`
	securities := make([]Security, 0, 2)
	err := parseSecuritiesResponse(&securities, []byte(income))
	if err == nil {
		t.Fatalf("Error: expecting non-nil error: got <nil> instead")
	}
}

func TestIndexParseSecurityMalformedArrayError(t *testing.T) {
	var incomeJSON = `{
		"data": [
		[5444,
	]
}
	`

	securities := make([]Security, 0, 2)
	if got, expected := parseSecurities(&securities, []byte(incomeJSON)), jsonparser.MalformedArrayError; got == nil || got != expected {
		t.Fatalf("Error: expecting:\n'%v'\ngot:\n'%v'\ninstead", expected, got)
	}
}

func TestIndexParseSecurityUnknownValueTypeError(t *testing.T) {
	var incomeJSON = `{
		"data": [
		[5444,]
	]
}
	`

	securities := make([]Security, 0, 2)
	if got, expected := parseSecurities(&securities, []byte(incomeJSON)), jsonparser.UnknownValueTypeError; got == nil || got != expected {
		t.Fatalf("Error: expecting:\n'%v'\ngot:\n'%v'\ninstead", expected, got)
	}
}
