package sportmonks

import (
	"context"
	"fmt"
	"net/url"
)

type (
	// MatchOdds provides a struct representation of a MatchOdds resource
	MatchOdds struct {
		ID                int               `json:"id"`
		Name              string            `json:"name"`
		Suspended         bool              `json:"suspended"`
		BookmakerOddsData bookmakerOddsData `json:"bookmaker"`
	}

	// BookmakerOdds provides a struct representation of a BookmakerOdds resource
	BookmakerOdds struct {
		ID       int      `json:"id"`
		Name     string   `json:"name"`
		OddsData oddsData `json:"odds"`
	}

	// Odds provides a struct representation of a Odds resource
	Odds struct {
		Value            string   `json:"value"`
		Handicap         *string  `json:"handicap"`
		Total            string   `json:"total"`
		Label            string   `json:"label"`
		Probability      string   `json:"probability"`
		Dp3              string   `json:"dp3"`
		American         int      `json:"american"`
		Fractional       *string  `json:"factional"`
		Winning          *string  `json:"winning"`
		Stop             bool     `json:"stop"`
		BookmakerEventID string   `json:"bookmaker_event_id"`
		LastUpdate       DateTime `json:"last_update"`
	}
)

// BookmakerOdds returns bookmaker specific odds for a match
func (m *MatchOdds) BookmakerOdds() []BookmakerOdds {
	return m.BookmakerOddsData.Data
}

// Odds returns odds data for a specific bookmaker
func (b *BookmakerOdds) Odds() []Odds {
	return b.OddsData.Data
}

// OddsByFixtureID fetches MatchOdds resources associated to a fixture.
func (c *HTTPClient) OddsByFixtureID(ctx context.Context, fixtureID int) ([]MatchOdds, *Meta, error) {
	path := fmt.Sprintf(oddsFixtureURI+"/%d", fixtureID)

	response := struct {
		Data []MatchOdds `json:"data"`
		Meta *Meta       `json:"meta"`
	}{}

	err := c.getResource(ctx, path, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// InPlayOddsByFixtureID fetches MatchOdds resources associated to a fixture that is in play.
func (c *HTTPClient) InPlayOddsByFixtureID(ctx context.Context, fixtureID int) ([]MatchOdds, *Meta, error) {
	path := fmt.Sprintf(inPlayOddsFixtureURI+"/%d", fixtureID)

	response := struct {
		Data []MatchOdds `json:"data"`
		Meta *Meta       `json:"meta"`
	}{}

	err := c.getResource(ctx, path, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// OddsByFixtureAndBookmaker fetches MatchOdds resources associated to a fixture and bookmaker.
func (c *HTTPClient) OddsByFixtureAndBookmaker(ctx context.Context, fixtureID, bookmakerID int) ([]MatchOdds, *Meta, error) {
	path := fmt.Sprintf(oddsFixtureURI+"/%d/bookmaker/%d", fixtureID, bookmakerID)

	response := struct {
		Data []MatchOdds `json:"data"`
		Meta *Meta       `json:"meta"`
	}{}

	err := c.getResource(ctx, path, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// OddsByFixtureAndMarket fetches MatchOdds resources associated to a fixture and market.
func (c *HTTPClient) OddsByFixtureAndMarket(ctx context.Context, fixtureID, marketID int) ([]MatchOdds, *Meta, error) {
	path := fmt.Sprintf(oddsFixtureURI+"/%d/market/%d", fixtureID, marketID)

	response := struct {
		Data []MatchOdds `json:"data"`
		Meta *Meta       `json:"meta"`
	}{}

	err := c.getResource(ctx, path, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
