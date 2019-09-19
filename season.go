package sportmonks

import "strconv"

const seasonUri = "/api/v2.0/seasons"

// Season struct
type Season struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	LeagueID        int    `json:"league_id"`
	IsCurrentSeason bool   `json:"is_current_season"`
	CurrentRoundID  *int   `json:"current_round_id"`
	CurrentStageID  *int   `json:"current_stage_id"`
	AggregatedGoalScorers struct {
		Data []GoalScorer `json:"data"`
	} `json:"aggregatedGoalscorers, omitempty"`
	AssistScorers struct {
		Data []AssistScorer `json:"data"`
	} `json:"assistscorers, omitempty"`
	CardScorers struct {
		Data []CardScorer `json:"data"`
	} `json:"cardscorers, omitempty"`
	Fixtures        struct {
		Data []Fixture `json:"data"`
	} `json:"fixtures, omitempty"`
	GoalScorers struct {
		Data []GoalScorer `json:"data"`
	} `json:"goalscorers, omitempty"`
	Groups struct {
		Data []Group `json:"data"`
	} `json:"groups, omitempty"`
	League struct {
		Data League `json:"data"`
	} `json:"league, omitempty"`
	Results struct {
		Data []Fixture `json:"data"`
	} `json:"results, omitempty"`
	Rounds struct {
		Data []Round `json:"data"`
	} `json:"rounds, omitempty"`
	Stages struct {
		Data []Stage `json:"data"`
	} `json:"stages, omitempty"`
	Upcoming struct {
		Data []Fixture `json:"data"`
	} `json:"upcoming, omitempty"`
}

// Seasons Response
type SeasonsResponse struct {
	Data []Season `json:"data"`
	Meta Meta     `json:"meta"`
}

// Season Response
type SeasonResponse struct {
	Data Season `json:"data"`
	Meta Meta     `json:"meta"`
}

// Make a request to retrieve multiple season resources. The request endpoint executed within this method
// is paginated, the second argument to this method allows the consumer to specify a page to request.
// Use the includes slice to enrich the response data, includes for this endpoint are:
// - aggregatedGoalscorers
// - assistscorers
// - cardscorers
// - fixtures
// - goalscorers
// - groups
// - league
// - results
// - rounds
// - stages
// - upcoming
func (c *Client) Seasons(page int, includes []string, retries int) (*SeasonsResponse, error) {
	url := c.BaseURL + seasonUri + "?api_token=" + c.ApiKey

	response := new(SeasonsResponse)

	err := c.sendRequest(url, includes, page, response)

	return response, err
}

// Retrieve a single continent season by ID. Use the includes slice to enrich the response data, includes
// for this endpoint are:
// - aggregatedGoalscorers
// - assistscorers
// - cardscorers
// - fixtures
// - goalscorers
// - groups
// - league
// - results
// - rounds
// - stages
// - upcoming
func (c *Client) SeasonById(id int, includes []string) (*SeasonResponse, error) {
	url := c.BaseURL + seasonUri + strconv.Itoa(id) + "?api_token=" + c.ApiKey

	response := new(SeasonResponse)

	err := c.sendRequest(url, includes, 0, response)

	return response, err
}
