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
	ID           int    `json:"id"`
	SportID      int    `json:"sport_id"`
	CountryID    int    `json:"country_id"`
	Name         string `json:"name"`
	Active       bool   `json:"active"`
	ShortCode    string `json:"short_code"`
	ImagePath    string `json:"image_path"`
	Type         string `json:"type"`
	SubType      string `json:"sub_type"`
	LastPlayedAt string `json:"last_played_at"`
	Category     int    `json:"category"`
	HasJerseys   bool   `json:"has_jerseys"`
}

// Leagues fetches League resources. The endpoint used within this method is paginated, to select the required
// page use the 'page' method argument. Page information including current page and total page are included
// within the Meta response. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) Leagues(ctx context.Context, page int, includes []string) ([]League, *Pagination, []Subscription, error) {
	values := url.Values{
		"page":    {strconv.Itoa(page)},
		"include": {strings.Join(includes, ";")},
	}

	response := struct {
		Data         []League       `json:"data"`
		Pagination   *Pagination    `json:"pagination"`
		Subscription []Subscription `json:"subscription"`
	}{}

	err := c.getResource(ctx, leaguesURI, values, &response)

	if err != nil {
		return nil, nil, nil, err
	}

	return response.Data, response.Pagination, response.Subscription, err
}

// LeagueByID fetches League resources by ID. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) LeagueByID(ctx context.Context, id int, includes []string) (*League, *Pagination, []Subscription, error) {
	path := fmt.Sprintf(leaguesURI+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ";")},
	}

	response := struct {
		Data         *League        `json:"data"`
		Pagination   *Pagination    `json:"pagination"`
		Subscription []Subscription `json:"subscription"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, nil, err
	}

	return response.Data, response.Pagination, response.Subscription, err
}
