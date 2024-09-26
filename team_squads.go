package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

// SquadPlayer provides a struct representation of a SquadPlayer resource.
type SquadPlayer struct {
	ID                 int               `json:"id"`
	TransferID         int               `json:"transfer_id"`
	PlayerID           int               `json:"player_id"`
	TeamID             int               `json:"team_id"`
	PositionID         int               `json:"position_id"`
	DetailedPositionID int               `json:"detailed_position_id"`
	Start              string            `json:"start"`
	End                string            `json:"end"`
	Captain            bool              `json:"captain"`
	JerseyNumber       int               `json:"jersey_number"`
	Position           *Position         `json:"position,omitempty"`
	DetailedPosition   *DetailedPosition `json:"detailedposition,omitempty"`
	Player             *Player           `json:"player,omitempty"`
}

// TeamSquad fetches SquadPlayer resources associated to season ID and team ID. Use the includes slice to enrich the response data.
func (c *HTTPClient) TeamSquad(ctx context.Context, seasonID, teamID int, includes []string) ([]SquadPlayer, *Meta, error) {
	path := fmt.Sprintf(teamSeasonSquadURI+"/%d/teams/%d", seasonID, teamID)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data []SquadPlayer `json:"data"`
		Meta *Meta         `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// CurrentSquad fetches SquadPlayer resources associated to team ID for the current season. Use the includes slice to enrich the response data.
func (c *HTTPClient) CurrentSquad(ctx context.Context, teamID int, includes []string) ([]SquadPlayer, *Meta, error) {
	path := fmt.Sprintf(teamSquadURI+"/%d", teamID)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data []SquadPlayer `json:"data"`
		Meta *Meta         `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
