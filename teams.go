package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

// Team provides a struct representation of a Team resource.
type Team struct {
	ID           int              `json:"id"`
	SportID      int              `json:"sport_id"`
	CountryID    int              `json:"country_id"`
	VenueID      int              `json:"venue_id"`
	Gender       string           `json:"gender"`
	Name         string           `json:"name"`
	ShortCode    string           `json:"short_code"`
	ImagePath    string           `json:"image_path"`
	Founded      int              `json:"founded"`
	Type         string           `json:"type"`
	Placeholder  bool             `json:"placeholder"`
	LastPlayedAt string           `json:"last_played_at"`
	Meta         *TeamFixtureMeta `json:"meta,omitempty"`
	Venue        *Venue           `json:"venue,omitempty"`
	Country      *Country         `json:"country,omitempty"`
	Coaches      []Coach          `json:"coaches,omitempty"`
	Seasons      []Season         `json:"seasons,omitempty"`
}

// TeamByID fetches a Team resource by ID. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) TeamByID(ctx context.Context, id int, includes []string, filters map[string][]int) (*Team, *ResponseDetails, error) {
	path := fmt.Sprintf(teamsURI+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ";")},
	}

	formatFilters(&values, filters)

	response := struct {
		Data         *Team          `json:"data"`
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

// TeamsBySeasonID fetches Team resources associated to a Season ID. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) TeamsBySeasonID(ctx context.Context, seasonID int, includes []string) ([]Team, *ResponseDetails, error) {
	path := fmt.Sprintf(teamsSeasonURI+"/%d", seasonID)

	values := url.Values{
		"include": {strings.Join(includes, ";")},
	}

	response := struct {
		Data         []Team         `json:"data"`
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
