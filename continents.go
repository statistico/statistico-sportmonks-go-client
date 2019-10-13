package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Continent struct
type Continent struct {
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	CountriesData CountriesData `json:"countries, omitempty"`
}

// Countries returns a Country struct slice associated to a Continent
func (c *Continent) Countries() []Country {
	return c.CountriesData.Data
}

// Continents returns a slice of Continent struct and supporting meta data. The endpoint used within this method
// is paginated, to select the required page use the 'page' method argument. Page information including current page
// and total page are included within the Meta response.

// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) Continents(ctx context.Context, page int, includes []string) ([]Continent, *Meta, error) {
	values := url.Values{
		"page":    {strconv.Itoa(page)},
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data []Continent `json:"data"`
		Meta *Meta       `json:"meta"`
	}{}

	err := c.getResource(ctx, continentsUri, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// ContinentById sends a request and returns a single Continent struct.

// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) ContinentById(ctx context.Context, id int, includes []string) (*Continent, *Meta, error) {
	path := fmt.Sprintf(continentsUri+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data *Continent `json:"data"`
		Meta *Meta      `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
