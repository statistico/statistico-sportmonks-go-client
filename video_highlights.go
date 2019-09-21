package sportmonks

const videoHighlightsUri = "/api/v2.0/highlights"

type VideoHighlights struct {
	FixtureID int         `json:"fixture_id"`
	EventID   *int `json:"event_id"`
	Location  string      `json:"location"`
	Type      string      `json:"type"`
	CreatedAt struct {
		Date         string `json:"date"`
		TimezoneType int    `json:"timezone_type"`
		Timezone     string `json:"timezone"`
	} `json:"created_at"`
	Fixture *Fixture`json:"fixture, omitempty"`
}

// Make a request to retrieve multiple video highlight resources. The request endpoint executed within this method
// is paginated, the first argument to this method allows the consumer to specify a page to request.
// Use the includes slice to enrich the response data.
func (c *ApiClient) VideoHighlights(page int, includes []string) ([]VideoHighlights, *Meta, error) {
	response := new(VideoHighlightsResponse)

	err := c.handlePaginatedRequest(videoHighlightsUri, includes, page, response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, &response.Meta, err
}
