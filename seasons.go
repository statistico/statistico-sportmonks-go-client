package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Season provides a struct representation of a Season resource.
type Season struct {
	ID                      int    `json:"id"`
	SportID                 int    `json:"sport_id"`
	LeagueID                int    `json:"league_id"`
	TieBreakerRuleID        int    `json:"tie_breaker_rule_id"`
	Name                    string `json:"name"`
	Finished                bool   `json:"finished"`
	Pending                 bool   `json:"pending"`
	IsCurrent               bool   `json:"is_current"`
	StartingAt              string `json:"starting_at"`
	EndingAt                string `json:"ending_at"`
	StandingsRecalculatedAt string `json:"standings_recalculated_at"`
	GamesInCurrentWeek      bool   `json:"games_in_current_week"`
}

// Seasons fetches Season resources. The endpoint used within this method is paginated, to select the required
// page use the 'page' method argument. Page information including current page and total page are included
// within the Meta response. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) Seasons(ctx context.Context, page int, includes []string) ([]Season, *Meta, error) {
	values := url.Values{
		"page":    {strconv.Itoa(page)},
		"include": {strings.Join(includes, ";")},
	}

	response := struct {
		Data []Season `json:"data"`
		Meta *Meta    `json:"meta"`
	}{}

	err := c.getResource(ctx, seasonsURI, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// SeasonByID fetches a Season resource by ID. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) SeasonByID(ctx context.Context, id int, includes []string) (*Season, *Meta, error) {
	path := fmt.Sprintf(seasonsURI+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ";")},
		"deleted": []string{"1"},
	}

	response := struct {
		Data *Season `json:"data"`
		Meta *Meta   `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
