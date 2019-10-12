package sportmonks

import (
	"context"
	"fmt"
	"net/url"
)

type Venue struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Surface  string  `json:"surface"`
	Address  *string `json:"address"`
	City     string  `json:"city"`
	Capacity int     `json:"capacity"`
	ImagePath    string  `json:"image_path"`
	Coordinates *string `json:"coordinates"`
}

// VenueById returns a single Venue struct and supporting meta data.
func (c *HTTPClient) VenueById(ctx context.Context, id int) (*Venue, *Meta, error) {
	path := fmt.Sprintf(venuesUri + "/%d", id)

	response := struct {
		Data *Venue `json:"data"`
		Meta *Meta    `json:"meta"`
	}{}

	err := c.getResource(ctx, path, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// VenuesBySeasonId returns a slice of Venue struct and supporting meta data associated to a Season
func (c *HTTPClient) VenuesBySeasonId(ctx context.Context, id int) ([]Venue, *Meta, error) {
	path := fmt.Sprintf(venuesSeasonUri + "/%d", id)

	response := struct {
		Data []Venue `json:"data"`
		Meta *Meta    `json:"meta"`
	}{}

	err := c.getResource(ctx, path, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
