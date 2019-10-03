package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const continentsUri = "/continents"

// Continent struct
type Continent struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Countries Countries `json:"countries, omitempty"`
}

// CountryData returns a Country slice
func (c *Continent) CountryData() []Country {
	return c.Countries.Data
}

func (c *HTTPClient) Continents(ctx context.Context, page int, includes []string) ([]Continent, *Meta, error) {
	values := url.Values{
		"include": {strings.Join(includes, ",")},
		"page":    {strconv.Itoa(page)},
	}

	response := struct {
		Data []Continent `json:"data"`
		Meta *Meta       `json:"meta"`
	}{}

	err := c.getResource(ctx, continentsUri, values, response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

func (c *HTTPClient) ContinentById(ctx context.Context, id int, includes []string) (*Continent, *Meta, error) {
	path := fmt.Sprintf(continentsUri+"/%d", id)

	response := struct {
		Data *Continent `json:"data"`
		Meta *Meta      `json:"meta"`
	}{}

	err := c.getResource(ctx, path, url.Values{}, response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
