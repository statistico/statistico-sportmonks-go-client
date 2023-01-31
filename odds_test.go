package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var oddsResponse = `{
	"data": [
		{
			"id": 38,
			"name": "Goals Over\/Under 2nd Half",
			"suspended": false,
			"bookmaker": {
				"data": [
					{
						"id": 1,
            			"name": "10Bet",
						"odds": {
							"data": [
								{
									"value": 1.91,
									"handicap": null,
									"total": "1.5",
									"label": "Over",
									"probability": "52.36%",
									"dp3": "1.910",
									"american": -110,
									"factional": null,
									"winning": null,
									"stop": true,
									"bookmaker_event_id": 40054117,
									"last_update": {
										"date": "2019-10-05 13:01:00.227530",
										"timezone_type": 3,
										"timezone": "UTC"
									}
								}
							]
						}
					}
				]
			}
		}
	]
}`

func TestOddsByFixtureID(t *testing.T) {
	t.Run("returns a slice of MatchOdds struct", func(t *testing.T) {
		url := defaultBaseURL + "/odds/fixture/11867285?api_token=api-key"

		server := mockResponseServer(t, oddsResponse, 200, url)

		client := newTestHTTPClient(server)

		odds, _, err := client.OddsByFixtureID(context.Background(), 11867285)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		bookmaker := odds[0].BookmakerOdds()
		odd := bookmaker[0].Odds()

		assertMatchOdds(t, &odds[0])
		assertBookmakerOdds(t, &bookmaker[0])
		assertOdds(t, &odd[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/odds/fixture/11867285?api_token=api-key"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		odds, _, err := client.OddsByFixtureID(context.Background(), 11867285)

		if odds != nil {
			t.Fatalf("Test failed, expected nil, got %+v", odds)
		}

		assertError(t, err)
	})
}

func TestInPlayOddsByFixtureID(t *testing.T) {
	t.Run("returns a slice of MatchOdds struct", func(t *testing.T) {
		url := defaultBaseURL + "/odds/inplay/fixture/11867285?api_token=api-key"

		server := mockResponseServer(t, oddsResponse, 200, url)

		client := newTestHTTPClient(server)

		odds, _, err := client.InPlayOddsByFixtureID(context.Background(), 11867285)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		bookmaker := odds[0].BookmakerOdds()
		odd := bookmaker[0].Odds()

		assertMatchOdds(t, &odds[0])
		assertBookmakerOdds(t, &bookmaker[0])
		assertOdds(t, &odd[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/odds/inplay/fixture/11867285?api_token=api-key"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		odds, _, err := client.InPlayOddsByFixtureID(context.Background(), 11867285)

		if odds != nil {
			t.Fatalf("Test failed, expected nil, got %+v", odds)
		}

		assertError(t, err)
	})
}

func TestOddsByFixtureAndBookmaker(t *testing.T) {
	t.Run("returns a slice of MatchOdds struct", func(t *testing.T) {
		url := defaultBaseURL + "/odds/fixture/11867285/bookmaker/1?api_token=api-key"

		server := mockResponseServer(t, oddsResponse, 200, url)

		client := newTestHTTPClient(server)

		odds, _, err := client.OddsByFixtureAndBookmaker(context.Background(), 11867285, 1)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		bookmaker := odds[0].BookmakerOdds()
		odd := bookmaker[0].Odds()

		assertMatchOdds(t, &odds[0])
		assertBookmakerOdds(t, &bookmaker[0])
		assertOdds(t, &odd[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/odds/fixture/11867285/bookmaker/1?api_token=api-key"

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
		url := defaultBaseURL + "/odds/fixture/11867285/market/1?api_token=api-key"

		server := mockResponseServer(t, oddsResponse, 200, url)

		client := newTestHTTPClient(server)

		odds, _, err := client.OddsByFixtureAndMarket(context.Background(), 11867285, 1)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		bookmaker := odds[0].BookmakerOdds()
		odd := bookmaker[0].Odds()

		assertMatchOdds(t, &odds[0])
		assertBookmakerOdds(t, &bookmaker[0])
		assertOdds(t, &odd[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/odds/fixture/11867285/market/1?api_token=api-key"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		odds, _, err := client.OddsByFixtureAndMarket(context.Background(), 11867285, 1)

		if odds != nil {
			t.Fatalf("Test failed, expected nil, got %+v", odds)
		}

		assertError(t, err)
	})
}

func assertMatchOdds(t *testing.T, odds *MatchOdds) {
	assert.Equal(t, 38, odds.ID)
	assert.Equal(t, "Goals Over/Under 2nd Half", odds.Name)
	assert.Equal(t, false, odds.Suspended)
}

func assertBookmakerOdds(t *testing.T, odds *BookmakerOdds) {
	assert.Equal(t, 1, odds.ID)
	assert.Equal(t, "10Bet", odds.Name)
}

func assertOdds(t *testing.T, odds *Odds) {
	assert.Equal(t, float32(1.91), odds.Value)
	assert.Nil(t, odds.Handicap)
	assert.Equal(t, "1.5", odds.Total)
	assert.Equal(t, "Over", odds.Label)
	assert.Equal(t, "52.36%", odds.Probability)
	assert.Equal(t, "1.910", odds.Dp3)
	assert.Equal(t, -110, odds.American)
	assert.Nil(t, odds.Fractional)
	assert.Nil(t, odds.Winning)
	assert.Equal(t, true, odds.Stop)
	assert.Equal(t, 40054117, odds.BookmakerEventID)
	assert.Equal(t, "2019-10-05 13:01:00.227530", odds.LastUpdate.Date)
	assert.Equal(t, 3, odds.LastUpdate.TimezoneType)
	assert.Equal(t, "UTC", odds.LastUpdate.Timezone)
}
