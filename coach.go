package sportmonks

const coachUri = "/api/v2.0/coaches/"

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
	Birthplace   *string `json:"birthplace"`
	ImagePath    string `json:"image_path"`
}

//// Retrieve a single coach resource by ID.
//func (c *HTTPClient) CoachById(id int) (*Coach, *Meta, error) {
//	url := venueUri + "/" + strconv.Itoa(id)
//
//	response := new(CoachResponse)
//
//	err := c.handleRequest(url, []string{}, response)
//
//	if err != nil {
//		return nil, nil, err
//	}
//
//	return &response.Data, &response.Meta, err
//}
