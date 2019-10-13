package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

// SquadPlayer struct
type SquadPlayer struct {
	PlayerID           int       `json:"player_id"`
	PositionID         int       `json:"position_id"`
	Number             int       `json:"number"`
	Captain            int       `json:"captain"`
	Injured            bool      `json:"injured"`
	Minutes            int       `json:"minutes"`
	Appearances        int       `json:"appearences"`
	Lineups            int       `json:"lineups"`
	SubstituteIn       int       `json:"substitute_in"`
	SubstituteOut      int       `json:"substitute_out"`
	SubstitutesOnBench int       `json:"substitutes_on_bench"`
	Goals              int       `json:"goals"`
	Assists            int       `json:"assists"`
	Saves              *int      `json:"saves"`
	InsideBoxSaves     *int      `json:"inside_box_saves"`
	Dispossessed       *int      `json:"dispossesed"`
	Interceptions      *int      `json:"interceptions"`
	YellowCards        int       `json:"yellowcards"`
	YellowRed          int       `json:"yellowred"`
	RedCards           int       `json:"redcards"`
	Tackles            *int      `json:"tackles"`
	Blocks             *int      `json:"blocks"`
	HitPost            *int      `json:"hit_post"`
	Fouls              Fouls     `json:"fouls"`
	Crosses            Crosses   `json:"crosses"`
	Dribbles           Dribbles  `json:"dribbles"`
	Duels              Duels     `json:"duels"`
	Passes             Passes    `json:"passes"`
	Penalties          Penalties `json:"penalties"`

	PlayerData PlayerData `json:"player"`
}

func (s *SquadPlayer) Player() *Player {
	return s.PlayerData.Data
}

// TeamSquad returns a slice of SquadPlayer struct associated to season ID and team ID. Use the includes slice to
// enrich the response data.
func (c *HTTPClient) TeamSquad(ctx context.Context, seasonId, teamId int, includes []string) ([]SquadPlayer, *Meta, error) {
	path := fmt.Sprintf(teamSquadUri+"/%d/team/%d", seasonId, teamId)

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
