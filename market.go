package sportmonks

import (
	"fmt"
)

const marketsUri = "/api/v.20/markets"

type Market struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Make a request to retrieve multiple market resources.
func (c *ApiClient) Markets() ([]Market, *Meta, error) {
	response := new(MarketsResponse)

	err := c.handleRequest(marketsUri, []string{}, response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, &response.Meta, err
}

// Retrieve a single market resource by ID.
func (c *ApiClient) MarketById(id int) (*Market, *Meta, error) {
	url := fmt.Sprintf(marketsUri + "/%d", id)

	response := new(MarketResponse)

	err := c.handleRequest(url, []string{}, response)

	if err != nil {
		return nil, nil, err
	}

	return &response.Data, &response.Meta, err
}
