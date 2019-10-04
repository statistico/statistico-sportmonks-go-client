package sportmonks

//import "strconv"

const leagueUri = "/api/v2.0/leagues"

// League struct
type League struct {
	ID              int    `json:"id"`
	Active          bool   `json:"active"`
	Type            string  `json:"type"`
	LegacyID        int    `json:"legacy_id"`
	CountryID       int    `json:"country_id"`
	LogoPath        string  `json:"logo_path"`
	Name            string `json:"name"`
	IsCup           bool   `json:"is_cup"`
	CurrentSeasonID int    `json:"current_season_id"`
	CurrentRoundID  int    `json:"current_round_id"`
	CurrentStageID  int    `json:"current_stage_id"`
	LiveStandings   bool   `json:"live_standings"`
	Coverage        LeagueCoverage `json:"coverage"`
	Country struct {
		Data *Country `json:"data"`
	} `json:"country, omitempty"`
	Season struct {
		Data *Season `json:"data"`
	} `json:"season"`
	Seasons struct {
		Data []Season `json:"data"`
	} `json:"seasons"`
}

//// Make a request to retrieve multiple league resources. The request endpoint executed within this method
//// is paginated, the first argument to this method allows the consumer to specify a page to request.
//// Use the includes slice to enrich the response data.
//func (c *HTTPClient) Leagues(page int, includes []string) ([]League, *Meta, error) {
//	response := new(LeaguesResponse)
//
//	err := c.handlePaginatedRequest(leagueUri, includes, page, response)
//
//	if err != nil {
//		return nil, nil, err
//	}
//
//	return response.Data, &response.Meta, err
//}
//
//// Retrieve a single league resource by ID. Use the includes slice to enrich the response data.
//func (c *HTTPClient) LeagueById(id int, includes []string) (*League, *Meta, error) {
//	url := leagueUri + "/" + strconv.Itoa(id)
//
//	response := new(LeagueResponse)
//
//	err := c.handleRequest(url, includes, response)
//
//	if err != nil {
//		return nil, nil, err
//	}
//
//	return &response.Data, &response.Meta, err
//}