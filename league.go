package statistico

//import "strconv"
//
//const leagueUri = "/api/v2.0/leagues"
//
// League struct
type League struct {
	ID              int    `json:"id"`
	LegacyID        int    `json:"legacy_id"`
	CountryID       int    `json:"country_id"`
	Name            string `json:"name"`
	IsCup           bool   `json:"is_cup"`
	CurrentSeasonID int    `json:"current_season_id"`
	CurrentRoundID  int    `json:"current_round_id"`
	CurrentStageID  int    `json:"current_stage_id"`
	LiveStandings   bool   `json:"live_standings"`
	Coverage        struct {
		Predictions      bool `json:"predications"`
		TopScorerGoals   bool `json:"topscorer_goals"`
		TopScorerAssists bool `json:"topscorer_assists"`
		TopScorerCards   bool `json:"topscorer_cards"`
	} `json:"coverage"`
	Country struct {
		Data Country `json:"data"`
	} `json:"country, omitempty"`
	Season struct {
		Data Season `json:"data"`
	} `json:"season"`
	Seasons struct {
		Data []Season `json:"data"`
	} `json:"seasons"`
}
//
//// Leagues Response
//type LeaguesResponse struct {
//	Data []League `json:"data"`
//	Meta Meta     `json:"meta"`
//}
//
//// League Response
//type LeagueResponse struct {
//	Data League `json:"data"`
//	Meta Meta     `json:"meta"`
//}
//
//// Make a request to retrieve multiple league resources. The request endpoint executed within this method
//// is paginated, the second argument to this method allows the consumer to specify a page to request.
//
//// Use the includes slice to enrich the response data.
//func (c *Client) Leagues(page int, includes []string, retries int) (*LeaguesResponse, error) {
//	url := c.BaseURL + leagueUri + "?api_token=" + c.ApiKey
//
//	response := new(LeaguesResponse)
//
//	err := c.sendRequest(url, includes, page, response)
//
//	return response, err
//}
//
//// Retrieve a single continent resource by ID. Use the includes slice to enrich the response data.
//// - seasons
//func (c *Client) LeagueById(id int, includes []string) (*LeagueResponse, error) {
//	url := c.BaseURL + leagueUri + strconv.Itoa(id) + "?api_token=" + c.ApiKey
//
//	response := new(LeagueResponse)
//
//	err := c.sendRequest(url, includes, 0, response)
//
//	return response, err
//}