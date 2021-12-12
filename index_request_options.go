package moexiss

import "net/url"

//Language represents a language of answers of MoEx ISS API
type Language string

const (
	LangUndefined Language = ""
	LangEn        Language = "en"
	LangRu        Language = "ru"
)

//string representations of Language values
func (s Language) String() string {
	switch s {
	case LangEn:
		return string(LangEn)
	case LangRu:
		return string(LangRu)
	default:
		return string(LangUndefined)
	}
}

//IndexRequestOptions contains options which using as arguments
//for building requests of 'Index'
//MoEx ISS API docs: https://iss.moex.com/iss/reference/28
type IndexRequestOptions struct {
	// Engines details
	enginesLang Language //`engines.lang` query parameter in url.URL

	// Markets details
	marketsLang Language //`markets.lang` query parameter in url.URL

	// Boards details
	boardsLang Language //`boards.lang` query parameter in url.URL

	// BoardGroups details
	boardGroupsLang     Language   //`boardgroups.lang` query parameter in url.URL
	boardGroupsEngine   EngineName //`boardgroups.engine` query parameter in url.URL
	boardGroupsIsTraded bool       //`boardgroups.is_traded` query parameter in url.URL

	// Durations details
	durationsLang Language //`durations.lang` query parameter in url.URL

	// SecurityTypes details
	securityTypesLang   Language   //`securitytypes.lang` query parameter in url.URL
	securityTypesEngine EngineName //`securitytypes.engine` query parameter in url.URL

	// SecurityGroups details
	securityGroupsLang         Language   //`securitygroups.lang` query parameter in url.URL
	securityGroupsEngine       EngineName //`securitygroups.trade_engine` query parameter in url.URL
	securityGroupsHideInactive bool       //`securitygroups.hide_inactive` query parameter in url.URL

	// SecurityCollections details
	securityCollectionsLang Language //`securitycollections.lang` query parameter in url.URL

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

//Board chains to type *IndexReqOptionsBuilder and returns *IndexReqOptionsBoardBuilder
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
func (e *IndexReqOptionsBoardGroupBuilder) WithEngine(engine EngineName) *IndexReqOptionsBoardGroupBuilder {
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
func (e *IndexReqOptionsSecurityTypeBuilder) WithEngine(engine EngineName) *IndexReqOptionsSecurityTypeBuilder {
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
func (e *IndexReqOptionsSecurityGroupBuilder) WithEngine(engine EngineName) *IndexReqOptionsSecurityGroupBuilder {
	e.options.securityGroupsEngine = engine
	return e
}

/* Options of SecurityCollection*/

//IndexReqOptionsSecurityCollectionBuilder facet of IndexReqOptionsBuilder
type IndexReqOptionsSecurityCollectionBuilder struct {
	IndexReqOptionsBuilder
}

//SecurityCollection chains to type *IndexReqOptionsBuilder and returns *IndexReqOptionsSecurityCollectionBuilder
func (b *IndexReqOptionsBuilder) SecurityCollection() *IndexReqOptionsSecurityCollectionBuilder {
	return &IndexReqOptionsSecurityCollectionBuilder{*b}
}

//Lang sets 'securityCollectionsLang' parameter to a request of directories of SecurityCollection
func (e *IndexReqOptionsSecurityCollectionBuilder) Lang(lang Language) *IndexReqOptionsSecurityCollectionBuilder {
	e.options.securityCollectionsLang = lang
	return e
}

//addIndexRequestOptions sets parameters into *url.URL
//from IndexRequestOptions struct and returns it back
func addIndexRequestOptions(url *url.URL, options IndexRequestOptions) *url.URL {
	q := url.Query()
	if options.enginesLang != LangUndefined {
		q.Set("engines.lang", options.enginesLang.String())
	}
	if options.marketsLang != LangUndefined {
		q.Set("markets.lang", options.marketsLang.String())
	}
	if options.boardsLang != LangUndefined {
		q.Set("boards.lang", options.boardsLang.String())
	}
	if options.boardGroupsLang != LangUndefined {
		q.Set("boardgroups.lang", options.boardGroupsLang.String())
	}
	if options.boardGroupsEngine != EngineUndefined {
		q.Set("boardgroups.engine", options.boardGroupsEngine.String())
	}
	if options.boardGroupsIsTraded {
		q.Set("boardgroups.is_traded", "1")
	}
	if options.durationsLang != LangUndefined {
		q.Set("durations.lang", options.durationsLang.String())
	}
	if options.securityTypesLang != LangUndefined {
		q.Set("securitytypes.lang", options.securityTypesLang.String())
	}
	if options.securityTypesEngine != EngineUndefined {
		q.Set("securitytypes.engine", options.securityTypesEngine.String())
	}
	if options.securityTypesLang != LangUndefined {
		q.Set("securitytypes.lang", options.securityTypesLang.String())
	}
	if options.securityGroupsLang != LangUndefined {
		q.Set("securitygroups.lang", options.securityGroupsLang.String())
	}
	if options.securityGroupsEngine != EngineUndefined {
		q.Set("securitygroups.trade_engine", options.securityGroupsEngine.String())
	}
	if options.securityGroupsHideInactive {
		q.Set("securitygroups.hide_inactive", "1")
	}
	if options.securityCollectionsLang != LangUndefined {
		q.Set("securitycollections.lang", options.securityCollectionsLang.String())
	}

	url.RawQuery = q.Encode()
	return url
}
