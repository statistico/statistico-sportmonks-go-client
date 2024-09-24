package sportmonks

import (
	"context"
	"fmt"
	"net/url"
)

// Coach provides a struct representation of a Coach resource
type Coach struct {
	ID            int    `json:"id"`
	PlayerID      int    `json:"player_id"`
	SportID       int    `json:"sport_id"`
	CountryID     int    `json:"country_id"`
	NationalityID int    `json:"nationality_id"`
	CityID        *int   `json:"city_id"`
	CommonName    string `json:"common_name"`
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	Name          string `json:"name"`
	DisplayName   string `json:"display_name"`
	ImagePath     string `json:"image_path"`
	Height        *int   `json:"height"`
	Weight        *int   `json:"weight"`
	DateOfBirth   string `json:"date_of_birth"`
	Gender        string `json:"gender"`
}

// CoachByID fetches a Coach resource by ID.
func (c *HTTPClient) CoachByID(ctx context.Context, id int) (*Coach, *Meta, error) {
	path := fmt.Sprintf(coachesURI+"/%d", id)

	response := struct {
		Data *Coach `json:"data"`
		Meta *Meta  `json:"meta"`
	}{}

	err := c.getResource(ctx, path, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
