package sportmonks

import (
	"context"
	"fmt"
	"net/url"
)

type Market struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Markets returns a slice of Market struct and supporting meta data
func (c *HTTPClient) Markets(ctx context.Context) ([]Market, *Meta, error) {
	response := struct {
		Data []Market `json:"data"`
		Meta *Meta    `json:"meta"`
	}{}

	err := c.getResource(ctx, marketsUri, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// MarketById returns a single Market struct and supporting meta data.
func (c *HTTPClient) MarketById(ctx context.Context, id int) (*Market, *Meta, error) {
	path := fmt.Sprintf(marketsUri+"/%d", id)

	response := struct {
		Data *Market `json:"data"`
		Meta *Meta   `json:"meta"`
	}{}

	err := c.getResource(ctx, path, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
