package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

// Player provides a struct representation of a Player resource
type Player struct {
	ID            int                   `json:"player_id"`
	TeamID        int                   `json:"team_id"`
	CountryID     int                   `json:"country_id"`
	PositionID    int                   `json:"position_id"`
	CommonName    string                `json:"common_name"`
	FullName      string                `json:"fullname"`
	FirstName     string                `json:"firstname"`
	LastName      string                `json:"lastname"`
	Nationality   string                `json:"nationality"`
	BirthDate     string                `json:"birthdate"`
	BirthCountry  string                `json:"birthcountry"`
	BirthPlace    string                `json:"birthplace"`
	Height        string                `json:"height"`
	Weight        string                `json:"weight"`
	ImagePath     string                `json:"image_path"`
	CountryData   countryData           `json:"country"`
	LineupData    playerLineupData      `json:"lineups, omitempty"`
	PositionData  positionData          `json:"position, omitempty"`
	SidelinedData sidelinedData         `json:"sidelined, omitempty"`
	StatsData     playerSeasonStatsData `json:"stats, omitempty"`
	TeamData      TeamData              `json:"team, omitempty"`
	TransfersData transfersData         `json:"transfers, omitempty"`
	TrophyData    trophyData            `json:"trophies, omitempty"`
}

// Country returns Country data.
func (p *Player) Country() *Country {
	return p.CountryData.Data
}

// Lineup return historic player stats information for individual fixtures.
func (p *Player) Lineup() []PlayerStats {
	return p.LineupData.Data
}

// Position returns position data.
func (p *Player) Position() *Position {
	return p.PositionData.Data
}

// Sidelined returns injury information.
func (p *Player) Sidelined() []Sidelined {
	return p.SidelinedData.Data
}

// Stats returns season stats.
func (p *Player) Stats() []PlayerSeasonStats {
	return p.StatsData.Data
}

// Team returns Team data.
func (p *Player) Team() *Team {
	return p.TeamData.Data
}

// Transfers returns transfer information.
func (p *Player) Transfers() []Transfer {
	return p.TransfersData.Data
}

// Trophies returns trophy information.
func (p *Player) Trophies() []Trophy {
	return p.TrophyData.Data
}

// PlayerByID fetches a Player resource by ID.
// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) PlayerByID(ctx context.Context, id int, includes []string) (*Player, *Meta, error) {
	path := fmt.Sprintf(playersURI+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data *Player `json:"data"`
		Meta *Meta   `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
