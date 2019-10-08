package sportmonks

const seasonsUri = "/seasons"

// Season struct
type Season struct {
	ID                         int                        `json:"id"`
	Name                       string                     `json:"name"`
	LeagueID                   int                        `json:"league_id"`
	IsCurrentSeason            bool                       `json:"is_current_season"`
	CurrentRoundID             *int                       `json:"current_round_id"`
	CurrentStageID             *int                       `json:"current_stage_id"`
	AggregatedAssistScorerData AggregatedAssistScorerData `json:"aggregatedAssistscorers, omitempty"`
	AggregatedCardScorerData   AggregatedCardScorerData   `json:"aggregatedCardscorers, omitempty"`
	AggregatedGoalScorerData   AggregatedGoalScorerData   `json:"aggregatedGoalscorers, omitempty"`
	AssistScorerData           AssistScorerData           `json:"assistscorers, omitempty"`
	CardScorerData             CardScorerData             `json:"cardscorers, omitempty"`
	FixturesData               FixturesData               `json:"fixtures, omitempty"`
	GoalScorerData             GoalScorerData             `json:"goalscorers, omitempty"`
	GroupsData                 GroupsData                 `json:"groups, omitempty"`
	LeagueData                 LeagueData                 `json:"league, omitempty"`
	ResultsData                FixturesData               `json:"results, omitempty"`
	RoundsData                 RoundsData                 `json:"rounds, omitempty"`
	StagesData                 StagesData                 `json:"stages, omitempty"`
	Upcoming                   FixturesData               `json:"upcoming, omitempty"`
}

//
//// Make a request to retrieve multiple season resources. The request endpoint executed within this method
//// is paginated, the first argument to this method allows the consumer to specify a page to request.
//// Use the includes slice to enrich the response data.
//func (c *HTTPClient) Seasons(page int, includes []string) ([]Season, *Meta, error) {
//	response := new(SeasonsResponse)
//
//	err := c.handlePaginatedRequest(seasonUri, includes, page, response)
//
//	if err != nil {
//		return nil, nil, err
//	}
//
//	return response.Data, &response.Meta, err
//}
//
//// Retrieve a single season resource by ID. Use the includes slice to enrich the response data.
//func (c *HTTPClient) SeasonById(id int, includes []string) (*Season, *Meta, error) {
//	url := seasonUri + "/" + strconv.Itoa(id)
//
//	response := new(SeasonResponse)
//
//	err := c.handleRequest(url, includes, response)
//
//	if err != nil {
//		return nil, nil, err
//	}
//
//	return &response.Data, &response.Meta, err
//}
