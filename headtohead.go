package sportmonks

type HeadToHead struct {
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

//const headToHeadUri = "/api/v2.0/head2head"
//
//// Retrieve a multiple fixture resources linked to fixtures between two teams. Use the includes slice to enrich
//// the response data.
//func (c *HTTPClient) HeadToHead(teamOne, teamTwo int, includes []string) ([]HeadToHead, *Meta, error) {
//	url := fmt.Sprintf(headToHeadUri + "/%d/%d", teamOne, teamTwo)
//
//	response := new(HeadToHeadResponse)
//
//	err := c.handleRequest(url, includes, response)
//
//	if err != nil {
//		return nil, nil, err
//	}
//
//	return response.Data, &response.Meta, err
//}