package sportmonks

const fixtureUri = "/api/v2.0/fixtures"
const fixtureDateUri = "/api/v2.0/fixtures/date"
const fixtureBetweenUri = "/api/v2.0/fixtures/between"
const fixtureMultiUri = "/api/v2.0/fixtures/between"
const fixtureUpdatesUri = "/api/v2.0/fixtures/between"

type Fixture struct {
	ID                    int                `json:"id"`
	LeagueID              int                `json:"league_id"`
	SeasonID              int                `json:"season_id"`
	StageID               *int               `json:"stage_id"`
	RoundID               *int               `json:"round_id"`
	GroupID               *int               `json:"group_id"`
	AggregateID           *int               `json:"aggregate_id"`
	VenueID               *int               `json:"venue_id"`
	RefereeID             *int               `json:"referee_id"`
	LocalTeamID           int                `json:"localteam_id"`
	VisitorTeamID         int                `json:"visitorteam_id"`
	WeatherReport         WeatherReport      `json:"weather_report"`
	Commentaries          *bool              `json:"commentaries"`
	Attendance            *int               `json:"attendance"`
	Pitch                 *string            `json:"pitch"`
	NeutralVenue          *bool              `json:"neutral_venue"`
	WinningOddsCalculated bool               `json:"winning_odds_calculated"`
	Formations            Formations         `json:"formations"`
	Scores                Scores             `json:"scores"`
	Time                  FixtureTime        `json:"time"`
	Coaches               Coaches            `json:"coaches"`
	Standings             Standings          `json:"standings"`
	Assistants            Assistants         `json:"assistants"`
	Leg                   *string            `json:"leg"`
	Colors                []TeamColors       `json:"colors"`
	Deleted               bool			     `json:"deleted"`
	Cards				  CardEvents         `json:"cards"`
	Corners               CornerEvents       `json:"corners"`
	Goals                 GoalEvents         `json:"goals"`
	Events                MatchEvents        `json:"events"`
	Lineup                Lineup             `json:"lineup"`
	Bench                 Lineup             `json:"bench"`
	LocalTeam struct {
		Data Team `json:"data"`
	} `json:"localTeam, omitempty"`
	Substitutions         SubstitutionEvents `json:"substitutions"`
	Sidelined struct {
		Data []Sidelined `json:"data"`
	} `json:"sidelined"`
	VisitorTeam struct {
		Data Team `json:"data"`
	} `json:"visitorTeam, omitempty"`
	Comments struct {
		Data []Commentary `json:"data"`
	} `json:"comments, omitempty"`
	TvStations struct {
		Data []TVStation `json:"data"`
	} `json:"tvstations"`
	Highlights struct {
		Data []VideoHighlights `json:"data"`
	} `json:"highlights, omitempty"`
	Round *Round `json:"round"`
	Stage *Stage `json:"stage"`
	Referee *Referee `json:"referee"`
	Venue *Venue `json:"venue"`
	Odds struct {
		Data []MatchOdds `json:"data"`
	} `json:"odds, omitempty"`
	InPlayOdds struct {
		Data []MatchOdds `json:"data"`
	} `json:"inplayOdds, omitempty"`
	FlatOdds struct {
		Data []MatchOdds `json:"data"`
	} `json:"flatOdds, omitempty"`
	LocalCoach *Coach `json:"localCoach"`
	VisitorCoach *Coach `json:"visitorCoach"`
	TeamStats             TeamsStats         `json:"stats"`
}

//// Retrieve a single fixture resource by ID. Use the includes slice to enrich the response data.
//func (c *HTTPClient) FixtureById(id int, includes []string) (*Fixture, *Meta, error) {
//	url := fixtureUri + "/" + strconv.Itoa(id)
//
//	response := new(FixtureResponse)
//
//	err := c.handleRequest(url, includes, response)
//
//	if err != nil {
//		return nil, nil, err
//	}
//
//	return &response.Data, &response.Meta, err
//}
//
//// Retrieve a multiple fixture resources by ID. Use the includes slice to enrich the response data.
//func (c *HTTPClient) FixturesById(ids []string, includes []string) ([]Fixture, *Meta, error) {
//	url := fixtureMultiUri + "/" + strings.Join(ids, ",")
//
//	response := new(FixturesResponse)
//
//	err := c.handleRequest(url, includes, response)
//
//	if err != nil {
//		return nil, nil, err
//	}
//
//	return response.Data, &response.Meta, err
//}
//
//// Retrieve a multiple fixture resources for a given date. The `date` string argument is expected in `YYYY-mm-dd` format
//// i.e. `2019-03-12`. Use the includes slice to enrich the response data.
//func (c *HTTPClient) FixturesByDate(date string, includes []string) ([]Fixture, *Meta, error) {
//	url := fixtureDateUri + "/" + date
//
//	response := new(FixturesResponse)
//
//	err := c.handleRequest(url, includes, response)
//
//	if err != nil {
//		return nil, nil, err
//	}
//
//	return response.Data, &response.Meta, err
//}
//
//// Retrieve a multiple fixture resources between two dates. The `dateFrom` and `dateTo` string arguments are expected
//// in `YYYY-mm-dd` format i.e. `2019-03-12`. Use the includes slice to enrich the response data.
//func (c *HTTPClient) FixturesBetween(dateFrom, dateTo string, includes []string) ([]Fixture, *Meta, error) {
//	url := fmt.Sprintf(fixtureBetweenUri + "/between/%s/%s", dateFrom, dateTo)
//
//	response := new(FixturesResponse)
//
//	err := c.handleRequest(url, includes, response)
//
//	if err != nil {
//		return nil, nil, err
//	}
//
//	return response.Data, &response.Meta, err
//}
//
//// Retrieve a multiple fixture resources between two dates for a given team. The `dateFrom` and `dateTo` string
//// arguments are expected in `YYYY-mm-dd` format i.e. `2019-03-12`. Use the includes slice to enrich the response data.
//func (c *HTTPClient) FixturesBetweenForTeam(dateFrom, dateTo string, teamId int, includes []string) ([]Fixture, *Meta, error) {
//	url := fmt.Sprintf(fixtureBetweenUri + "/between/%s/%s/%d", dateFrom, dateTo, teamId)
//
//	response := new(FixturesResponse)
//
//	err := c.handleRequest(url, includes, response)
//
//	if err != nil {
//		return nil, nil, err
//	}
//
//	return response.Data, &response.Meta, err
//}
//
//// Retrieve a multiple fixture resources that have recently been updated. Use the includes slice to enrich the response
//// data.
//func (c *HTTPClient) FixturesLastUpdated(includes []string) ([]Fixture, *Meta, error) {
//	response := new(FixturesResponse)
//
//	err := c.handleRequest(fixtureUpdatesUri, includes, response)
//
//	if err != nil {
//		return nil, nil, err
//	}
//
//	return response.Data, &response.Meta, err
//}