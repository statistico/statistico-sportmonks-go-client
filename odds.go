package sportmonks

import (
	"context"
	"fmt"
	"net/url"
)

type (
	MatchOdds struct {
		ID                int               `json:"id"`
		Name              string            `json:"name"`
		Suspended         bool               `json:"suspended"`
		BookmakerOddsData BookmakerOddsData `json:"bookmaker"`
	}

	BookmakerOdds struct {
		ID       int      `json:"id"`
		Name     string   `json:"name"`
		OddsData OddsData `json:"odds"`
	}

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
		BookmakerEventID string      `json:"bookmaker_event_id"`
		LastUpdate       DateTime `json:"last_update"`
	}
)

func (m *MatchOdds) BookmakerOdds() []BookmakerOdds {
	return m.BookmakerOddsData.Data
}

func (b *BookmakerOdds) Odds() []Odds {
	return b.OddsData.Data
}

// OddsByFixtureId returns a slice of MatchOdds struct associated to a fixture.
func (c *HTTPClient) OddsByFixtureId(ctx context.Context, fixtureId int) ([]MatchOdds, *Meta, error) {
	path := fmt.Sprintf(oddsFixtureUri + "/%d", fixtureId)

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

// InPlayOddsByFixtureId returns a slice of MatchOdds struct associated to a fixture.
func (c *HTTPClient) InPlayOddsByFixtureId(ctx context.Context, fixtureId int) ([]MatchOdds, *Meta, error) {
	path := fmt.Sprintf(inPlayOddsFixtureUri + "/%d", fixtureId)

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

// OddsByFixtureAndBookmaker returns a slice of MatchOdds struct associated to a fixture and bookmaker.
func (c *HTTPClient) OddsByFixtureAndBookmaker(ctx context.Context, fixtureId, bookmakerId int) ([]MatchOdds, *Meta, error) {
	path := fmt.Sprintf(oddsFixtureUri + "/%d/bookmaker/%d", fixtureId, bookmakerId)

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

// OddsByFixtureAndMarket returns a slice of MatchOdds struct associated to a fixture and market.
func (c *HTTPClient) OddsByFixtureAndMarket(ctx context.Context, fixtureId, marketId int) ([]MatchOdds, *Meta, error) {
	path := fmt.Sprintf(oddsFixtureUri + "/%d/market/%d", fixtureId, marketId)

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
