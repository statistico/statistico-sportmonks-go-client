package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

const playersUri = "/players"

type Player struct {
	ID     int    `json:"player_id"`
	TeamID       int    `json:"team_id"`
	CountryID    int    `json:"country_id"`
	PositionID   int    `json:"position_id"`
	CommonName   string `json:"common_name"`
	FullName     string `json:"fullname"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Nationality  string `json:"nationality"`
	BirthDate    string `json:"birthdate"`
	BirthCountry string `json:"birthcountry"`
	BirthPlace   string `json:"birthplace"`
	Height       string `json:"height"`
	Weight       string `json:"weight"`
	ImagePath    string `json:"image_path"`

	CountryData      CountryData `json:"country"`
	LineupData 	PlayerLineupData `json:"lineups, omitempty"`
	PositionData PositionData `json:"position, omitempty"`
	SidelinedData SidelinedData `json:"sidelined, omitempty"`
	StatsData PlayerSeasonStatsData `json:"stats, omitempty"`
	TeamData TeamData `json:"team, omitempty"`
	TransfersData TransfersData `json:"transfers, omitempty"`
	TrophyData TrophyData`json:"trophies, omitempty"`
}

func (p *Player) Country() *Country {
	return p.CountryData.Data
}

func (p *Player) Lineup() []PlayerMatchStats {
	return p.LineupData.Data
}

func (p *Player) Position() *Position {
	return p.PositionData.Data
}

func (p *Player) Sidelined() []Sidelined {
	return p.SidelinedData.Data
}

func (p *Player) Stats() []PlayerSeasonStats {
	return p.StatsData.Data
}

func (p *Player) Team() *Team {
	return p.TeamData.Data
}

func (p *Player) Transfers() []Transfer {
	return p.TransfersData.Data
}

func (p *Player) Trophies() []Trophy {
	return p.TrophyData.Data
}

// MarketById returns a single Market struct and supporting meta data. Use the includes slice of string to enrich
// the response data.
func (c *HTTPClient) PlayerById(ctx context.Context, id int, includes []string) (*Player, *Meta, error) {
	path := fmt.Sprintf(playersUri + "/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data *Player `json:"data"`
		Meta *Meta    `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
