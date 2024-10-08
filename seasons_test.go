package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var seasonsResponse = `{
	"data": [
		{
		  "id": 2,
		  "sport_id": 1,
		  "league_id": 8,
		  "tie_breaker_rule_id": 1526,
		  "name": "2010/2011",
		  "finished": true,
		  "pending": false,
		  "is_current": false,
		  "starting_at": "2010-08-14",
		  "ending_at": "2011-05-22",
		  "standings_recalculated_at": "2023-05-24 08:28:07",
		  "games_in_current_week": false
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
		"requested_entity": "Season"
	},
	"timezone": "UTC"
}`

var seasonsIncludesResponse = `{
	"data": [
		{
		  "id": 2,
		  "sport_id": 1,
		  "league_id": 8,
		  "tie_breaker_rule_id": 1526,
		  "name": "2010/2011",
		  "finished": true,
		  "pending": false,
		  "is_current": false,
		  "starting_at": "2010-08-14",
		  "ending_at": "2011-05-22",
		  "standings_recalculated_at": "2023-05-24 08:28:07",
		  "games_in_current_week": false,
		  "league": {
			  "id": 8,
			  "sport_id": 1,
			  "country_id": 462,
			  "name": "Premier League",
			  "active": true,
			  "short_code": "UK PL",
			  "image_path": "https://cdn.sportmonks.com/images/soccer/leagues/8/8.png",
			  "type": "league",
			  "sub_type": "domestic",
			  "last_played_at": "2024-09-22 15:30:00",
			  "category": 1,
			  "has_jerseys": false
          }
		}
	]
}`

var seasonResponse = `{
	"data": {
		  "id": 2,
		  "sport_id": 1,
		  "league_id": 8,
		  "tie_breaker_rule_id": 1526,
		  "name": "2010/2011",
		  "finished": true,
		  "pending": false,
		  "is_current": false,
		  "starting_at": "2010-08-14",
		  "ending_at": "2011-05-22",
		  "standings_recalculated_at": "2023-05-24 08:28:07",
		  "games_in_current_week": false
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
		"requested_entity": "Season"
	},
	"timezone": "UTC"
}`

var seasonIncludesResponse = `{
	"data": {
      "id": 2,
      "sport_id": 1,
      "league_id": 8,
      "tie_breaker_rule_id": 1526,
      "name": "2010/2011",
      "finished": true,
      "pending": false,
      "is_current": false,
      "starting_at": "2010-08-14",
      "ending_at": "2011-05-22",
      "standings_recalculated_at": "2023-05-24 08:28:07",
      "games_in_current_week": false,
	  "league": {
		  "id": 8,
		  "sport_id": 1,
		  "country_id": 462,
		  "name": "Premier League",
		  "active": true,
		  "short_code": "UK PL",
		  "image_path": "https://cdn.sportmonks.com/images/soccer/leagues/8/8.png",
		  "type": "league",
		  "sub_type": "domestic",
		  "last_played_at": "2024-09-22 15:30:00",
		  "category": 1,
		  "has_jerseys": false
	  }
    }
}`

func TestSeasons(t *testing.T) {
	t.Run("returns a slice of Season struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/seasons?api_token=api-key&include=&page=1"

		server := mockResponseServer(t, seasonsResponse, 200, url)

		client := newTestHTTPClient(server)

		seasons, _, err := client.Seasons(context.Background(), 1, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertSeason(t, &seasons[0])
	})

	t.Run("returns a slice of Season struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/football/seasons?api_token=api-key&include=league%3Bgoalscorers%3Brounds&page=1"

		server := mockResponseServer(t, seasonsIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		seasons, _, err := client.Seasons(context.Background(), 1, []string{"league", "goalscorers", "rounds"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertSeason(t, &seasons[0])
		assertLeague(t, seasons[0].League)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/seasons?api_token=api-key&include=&page=1"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		seasons, _, err := client.Seasons(context.Background(), 1, []string{})

		if seasons != nil {
			t.Fatalf("Test failed, expected nil, got %+v", seasons)
		}

		assertError(t, err)
	})

	t.Run("can handle response details", func(t *testing.T) {
		url := defaultBaseURL + "/football/seasons?api_token=api-key&include=&page=1"

		server := mockResponseServer(t, seasonsResponse, 200, url)

		client := newTestHTTPClient(server)

		_, details, _ := client.Seasons(context.Background(), 1, []string{})

		assertResponseDetails(t, details, "Season")
	})
}

func TestSeasonByID(t *testing.T) {
	t.Run("returns a single Season struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/seasons/55?api_token=api-key&deleted=1&include="

		server := mockResponseServer(t, seasonResponse, 200, url)

		client := newTestHTTPClient(server)

		season, _, err := client.SeasonByID(context.Background(), 55, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertSeason(t, season)
	})

	t.Run("returns a single Season struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/football/seasons/55?api_token=api-key&deleted=1&include=league%3Bgoalscorers%3Brounds%3Bresults"

		server := mockResponseServer(t, seasonIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		season, _, err := client.SeasonByID(context.Background(), 55, []string{"league", "goalscorers", "rounds", "results"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertSeason(t, season)
		assertLeague(t, season.League)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/seasons/55?api_token=api-key&deleted=1&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		season, _, err := client.SeasonByID(context.Background(), 55, []string{})

		if season != nil {
			t.Fatalf("Test failed, expected nil, got %+v", season)
		}

		assertError(t, err)
	})

	t.Run("can handle response details response", func(t *testing.T) {
		url := defaultBaseURL + "/football/seasons/55?api_token=api-key&deleted=1&include="

		server := mockResponseServer(t, seasonResponse, 200, url)

		client := newTestHTTPClient(server)

		_, details, _ := client.SeasonByID(context.Background(), 55, []string{})

		assertResponseDetails(t, details, "Season")
	})
}

func assertSeason(t *testing.T, season *Season) {
	assert.Equal(t, 2, season.ID)
	assert.Equal(t, 1, season.SportID)
	assert.Equal(t, 8, season.LeagueID)
	assert.Equal(t, 1526, season.TieBreakerRuleID)
	assert.Equal(t, "2010/2011", season.Name)
	assert.Equal(t, true, season.Finished)
	assert.Equal(t, false, season.Pending)
	assert.Equal(t, false, season.IsCurrent)
	assert.Equal(t, "2010-08-14", season.StartingAt)
	assert.Equal(t, "2011-05-22", season.EndingAt)
	assert.Equal(t, "2023-05-24 08:28:07", season.StandingsRecalculatedAt)
	assert.Equal(t, false, season.GamesInCurrentWeek)
}
