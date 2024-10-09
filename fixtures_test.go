package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var fixtureResponse = `{
	"data": {
      "id": 463,
      "sport_id": 1,
      "league_id": 8,
      "season_id": 2,
      "stage_id": 2,
      "group_id": null,
      "aggregate_id": null,
      "round_id": 43,
      "state_id": 5,
      "venue_id": 209,
      "name": "Tottenham Hotspur vs Manchester City",
      "starting_at": "2010-08-14 11:45:00",
      "result_info": "Game ended in draw.",
      "leg": "1/1",
      "details": null,
      "length": 90,
      "placeholder": false,
      "has_odds": false,
      "has_premium_odds": false,
      "starting_at_timestamp": 1281786300,
      "round": {
		  "id": 43,
		  "sport_id": 1,
		  "league_id": 8,
		  "season_id": 2,
		  "stage_id": 2,
		  "name": "1",
		  "finished": true,
		  "is_current": false,
		  "starting_at": "2010-08-14",
		  "ending_at": "2010-08-16",
		  "games_in_current_week": false
      },
      "scores": [
		{
			"id": 14902138,
			"fixture_id": 19134492,
			"type_id": 1525,
			"participant_id": 1,
			"score": {
				"goals": 0,
				"participant": "home"
			},
			"description": "CURRENT"
		}
      ]
    },
	"subscription": [
		{
			"meta": {
				"trial_ends_at": null,
				"ends_at": "2024-10-26 12:06:34",
				"current_timestamp": 1728372666
			},
			"plans": [
				{
					"plan": "Joe Sweeny Custom Plan",
					"sport": "Football",
					"category": "Advanced"
				}
			],
			"add_ons": [],
			"widgets": []
		}
	],
	"rate_limit": {
		"resets_in_seconds": 3386,
		"remaining": 2997,
		"requested_entity": "Fixture"
	},
	"timezone": "UTC"
}`

var fixturesResponse = `{
	"data": [
		{
			  "id": 463,
			  "sport_id": 1,
			  "league_id": 8,
			  "season_id": 2,
			  "stage_id": 2,
			  "group_id": null,
			  "aggregate_id": null,
			  "round_id": 43,
			  "state_id": 5,
			  "venue_id": 209,
			  "name": "Tottenham Hotspur vs Manchester City",
			  "starting_at": "2010-08-14 11:45:00",
			  "result_info": "Game ended in draw.",
			  "leg": "1/1",
			  "details": null,
			  "length": 90,
			  "placeholder": false,
			  "has_odds": false,
			  "has_premium_odds": false,
			  "starting_at_timestamp": 1281786300,
			  "round": {
				  "id": 43,
				  "sport_id": 1,
				  "league_id": 8,
				  "season_id": 2,
				  "stage_id": 2,
				  "name": "1",
				  "finished": true,
				  "is_current": false,
				  "starting_at": "2010-08-14",
				  "ending_at": "2010-08-16",
				  "games_in_current_week": false
			  }
		}
	],
	"pagination": {
		"count": 2,
		"per_page": 25,
		"current_page": 1,
		"next_page": null,
		"has_more": false
	},
	"subscription": [
		{
			"meta": {
				"trial_ends_at": null,
				"ends_at": "2024-10-26 12:06:34",
				"current_timestamp": 1728372666
			},
			"plans": [
				{
					"plan": "Joe Sweeny Custom Plan",
					"sport": "Football",
					"category": "Advanced"
				}
			],
			"add_ons": [],
			"widgets": []
		}
	],
	"rate_limit": {
		"resets_in_seconds": 3386,
		"remaining": 2997,
		"requested_entity": "Fixture"
	},
	"timezone": "UTC"
}`

