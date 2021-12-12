package moexiss

//Language represents of the language of answers of MoEx ISS API
type Language string

const (
	UndefinedLanguage Language = ""
	EngLanguage       Language = "en"
	RusLanguage       Language = "ru"
)

//string representations of Language values
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

//IndexRequestOptions contains options which using as arguments
//for building requests of 'Index'
type IndexRequestOptions struct {
	// Engines details
	enginesLang Language

	// Markets details
	marketsLang Language

	// Boards details
	boardsLang Language

	// BoardGroups details
	boardGroupsLang     Language
	boardGroupsEngine   string
	boardGroupsIsTraded bool

	// Durations details
	durationsLang Language

	// SecurityTypes details
	securityTypesLang   Language
	securityTypesEngine string

	// SecurityGroups details
	securityGroupsLang         Language
	securityGroupsEngine       string
	securityGroupsHideInactive bool
}

//IndexReqOptionsBuilder represents a builder of IndexRequestOptions struct
type IndexReqOptionsBuilder struct {
	options *IndexRequestOptions
}

//NewIndexReqOptionsBuilder is a constructor for IndexReqOptionsBuilder
func NewIndexReqOptionsBuilder() *IndexReqOptionsBuilder {
	return &IndexReqOptionsBuilder{options: &IndexRequestOptions{}}
}

//Build builds IndexRequestOptions from IndexReqOptionsBuilder
func (b *IndexReqOptionsBuilder) Build() *IndexRequestOptions {
	return b.options
}

/* Options of Engine*/

//IndexReqOptionsEngineBuilder facet of IndexReqOptionsBuilder
type IndexReqOptionsEngineBuilder struct {
	IndexReqOptionsBuilder
}

//Engine chains to type *IndexReqOptionsBuilder and returns *IndexReqOptionsEngineBuilder
func (b *IndexReqOptionsBuilder) Engine() *IndexReqOptionsEngineBuilder {
	return &IndexReqOptionsEngineBuilder{*b}
}

//Lang sets 'enginesLang' parameter to a request of directories of Engine
func (e *IndexReqOptionsEngineBuilder) Lang(lang Language) *IndexReqOptionsEngineBuilder {
	e.options.enginesLang = lang
	return e
}

/* Options of Market*/

//IndexReqOptionsMarketBuilder facet of IndexReqOptionsBuilder
type IndexReqOptionsMarketBuilder struct {
	IndexReqOptionsBuilder
}

//Market chains to type *IndexReqOptionsBuilder and returns *IndexReqOptionsMarketBuilder
func (b *IndexReqOptionsBuilder) Market() *IndexReqOptionsMarketBuilder {
	return &IndexReqOptionsMarketBuilder{*b}
}

//Lang sets 'marketsLang' parameter to a request of directories of Market
func (e *IndexReqOptionsMarketBuilder) Lang(lang Language) *IndexReqOptionsMarketBuilder {
	e.options.marketsLang = lang
	return e
}

/* Options of Board*/

//IndexReqOptionsBoardBuilder facet of IndexReqOptionsBuilder
type IndexReqOptionsBoardBuilder struct {
	IndexReqOptionsBuilder
}

//Engine chains to type *IndexReqOptionsBuilder and returns *IndexReqOptionsBoardBuilder
func (b *IndexReqOptionsBuilder) Board() *IndexReqOptionsBoardBuilder {
	return &IndexReqOptionsBoardBuilder{*b}
}

//Lang sets 'boardsLang' parameter to a request of directories of Board
func (e *IndexReqOptionsBoardBuilder) Lang(lang Language) *IndexReqOptionsBoardBuilder {
	e.options.boardsLang = lang
	return e
}

/* Options of BoardGroup*/

//IndexReqOptionsBoardGroupBuilder facet of IndexReqOptionsBuilder
type IndexReqOptionsBoardGroupBuilder struct {
	IndexReqOptionsBuilder
}

//BoardGroup chains to type *IndexReqOptionsBuilder and returns *IndexReqOptionsBoardGroupBuilder
func (b *IndexReqOptionsBuilder) BoardGroup() *IndexReqOptionsBoardGroupBuilder {
	return &IndexReqOptionsBoardGroupBuilder{*b}
}

