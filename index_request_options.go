package moexiss

//Language represents of the language of answers of MoEx ISS API
type Language string

const (
	UndefinedLanguage Language = ""
	EngLanguage       Language = "en"
	RusLanguage       Language = "ru"
)

func (s Language) String() string {
	switch s {
	case EngLanguage:
		return string(EngLanguage)
	case RusLanguage:
		return string(RusLanguage)
	default:
		return string(UndefinedLanguage)
	}
}

type IndexRequestOptions struct {
	// Engines details
	enginesLang Language
	// Markets details
	marketsLang Language
}

//DirReqOptBuilder represents a builder of IndexRequestOptions struct
type IndexReqOptionsBuilder struct {
	options *IndexRequestOptions
}

//NewIndexReqOptionsBuilder is a constructor for IndexReqOptionsBuilder
func NewIndexReqOptionsBuilder() *IndexReqOptionsBuilder {
	return &IndexReqOptionsBuilder{options: &IndexRequestOptions{}}
}

//IndexReqOptionsEngineBuilder facet of IndexReqOptionsBuilder
type IndexReqOptionsEngineBuilder struct {
	IndexReqOptionsBuilder
}

//IndexReqOptionsMarketBuilder facet of IndexReqOptionsBuilder
type IndexReqOptionsMarketBuilder struct {
	IndexReqOptionsBuilder
}

//Engine chains to type *IndexReqOptionsBuilder and returns a *IndexReqOptionsEngineBuilder
func (b *IndexReqOptionsBuilder) Engine() *IndexReqOptionsEngineBuilder {
	return &IndexReqOptionsEngineBuilder{*b}
}

//Engine chains to type *IndexReqOptionsBuilder and returns a *IndexReqOptionsMarketBuilder
func (b *IndexReqOptionsBuilder) Market() *IndexReqOptionsMarketBuilder {
	return &IndexReqOptionsMarketBuilder{*b}
}

//Lang sets 'enginesLang' parameter to a request of directories of Engine
func (e *IndexReqOptionsEngineBuilder) Lang(lang Language) *IndexReqOptionsEngineBuilder {
	e.options.enginesLang = lang
	return e
}

//Lang sets 'marketsLang' parameter to a request of directories of Market
func (e *IndexReqOptionsMarketBuilder) Lang(lang Language) *IndexReqOptionsMarketBuilder {
	e.options.marketsLang = lang
	return e
}

//Build builds IndexRequestOptions from IndexReqOptionsBuilder
func (b *IndexReqOptionsBuilder) Build() *IndexRequestOptions {
	return b.options
}
