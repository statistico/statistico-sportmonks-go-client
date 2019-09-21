package sportmonks

import "fmt"

const commentaryFixtureUri = "/api/v2.0/commentaries/fixture"

type Commentary struct {
	FixtureID   int         `json:"fixture_id"`
	Important   bool        `json:"important"`
	Order       int         `json:"order"`
	Goal        bool        `json:"goal"`
	Minute      int         `json:"minute"`
	ExtraMinute *int `json:"extra_minute"`
	Comment     string      `json:"comment"`
}

// Make a request to retrieve multiple commentary resources for a given fixture ID.
func (c *ApiClient) CommentariesByFixtureId(fixtureId int) ([]Commentary, *Meta, error) {
	url := fmt.Sprintf(commentaryFixtureUri +"/%d", fixtureId)

	response := new(CommentariesResponse)

	err := c.handleRequest(url, []string{}, response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, &response.Meta, err
}