package sportmonks

import (
	"fmt"
)

const tvStationUri = "/api/v2.0/tvstations/fixture"

type TVStation struct {
	FixtureID int    `json:"fixture_id"`
	TVStation string `json:"tvstation"`
}

// Retrieve a multiple tv station resources by fixture ID.
func (c *ApiClient) TvStationsByFixtureId(fixtureId int) ([]TVStation, *Meta, error) {
	url := fmt.Sprintf(tvStationUri +"/%d", fixtureId)

	response := new(TVStationResponse)

	err := c.handleRequest(url, []string{}, response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, &response.Meta, err
}
