package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"
)

const dateFormat = "2006-01-02"

// Fixture provides a struct representation of a Fixture resource
type Fixture struct {
	ID                    int                     `json:"id"`
	LeagueID              int                     `json:"league_id"`
	SeasonID              int                     `json:"season_id"`
	StageID               *int                    `json:"stage_id"`
	RoundID               *int                    `json:"round_id"`
	GroupID               *int                    `json:"group_id"`
	AggregateID           *int                    `json:"aggregate_id"`
	VenueID               *int                    `json:"venue_id"`
	RefereeID             *int                    `json:"referee_id"`
	LocalTeamID           int                     `json:"localteam_id"`
	VisitorTeamID         int                     `json:"visitorteam_id"`
	WinnerTeamID          *int                    `json:"winner_team_id"`
	WeatherReport         WeatherReport           `json:"weather_report"`
	Commentaries          *bool                   `json:"commentaries"`
	Attendance            *int                    `json:"attendance"`
	Pitch                 *string                 `json:"pitch"`
	Details               *string                 `json:"details"`
	NeutralVenue          *bool                   `json:"neutral_venue"`
	WinningOddsCalculated bool                    `json:"winning_odds_calculated"`
	Formations            Formations              `json:"formations"`
	Scores                Scores                  `json:"scores"`
	Time                  FixtureTime             `json:"time"`
	Coaches               Coaches                 `json:"coaches"`
	Standings             TeamStandings           `json:"standings"`
	Assistants            Assistants              `json:"assistants"`
	Leg                   *string                 `json:"leg"`
	Colors                TeamColors              `json:"colors"`
	Deleted               bool                    `json:"deleted"`
	CardEvents            cardEventsData          `json:"cards"`
	CornerEvents          cornerEventsData        `json:"corners"`
	GoalEvents            GoalEvents              `json:"goals"`
	MatchEvents           MatchEvents             `json:"events"`
	LineupData            playerLineupData        `json:"lineup"`
	BenchData             playerLineupData        `json:"bench"`
	LocalTeamData         TeamData                `json:"localTeam, omitempty"`
	VisitorTeamData       TeamData                `json:"visitorTeam, omitempty"`
	SubstitutionData      SubstitutionEvents      `json:"substitutions"`
	SidelinedData         sidelinedData           `json:"sidelined"`
	CommentsData          MatchCommentary         `json:"comments, omitempty"`
	TVStationData         matchTVStations         `json:"tvstations"`
	HighlightData         matchHighlightsData     `json:"highlights, omitempty"`
	LeagueData            leagueData              `json:"league, omitempty"`
	RoundData             roundData               `json:"round, omitempty"`
	StageData             stageData               `json:"stage, omitempty"`
	RefereeData           matchOfficialData       `json:"referee, omitempty"`
	FirstAssistantData    matchOfficialData       `json:"firstAssistant, omitempty"`
	SecondAssistantData   matchOfficialData       `json:"secondAssistant, omitempty"`
	FourthOfficialData    matchOfficialData       `json:"fourthOfficial, omitempty"`
	VenueData             venueData               `json:"venue, omitempty"`
	OddsData              aggregatedMatchOddsData `json:"odds, omitempty"`
	InPlayOddsData        aggregatedMatchOddsData `json:"inplayOdds, omitempty"`
	FlatOddsData          aggregatedMatchOddsData `json:"flatOdds, omitempty"`
	LocalCoachData        coachData               `json:"localCoach, omitempty"`
	VisitorCoachData      coachData               `json:"visitorCoach, omitempty"`
	TeamStatsData         TeamsStats              `json:"stats, omitempty"`
}

// Bench returns PlayerStats data for players on the bench for a fixture
func (f *Fixture) Bench() []PlayerStats {
	return f.BenchData.Data
}

// Cards returns CardEvents for a fixture
func (f *Fixture) Cards() []CardEvent {
	return f.CardEvents.Data
}

// Commentary returns Commentary data for a fixture
func (f *Fixture) Commentary() []Commentary {
	return f.CommentsData.Data
}

// Corners returns CornerEvents for a fixture
func (f *Fixture) Corners() []CornerEvent {
	return f.CornerEvents.Data
}

// Events returns all events for a fixture
func (f *Fixture) Events() []MatchEvent {
	return f.MatchEvents.Data
}

// FirstAssistant return first assistant data for a fixture
func (f *Fixture) FirstAssistant() *MatchOfficial {
	return f.FirstAssistantData.Data
}

// FlatOdds returns flat odds data for a fixture
func (f *Fixture) FlatOdds() []MatchOdds {
	return f.FlatOddsData.Data
}

// FourthOfficial return fourth official data for a fixture
func (f *Fixture) FourthOfficial() *MatchOfficial {
	return f.FourthOfficialData.Data
}

// Goals returns GoalEvents for a fixture
func (f *Fixture) Goals() []GoalEvent {
	return f.GoalEvents.Data
}

// Highlights returns video highlights data for a fixture
func (f *Fixture) Highlights() []VideoHighlights {
	return f.HighlightData.Data
}

