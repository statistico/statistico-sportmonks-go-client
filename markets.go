package sportmonks

import (
	"context"
	"fmt"
	"net/url"
)

// Market provides a struct representation of a Market resource.
type Market struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Markets fetches Market resources.
func (c *HTTPClient) Markets(ctx context.Context) ([]Market, *Meta, error) {
	response := struct {
		Data []Market `json:"data"`
		Meta *Meta    `json:"meta"`
	}{}

	err := c.getResource(ctx, marketsURI, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// MarketByID fetches Market resource by ID.
func (c *HTTPClient) MarketByID(ctx context.Context, id int) (*Market, *Meta, error) {
	path := fmt.Sprintf(marketsURI+"/%d", id)

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
