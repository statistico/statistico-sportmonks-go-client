package sportmonks

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

// LiveScores returns a slice of Fixture struct and supporting meta data. The endpoint used within this method
// is paginated, to select the required page use the 'page' method argument. Page information including current page
// and total page are included within the Meta response.

// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) LiveScores(ctx context.Context, page int, includes []string) ([]Fixture, *Meta, error) {
	values := url.Values{
		"page":    {strconv.Itoa(page)},
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data []Fixture `json:"data"`
		Meta *Meta     `json:"meta"`
	}{}

	err := c.getResource(ctx, liveScoresUri, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// LiveScoresInPlay returns a slice of Fixture struct and supporting meta data. Use the includes slice of string to
// enrich the response data.
func (c *HTTPClient) LiveScoresInPlay(ctx context.Context, includes []string) ([]Fixture, *Meta, error) {
	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data []Fixture `json:"data"`
		Meta *Meta     `json:"meta"`
	}{}

	err := c.getResource(ctx, liveScoresInPlayUri, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
