package sportmonks

import "strconv"

const roundUri = "/api/v2.0/rounds/"
const roundSeasonUri = "/api/v2.0/rounds/season/"

type Round struct {
	ID       int    `json:"id"`
	Name     int    `json:"name"`
	LeagueID int    `json:"league_id"`
	SeasonID int    `json:"season_id"`
	StageID  int    `json:"stage_id"`
	Start    string `json:"start"`
	End      string `json:"end"`
	Fixtures struct {
		Data []Fixture `json:"data"`
	} `json:"fixtures, omitempty"`
	League  *League `json:"league"`
	Results struct {
		Data []Fixture `json:"data"`
	} `json:"results, omitempty"`
	Season *Season `json:"season"`
}

// Retrieve a single round resource by ID. Use the includes slice to enrich the response data.
func (c *ApiClient) RoundById(id int, includes []string) (*Round, *Meta, error) {
	url := roundUri + "/" + strconv.Itoa(id)

	response := new(RoundResponse)

	err := c.handleRequest(url, includes, response)

	if err != nil {
		return nil, nil, err
	}

	return &response.Data, &response.Meta, err
}

// Make a request to retrieve multiple round resources for a given season. Use the includes slice to enrich the
// response data.
func (c *ApiClient) RoundsBySeasonId(id int, includes []string) ([]Round, *Meta, error) {
	url := roundSeasonUri + "/" + strconv.Itoa(id)

	response := new(RoundsResponse)

	err := c.handleRequest(url, includes, response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, &response.Meta, err
}