package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

const roundsUri = "/rounds"
const roundsSeasonUri = "/rounds/season"

type Round struct {
	ID       int    `json:"id"`
	Name     int    `json:"name"`
	LeagueID int    `json:"league_id"`
	SeasonID int    `json:"season_id"`
	StageID  int    `json:"stage_id"`
	Start    string `json:"start"`
	End      string `json:"end"`
	FixtureData Fixtures `json:"fixtures, omitempty"`
	LeagueData  LeagueData `json:"league"`
	ResultsData Fixtures `json:"results, omitempty"`
	SeasonData  SeasonData `json:"season"`
}

// Fixtures returns a slice of Fixture struct associated to a Round
func (r *Round) Fixtures() []Fixture {
	return r.FixtureData.Data
}

// League returns a League struct associated to a Round
func (r *Round) League() *League {
	return r.LeagueData.Data
}

// Results returns a slice of completed Fixture struct associated to a Round
func (r *Round) Results() []Fixture {
	return r.ResultsData.Data
}

// Season returns a League struct associated to a Round
func (r *Round) Season() *Season {
	return r.SeasonData.Data
}

// RoundById returns a single Round struct. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) RoundById(ctx context.Context, id int, includes []string) (*Round, *Meta, error) {
	path := fmt.Sprintf(roundsUri+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data *Round 	`json:"data"`
		Meta *Meta      `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// RoundsBySeasonId returns a single Round struct associated to a Season. Use the includes slice of string to enrich
// the response data.
func (c *HTTPClient) RoundsBySeasonId(ctx context.Context, id int, includes []string) ([]Round, *Meta, error) {
	path := fmt.Sprintf(roundsSeasonUri+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data []Round 	`json:"data"`
		Meta *Meta      `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
