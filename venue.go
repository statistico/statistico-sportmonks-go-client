package sportmonks

const venueUri = "/api/v2.0/venues/"
const venueSeasonUri = "/api/v2.0/venues/season/"

type Venue struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Surface  string  `json:"surface"`
	Address  *string `json:"address"`
	City     string  `json:"city"`
	Capacity int     `json:"capacity"`
	ImagePath    string  `json:"image_path"`
	Coordinates *string `json:"coordinates"`
}

//// Retrieve a single venue resource by ID.
//func (c *HTTPClient) VenueById(id int) (*Venue, *Meta, error) {
//	url := venueUri + "/" + strconv.Itoa(id)
//
//	response := new(VenueResponse)
//
//	err := c.handleRequest(url, []string{}, response)
//
//	if err != nil {
//		return nil, nil, err
//	}
//
//	return &response.Data, &response.Meta, err
//}
//
//// Make a request to retrieve multiple venue resources for a given season.
//func (c *HTTPClient) VenuesBySeasonId(id int) ([]Venue, *Meta, error) {
//	url := venueSeasonUri + "/" + strconv.Itoa(id)
//
//	response := new(VenuesResponse)
//
//	err := c.handleRequest(url, []string{}, response)
//
//	if err != nil {
//		return nil, nil, err
//	}
//
//	return response.Data, &response.Meta, err
//}
