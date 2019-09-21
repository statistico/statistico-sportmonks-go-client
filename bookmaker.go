package sportmonks

import (
	"fmt"
)

const bookmakerUri = "/api/v2.0/bookmakers"

type Bookmaker struct {
	ID   int         `json:"id"`
	Name string      `json:"name"`
	Logo *string     `json:"logo"`
}

// Make a request to retrieve multiple bookmaker resources.
func (c *ApiClient) Bookmaker(page int, includes []string) ([]Continent, *Meta, error) {
	response := new(ContinentsResponse)

	err := c.handlePaginatedRequest(bookmakerUri, includes, page, response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, &response.Meta, err
}

// Retrieve a single bookmaker resource by ID. Use the includes slice to enrich the response data.
func (c *ApiClient) BookmakerById(id int, includes []string) (*Bookmaker, *Meta, error) {
	url := fmt.Sprintf(bookmakerUri +"/%d", id)

	response := new(BookmakerResponse)

	err := c.handleRequest(url, includes, response)

	if err != nil {
		return nil, nil, err
	}

	return &response.Data, &response.Meta, err
}