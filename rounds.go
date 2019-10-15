package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

// Round provides a struct representation of a Round resource
type Round struct {
	ID           int          `json:"id"`
	Name         int          `json:"name"`
	LeagueID     int          `json:"league_id"`
	SeasonID     int          `json:"season_id"`
	StageID      int          `json:"stage_id"`
	Start        string       `json:"start"`
	End          string       `json:"end"`
	FixturesData fixturesData `json:"fixtures, omitempty"`
	LeagueData   leagueData   `json:"league"`
	ResultsData  fixturesData `json:"results, omitempty"`
	SeasonData   seasonData   `json:"season"`
}

// Fixtures returns a slice of Fixture struct associated to a Round
func (r *Round) Fixtures() []Fixture {
	return r.FixturesData.Data
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

// RoundByID fetches a Round resource by ID.
// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) RoundByID(ctx context.Context, id int, includes []string) (*Round, *Meta, error) {
	path := fmt.Sprintf(roundsURI+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data *Round `json:"data"`
		Meta *Meta  `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// RoundsBySeasonID fetches a Round resource associated to a Season by Season ID.
// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) RoundsBySeasonID(ctx context.Context, id int, includes []string) ([]Round, *Meta, error) {
	path := fmt.Sprintf(roundsSeasonURI+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data []Round `json:"data"`
		Meta *Meta   `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
