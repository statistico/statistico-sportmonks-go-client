package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Continent provides a struct representation of a Continent resource.
type Continent struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Countries []Country `json:"countries,omitempty"`
}

// Continents fetches Continent resources. The endpoint used within this method is paginated, to select the required
// page use the 'page' method argument. Page information including current page and total page are included within
// the Meta struct. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) Continents(ctx context.Context, page int, includes []string) ([]Continent, *Meta, error) {
	values := url.Values{
		"page":    {strconv.Itoa(page)},
		"include": {strings.Join(includes, ";")},
	}

	response := struct {
		Data []Continent `json:"data"`
		Meta *Meta       `json:"meta"`
	}{}

	err := c.getResource(ctx, continentsURI, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// ContinentByID fetches a Continent resource by ID. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) ContinentByID(ctx context.Context, id int, includes []string) (*Continent, *Meta, error) {
	path := fmt.Sprintf(continentsURI+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ";")},
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
