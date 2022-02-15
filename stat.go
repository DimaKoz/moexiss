package moexiss

import (
	"path"
	"unicode/utf8"
)

const (
	statsPartsUrl = "secstats.json"
)

// StatsService gets intermediate day summary
// from the MoEx ISS API.
//
// MoEx ISS API docs: https://iss.moex.com/iss/reference/823
type StatsService service

// getUrl provides an url to get intermediate day summary
// opt *StatRequestOptions can be nil, it is safe
func (s *StatsService) getUrl(engine EngineName, market string, opt *StatRequestOptions) (string, error) {
	if engine == EngineUndefined {
		return "", ErrBadEngineParameter
	}
	marketMinLen := 3
	if market == "" || utf8.RuneCountInString(market) < marketMinLen {
		return "", ErrBadMarketParameter
	}

	url, _ := s.client.BaseURL.Parse("engines")

	url.Path = path.Join(url.Path, engine.String(), "markets", market, statsPartsUrl)
	gotURL := addStatRequestOptions(url, opt)
	return gotURL.String(), nil
}
