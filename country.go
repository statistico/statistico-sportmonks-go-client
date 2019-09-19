package sportmonks_client

const countryUri = "/api/v2.0/countries"

func (c *Client) Countries(page int, includes []string, retries int) (*CountriesResponse, error) {
	url := c.BaseURL + countryUri + "?api_token=" + c.ApiKey

	req, err := buildRequest("GET", url, nil, page, includes)

	if err != nil {
		return nil, err
	}

	countries := new(CountriesResponse)

	if err := c.sendRequest(req, &countries, retries); err != nil {
		return nil, err
	}

	return countries, err
}
