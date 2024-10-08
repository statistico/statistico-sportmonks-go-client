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
	ID                      int       `json:"id"`
	SportID                 int       `json:"sport_id"`
	LeagueID                int       `json:"league_id"`
	TieBreakerRuleID        int       `json:"tie_breaker_rule_id"`
	Name                    string    `json:"name"`
	Finished                bool      `json:"finished"`
	Pending                 bool      `json:"pending"`
	IsCurrent               bool      `json:"is_current"`
	StartingAt              string    `json:"starting_at"`
	EndingAt                string    `json:"ending_at"`
	StandingsRecalculatedAt string    `json:"standings_recalculated_at"`
	GamesInCurrentWeek      bool      `json:"games_in_current_week"`
	League                  *League   `json:"league,omitempty"`
	Teams                   []Team    `json:"teams,omitempty"`
	Stages                  []Stage   `json:"stages,omitempty"`
	Fixtures                []Fixture `json:"fixtures,omitempty"`
}

// Seasons fetches Season resources. The endpoint used within this method is paginated, to select the required
// page use the 'page' method argument. Pagination information including current page and count are included within
// // the Pagination struct with the ResponseDetails struct. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) Seasons(ctx context.Context, page int, includes []string) ([]Season, *ResponseDetails, error) {
	values := url.Values{
		"page":    {strconv.Itoa(page)},
		"include": {strings.Join(includes, ";")},
	}

	response := struct {
		Data         []Season       `json:"data"`
		Pagination   Pagination     `json:"pagination"`
		Subscription []Subscription `json:"subscription"`
		RateLimit    RateLimit      `json:"rate_limit"`
		TimeZone     string         `json:"timezone"`
	}{}

	err := c.getResource(ctx, seasonsURI, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, &ResponseDetails{
		Pagination:   &response.Pagination,
		Subscription: response.Subscription,
		RateLimit:    response.RateLimit,
		TimeZone:     response.TimeZone,
	}, err
}

// SeasonByID fetches a Season resource by ID. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) SeasonByID(ctx context.Context, id int, includes []string) (*Season, *ResponseDetails, error) {
	path := fmt.Sprintf(seasonsURI+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ";")},
		"deleted": []string{"1"},
	}

	response := struct {
		Data         *Season        `json:"data"`
		Subscription []Subscription `json:"subscription"`
		RateLimit    RateLimit      `json:"rate_limit"`
		TimeZone     string         `json:"timezone"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, &ResponseDetails{
		Subscription: response.Subscription,
		RateLimit:    response.RateLimit,
		TimeZone:     response.TimeZone,
	}, err
}
