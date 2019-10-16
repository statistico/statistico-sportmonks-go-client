package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

// Stage provides a struct representation of a Stage resource
type Stage struct {
	ID           int          `json:"id"`
	Name         string       `json:"name"`
	Type         *string      `json:"type"`
	LeagueID     int          `json:"league_id"`
	SeasonID     int          `json:"season_id"`
	SortOrder    *string      `json:"sort_order"`
	HasStandings *bool        `json:"has_standings"`
	FixturesData fixturesData `json:"fixtures,omitempty"`
	LeagueData   leagueData   `json:"league,omitempty"`
	ResultsData  fixturesData `json:"results,omitempty"`
	SeasonData   seasonData   `json:"season,omitempty"`
}

// Fixtures returns fixture data.
func (s *Stage) Fixtures() []Fixture {
	return s.FixturesData.Data
}

// League returns league data.
func (s *Stage) League() *League {
	return s.LeagueData.Data
}

// Results returns completed fixture data.
func (s *Stage) Results() []Fixture {
	return s.ResultsData.Data
}

// Season returns a season data.
func (s *Stage) Season() *Season {
	return s.SeasonData.Data
}

// StageByID fetches a Stage resource by ID.
// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) StageByID(ctx context.Context, id int, includes []string) (*Stage, *Meta, error) {
	path := fmt.Sprintf(stagesURI+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data *Stage `json:"data"`
		Meta *Meta  `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// StagesBySeasonID fetches a Stage resources by a season ID.
// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) StagesBySeasonID(ctx context.Context, id int, includes []string) ([]Stage, *Meta, error) {
	path := fmt.Sprintf(stagesSeasonURI+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data []Stage `json:"data"`
		Meta *Meta   `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
