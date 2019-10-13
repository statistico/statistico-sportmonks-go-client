package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"
)

const dateFormat = "2006-01-02"

type Fixture struct {
	ID                    int                 `json:"id"`
	LeagueID              int                 `json:"league_id"`
	SeasonID              int                 `json:"season_id"`
	StageID               *int                `json:"stage_id"`
	RoundID               *int                `json:"round_id"`
	GroupID               *int                `json:"group_id"`
	AggregateID           *int                `json:"aggregate_id"`
	VenueID               *int                `json:"venue_id"`
	RefereeID             *int                `json:"referee_id"`
	LocalTeamID           int                 `json:"localteam_id"`
	VisitorTeamID         int                 `json:"visitorteam_id"`
	WinnerTeamID          *int                `json:"winner_team_id"`
	WeatherReport         WeatherReport       `json:"weather_report"`
	Commentaries          *bool               `json:"commentaries"`
	Attendance            *int                `json:"attendance"`
	Pitch                 *string             `json:"pitch"`
	Details               *string             `json:"details"`
	NeutralVenue          *bool               `json:"neutral_venue"`
	WinningOddsCalculated bool                `json:"winning_odds_calculated"`
	Formations            Formations          `json:"formations"`
	Scores                Scores              `json:"scores"`
	Time                  FixtureTime         `json:"time"`
	Coaches               Coaches             `json:"coaches"`
	Standings             TeamStandings       `json:"standings"`
	Assistants            Assistants          `json:"assistants"`
	Leg                   *string             `json:"leg"`
	Colors                TeamColors          `json:"colors"`
	Deleted               bool                `json:"deleted"`
	CardEvents            CardEvents          `json:"cards"`
	CornerEvents          CornerEvents        `json:"corners"`
	GoalEvents            GoalEvents          `json:"goals"`
	MatchEvents           MatchEvents         `json:"events"`
	LineupData            PlayerLineupData    `json:"lineup"`
	BenchData             PlayerLineupData    `json:"bench"`
	LocalTeamData         TeamData            `json:"localTeam, omitempty"`
	VisitorTeamData       TeamData            `json:"visitorTeam, omitempty"`
	SubstitutionData      SubstitutionEvents  `json:"substitutions"`
	SidelinedData         SidelinedData       `json:"sidelined"`
	CommentsData          MatchCommentary     `json:"comments, omitempty"`
	TVStationData         MatchTVStations     `json:"tvstations"`
	HighlightData         MatchHighlights     `json:"highlights, omitempty"`
	LeagueData            LeagueData          `json:"league, omitempty"`
	RoundData             RoundData           `json:"round, omitempty"`
	StageData             StageData           `json:"stage, omitempty"`
	RefereeData           MatchOfficialData   `json:"referee, omitempty"`
	FirstAssistantData    MatchOfficialData   `json:"firstAssistant, omitempty"`
	SecondAssistantData   MatchOfficialData   `json:"secondAssistant, omitempty"`
	FourthOfficialData    MatchOfficialData   `json:"fourthOfficial, omitempty"`
	VenueData             VenueData           `json:"venue, omitempty"`
	OddsData              AggregatedMatchOdds `json:"odds, omitempty"`
	InPlayOddsData        AggregatedMatchOdds `json:"inplayOdds, omitempty"`
	FlatOddsData          AggregatedMatchOdds `json:"flatOdds, omitempty"`
	LocalCoachData        CoachData           `json:"localCoach, omitempty"`
	VisitorCoachData      CoachData           `json:"visitorCoach, omitempty"`
	TeamStatsData         TeamsStats          `json:"stats, omitempty"`
}

func (f *Fixture) Bench() []PlayerStats {
	return f.BenchData.Data
}

func (f *Fixture) Cards() []CardEvent {
	return f.CardEvents.Data
}

func (f *Fixture) Commentary() []Commentary {
	return f.CommentsData.Data
}

func (f *Fixture) Corners() []CornerEvent {
	return f.CornerEvents.Data
}

func (f *Fixture) Events() []MatchEvent {
	return f.MatchEvents.Data
}

func (f *Fixture) FirstAssistant() *MatchOfficial {
	return f.FirstAssistantData.Data
}

func (f *Fixture) FlatOdds() []MatchOdds {
	return f.FlatOddsData.Data
}

func (f *Fixture) FourthOfficial() *MatchOfficial {
	return f.FourthOfficialData.Data
}

func (f *Fixture) Goals() []GoalEvent {
	return f.GoalEvents.Data
}

func (f *Fixture) Highlights() []VideoHighlights {
	return f.HighlightData.Data
}

