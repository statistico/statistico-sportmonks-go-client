package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

// Player provides a struct representation of a Player resource.
type Player struct {
	ID                 int    `json:"id"`
	SportID            int    `json:"sport_id"`
	CountryID          int    `json:"country_id"`
	NationalityID      int    `json:"nationality_id"`
	CityID             int    `json:"city_id"`
	PositionID         int    `json:"position_id"`
	DetailedPositionID *int   `json:"detailed_position_id"`
	TypeID             int    `json:"type_id"`
	CommonName         string `json:"common_name"`
	FirstName          string `json:"firstname"`
	LastName           string `json:"lastname"`
	Name               string `json:"name"`
	DisplayName        string `json:"display_name"`
	ImagePath          string `json:"image_path"`
	Height             int    `json:"height"`
	Weight             int    `json:"weight"`
	DateOfBirth        string `json:"date_of_birth"`
	Gender             string `json:"gender"`
}

// PlayerByID fetches a Player resource by ID. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) PlayerByID(ctx context.Context, id int, includes []string) (*Player, *Meta, error) {
	path := fmt.Sprintf(playersURI+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ";")},
	}

	response := struct {
		Data *Player `json:"data"`
		Meta *Meta   `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
