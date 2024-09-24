package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

// TopScorer provides a struct representation of a TopScorer resource.
type TopScorer struct {
	ID            int `json:"id"`
	SeasonID      int `json:"season_id"`
	PlayerID      int `json:"player_id"`
	TypeID        int `json:"type_id"`
	Position      int `json:"position"`
	Total         int `json:"total"`
	ParticipantID int `json:"participant_id"`
	//Season        *Season `json:"season,omitempty"`
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
	PlayerData PlayerData `json:"player"`
	TeamData   TeamData   `json:"team"`
}

// Player returns the player data associated to an assist scorer record.
func (a *AssistScorer) Player() *Player {
	return a.PlayerData.Data
}

// Team returns the team data associated to an assist scorer record.
func (a *AssistScorer) Team() *Team {
	return a.TeamData.Data
}

// CardScorer provides a struct representation of a CardScorer resource.
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

// Player returns the player data associated to an card scorer record.
func (c *CardScorer) Player() *Player {
	return c.PlayerData.Data
}

// Team returns the team data associated to an card scorer record.
func (c *CardScorer) Team() *Team {
	return c.TeamData.Data
}

// GoalScorer provides a struct representation of a GoalScorer resource.
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

// Player returns the player data associated to a goalscorer record.
func (g *GoalScorer) Player() *Player {
	return g.PlayerData.Data
}

// Team returns the team data associated to a goalscorer record.
func (g *GoalScorer) Team() *Team {
	return g.TeamData.Data
}

// TopScorersBySeasonID fetches a TopScorers resource for a season by ID. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) TopScorersBySeasonID(ctx context.Context, seasonID int, includes []string, filters map[string][]int) ([]TopScorer, *Meta, error) {
	path := fmt.Sprintf(topScorersSeasonURI+"/%d", seasonID)

	values := url.Values{
		"include": {strings.Join(includes, ";")},
	}

	formatFilters(&values, filters)

	response := struct {
		Data []TopScorer `json:"data"`
		Meta *Meta       `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