func (f *Fixture) InPlayOdds() []MatchOdds {
	return f.InPlayOddsData.Data
}

func (f *Fixture) League() *League {
	return f.LeagueData.Data
}

func (f *Fixture) Lineups() []PlayerStats {
	return f.LineupData.Data
}

func (f *Fixture) LocalCoach() *Coach {
	return f.LocalCoachData.Data
}

func (f *Fixture) LocalTeam() *Team {
	return f.LocalTeamData.Data
}

func (f *Fixture) Odds() []MatchOdds {
	return f.OddsData.Data
}

func (f *Fixture) Referee() *MatchOfficial {
	return f.RefereeData.Data
}

func (f *Fixture) Round() *Round {
	return f.RoundData.Data
}

func (f *Fixture) SecondAssistant() *MatchOfficial {
	return f.SecondAssistantData.Data
}

func (f *Fixture) Sidelined() []Sidelined {
	return f.SidelinedData.Data
}

func (f *Fixture) Stage() *Stage {
	return f.StageData.Data
}

func (f *Fixture) Substitutions() []SubstitutionEvent {
	return f.SubstitutionData.Data
}

func (f *Fixture) TeamStats() []TeamStats {
	return f.TeamStatsData.Data
}

func (f *Fixture) TVStations() []TVStation {
	return f.TVStationData.Data
}

func (f *Fixture) Venue() *Venue {
	return f.VenueData.Data
}

func (f *Fixture) VisitorCoach() *Coach {
	return f.VisitorCoachData.Data
}

func (f *Fixture) VisitorTeam() *Team {
	return f.VisitorTeamData.Data
}

// FixtureById returns a single Fixture struct. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) FixtureById(ctx context.Context, id int, includes []string) (*Fixture, *Meta, error) {
	path := fmt.Sprintf(fixturesUri+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data *Fixture `json:"data"`
		Meta *Meta    `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// FixturesById returns a slice of Fixture struct. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) FixturesById(ctx context.Context, ids []int, includes []string) ([]Fixture, *Meta, error) {
	str := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ","), "[]")

	path := fmt.Sprintf(fixturesMultiUri+"/%s", str)

	return multipleFixtureResponse(c, ctx, path, includes)
}

// FixturesByDate returns a slice of Fixture struct for a given date.

// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) FixturesByDate(ctx context.Context, date time.Time, includes []string) ([]Fixture, *Meta, error) {
	path := fmt.Sprintf(fixturesDateUri + "/" + date.Format("2006-01-02"))

	return multipleFixtureResponse(c, ctx, path, includes)
}

// FixturesByDate returns a slice of Fixture struct for a between two dates. Use the includes slice of string to enrich
// the response data.
func (c *HTTPClient) FixturesBetween(ctx context.Context, from, to time.Time, includes []string) ([]Fixture, *Meta, error) {
	path := fmt.Sprintf(fixturesBetweenUri+"/%s/%s", from.Format(dateFormat), to.Format(dateFormat))

	return multipleFixtureResponse(c, ctx, path, includes)
}

// FixturesByDate returns a slice of Fixture struct for a between two dates for a given team ID. Use the includes slice
// of string to enrich the response data.
func (c *HTTPClient) FixturesBetweenForTeam(ctx context.Context, from, to time.Time, teamId int, includes []string) ([]Fixture, *Meta, error) {
	path := fmt.Sprintf(fixturesBetweenUri+"/%s/%s/%d", from.Format(dateFormat), to.Format(dateFormat), teamId)

	return multipleFixtureResponse(c, ctx, path, includes)

}

// FixturesLastUpdated returns a slice of Fixture struct of fixtures that were most recently updated. Use the includes
// slice of string to enrich the response data.
func (c *HTTPClient) FixturesLastUpdated(ctx context.Context, includes []string) ([]Fixture, *Meta, error) {
	return multipleFixtureResponse(c, ctx, fixturesLastUpdatedUri, includes)
}

// HeadToHead returns a slice of Fixture struct of results between two teams. Use the includes slice of string to
// enrich the response data.
func (c *HTTPClient) HeadToHead(ctx context.Context, idOne, idTwo int, includes []string) ([]Fixture, *Meta, error) {
	path := fmt.Sprintf(headToHeadUri+"/%d/%d", idOne, idTwo)

	return multipleFixtureResponse(c, ctx, path, includes)
}

func multipleFixtureResponse(client *HTTPClient, ctx context.Context, path string, includes []string) ([]Fixture, *Meta, error) {
	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data []Fixture `json:"data"`
		Meta *Meta     `json:"meta"`
	}{}

	err := client.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
