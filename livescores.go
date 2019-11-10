package sportmonks

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

// LiveScores fetches multiple Fixture resources for live fixtures. The endpoint used within this method
// is paginated, to select the required page use the 'page' method argument. Page information including current page
// and total page are included within the Meta response. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) LiveScores(ctx context.Context, page int, includes []string, filters map[string][]int) ([]Fixture, *Meta, error) {
	values := url.Values{
		"page":    {strconv.Itoa(page)},
		"include": {strings.Join(includes, ",")},
	}

	formatFilters(&values, filters)

	response := struct {
		Data []Fixture `json:"data"`
		Meta *Meta     `json:"meta"`
	}{}

	err := c.getResource(ctx, liveScoresURI, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// LiveScoresInPlay fetches multiple Fixture resources for in play fixtures. Use the includes slice of string to enrich the
// response data.
func (c *HTTPClient) LiveScoresInPlay(ctx context.Context, includes []string, filters map[string][]int) ([]Fixture, *Meta, error) {
	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data []Fixture `json:"data"`
		Meta *Meta     `json:"meta"`
	}{}

	formatFilters(&values, filters)

	err := c.getResource(ctx, liveScoresInPlayURI, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
