package sportmonks

import "strconv"

const venueSeasonUri = "/api/v2.0/venues/seasons/"

func (c *Client) BySeasonId(seasonId int) (*VenuesResponse, error) {
	url := c.BaseURL + venueSeasonUri + strconv.Itoa(seasonId) + "?api_token=" + c.ApiKey

	req, err := buildRequest("GET", url, nil, 0, []string{})

	if err != nil {
		return nil, err
	}

	r := new(VenuesResponse)

	err = c.sendRequest(req, &r)

	return r, err
}