//Lang sets 'boardGroupsLang' parameter to a request of directories of BoardGroup
func (e *IndexReqOptionsBoardGroupBuilder) Lang(lang Language) *IndexReqOptionsBoardGroupBuilder {
	e.options.boardGroupsLang = lang
	return e
}

//IsTraded sets 'boardGroupsIsTraded' parameter to a request of directories of BoardGroup
func (e *IndexReqOptionsBoardGroupBuilder) IsTraded(isTrading bool) *IndexReqOptionsBoardGroupBuilder {
	e.options.boardGroupsIsTraded = isTrading
	return e
}

//WithEngine sets 'boardGroupsEngine' parameter to a request of directories of BoardGroup
func (e *IndexReqOptionsBoardGroupBuilder) WithEngine(engine string) *IndexReqOptionsBoardGroupBuilder {
	e.options.boardGroupsEngine = engine
	return e
}

/* Options of Duration*/

//IndexReqOptionsDurationBuilder facet of IndexReqOptionsBuilder
type IndexReqOptionsDurationBuilder struct {
	IndexReqOptionsBuilder
}

//Duration chains to type *IndexReqOptionsBuilder and returns *IndexReqOptionsDurationBuilder
func (b *IndexReqOptionsBuilder) Duration() *IndexReqOptionsDurationBuilder {
	return &IndexReqOptionsDurationBuilder{*b}
}

//Lang sets 'durationsLang' parameter to a request of directories of Duration
//and returns *IndexReqOptionsDurationBuilder
func (e *IndexReqOptionsDurationBuilder) Lang(lang Language) *IndexReqOptionsDurationBuilder {
	e.options.durationsLang = lang
	return e
}

/* Options of SecurityType*/

//IndexReqOptionsSecurityTypeBuilder facet of IndexReqOptionsBuilder
type IndexReqOptionsSecurityTypeBuilder struct {
	IndexReqOptionsBuilder
}

//SecurityType chains to type *IndexReqOptionsBuilder and returns *IndexReqOptionsSecurityTypeBuilder
func (b *IndexReqOptionsBuilder) SecurityType() *IndexReqOptionsSecurityTypeBuilder {
	return &IndexReqOptionsSecurityTypeBuilder{*b}
}

//Lang sets 'securityTypesLang' parameter to a request of directories of SecurityType
func (e *IndexReqOptionsSecurityTypeBuilder) Lang(lang Language) *IndexReqOptionsSecurityTypeBuilder {
	e.options.securityTypesLang = lang
	return e
}

//WithEngine sets 'securityTypesEngine' parameter to a request of directories of SecurityType
func (e *IndexReqOptionsSecurityTypeBuilder) WithEngine(engine string) *IndexReqOptionsSecurityTypeBuilder {
	e.options.securityTypesEngine = engine
	return e
}

/* Options of SecurityGroup*/

//IndexReqOptionsSecurityGroupBuilder facet of IndexReqOptionsBuilder
type IndexReqOptionsSecurityGroupBuilder struct {
	IndexReqOptionsBuilder
}

//SecurityGroup chains to type *IndexReqOptionsBuilder and returns *IndexReqOptionsSecurityGroupBuilder
func (b *IndexReqOptionsBuilder) SecurityGroup() *IndexReqOptionsSecurityGroupBuilder {
	return &IndexReqOptionsSecurityGroupBuilder{*b}
}

//Lang sets 'securityGroupsLang' parameter to a request of directories of SecurityGroup
func (e *IndexReqOptionsSecurityGroupBuilder) Lang(lang Language) *IndexReqOptionsSecurityGroupBuilder {
	e.options.securityGroupsLang = lang
	return e
}

//HideInactive sets 'securityGroupsHideInactive' parameter to a request of directories of SecurityGroup
func (e *IndexReqOptionsSecurityGroupBuilder) HideInactive(isHiding bool) *IndexReqOptionsSecurityGroupBuilder {
	e.options.securityGroupsHideInactive = isHiding
	return e
}

//WithEngine sets 'securityGroupsEngine' parameter to a request of directories of SecurityGroup
func (e *IndexReqOptionsSecurityGroupBuilder) WithEngine(engine string) *IndexReqOptionsSecurityGroupBuilder {
	e.options.securityGroupsEngine = engine
	return e
}
