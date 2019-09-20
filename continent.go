package statistico

import (
	"strconv"
)

const continentUri = "/api/v2.0/continents"

// Continent struct
type Continent struct {
	ID   int64    `json:"id"`
	Name string `json:"name"`
	Countries struct {
		Data []Country `json:"data"`
	} `json:"countries, omitempty"`
}

// Make a request to retrieve multiple continent resources. The request endpoint executed within this method
// is paginated, the first argument to this method allows the consumer to specify a page to request.
// Use the includes slice to enrich the response data.
func (c *SportMonksClient) Continents(page int, includes []string) ([]Continent, *Meta, error) {
	response := new(ContinentsResponse)

	err := c.handlePaginatedRequest(continentUri, includes, page, response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, &response.Meta, err
}

// Retrieve a single continent resource by ID. Use the includes slice to enrich the response data.
func (c *SportMonksClient) ContinentById(id int, includes []string) (*Continent, *Meta, error) {
	url := continentUri + "/" + strconv.Itoa(id)

	response := new(ContinentResponse)

	err := c.handleRequest(url, includes, response)

	if err != nil {
		return nil, nil, err
	}

	return &response.Data, &response.Meta, err
}
