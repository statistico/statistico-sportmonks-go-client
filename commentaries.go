package sportmonks

import (
	"context"
	"fmt"
	"net/url"
)

const commentariesFixtureUri = "/commentaries/fixture"

type Commentary struct {
	FixtureID   int         `json:"fixture_id"`
	Important   bool        `json:"important"`
	Order       int         `json:"order"`
	Goal        bool        `json:"goal"`
	Minute      int         `json:"minute"`
	ExtraMinute *int `json:"extra_minute"`
	Comment     string      `json:"comment"`
}

// CommentariesByFixtureId returns a slice of Commentary associated to a fixture
func (c *HTTPClient) CommentariesByFixtureId(ctx context.Context, fixtureId int) ([]Commentary, *Meta, error) {
	path := fmt.Sprintf(commentariesFixtureUri +"/%d", fixtureId)

	response := struct {
		Data []Commentary `json:"data"`
		Meta *Meta `json:"meta"`
	}{}

	err := c.getResource(ctx, path, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
