package sportmonks

import "strconv"

const countryUri = "/api/v2.0/countries"

// Country struct
type Country struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Extra struct {
		Continent   string      `json:"continent"`
		SubRegion   string      `json:"sub_region"`
		WorldRegion string      `json:"world_region"`
		FIFA        *string 	`json:"fifa,string"`
		ISO         string      `json:"iso"`
		Longitude   string      `json:"longitude"`
		Latitude    string      `json:"latitude"`
	} `json:"extra, omitempty"`
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

type CountryResponse struct {
	Data Country 	`json:"data"`
	Meta Meta       `json:"meta"`
}

// Make a request to retrieve multiple country resources. The request endpoint executed within this method
// is paginated, the second argument to this method allows the consumer to specify a page to request.
// Use the includes slice to enrich the response data, includes for this endpoint are:
// - continent
// - leagues
func (c *Client) Countries(includes []string, page int) (*CountriesResponse, error) {
	url := c.BaseURL + countryUri + "?api_token=" + c.ApiKey

	response := new(CountriesResponse)

	err := c.sendRequest(url, includes, page, response)

	return response, err
}

// Retrieve a single continent resource by ID. Use the includes slice to enrich the response data, includes
// for this endpoint are:
// - countries
func (c *Client) CountryById(id int, includes []string) (*CountryResponse, error) {
	url := c.BaseURL + countryUri + strconv.Itoa(id) + "?api_token=" + c.ApiKey

	response := new(CountryResponse)

	err := c.sendRequest(url, includes, 0, response)

	return response, err
}
