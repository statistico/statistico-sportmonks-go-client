package sportmonks

import (
	"context"
	"fmt"
	"net/url"
)

// TVStation provides a struct representation of a TVStation resource.
type TVStation struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	ImagePath string `json:"image_path"`
	Type      string `json:"type"`
	RelatedID *int   `json:"related_id"`
}

// TVStationsByFixtureID fetches TVStation resources for a fixture ID.
func (c *HTTPClient) TVStationsByFixtureID(ctx context.Context, fixtureID int) ([]TVStation, *Meta, error) {
	path := fmt.Sprintf(tvStationsURI+"/%d", fixtureID)

	response := struct {
		Data []TVStation `json:"data"`
		Meta *Meta       `json:"meta"`
	}{}

	err := c.getResource(ctx, path, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
