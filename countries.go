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
	//Leagues      []League  `json:"leagues,omitempty"`
}

// Countries fetches Country resources. The endpoint used within this method is paginated, to select the required
// page use the 'page' method argument. Page information including current page and total page are included within
// the Meta struct. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) Countries(ctx context.Context, page int, includes []string) ([]Country, *Meta, error) {
	values := url.Values{
		"page":    {strconv.Itoa(page)},
		"include": {strings.Join(includes, ";")},
	}

	response := struct {
		Data []Country `json:"data"`
		Meta *Meta     `json:"meta"`
	}{}

	err := c.getResource(ctx, countriesURI, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// CountryByID fetches a Country resource by ID. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) CountryByID(ctx context.Context, id int, includes []string) (*Country, *Meta, error) {
	path := fmt.Sprintf(countriesURI+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ";")},
	}

	response := struct {
		Data *Country `json:"data"`
		Meta *Meta    `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
