package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

// Stage provides a struct representation of a Stage resource.
type Stage struct {
	ID                 int    `json:"id"`
	SportID            int    `json:"sport_id"`
	LeagueID           int    `json:"league_id"`
	SeasonID           int    `json:"season_id"`
	TypeID             int    `json:"type_id"`
	Name               string `json:"name"`
	SortOrder          int    `json:"sort_order"`
	Finished           bool   `json:"finished"`
	IsCurrent          bool   `json:"is_current"`
	StartingAt         string `json:"starting_at"`
	EndingAt           string `json:"ending_at"`
	GamesInCurrentWeek bool   `json:"games_in_current_week"`
	TieBreakerRuleID   *int   `json:"tie_breaker_rule_id"` // Using pointer to handle null value
}

// StageByID fetches a Stage resource by ID. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) StageByID(ctx context.Context, id int, includes []string) (*Stage, *Meta, error) {
	path := fmt.Sprintf(stagesURI+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ";")},
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

// StagesBySeasonID fetches a Stage resources by a season ID. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) StagesBySeasonID(ctx context.Context, id int, includes []string) ([]Stage, *Meta, error) {
	path := fmt.Sprintf(stagesSeasonURI+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ";")},
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
