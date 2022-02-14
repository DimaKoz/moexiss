package moexiss

// EngineName represents the known Engine names of MoEx ISS API
type EngineName string

// A section of EngineName values
const (
	EngineUndefined     EngineName = ""
	EngineStock         EngineName = "stock"
	EngineState         EngineName = "state"
	EngineCurrency      EngineName = "currency"
	EngineFutures       EngineName = "futures"
	EngineCommodity     EngineName = "commodity"
	EngineInterventions EngineName = "interventions"
	EngineOffBoard      EngineName = "offboard"
	EngineAgro          EngineName = "agro"
)

//string representations of EngineName values
func (s EngineName) String() string {
	return string(s)
}
