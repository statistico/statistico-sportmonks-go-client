package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

// Standings provides a struct representation of a Standings resource.
type Standings struct {
	Name               string             `json:"name"`
	LeagueID           int                `json:"league_id"`
	SeasonID           int                `json:"season_id"`
	RoundID            int                `json:"round_id"`
	RoundName          int                `json:"round_name"`
	Type               string             `json:"type"`
	StageID            int                `json:"stage_id"`
	StageName          string             `json:"stage_name"`
	Resource           string             `json:"resource"`
	LeagueStandingData LeagueStandingData `json:"standings"`
}

// LeagueStandings returns league standing resources describing team specific league information.
func (s *Standings) LeagueStandings() []LeagueStanding {
	return s.LeagueStandingData.Data
}

// LeagueStanding provides a struct representation of a LeagueStanding resource.
type LeagueStanding struct {
	Position   int                  `json:"position"`
	TeamID     int                  `json:"team_id"`
	TeamName   string               `json:"team_name"`
	RoundID    int                  `json:"round_id"`
	RoundName  int                  `json:"round_name"`
	GroupID    *int                 `json:"group_id"`
	GroupName  *int                 `json:"group_name"`
	Overall    TeamLeagueStats      `json:"overall"`
	Home       TeamLeagueStats      `json:"home"`
	Away       TeamLeagueStats      `json:"away"`
	Total      TeamLeagueTotalStats `json:"total"`
	Result     *string              `json:"result"`
	Points     int                  `json:"points"`
	RecentForm string               `json:"recent_form"`
	Status     string               `json:"status"`

	LeagueData LeagueData `json:"league"`
	RoundData  RoundData  `json:"round"`
	SeasonData SeasonData `json:"season"`
	StagesData StagesData `json:"stages"`
	TeamData   TeamData   `json:"team"`
}

// League returns league data.
func (s *LeagueStanding) League() *League {
	return s.LeagueData.Data
}

// Round returns round data.
func (s *LeagueStanding) Round() *Round {
	return s.RoundData.Data
}

// Season returns season data.
func (s *LeagueStanding) Season() *Season {
	return s.SeasonData.Data
}

// Stages returns stages data.
func (s *LeagueStanding) Stages() []Stage {
	return s.StagesData.Data
}

// Team return team data.
func (s *LeagueStanding) Team() *Team {
	return s.TeamData.Data
}

// StandingsBySeasonID fetches a Standings resource for a Season ID. Use the includes slice to enrich the response data.
func (c *HTTPClient) StandingsBySeasonID(ctx context.Context, seasonID int, includes []string, filters map[string][]int) ([]Standings, *Meta, error) {
	path := fmt.Sprintf(leagueStandingsURI+"/%d", seasonID)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	formatFilters(&values, filters)

	response := struct {
		Data []Standings `json:"data"`
		Meta *Meta       `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// LiveStandingsBySeasonID fetches a Standings resource for a Season ID for present moment standings. Use the includes slice to enrich
// the response data.
func (c *HTTPClient) LiveStandingsBySeasonID(ctx context.Context, seasonID int, filters map[string][]int) ([]LiveStandings, *Meta, error) {
	path := fmt.Sprintf(liveLeagueStandingsURI+"/%d", seasonID)

	values := url.Values{}

	formatFilters(&values, filters)

	response := struct {
		Data []LiveStandings `json:"data"`
		Meta *Meta           `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
