package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

// TopScorers provides a struct representation of a TopScorers resource
type TopScorers struct {
	ID               int              `json:"id"`
	Name             string           `json:"name"`
	LeagueID         int              `json:"league_id"`
	IsCurrentSeason  bool             `json:"is_current_season"`
	CurrentRoundID   *int             `json:"current_round_id"`
	CurrentStageID   *int             `json:"current_stage_id"`
	AssistScorerData assistScorerData `json:"assistscorers"`
	CardScorerData   cardScorerData   `json:"cardscorers"`
	GoalScorerData   goalScorerData   `json:"goalscorers"`
}

// AssistScorers returns assist scorer data.
func (t *TopScorers) AssistScorers() []AssistScorer {
	return t.AssistScorerData.Data
}

// CardScorers returns card scorer data.
func (t *TopScorers) CardScorers() []CardScorer {
	return t.CardScorerData.Data
}

// GoalScorers returns goal scorer data.
func (t *TopScorers) GoalScorers() []GoalScorer {
	return t.GoalScorerData.Data
}

// AggregatedTopScorers provides a struct representation of a AggregatedTopScorers resource
type AggregatedTopScorers struct {
	ID               int              `json:"id"`
	Name             string           `json:"name"`
	LeagueID         int              `json:"league_id"`
	IsCurrentSeason  bool             `json:"is_current_season"`
	CurrentRoundID   *int             `json:"current_round_id"`
	CurrentStageID   *int             `json:"current_stage_id"`
	AssistScorerData assistScorerData `json:"aggregatedAssistscorers"`
	CardScorerData   cardScorerData   `json:"aggregatedCardscorers"`
	GoalScorerData   goalScorerData   `json:"aggregatedGoalscorers"`
}

// AssistScorers returns aggregated assist scorer data.
func (a *AggregatedTopScorers) AssistScorers() []AssistScorer {
	return a.AssistScorerData.Data
}

// CardScorers returns aggregated card scorer data.
func (a *AggregatedTopScorers) CardScorers() []CardScorer {
	return a.CardScorerData.Data
}

// GoalScorers returns aggregated goal scorer data.
func (a *AggregatedTopScorers) GoalScorers() []GoalScorer {
	return a.GoalScorerData.Data
}

// AssistScorer provides a struct representation of a AssistScorer resource
type AssistScorer struct {
	Position   int        `json:"position"`
	SeasonID   int        `json:"season_id"`
	PlayerID   int        `json:"player_id"`
	TeamID     int        `json:"team_id"`
	StageID    int        `json:"stage_id"`
	Assists    int        `json:"assists"`
	Type       string     `json:"type"`
	PlayerData playerData `json:"player"`
	TeamData   teamData   `json:"team"`
}

// Player returns the player data associated to an assist scorer record.
func (a *AssistScorer) Player() *Player {
	return a.PlayerData.Data
}

// Team returns the team data associated to an assist scorer record.
func (a *AssistScorer) Team() *Team {
	return a.TeamData.Data
}

// CardScorer provides a struct representation of a CardScorer resource
type CardScorer struct {
	Position    int        `json:"position"`
	SeasonID    int        `json:"season_id"`
	PlayerID    int        `json:"player_id"`
	TeamID      int        `json:"team_id"`
	StageID     int        `json:"stage_id"`
	YellowCards int        `json:"yellowcards"`
	RedCards    int        `json:"redcards"`
	Type        string     `json:"type"`
	PlayerData  playerData `json:"player"`
	TeamData    teamData   `json:"team"`
}

// Player returns the player data associated to an card scorer record.
func (c *CardScorer) Player() *Player {
	return c.PlayerData.Data
}

// Team returns the team data associated to an card scorer record.
func (c *CardScorer) Team() *Team {
	return c.TeamData.Data
}

// GoalScorer provides a struct representation of a GoalScorer resource
type GoalScorer struct {
	Position     int        `json:"position"`
	SeasonID     int        `json:"season_id"`
	PlayerID     int        `json:"player_id"`
	TeamID       int        `json:"team_id"`
	StageID      int        `json:"stage_id"`
	Goals        int        `json:"goals"`
	PenaltyGoals int        `json:"penalty_goals"`
	Type         string     `json:"type"`
	PlayerData   playerData `json:"player"`
	TeamData     teamData   `json:"team"`
}

// Player returns the player data associated to an goal scorer record.
func (g *GoalScorer) Player() *Player {
	return g.PlayerData.Data
}

// Team returns the team data associated to an goal scorer record.
func (g *GoalScorer) Team() *Team {
	return g.TeamData.Data
}

// TopScorersBySeasonID fetches a TopScorers resource for a season by ID.
// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) TopScorersBySeasonID(ctx context.Context, seasonID int, includes []string, filters map[string][]int) (*TopScorers, *Meta, error) {
	path := fmt.Sprintf(topScorersSeasonURI+"/%d", seasonID)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	formatFilters(&values, filters)

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

// AggregatedTopScorersBySeasonID fetches an AggregatedTopScorers resource for a season by ID.
// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) AggregatedTopScorersBySeasonID(ctx context.Context, seasonID int, includes []string) (*AggregatedTopScorers, *Meta, error) {
	path := fmt.Sprintf(topScorersSeasonURI+"/%d/aggregated", seasonID)

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
