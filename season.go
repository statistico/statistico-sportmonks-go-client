package sportmonks

const seasonUri = "/api/v2.0/seasons"
const IncludeFixtures = "fixtures"
const IncludeResults  = "results"

func (c *Client) Seasons(page int, includes []string) (*SeasonsResponse, error) {
	url := c.BaseURL + seasonUri + "?api_token=" + c.ApiKey

	req, err := buildRequest("GET", url, nil, page, includes)

	if err != nil {
		return nil, err
	}

	s := new(SeasonsResponse)

	if err := c.sendRequest(req, &s); err != nil {
		return nil, err
	}

	return s, err
}