package sportmonks

const countryUri = "/api/v2.0/countries"

// Country struct
type Country struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Extra struct {
		Continent   string      `json:"continent"`
		SubRegion   string      `json:"sub_region"`
		WorldRegion string      `json:"world_region"`
		Fifa        interface{} `json:"fifa,string"`
		ISO         string      `json:"iso"`
		Longitude   string      `json:"longitude"`
		Latitude    string      `json:"latitude"`
	} `json:"extra"`
	Continent struct {
		Data Continent `json:"data"`
	} `json:"continent, omitempty"`
	Leagues struct {
		Data []League `json:"data"`
	} `json:"leagues, omitempty"`
}

// CountriesResponse struct
type CountriesResponse struct {
	Data []Country `json:"data"`
	Meta Meta      `json:"meta"`
}

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
