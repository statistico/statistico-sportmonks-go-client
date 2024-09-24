package sportmonks

import (
	"context"
	"fmt"
	"net/url"
)

// Bookmaker provides a struct representation of a Bookmaker resource
type Bookmaker struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LegacyID int    `json:"legacy_id"`
}

// Bookmakers fetches a slice of Bookmaker resource struct.
func (c *HTTPClient) Bookmakers(ctx context.Context) ([]Bookmaker, *Meta, error) {
	response := struct {
		Data []Bookmaker `json:"data"`
		Meta *Meta       `json:"meta"`
	}{}

	err := c.getResource(ctx, bookmakersURI, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// BookmakerByID fetches a single Bookmaker resource by ID.
func (c *HTTPClient) BookmakerByID(ctx context.Context, id int) (*Bookmaker, *Meta, error) {
	path := fmt.Sprintf(bookmakersURI+"/%d", id)

	response := struct {
		Data *Bookmaker `json:"data"`
		Meta *Meta      `json:"meta"`
	}{}

	err := c.getResource(ctx, path, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
