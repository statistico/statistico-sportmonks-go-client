package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

const stagesUri = "/stages"
const stagesSeasonUri = "/stages/season"

type Stage struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	Type         *string      `json:"type"`
	LeagueID     int         `json:"league_id"`
	SeasonID     int         `json:"season_id"`
	SortOrder    *string     `json:"sort_order"`
	HasStandings *bool        `json:"has_standings"`
	FixturesData FixturesData `json:"fixtures, omitempty"`
	LeagueData LeagueData `json:"league, omitempty"`
	ResultsData FixturesData`json:"results, omitempty"`
	SeasonData SeasonData`json:"season, omitempty"`
}

// Fixtures returns a slice of Fixture struct associated to a Stage.
func (s *Stage) Fixtures() []Fixture {
	return s.FixturesData.Data
}

// League returns a League struct associated to a Stage.
func (s *Stage) League() *League {
	return s.LeagueData.Data
}

// Results returns a slice of completed Fixture struct associated to a Stage.
func (s *Stage) Results() []Fixture {
	return s.ResultsData.Data
}

// Season returns a Season struct associated to a Stage.
func (s *Stage) Season() *Season {
	return s.SeasonData.Data
}

// StageById returns a Stage struct and associated meta data.
func (c *HTTPClient) StageById(ctx context.Context, id int, includes []string) (*Stage, *Meta, error) {
	path := fmt.Sprintf(stagesUri+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data *Stage `json:"data"`
		Meta *Meta      `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// StagesBySeasonId returns a slice of Stage struct and supporting meta data. Use the includes slice to
// enrich the response data.
func (c *HTTPClient) StagesBySeasonId(ctx context.Context, id int, includes []string) ([]Stage, *Meta, error) {
	path := fmt.Sprintf(stagesSeasonUri + "/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data []Stage `json:"data"`
		Meta *Meta      `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}