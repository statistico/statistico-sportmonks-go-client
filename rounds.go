package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

// Round provides a struct representation of a Round resource.
type Round struct {
	ID                 int    `json:"id"`
	SportID            int    `json:"sport_id"`
	LeagueID           int    `json:"league_id"`
	SeasonID           int    `json:"season_id"`
	StageID            int    `json:"stage_id"`
	Name               string `json:"name"`
	Finished           bool   `json:"finished"`
	IsCurrent          bool   `json:"is_current"`
	StartingAt         string `json:"starting_at"`
	EndingAt           string `json:"ending_at"`
	GamesInCurrentWeek bool   `json:"games_in_current_week"`
}

// RoundByID fetches a Round resource by ID. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) RoundByID(ctx context.Context, id int, includes []string) (*Round, *Meta, error) {
	path := fmt.Sprintf(roundsURI+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ";")},
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

// RoundsBySeasonID fetches a Round resource associated to a Season by Season ID. Use the includes slice of string
// to enrich the response data.
func (c *HTTPClient) RoundsBySeasonID(ctx context.Context, id int, includes []string) ([]Round, *Meta, error) {
	path := fmt.Sprintf(roundsSeasonURI+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ";")},
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
