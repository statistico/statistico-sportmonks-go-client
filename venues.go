package sportmonks

import (
	"context"
	"fmt"
	"net/url"
)

// Venue provides a struct representation of a Venue resource
type Venue struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Surface     string  `json:"surface"`
	Address     *string `json:"address"`
	City        string  `json:"city"`
	Capacity    int     `json:"capacity"`
	ImagePath   string  `json:"image_path"`
	Coordinates *string `json:"coordinates"`
}

// VenueByID fetches a Venue resource by ID.
func (c *HTTPClient) VenueByID(ctx context.Context, id int) (*Venue, *Meta, error) {
	path := fmt.Sprintf(venuesURI+"/%d", id)

	response := struct {
		Data *Venue `json:"data"`
		Meta *Meta  `json:"meta"`
	}{}

	err := c.getResource(ctx, path, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// VenuesBySeasonID fetches a Venue resource by season ID.
func (c *HTTPClient) VenuesBySeasonID(ctx context.Context, id int) ([]Venue, *Meta, error) {
	path := fmt.Sprintf(venuesSeasonURI+"/%d", id)

	response := struct {
		Data []Venue `json:"data"`
		Meta *Meta   `json:"meta"`
	}{}

	err := c.getResource(ctx, path, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
