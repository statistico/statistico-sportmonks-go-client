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
	AggregatedAssistScorers struct {
		Data []AssistScorer `json:"data"`
	} `json:"aggregatedAssistscorers, omitempty"`
	AggregatedCardScorers struct {
		Data []CardScorer `json:"data"`
	} `json:"aggregatedCardscorers, omitempty"`
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
		Data *League `json:"data"`
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

// Make a request to retrieve multiple season resources. The request endpoint executed within this method
// is paginated, the first argument to this method allows the consumer to specify a page to request.
// Use the includes slice to enrich the response data.
func (c *ApiClient) Seasons(page int, includes []string) ([]Season, *Meta, error) {
	response := new(SeasonsResponse)

	err := c.handlePaginatedRequest(seasonUri, includes, page, response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, &response.Meta, err
}

// Retrieve a single season resource by ID. Use the includes slice to enrich the response data.
func (c *ApiClient) SeasonById(id int, includes []string) (*Season, *Meta, error) {
	url := seasonUri + "/" + strconv.Itoa(id)

	response := new(SeasonResponse)

	err := c.handleRequest(url, includes, response)

	if err != nil {
		return nil, nil, err
	}

	return &response.Data, &response.Meta, err
}