func TestFixtureByID(t *testing.T) {
	t.Run("returns a single Fixture struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/fixtures/11867285?api_token=api-key&include="

		server := mockResponseServer(t, fixtureResponse, 200, url)

		client := newTestHTTPClient(server)

		fixture, _, err := client.FixtureByID(context.Background(), 11867285, []string{}, map[string][]int{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertFixture(t, fixture)
	})

	t.Run("returns a single Fixture struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/football/fixtures/11867285?api_token=api-key&include=round%3Bscores"

		server := mockResponseServer(t, fixtureResponse, 200, url)

		client := newTestHTTPClient(server)

		fixture, _, err := client.FixtureByID(
			context.Background(),
			11867285,
			[]string{"round", "scores"},
			map[string][]int{},
		)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertFixture(t, fixture)
		assertRound(t, fixture.Round)
		assertScore(t, fixture.Scores[0])
	})

	t.Run("returns a single Fixture struct with includes data and filter parameters", func(t *testing.T) {
		url := defaultBaseURL + "/football/fixtures/11867285?api_token=api-key&include=round%3Bstage%3Bgoals&leagues=8%2C10"

		server := mockResponseServer(t, fixtureResponse, 200, url)

		client := newTestHTTPClient(server)

		fixture, _, err := client.FixtureByID(
			context.Background(),
			11867285,
			[]string{"round", "stage", "goals"},
			map[string][]int{
				"leagues": {
					8,
					10,
				},
			},
		)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertFixture(t, fixture)
		assertRound(t, fixture.Round)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/fixtures/11867285?api_token=api-key&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		fixture, _, err := client.FixtureByID(context.Background(), 11867285, []string{}, map[string][]int{})

		if fixture != nil {
			t.Fatalf("Test failed, expected nil, got %+v", fixture)
		}

		assertError(t, err)
	})

	t.Run("can handle response details response", func(t *testing.T) {
		url := defaultBaseURL + "/football/fixtures/11867285?api_token=api-key&include="

		server := mockResponseServer(t, fixtureResponse, 200, url)

		client := newTestHTTPClient(server)

		_, details, _ := client.FixtureByID(context.Background(), 11867285, []string{}, map[string][]int{})

		assertResponseDetails(t, details, "Fixture")
	})
}

