package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

type TopScorers struct {
	ID               int              `json:"id"`
	Name             string           `json:"name"`
	LeagueID         int              `json:"league_id"`
	IsCurrentSeason  bool             `json:"is_current_season"`
	CurrentRoundID   *int             `json:"current_round_id"`
	CurrentStageID   *int             `json:"current_stage_id"`
	AssistScorerData AssistScorerData `json:"assistscorers"`
	CardScorerData   CardScorerData   `json:"cardscorers"`
	GoalScorerData   GoalScorerData   `json:"goalscorers"`
}

func (t *TopScorers) AssistScorers() []AssistScorer {
	return t.AssistScorerData.Data
}

func (t *TopScorers) CardScorers() []CardScorer {
	return t.CardScorerData.Data
}

func (t *TopScorers) GoalScorers() []GoalScorer {
	return t.GoalScorerData.Data
}

type AggregatedTopScorers struct {
	ID               int              `json:"id"`
	Name             string           `json:"name"`
	LeagueID         int              `json:"league_id"`
	IsCurrentSeason  bool             `json:"is_current_season"`
	CurrentRoundID   *int             `json:"current_round_id"`
	CurrentStageID   *int             `json:"current_stage_id"`
	AssistScorerData AssistScorerData `json:"aggregatedAssistscorers"`
	CardScorerData   CardScorerData   `json:"aggregatedCardscorers"`
	GoalScorerData   GoalScorerData   `json:"aggregatedGoalscorers"`
}

func (a *AggregatedTopScorers) AssistScorers() []AssistScorer {
	return a.AssistScorerData.Data
}

func (a *AggregatedTopScorers) CardScorers() []CardScorer {
	return a.CardScorerData.Data
}

func (a *AggregatedTopScorers) GoalScorers() []GoalScorer {
	return a.GoalScorerData.Data
}

type AssistScorer struct {
	Position   int        `json:"position"`
	SeasonID   int        `json:"season_id"`
	PlayerID   int        `json:"player_id"`
	TeamID     int        `json:"team_id"`
	StageID    int        `json:"stage_id"`
	Assists    int        `json:"assists"`
	Type       string     `json:"type"`
	PlayerData PlayerData `json:"player"`
	TeamData   TeamData   `json:"team"`
}

func (a *AssistScorer) Player() *Player {
	return a.PlayerData.Data
}

func (a *AssistScorer) Team() *Team {
	return a.TeamData.Data
}

type CardScorer struct {
	Position    int        `json:"position"`
	SeasonID    int        `json:"season_id"`
	PlayerID    int        `json:"player_id"`
	TeamID      int        `json:"team_id"`
	StageID     int        `json:"stage_id"`
	YellowCards int        `json:"yellowcards"`
	RedCards    int        `json:"redcards"`
	Type        string     `json:"type"`
	PlayerData  PlayerData `json:"player"`
	TeamData    TeamData   `json:"team"`
}

func (c *CardScorer) Player() *Player {
	return c.PlayerData.Data
}

func (c *CardScorer) Team() *Team {
	return c.TeamData.Data
}

type GoalScorer struct {
	Position     int        `json:"position"`
	SeasonID     int        `json:"season_id"`
	PlayerID     int        `json:"player_id"`
	TeamID       int        `json:"team_id"`
	StageID      int        `json:"stage_id"`
	Goals        int        `json:"goals"`
	PenaltyGoals int        `json:"penalty_goals"`
	Type         string     `json:"type"`
	PlayerData   PlayerData `json:"player"`
	TeamData     TeamData   `json:"team"`
}

func (g *GoalScorer) Player() *Player {
	return g.PlayerData.Data
}

func (g *GoalScorer) Team() *Team {
	return g.TeamData.Data
}

// TopScorersBySeasonId returns a slice of Fixture struct. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) TopScorersBySeasonId(ctx context.Context, seasonId int, includes []string) (*TopScorers, *Meta, error) {
	path := fmt.Sprintf(topScorersSeasonUri+"/%d", seasonId)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data *TopScorers `json:"data"`
		Meta *Meta       `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// AggregatedTopScorersBySeasonId returns a slice of Fixture struct. Use the includes slice of string to enrich
// the response data.
func (c *HTTPClient) AggregatedTopScorersBySeasonId(ctx context.Context, seasonId int, includes []string) (*AggregatedTopScorers, *Meta, error) {
	path := fmt.Sprintf(topScorersSeasonUri+"/%d/aggregated", seasonId)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data *AggregatedTopScorers `json:"data"`
		Meta *Meta                 `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
