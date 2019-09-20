package sportmonks

const playerUri = "/api/v2.0/players/"

type Player struct {
	PlayerID     int    `json:"player_id"`
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
	Birthplace   string `json:"birthplace"`
	Height       string `json:"height"`
	Weight       string `json:"weight"`
	ImagePath    string `json:"image_path"`
	Position struct {
		Data *Position `json:"data"`
	} `json:"position, omitempty"`
	Sidelined struct {
		Data []Sidelined `json:"data"`
	} `json:"sidelined, omitempty"`
	Stats struct {
		Data []PlayerStats `json:"data"`
	} `json:"stats, omitempty"`
	Team struct {
		Data *Team `json:"data"`
	} `json:"team, omitempty"`
	Transfers struct {
		Data []Transfer `json:"data"`
	} `json:"transfers, omitempty"`
	Trophies struct {
		Data []Trophy `json:"data"`
	} `json:"trophies, omitempty"`
}
