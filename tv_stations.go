package sportmonks

import (
	"context"
	"fmt"
	"net/url"
)

type TVStation struct {
	FixtureID int    `json:"fixture_id"`
	TVStation string `json:"tvstation"`
}

// TvStationsByFixtureId returns a slice of TVStation struct and supporting meta data for a Fixture
func (c *HTTPClient) TVStationsByFixtureId(ctx context.Context, fixtureId int) ([]TVStation, *Meta, error) {
	path := fmt.Sprintf(tvStationsUri + "/%d", fixtureId)

	response := struct {
		Data []TVStation `json:"data"`
		Meta *Meta    `json:"meta"`
	}{}

	err := c.getResource(ctx, path, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
