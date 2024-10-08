package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Country provides a struct representation of a Country resource.
type Country struct {
	ID           int        `json:"id"`
	ContinentID  int        `json:"continent_id"`
	Name         string     `json:"name"`
	OfficialName string     `json:"official_name"`
	FifaName     string     `json:"fifa_name"`
	Iso2         string     `json:"iso2"`
	Iso3         string     `json:"iso3"`
	Latitude     string     `json:"latitude"`
	Longitude    string     `json:"longitude"`
	Borders      []string   `json:"borders"`
	ImagePath    string     `json:"image_path"`
	Continent    *Continent `json:"continent,omitempty"`
	Leagues      []League   `json:"leagues,omitempty"`
}

// Countries fetches Country resources. The endpoint used within this method is paginated, to select the required
// page use the 'page' method argument. Pagination information including current page and count are included within
// the Pagination struct with the ResponseDetails struct. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) Countries(ctx context.Context, page int, includes []string) ([]Country, *ResponseDetails, error) {
	values := url.Values{
		"page":    {strconv.Itoa(page)},
		"include": {strings.Join(includes, ";")},
	}

	response := struct {
		Data         []Country      `json:"data"`
		Pagination   Pagination     `json:"pagination"`
		Subscription []Subscription `json:"subscription"`
		RateLimit    RateLimit      `json:"rate_limit"`
		TimeZone     string         `json:"timezone"`
	}{}

	err := c.getResource(ctx, countriesURI, values, &response)

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

// CountryByID fetches a Country resource by ID. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) CountryByID(ctx context.Context, id int, includes []string) (*Country, *ResponseDetails, error) {
	path := fmt.Sprintf(countriesURI+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ";")},
	}

	response := struct {
		Data         *Country       `json:"data"`
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
