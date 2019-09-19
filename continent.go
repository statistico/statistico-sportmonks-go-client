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

// Make a request to retrieve continents resources. The request endpoint executed within this method
// is paginated, the second argument to this method allows the consumer to specify a page to request
func (c *Client) Continent(includes []string, page int) (*ContinentsResponse, error) {
	str := c.BaseURL + continentUri + "?api_token=" + c.ApiKey

	url := buildUrl(str, includes, page)

	response := new(ContinentsResponse)

	err := c.sendRequest(url, response)

	return response, err
}

// Retrieve a single continent resource by ID
func (c *Client) ContinentById(id int) (*ContinentResponse, error) {
	str := c.BaseURL + continentUri + strconv.Itoa(id) + "?api_token=" + c.ApiKey

	url := buildUrl(str, []string{}, 0)

	response := new(ContinentResponse)

	err := c.sendRequest(url, &response)

	return response, err
}
