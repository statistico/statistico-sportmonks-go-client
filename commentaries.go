package sportmonks

import (
	"context"
	"fmt"
	"net/url"
)

// Commentary provides a struct representation of a Commentary resource.
type Commentary struct {
	FixtureID   int    `json:"fixture_id"`
	Important   bool   `json:"important"`
	Order       int    `json:"order"`
	Goal        bool   `json:"goal"`
	Minute      int    `json:"minute"`
	ExtraMinute *int   `json:"extra_minute"`
	Comment     string `json:"comment"`
}

// CommentariesByFixtureID fetches Commentary resources associated to a fixture.
func (c *HTTPClient) CommentariesByFixtureID(ctx context.Context, fixtureID int) ([]Commentary, *Meta, error) {
	path := fmt.Sprintf(commentariesFixtureURI+"/%d", fixtureID)

	response := struct {
		Data []Commentary `json:"data"`
		Meta *Meta        `json:"meta"`
	}{}

	err := c.getResource(ctx, path, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
