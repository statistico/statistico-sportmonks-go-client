package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Team provides a struct representation of a Team resource.
type Team struct {
	ID           int    `json:"id"`
	SportID      int    `json:"sport_id"`
	CountryID    int    `json:"country_id"`
	VenueID      int    `json:"venue_id"`
	Gender       string `json:"gender"`
	Name         string `json:"name"`
	ShortCode    string `json:"short_code"`
	ImagePath    string `json:"image_path"`
	Founded      int    `json:"founded"`
	Type         string `json:"type"`
	Placeholder  bool   `json:"placeholder"`
	LastPlayedAt string `json:"last_played_at"`
	Venue        *Venue `json:"venue,omitempty"`
}

// TeamByID fetches a Team resource by ID. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) TeamByID(ctx context.Context, id int, includes []string, filters map[string][]int) (*Team, *Meta, error) {
	path := fmt.Sprintf(teamsURI+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ";")},
	}

	formatFilters(&values, filters)

	response := struct {
		Data *Team `json:"data"`
		Meta *Meta `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// TeamsBySeasonID fetches Team resources associated to a season ID. The endpoint used within this method is paginated,
// to select the required page use the 'page' method argument. Page information including current page and total page
// are included within the Meta response. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) TeamsBySeasonID(ctx context.Context, seasonID, page int, includes []string) ([]Team, *Meta, error) {
	path := fmt.Sprintf(teamsSeasonURI+"/%d", seasonID)

	values := url.Values{
		"include": {strings.Join(includes, ";")},
		"page":    {strconv.Itoa(page)},
	}

	response := struct {
		Data []Team `json:"data"`
		Meta *Meta  `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
