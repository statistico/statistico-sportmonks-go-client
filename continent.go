package sportmonks

import (
	"strconv"
)

const continentUri = "/api/v2.0/continents"

// Continent struct
type Continent struct {
	ID   int64    `json:"id"`
	Name string `json:"name"`
	Countries struct {
		Data Country `json:"data"`
	} `json:"countries, omitempty"`
}

// ContinentsResponse struct
type ContinentsResponse struct {
	Data []Continent `json:"data"`
	Meta Meta `json:"meta"`
}

// ContinentResponse struct
type ContinentResponse struct {
	Data Continent `json:"data"`
	Meta Meta `json:"meta"`
}

// Make a request to retrieve multiple continent resources. The request endpoint executed within this method
// is paginated, the second argument to this method allows the consumer to specify a page to request.
// Use the includes slice to enrich the response data, includes for this endpoint are:
// - countries
func (c *Client) Continent(includes []string, page int) (*ContinentsResponse, error) {
	str := c.BaseURL + continentUri + "?api_token=" + c.ApiKey

	url := buildUrl(str, includes, page)

	response := new(ContinentsResponse)

	err := c.sendRequest(url, response)

	return response, err
}

// Retrieve a single continent resource by ID. Use the includes slice to enrich the response data, includes
// for this endpoint are:
// - countries
func (c *Client) ContinentById(id int, includes []string) (*ContinentResponse, error) {
	str := c.BaseURL + continentUri + strconv.Itoa(id) + "?api_token=" + c.ApiKey

	url := buildUrl(str, includes, 0)

	response := new(ContinentResponse)

	err := c.sendRequest(url, &response)

	return response, err
}
