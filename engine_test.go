package moexiss

import "testing"

func TestEngineName_String(t *testing.T) {
	testCases := []struct {
		engineName EngineName
		expected   string
	}{
		{EngineUndefined, ""},
		{EngineStock, "stock"},
		{EngineState, "state"},
		{EngineCurrency, "currency"},
		{EngineFutures, "futures"},
		{EngineCommodity, "commodity"},
		{EngineInterventions, "interventions"},
		{EngineOffBoard, "offboard"},
		{EngineAgro, "agro"},
	}
	for _, testCase := range testCases {
		if got, expected := testCase.engineName.String(), testCase.expected; got != expected {
			t.Fatalf("Error: expecting `%s` Engine \ngot `%s` Engine \ninstead", expected, got)
		}
	}
}
