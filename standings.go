package sportmonks

//import "fmt"

const leagueStandingsUri = "/api/v2.0/standings/season"
const liveLeagueStandingsUri = "/api/v2.0/standings/season/live"

type LeagueStandings struct {
	Name      string `json:"name"`
	LeagueID  int    `json:"league_id"`
	SeasonID  int    `json:"season_id"`
	RoundID   int    `json:"round_id"`
	RoundName int    `json:"round_name"`
	Type      string `json:"type"`
	StageID   int    `json:"stage_id"`
	StageName string `json:"stage_name"`
	Resource  string `json:"resource"`
	Standings struct {
		Data []TeamLeagueStanding `json:"data"`
	} `json:"standings"`
	Season struct {
		Data *Season `json:"data"`
	} `json:"season"`
}

//// Retrieve a league standings resource by season ID. Use the includes slice to enrich the response data.
//func (c *HTTPClient) LeagueStandingsBySeasonId(seasonId int, include []string) (*LeagueStandings, *Meta, error) {
//	url := fmt.Sprintf(leagueStandingsUri +"/%d", seasonId)
//
//	response := new(LeagueStandingsResponse)
//
//	err := c.handleRequest(url, include, response)
//
//	if err != nil {
//		return nil, nil, err
//	}
//
//	return &response.Data, &response.Meta, err
//}
//
//// Retrieve live league standings resource by season ID. Use the includes slice to enrich the response data.
//func (c *HTTPClient) LiveLeagueStandingsBySeasonId(seasonId int, include []string) (*LeagueStandings, *Meta, error) {
//	url := fmt.Sprintf(liveLeagueStandingsUri +"/%d", seasonId)
//
//	response := new(LeagueStandingsResponse)
//
//	err := c.handleRequest(url, include, response)
//
//	if err != nil {
//		return nil, nil, err
//	}
//
//	return &response.Data, &response.Meta, err
//}