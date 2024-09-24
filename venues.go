package sportmonks

import (
	"context"
	"fmt"
	"net/url"
)

type Venue struct {
	ID           int     `json:"id"`
	CountryID    int     `json:"country_id"`
	CityID       int     `json:"city_id"`
	Name         string  `json:"name"`
	Address      string  `json:"address"`
	Zipcode      *string `json:"zipcode"`
	Latitude     string  `json:"latitude"`
	Longitude    string  `json:"longitude"`
	Capacity     int     `json:"capacity"`
	ImagePath    string  `json:"image_path"`
	CityName     string  `json:"city_name"`
	Surface      string  `json:"surface"`
	NationalTeam bool    `json:"national_team"`
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