func TestFixturesByID(t *testing.T) {
	t.Run("returns slice of Fixture struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/fixtures/multi/11867285,555?api_token=api-key&include="

		server := mockResponseServer(t, fixturesResponse, 200, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.FixturesByID(
			context.Background(),
			[]int{11867285, 555},
			[]string{},
			map[string][]int{},
		)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertFixture(t, &fixtures[0])
	})

	t.Run("returns slice of Fixture struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/football/fixtures/multi/11867285,555?api_token=api-key&include=round%3Bstage%3Bgoals"

		server := mockResponseServer(t, fixturesResponse, 200, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.FixturesByID(
			context.Background(),
			[]int{11867285, 555},
			[]string{"round", "stage", "goals"},
			map[string][]int{},
		)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertFixture(t, &fixtures[0])
		assertRound(t, fixtures[0].Round)
	})

	t.Run("returns slice of Fixture struct with includes data and filter parameters", func(t *testing.T) {
		url := defaultBaseURL + "/football/fixtures/multi/11867285,555?api_token=api-key&include=round%3Bstage%3Bgoals&leagues=8%2C10"

		server := mockResponseServer(t, fixturesResponse, 200, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.FixturesByID(
			context.Background(),
			[]int{11867285, 555},
			[]string{"round", "stage", "goals"},
			map[string][]int{
				"leagues": {
					8,
					10,
				},
			},
		)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertFixture(t, &fixtures[0])
		assertRound(t, fixtures[0].Round)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/fixtures/multi/11867285,555?api_token=api-key&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.FixturesByID(context.Background(), []int{11867285, 555}, []string{}, map[string][]int{})

		if fixtures != nil {
			t.Fatalf("Test failed, expected nil, got %+v", fixtures)
		}

		assertError(t, err)
	})

	t.Run("can handle response details response", func(t *testing.T) {
		url := defaultBaseURL + "/football/fixtures/multi/11867285,555?api_token=api-key&include="

		server := mockResponseServer(t, fixturesResponse, 200, url)

		client := newTestHTTPClient(server)

		_, details, _ := client.FixturesByID(
			context.Background(),
			[]int{11867285, 555},
			[]string{},
			map[string][]int{},
		)

		assertResponseDetails(t, details, "Fixture")
	})
}

func TestFixturesByDate(t *testing.T) {
	str := "2014-11-12T11:45:26.371Z"
	d, err := time.Parse(time.RFC3339, str)

	if err != nil {
		t.Fatalf("Test failed, expected nil, got %+v", err.Error())
	}

	t.Run("returns slice of Fixture struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/fixtures/date/2014-11-12?api_token=api-key&include="

		server := mockResponseServer(t, fixturesResponse, 200, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.FixturesByDate(context.Background(), d, []string{}, map[string][]int{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertFixture(t, &fixtures[0])
	})

	t.Run("returns slice of Fixture struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/football/fixtures/date/2014-11-12?api_token=api-key&include=round%3Bstage%3Bgoals"

		server := mockResponseServer(t, fixturesResponse, 200, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.FixturesByDate(
			context.Background(),
			d,
			[]string{"round", "stage", "goals"},
			map[string][]int{},
		)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertFixture(t, &fixtures[0])
		assertRound(t, fixtures[0].Round)
	})

	t.Run("returns slice of Fixture struct with includes data and filter parameters", func(t *testing.T) {
		url := defaultBaseURL + "/football/fixtures/date/2014-11-12?api_token=api-key&include=round%3Bstage%3Bgoals&markets=8%2C10"

		server := mockResponseServer(t, fixturesResponse, 200, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.FixturesByDate(
			context.Background(),
			d,
			[]string{"round", "stage", "goals"},
			map[string][]int{
				"markets": {
					8,
					10,
				},
			},
		)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertFixture(t, &fixtures[0])
		assertRound(t, fixtures[0].Round)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/fixtures/date/2014-11-12?api_token=api-key&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.FixturesByDate(context.Background(), d, []string{}, map[string][]int{})

		if fixtures != nil {
			t.Fatalf("Test failed, expected nil, got %+v", fixtures)
		}

		assertError(t, err)
	})

	t.Run("can handle response details response", func(t *testing.T) {
		url := defaultBaseURL + "/football/fixtures/date/2014-11-12?api_token=api-key&include="

		server := mockResponseServer(t, fixturesResponse, 200, url)

		client := newTestHTTPClient(server)

		_, details, _ := client.FixturesByDate(context.Background(), d, []string{}, map[string][]int{})

		assertResponseDetails(t, details, "Fixture")
	})
}

func TestFixturesBetween(t *testing.T) {
	dateFrom, err := time.Parse(time.RFC3339, "2014-11-12T11:45:26.371Z")

	if err != nil {
		t.Fatalf("Test failed, expected nil, got %+v", err.Error())
	}

	dateTo, err := time.Parse(time.RFC3339, "2014-12-12T11:45:26.371Z")

	if err != nil {
		t.Fatalf("Test failed, expected nil, got %+v", err.Error())
	}

	t.Run("returns slice of Fixture struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/fixtures/between/2014-11-12/2014-12-12?api_token=api-key&include="

		server := mockResponseServer(t, fixturesResponse, 200, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.FixturesBetween(context.Background(), dateFrom, dateTo, []string{}, map[string][]int{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertFixture(t, &fixtures[0])
	})

	t.Run("returns slice of Fixture struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/football/fixtures/between/2014-11-12/2014-12-12?api_token=api-key&include=round%3Bstage%3Bgoals"

		server := mockResponseServer(t, fixturesResponse, 200, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.FixturesBetween(
			context.Background(),
			dateFrom,
			dateTo,
			[]string{"round", "stage", "goals"},
			map[string][]int{},
		)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertFixture(t, &fixtures[0])
		assertRound(t, fixtures[0].Round)
	})

	t.Run("returns slice of Fixture struct with includes data and filter parameters", func(t *testing.T) {
		url := defaultBaseURL + "/football/fixtures/between/2014-11-12/2014-12-12?api_token=api-key&include=round%3Bstage%3Bgoals&leagues=8%2C10"

		server := mockResponseServer(t, fixturesResponse, 200, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.FixturesBetween(
			context.Background(),
			dateFrom,
			dateTo,
			[]string{"round", "stage", "goals"},
			map[string][]int{
				"leagues": {
					8,
					10,
				},
			},
		)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertFixture(t, &fixtures[0])
		assertRound(t, fixtures[0].Round)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/fixtures/between/2014-11-12/2014-12-12?api_token=api-key&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.FixturesBetween(context.Background(), dateFrom, dateTo, []string{}, map[string][]int{})

		if fixtures != nil {
			t.Fatalf("Test failed, expected nil, got %+v", fixtures)
		}

		assertError(t, err)
	})

	t.Run("can handle response details response", func(t *testing.T) {
		url := defaultBaseURL + "/football/fixtures/between/2014-11-12/2014-12-12?api_token=api-key&include="

		server := mockResponseServer(t, fixturesResponse, 200, url)

		client := newTestHTTPClient(server)

		_, details, _ := client.FixturesBetween(context.Background(), dateFrom, dateTo, []string{}, map[string][]int{})

		assertResponseDetails(t, details, "Fixture")
	})
}

func TestFixturesBetweenForTeam(t *testing.T) {
	dateFrom, err := time.Parse(time.RFC3339, "2014-11-12T11:45:26.371Z")

	if err != nil {
		t.Fatalf("Test failed, expected nil, got %+v", err.Error())
	}

	dateTo, err := time.Parse(time.RFC3339, "2014-12-12T11:45:26.371Z")

	if err != nil {
		t.Fatalf("Test failed, expected nil, got %+v", err.Error())
	}

	t.Run("returns slice of Fixture struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/fixtures/between/2014-11-12/2014-12-12/1?api_token=api-key&include="

		server := mockResponseServer(t, fixturesResponse, 200, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.FixturesBetweenForTeam(
			context.Background(),
			dateFrom,
			dateTo,
			1,
			[]string{},
			map[string][]int{},
		)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertFixture(t, &fixtures[0])
	})

	t.Run("returns slice of Fixture struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/football/fixtures/between/2014-11-12/2014-12-12/1?api_token=api-key&include=round%3Bstage%3Bgoals"

		server := mockResponseServer(t, fixturesResponse, 200, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.FixturesBetweenForTeam(
			context.Background(),
			dateFrom,
			dateTo,
			1,
			[]string{"round", "stage", "goals"},
			map[string][]int{},
		)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertFixture(t, &fixtures[0])
		assertRound(t, fixtures[0].Round)
	})

	t.Run("returns slice of Fixture struct with includes data and filter parameters", func(t *testing.T) {
		url := defaultBaseURL + "/football/fixtures/between/2014-11-12/2014-12-12/1?api_token=api-key&include=round%3Bstage%3Bgoals%3Aorder%28starting_at%7Casc%29&leagues=8%2C10"

		server := mockResponseServer(t, fixturesResponse, 200, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.FixturesBetweenForTeam(
			context.Background(),
			dateFrom,
			dateTo,
			1,
			[]string{"round", "stage", "goals:order(starting_at|asc)"},
			map[string][]int{
				"leagues": {
					8,
					10,
				},
			},
		)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertFixture(t, &fixtures[0])
		assertRound(t, fixtures[0].Round)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/fixtures/between/2014-11-12/2014-12-12/1?api_token=api-key&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.FixturesBetweenForTeam(
			context.Background(),
			dateFrom,
			dateTo,
			1,
			[]string{},
			map[string][]int{},
		)

		if fixtures != nil {
			t.Fatalf("Test failed, expected nil, got %+v", fixtures)
		}

		assertError(t, err)
	})

	t.Run("can handle response details response", func(t *testing.T) {
		url := defaultBaseURL + "/football/fixtures/between/2014-11-12/2014-12-12?api_token=api-key&include="

		server := mockResponseServer(t, fixturesResponse, 200, url)

		client := newTestHTTPClient(server)

		_, details, _ := client.FixturesBetween(context.Background(), dateFrom, dateTo, []string{}, map[string][]int{})

		assertResponseDetails(t, details, "Fixture")
	})
}

func assertFixture(t *testing.T, fixture *Fixture) {
	assert.Equal(t, 463, fixture.ID)
	assert.Equal(t, 1, fixture.SportID)
	assert.Equal(t, 8, fixture.LeagueID)
	assert.Equal(t, 2, fixture.SeasonID)
	assert.Equal(t, 2, fixture.StageID)
	assert.Nil(t, fixture.GroupID)
	assert.Nil(t, fixture.AggregateID)
	assert.Equal(t, 43, fixture.RoundID)
	assert.Equal(t, 5, fixture.StateID)
	assert.Equal(t, 209, fixture.VenueID)
	assert.Equal(t, "Tottenham Hotspur vs Manchester City", fixture.Name)
	assert.Equal(t, "2010-08-14 11:45:00", fixture.StartingAt)
	assert.Equal(t, "Game ended in draw.", fixture.ResultInfo)
	assert.Equal(t, "1/1", fixture.Leg)
	assert.Nil(t, fixture.Details)
	assert.Equal(t, 90, fixture.Length)
	assert.Equal(t, false, fixture.Placeholder)
	assert.Equal(t, false, fixture.HasOdds)
	assert.Equal(t, false, fixture.HasPremiumOdds)
	assert.Equal(t, int64(1281786300), fixture.StartingAtTimestamp)
}

func assertScore(t *testing.T, score Score) {
	assert.Equal(t, 14902138, score.ID)
	assert.Equal(t, 19134492, score.FixtureID)
	assert.Equal(t, 1525, score.TypeID)
	assert.Equal(t, 1, score.ParticipantID)
	assert.Equal(t, 0, score.ScoreData.Goals)
	assert.Equal(t, "home", score.ScoreData.Participant)
	assert.Equal(t, "CURRENT", score.Description)
}
