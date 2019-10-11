package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

//import "fmt"

const leagueStandingsUri = "/standings/season"
const liveLeagueStandingsUri = "/standings/season/live"

type Standings struct {
	Name      string `json:"name"`
	LeagueID  int    `json:"league_id"`
	SeasonID  int    `json:"season_id"`
	RoundID   int    `json:"round_id"`
	RoundName int    `json:"round_name"`
	Type      string `json:"type"`
	StageID   int    `json:"stage_id"`
	StageName string `json:"stage_name"`
	Resource  string `json:"resource"`
	LeagueStandingData LeagueStandingData`json:"standings"`
}

func (s *Standings) LeagueStandings() []LeagueStanding {
	return s.LeagueStandingData.Data
}

type LeagueStanding struct {
	Position  int         `json:"position"`
	TeamID    int         `json:"team_id"`
	TeamName  string      `json:"team_name"`
	RoundID   int         `json:"round_id"`
	RoundName int         `json:"round_name"`
	GroupID   *int 		   `json:"group_id"`
	GroupName *int `json:"group_name"`
	Overall  	TeamLeagueStats `json:"overall"`
	Home 		TeamLeagueStats `json:"home"`
	Away 		TeamLeagueStats `json:"away"`
	Total 		TeamLeagueTotalStats `json:"total"`
	Result     *string `json:"result"`
	Points     int    `json:"points"`
	RecentForm string `json:"recent_form"`
	Status     string `json:"status"`

	LeagueData LeagueData `json:"league"`
	RoundData RoundData `json:"round"`
	SeasonData SeasonData `json:"season"`
	StagesData StagesData `json:"stages"`
	TeamData TeamData `json:"team"`
}

func (s *LeagueStanding) League() *League {
	return s.LeagueData.Data
}

func (s *LeagueStanding) Round() *Round {
	return s.RoundData.Data
}

func (s *LeagueStanding) Season() *Season {
	return s.SeasonData.Data
}

func (s *LeagueStanding) Stages() []Stage {
	return s.StagesData.Data
}

func (s *LeagueStanding) Team() *Team {
	return s.TeamData.Data
}

// StandingsBySeasonId returns a slice of Standings struct for a Season. Use the includes slice to enrich the response data.
func (c *HTTPClient) StandingsBySeasonId(ctx context.Context, seasonId int, includes []string) ([]Standings, *Meta, error) {
	path := fmt.Sprintf(leagueStandingsUri +"/%d", seasonId)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data []Standings `json:"data"`
		Meta *Meta    `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// LiveStandingsBySeasonId returns a slice of Standings struct for a Season with supporting meta data.
func (c *HTTPClient) LiveStandingsBySeasonId(ctx context.Context, seasonId int) ([]LiveStandings, *Meta, error) {
	path := fmt.Sprintf(liveLeagueStandingsUri +"/%d", seasonId)

	response := struct {
		Data []LiveStandings `json:"data"`
		Meta *Meta    `json:"meta"`
	}{}

	err := c.getResource(ctx, path, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
