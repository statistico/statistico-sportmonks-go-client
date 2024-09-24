package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// League provides a struct representation of a League resource.
type League struct {
	ID              int            `json:"id"`
	Active          bool           `json:"active"`
	Type            string         `json:"type"`
	LegacyID        int            `json:"legacy_id"`
	CountryID       int            `json:"country_id"`
	LogoPath        string         `json:"logo_path"`
	Name            string         `json:"name"`
	IsCup           bool           `json:"is_cup"`
	CurrentSeasonID int            `json:"current_season_id"`
	CurrentRoundID  int            `json:"current_round_id"`
	CurrentStageID  int            `json:"current_stage_id"`
	LiveStandings   bool           `json:"live_standings"`
	Coverage        LeagueCoverage `json:"coverage"`
	Country         *Country       `json:"country,omitempty"`
	Season          SeasonData     `json:"season,omitempty"`
	Seasons         SeasonsData    `json:"seasons,omitempty"`
}

// SeasonData returns the current Season struct associated to a League.
func (l *League) SeasonData() *Season {
	return l.Season.Data
}

// SeasonsData returns a slice of Season struct associated to a League.
func (l *League) SeasonsData() []Season {
	return l.Seasons.Data
}

// Leagues fetches League resources. The endpoint used within this method is paginated, to select the required
// page use the 'page' method argument. Page information including current page and total page are included
// within the Meta response. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) Leagues(ctx context.Context, page int, includes []string) ([]League, *Meta, error) {
	values := url.Values{
		"page":    {strconv.Itoa(page)},
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data []League `json:"data"`
		Meta *Meta    `json:"meta"`
	}{}

	err := c.getResource(ctx, leaguesURI, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// LeagueByID fetches League resources by ID. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) LeagueByID(ctx context.Context, id int, includes []string) (*League, *Meta, error) {
	path := fmt.Sprintf(leaguesURI+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data *League `json:"data"`
		Meta *Meta   `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
