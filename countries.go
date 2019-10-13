package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Country struct
type Country struct {
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	Extra         CountryExtra  `json:"extra"`
	ContinentData ContinentData `json:"continent, omitempty"`
	LeaguesData   LeaguesData   `json:"leagues, omitempty"`
}

// Continent returns a Continent struct associated to a Country
func (c *Country) Continent() *Continent {
	return c.ContinentData.Data
}

// Leagues returns a League struct slice associated to a Country
func (c *Country) Leagues() []League {
	return c.LeaguesData.Data
}

// Countries returns a slice of Country struct and supporting meta data. The endpoint used within this method
// is paginated, to select the required page use the 'page' method argument. Page information including current page
// and total page are included within the Meta response.

// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) Countries(ctx context.Context, page int, includes []string) ([]Country, *Meta, error) {
	values := url.Values{
		"page":    {strconv.Itoa(page)},
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data []Country `json:"data"`
		Meta *Meta     `json:"meta"`
	}{}

	err := c.getResource(ctx, countriesUri, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// CountryById sends a request and returns a single Country struct.

// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) CountryById(ctx context.Context, id int, includes []string) (*Country, *Meta, error) {
	path := fmt.Sprintf(countriesUri+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
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
