package sportmonks

import "strconv"

const teamSeasonUri = "/api/v2.0/teams/season/"

func (c Client) TeamsBySeasonId(seasonId int, retries int) (*TeamsResponse, error) {
	url := c.BaseURL + teamSeasonUri + strconv.Itoa(seasonId) + "?api_token=" + c.ApiKey

	req, err := buildRequest("GET", url, nil, 0, []string{})

	if err != nil {
		return nil, err
	}

	r := new(TeamsResponse)

	err = c.sendRequest(req, &r,  retries)

	return r, err
}
