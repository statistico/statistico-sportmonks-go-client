package sportmonks

const leagueUri = "/api/v2.0/leagues"

func (c *Client) Leagues(page int, includes []string, retries int) (*LeaguesResponse, error) {
	url := c.BaseURL + leagueUri + "?api_token=" + c.ApiKey

	req, err := buildRequest("GET", url, nil, page, includes)

	if err != nil {
		return nil, err
	}

	leagues := new(LeaguesResponse)

	if err := c.sendRequest(req, &leagues, retries); err != nil {
		return nil, err
	}

	return leagues, err
}
