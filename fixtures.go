package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

const fixturesUri = "/fixtures"
const fixtureDateUri = "/api/v2.0/fixtures/date"
const fixtureBetweenUri = "/api/v2.0/fixtures/between"
const fixturesMultiUri = "/fixtures/multi"
const fixtureUpdatesUri = "/api/v2.0/fixtures/between"

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
	Standings             Standings           `json:"standings"`
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

// FixtureById sends a request and returns a single Fixture struct.

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

// FixturesById returns a slice of Fixture resource struct and supporting meta data.

// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) FixturesById(ctx context.Context, ids []int, includes []string) ([]Fixture, *Meta, error) {
	str := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ","), "[]")

	path := fmt.Sprintf(fixturesMultiUri +"/%s", str)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data []Fixture `json:"data"`
		Meta *Meta    `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

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
