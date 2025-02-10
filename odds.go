package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

type (
	// Odds provides a struct representation of an Odds resource.
	Odds struct {
		ID                    int        `json:"id"`
		FixtureID             int        `json:"fixture_id"`
		MarketID              int        `json:"market_id"`
		BookmakerID           int        `json:"bookmaker_id"`
		Label                 string     `json:"label"`
		Value                 string     `json:"value"`
		Name                  *string    `json:"name"`
		SortOrder             *int       `json:"sort_order"`
		MarketDescription     string     `json:"market_description"`
		Probability           string     `json:"probability"`
		DP3                   string     `json:"dp3"`
		Fractional            string     `json:"fractional"`
		American              string     `json:"american"`
		Winning               bool       `json:"winning"`
		Stopped               bool       `json:"stopped"`
		Total                 *string    `json:"total"`
		Handicap              *string    `json:"handicap"`
		Participants          *string    `json:"participants"`
		CreatedAt             time.Time  `json:"created_at"`
		OriginalLabel         *string    `json:"original_label"`
		LatestBookmakerUpdate string     `json:"latest_bookmaker_update"`
		Bookmaker             *Bookmaker `json:"bookmaker"`
		Market                *Market    `json:"market"`
	}
)

// OddsByFixtureID fetches MatchOdds resources associated to a fixture.
func (c *HTTPClient) OddsByFixtureID(ctx context.Context, fixtureID int) ([]Odds, *Meta, error) {
	path := fmt.Sprintf(oddsFixtureURI+"/%d", fixtureID)

	response := struct {
		Data []Odds `json:"data"`
		Meta *Meta  `json:"meta"`
	}{}

	err := c.getResource(ctx, path, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// OddsByFixtureAndBookmaker fetches MatchOdds resources associated to a fixture and bookmaker.
func (c *HTTPClient) OddsByFixtureAndBookmaker(ctx context.Context, fixtureID, bookmakerID int) ([]Odds, *Meta, error) {
	path := fmt.Sprintf(oddsFixtureURI+"/%d/bookmakers/%d", fixtureID, bookmakerID)

	response := struct {
		Data []Odds `json:"data"`
		Meta *Meta  `json:"meta"`
	}{}

	err := c.getResource(ctx, path, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// OddsByFixtureAndMarket fetches MatchOdds resources associated to a fixture and market.
func (c *HTTPClient) OddsByFixtureAndMarket(ctx context.Context, fixtureID, marketID int) ([]Odds, *Meta, error) {
	path := fmt.Sprintf(oddsFixtureURI+"/%d/markets/%d", fixtureID, marketID)

	response := struct {
		Data []Odds `json:"data"`
		Meta *Meta  `json:"meta"`
	}{}

	err := c.getResource(ctx, path, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
