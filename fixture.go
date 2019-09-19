package sportmonks_client

import "strconv"

const fixtureUri = "/api/v2.0/fixtures/"

func (c Client) FixtureById(id int, includes []string, retries int) (*FixtureResponse, error) {
	url := c.BaseURL + fixtureUri + strconv.Itoa(id) + "?api_token=" + c.ApiKey

	req, err := buildRequest("GET", url, nil, 0, includes)

	if err != nil {
		return nil, err
	}

	r := new(FixtureResponse)

	err = c.sendRequest(req, &r, retries)

	return r, err
}
