package sportmonks

const fixtureUri = "/api/v2.0/fixtures/"

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
	WinningOddsCalculated bool               `json:"winning_odds_calculated"`
	Formations            Formations         `json:"formations"`
	Scores                Scores             `json:"scores"`
	Time                  FixtureTime        `json:"time"`
	Coaches               Coaches            `json:"coaches"`
	Standings             Standings          `json:"standings"`
	Assistants            Assistants         `json:"assistants"`
	Leg                   *string            `json:"leg"`
	Lineup                Lineup             `json:"lineup"`
	Bench                 Lineup             `json:"bench"`
	TeamStats             TeamsStats         `json:"stats"`
	Goals                 GoalEvents         `json:"goals"`
	Subs                  SubstitutionEvents `json:"substitutions"`
}

//// Retrieve a single continent fixture by ID. Use the includes slice to enrich the response data.
//func (c *ApiClient) FixtureById(id int, includes []string) (*Fixture, *Meta, error) {
//	url := continentUri + "/" + strconv.Itoa(id)
//
//	response := new(ContinentResponse)
//
//	err := c.handleRequest(url, includes, response)
//
//	if err != nil {
//		return nil, nil, err
//	}
//
//	return &response.Data, &response.Meta, err
//}
