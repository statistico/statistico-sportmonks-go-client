package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var oddsResponse = `{
	"data": [
		{
			"id": 151577019200,
			"fixture_id": 19155301,
			"market_id": 126,
			"bookmaker_id": 2,
			"label": "1",
			"value": "3.75",
			"name": "1",
			"sort_order": 0,
			"market_description": "Winning Margin",
			"probability": "26.67%",
			"dp3": "3.750",
			"fractional": "15\/4",
			"american": "275",
			"winning": false,
			"stopped": false,
			"total": null,
			"handicap": "",
			"participants": null,
			"created_at": "2025-01-26T14:57:19.000000Z",
			"original_label": null,
			"latest_bookmaker_update": "2025-02-10 14:10:51",
			"bookmaker": {
				"id": 2,
				"legacy_id": 2,
				"name": "bet365"
			},
			"market": {
				"id": 126,
				"legacy_id": 136830688,
				"name": "Winning Margin",
				"developer_name": "WINNING_MARGIN",
				"has_winning_calculations": false
			}
		}
	]
}`

func TestOddsByFixtureID(t *testing.T) {
	t.Run("returns a slice of MatchOdds struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/odds/pre-match/fixtures/11867285?api_token=api-key"

		server := mockResponseServer(t, oddsResponse, 200, url)

		client := newTestHTTPClient(server)

		odds, _, err := client.OddsByFixtureID(context.Background(), 11867285)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertOdds(t, odds[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/odds/pre-match/fixtures/11867285?api_token=api-key"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		odds, _, err := client.OddsByFixtureID(context.Background(), 11867285)

		if odds != nil {
			t.Fatalf("Test failed, expected nil, got %+v", odds)
		}

		assertError(t, err)
	})
}

func TestOddsByFixtureAndBookmaker(t *testing.T) {
	t.Run("returns a slice of MatchOdds struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/odds/pre-match/fixtures/11867285/bookmakers/1?api_token=api-key"

		server := mockResponseServer(t, oddsResponse, 200, url)

		client := newTestHTTPClient(server)

		odds, _, err := client.OddsByFixtureAndBookmaker(context.Background(), 11867285, 1)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertOdds(t, odds[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/odds/pre-match/fixtures/11867285/bookmakers/1?api_token=api-key"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		odds, _, err := client.OddsByFixtureAndBookmaker(context.Background(), 11867285, 1)

		if odds != nil {
			t.Fatalf("Test failed, expected nil, got %+v", odds)
		}

		assertError(t, err)
	})
}

func TestOddsByFixtureAndMarket(t *testing.T) {
	t.Run("returns a slice of MatchOdds struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/odds/pre-match/fixtures/11867285/markets/1?api_token=api-key"

		server := mockResponseServer(t, oddsResponse, 200, url)

		client := newTestHTTPClient(server)

		odds, _, err := client.OddsByFixtureAndMarket(context.Background(), 11867285, 1)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertOdds(t, odds[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/odds/pre-match/fixtures/11867285/markets/1?api_token=api-key"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		odds, _, err := client.OddsByFixtureAndMarket(context.Background(), 11867285, 1)

		if odds != nil {
			t.Fatalf("Test failed, expected nil, got %+v", odds)
		}

		assertError(t, err)
	})
}

func assertOdds(t *testing.T, odds Odds) {
	assert.Equal(t, 151577019200, odds.ID)
	assert.Equal(t, 19155301, odds.FixtureID)
	assert.Equal(t, 126, odds.MarketID)
	assert.Equal(t, 2, odds.BookmakerID)
	assert.Equal(t, "1", odds.Label)
	assert.Equal(t, "3.75", odds.Value)
	assert.Equal(t, "1", *odds.Name)
	assert.Equal(t, 0, *odds.SortOrder)
	assert.Equal(t, "Winning Margin", odds.MarketDescription)
	assert.Equal(t, "26.67%", odds.Probability)
	assert.Equal(t, "3.750", odds.DP3)
	assert.Equal(t, "15/4", odds.Fractional)
	assert.Equal(t, "275", odds.American)
	assert.Equal(t, false, odds.Winning)
	assert.Equal(t, false, odds.Stopped)
	assert.Nil(t, odds.Total)
	assert.Equal(t, "", *odds.Handicap)
	assert.Nil(t, odds.Participants)
	assert.Equal(t, "2025-01-26 14:57:19 +0000 UTC", odds.CreatedAt.String())
	assert.Nil(t, odds.OriginalLabel)
	assert.Equal(t, "2025-02-10 14:10:51", odds.LatestBookmakerUpdate)
	assert.Equal(t, 2, odds.Bookmaker.ID)
	assert.Equal(t, 2, odds.Bookmaker.LegacyID)
	assert.Equal(t, "bet365", odds.Bookmaker.Name)
	assert.Equal(t, 126, odds.Market.ID)
	assert.Equal(t, 136830688, odds.Market.LegacyID)
	assert.Equal(t, "Winning Margin", odds.Market.Name)
	assert.Equal(t, "WINNING_MARGIN", odds.Market.DeveloperName)
	assert.Equal(t, false, odds.Market.HasWinningCalculations)
}
