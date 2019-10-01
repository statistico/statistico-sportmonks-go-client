package sportmonks

import "strconv"

const topScorersSeasonUri = "/api/v2.0/topscorers/season"

type TopScorers struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	LeagueID        int    `json:"league_id"`
	IsCurrentSeason bool   `json:"is_current_season"`
	CurrentRoundID  int    `json:"current_round_id"`
	CurrentStageID  int    `json:"current_stage_id"`
	Goalscorers     struct {
		Data []GoalScorer `json:"data"`
	} `json:"goalscorers"`
	AssistScorers struct {
		Data []AssistScorer `json:"data"`
	} `json:"assistscorers"`
	CardScorers struct {
		Data []CardScorer `json:"data"`
	} `json:"cardscorers"`
}

type AggregatedTopScorers struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	LeagueID        int    `json:"league_id"`
	IsCurrentSeason bool   `json:"is_current_season"`
	CurrentRoundID  int    `json:"current_round_id"`
	CurrentStageID  int    `json:"current_stage_id"`
	Goalscorers     struct {
		Data []GoalScorer `json:"data"`
	} `json:"aggregatedGoalscorers"`
	AssistScorers struct {
		Data []AssistScorer `json:"data"`
	} `json:"aggregatedAssistscorers"`
	CardScorers struct {
		Data []CardScorer `json:"data"`
	} `json:"aggregatedCardscorers"`
}

// Retrieve a single top scorers resource by season ID. Use the includes slice to enrich the response data.
func (c *ApiClient) TopScorersBySeasonId(seasonId int, includes []string) (*TopScorers, *Meta, error) {
	url := topScorersSeasonUri + "/" + strconv.Itoa(seasonId)

	response := new(TopScorersResponse)

	err := c.handleRequest(url, includes, response)

	if err != nil {
		return nil, nil, err
	}

	return &response.Data, &response.Meta, err
}

// Retrieve a single aggregated top scorers resource by season ID. Use the includes slice to enrich the response data.
func (c *ApiClient) AggregatedTopScorersBySeasonId(seasonId int, includes []string) (*AggregatedTopScorers, *Meta, error) {
	url := topScorersSeasonUri + "/" + strconv.Itoa(seasonId) + "/aggregated"

	response := new(AggregatedTopScorersResponse)

	err := c.handleRequest(url, includes, response)

	if err != nil {
		return nil, nil, err
	}

	return &response.Data, &response.Meta, err
}