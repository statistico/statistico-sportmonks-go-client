package sportmonks

// Stage struct
type Stage struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	LeagueID     int         `json:"league_id"`
	SeasonID     int         `json:"season_id"`
	SortOrder    *string     `json:"sort_order"`
	HasStandings bool        `json:"has_standings"`
}
