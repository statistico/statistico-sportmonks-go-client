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
	Details                 *string             `json:"details"`
	NeutralVenue          *bool               `json:"neutral_venue"`
	WinningOddsCalculated bool                `json:"winning_odds_calculated"`
	Formations            Formations          `json:"formations"`
	Scores                Scores              `json:"scores"`
	Time                  FixtureTime         `json:"time"`
	Coaches               Coaches             `json:"coaches"`
	Standings             TeamStandings           `json:"standings"`
	Leg                   *string             `json:"leg"`
	Colors                TeamColors        	`json:"colors"`
	Deleted               bool                `json:"deleted"`
	Cards                 CardEvents          `json:"cards"`
	Corners               CornerEvents        `json:"corners"`
	Goals                 GoalEvents          `json:"goals"`
	Events                MatchEvents         `json:"events"`
	Lineup                Lineup              `json:"lineup"`
	Bench                 Lineup              `json:"bench"`
	LocalTeam             TeamData            `json:"localTeam, omitempty"`
	VisitorTeam           TeamData            `json:"visitorTeam, omitempty"`
	Substitutions         SubstitutionEvents  `json:"substitutions"`
	Sidelined             SidelinedData       `json:"sidelined"`
	Comments              MatchCommentary     `json:"comments, omitempty"`
	TVStations            MatchTVStations     `json:"tvstations"`
	Highlights            MatchHighlights     `json:"highlights, omitempty"`
	Round                 *Round              `json:"round, omitempty"`
	Stage                 *Stage              `json:"stage, omitempty"`
	Referee               *Referee            `json:"referee, omitempty"`
	Assistants            Assistants          `json:"assistants"`
	Venue                 *Venue              `json:"venue, omitempty"`
	Odds                  AggregatedMatchOdds `json:"odds, omitempty"`
	InPlayOdds            AggregatedMatchOdds `json:"inplayOdds, omitempty"`
	FlatOdds              AggregatedMatchOdds `json:"flatOdds, omitempty"`
	LocalCoach            *Coach              `json:"localCoach, omitempty"`
	VisitorCoach          *Coach              `json:"visitorCoach, omitempty"`
	TeamStats             TeamsStats          `json:"stats, omitempty"`
}

// FixtureById returns a single Fixture struct.

// Use the includes slice of string to enrich the response data.
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

// FixturesById returns a slice of Fixture struct.

// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) FixturesById(ctx context.Context, ids []int, includes []string) ([]Fixture, *Meta, error) {
	str := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ","), "[]")

	path := fmt.Sprintf(fixturesMultiUri +"/%s", str)

	return multipleFixtureResponse(c, ctx, path, includes)
}

// FixturesByDate returns a slice of Fixture struct for a given date.

// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) FixturesByDate(ctx context.Context, date time.Time, includes []string) ([]Fixture, *Meta, error) {
	path := fmt.Sprintf(fixturesDateUri + "/" + date.Format("2006-01-02"))

	return multipleFixtureResponse(c, ctx, path, includes)
}

// FixturesByDate returns a slice of Fixture struct for a between two dates.

// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) FixturesBetween(ctx context.Context, from, to time.Time, includes []string) ([]Fixture, *Meta, error) {
	path := fmt.Sprintf(fixturesBetweenUri + "/%s/%s", from.Format(dateFormat), to.Format(dateFormat))

	return multipleFixtureResponse(c, ctx, path, includes)
}

// FixturesByDate returns a slice of Fixture struct for a between two dates for a given team ID.

// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) FixturesBetweenForTeam(ctx context.Context, from, to time.Time, teamId int, includes []string) ([]Fixture, *Meta, error) {
	path := fmt.Sprintf(fixturesBetweenUri + "/%s/%s/%d", from.Format(dateFormat), to.Format(dateFormat), teamId)

	return multipleFixtureResponse(c, ctx, path, includes)

}

// FixturesLastUpdated returns a slice of Fixture struct of fixtures that were most recently updated.

// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) FixturesLastUpdated(ctx context.Context, includes []string) ([]Fixture, *Meta, error) {
	return multipleFixtureResponse(c, ctx, fixturesLastUpdatedUri, includes)
}

// HeadToHead returns a slice of Fixture struct of results between two teams.

// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) HeadToHead(ctx context.Context, idOne, idTwo int, includes []string) ([]Fixture, *Meta, error) {
	path := fmt.Sprintf(headToHeadUri + "/%d/%d", idOne, idTwo)

	return multipleFixtureResponse(c, ctx, path, includes)
}

func multipleFixtureResponse(client *HTTPClient, ctx context.Context, path string, includes []string) ([]Fixture, *Meta, error) {
	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data []Fixture `json:"data"`
		Meta *Meta    `json:"meta"`
	}{}

	err := client.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}