// InPlayOdds returns in play odds data for a fixture
func (f *Fixture) InPlayOdds() []MatchOdds {
	return f.InPlayOddsData.Data
}

// League returns league data for a fixture
func (f *Fixture) League() *League {
	return f.LeagueData.Data
}

// Lineups returns PlayerStats data for starting players for a fixture
func (f *Fixture) Lineups() []PlayerStats {
	return f.LineupData.Data
}

// LocalCoach returns coach data for the home team
func (f *Fixture) LocalCoach() *Coach {
	return f.LocalCoachData.Data
}

// LocalTeam return team data for the home team
func (f *Fixture) LocalTeam() *Team {
	return f.LocalTeamData.Data
}

// Odds returns bookmaker odds data for a fixture
func (f *Fixture) Odds() []MatchOdds {
	return f.OddsData.Data
}

// Referee returns referee data for a fixture
func (f *Fixture) Referee() *MatchOfficial {
	return f.RefereeData.Data
}

// Round returns round data for a fixture
func (f *Fixture) Round() *Round {
	return f.RoundData.Data
}

// SecondAssistant return second assistant data for a fixture
func (f *Fixture) SecondAssistant() *MatchOfficial {
	return f.SecondAssistantData.Data
}

// Sidelined returns player data for injured players
func (f *Fixture) Sidelined() []Sidelined {
	return f.SidelinedData.Data
}

// Stage returns round stage for a fixture
func (f *Fixture) Stage() *Stage {
	return f.StageData.Data
}

// Substitutions returns SubstitutionEvents for a fixture
func (f *Fixture) Substitutions() []SubstitutionEvent {
	return f.SubstitutionData.Data
}

// TeamStats returns stats data for home and away teams
func (f *Fixture) TeamStats() []TeamStats {
	return f.TeamStatsData.Data
}

// TVStations return tv station data for a fixture
func (f *Fixture) TVStations() []TVStation {
	return f.TVStationData.Data
}

// Venue returns round venue for a fixture
func (f *Fixture) Venue() *Venue {
	return f.VenueData.Data
}

// VisitorCoach returns coach data for the away team
func (f *Fixture) VisitorCoach() *Coach {
	return f.VisitorCoachData.Data
}

// VisitorTeam return teams data for the away team
func (f *Fixture) VisitorTeam() *Team {
	return f.VisitorTeamData.Data
}

// FixtureByID fetches a Fixture resource by ID.
// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) FixtureByID(ctx context.Context, id int, includes []string) (*Fixture, *Meta, error) {
	path := fmt.Sprintf(fixturesURI+"/%d", id)

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

// FixturesByID fetches multiple Fixture resources by their IDS.
// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) FixturesByID(ctx context.Context, ids []int, includes []string) ([]Fixture, *Meta, error) {
	str := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ","), "[]")

	path := fmt.Sprintf(fixturesMultiURI+"/%s", str)

	return multipleFixtureResponse(ctx, c, path, includes)
}

// FixturesByDate fetches multiple Fixture resources for a given date.
// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) FixturesByDate(ctx context.Context, date time.Time, includes []string) ([]Fixture, *Meta, error) {
	path := fmt.Sprintf(fixturesDateURI + "/" + date.Format("2006-01-02"))

	return multipleFixtureResponse(ctx, c, path, includes)
}

// FixturesBetween fetches multiple Fixture resources for between two dates.
// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) FixturesBetween(ctx context.Context, from, to time.Time, includes []string) ([]Fixture, *Meta, error) {
	path := fmt.Sprintf(fixturesBetweenURI+"/%s/%s", from.Format(dateFormat), to.Format(dateFormat))

	return multipleFixtureResponse(ctx, c, path, includes)
}

// FixturesBetweenForTeam fetches multiple Fixture resources for between two dates for a given team ID.
// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) FixturesBetweenForTeam(ctx context.Context, from, to time.Time, teamID int, includes []string) ([]Fixture, *Meta, error) {
	path := fmt.Sprintf(fixturesBetweenURI+"/%s/%s/%d", from.Format(dateFormat), to.Format(dateFormat), teamID)

	return multipleFixtureResponse(ctx, c, path, includes)

}

// FixturesLastUpdated fetches multiple Fixture resources of fixtures that were most recently updated.
// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) FixturesLastUpdated(ctx context.Context, includes []string) ([]Fixture, *Meta, error) {
	return multipleFixtureResponse(ctx, c, fixturesLastUpdatedURI, includes)
}

// HeadToHead fetches multiple Fixture resources of results between two teams.
// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) HeadToHead(ctx context.Context, idOne, idTwo int, includes []string) ([]Fixture, *Meta, error) {
	path := fmt.Sprintf(headToHeadURI+"/%d/%d", idOne, idTwo)

	return multipleFixtureResponse(ctx, c, path, includes)
}

func multipleFixtureResponse(ctx context.Context, client *HTTPClient, path string, includes []string) ([]Fixture, *Meta, error) {
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
