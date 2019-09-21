package sportmonks

import (
	"fmt"
	"strconv"
)

const stageUri = "/api/v2.0/stages"
const stageSeasonUri = "/api/v2.0/stages/season"

type Stage struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	Type         *string      `json:"type"`
	LeagueID     int         `json:"league_id"`
	SeasonID     int         `json:"season_id"`
	SortOrder    *string     `json:"sort_order"`
	HasStandings *bool        `json:"has_standings"`
	Fixtures        struct {
		Data []Fixture `json:"data"`
	} `json:"fixtures, omitempty"`
	League struct {
		Data *League `json:"data"`
	} `json:"league, omitempty"`
	Results struct {
		Data []Fixture `json:"data"`
	} `json:"results, omitempty"`
	Season *Season `json:"season"`
}

// Retrieve a single stage resource by ID. Use the includes slice to enrich the response data.
func (c *ApiClient) StageById(id int, includes []string) (*Stage, *Meta, error) {
	url := fmt.Sprintf(stageUri +"/%s", strconv.Itoa(id))

	response := new(StageResponse)

	err := c.handleRequest(url, includes, response)

	if err != nil {
		return nil, nil, err
	}

	return &response.Data, &response.Meta, err
}

// Make a request to retrieve multiple venue resources for a given season. Use the includes slice to enrich the
// response data.
func (c *ApiClient) StagesBySeasonId(id int, includes []string) ([]Stage, *Meta, error) {
	url := stageSeasonUri + "/" + strconv.Itoa(id)

	response := new(StagesSeasonResponse)

	err := c.handleRequest(url, []string{}, response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, &response.Meta, err
}