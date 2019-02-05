package sportmonks

import "strconv"

const playerUri = "/api/v2.0/players/"

func (c Client) PlayerById(id int, includes []string, retries int) (*PlayerResponse, error) {
	url := c.BaseURL + playerUri + strconv.Itoa(id) + "?api_token=" + c.ApiKey

	req, err := buildRequest("GET", url, nil, 0, includes)

	if err != nil {
		return nil, err
	}

	r := new(PlayerResponse)

	err = c.sendRequest(req, &r, retries)

	return r, err
}
