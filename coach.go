package sportmonks

type Coach struct {
	CoachID      int    `json:"coach_id"`
	TeamID       int    `json:"team_id"`
	CountryID    int    `json:"country_id"`
	CommonName   string `json:"common_name"`
	FullName     string `json:"fullname"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Nationality  string `json:"nationality"`
	BirthDate    string `json:"birthdate"`
	BirthCountry string `json:"birthcountry"`
	Birthplace   string `json:"birthplace"`
	ImagePath    string `json:"image_path"`
}