package sportmonks

import (
	"context"
	"testing"
)

var liveScoresResponse = `{
	"data": [
		{
			"id": 11867285,
			"league_id": 8,
			"season_id": 16036,
			"stage_id": 77443862,
			"round_id": 169657,
			"group_id": null,
			"aggregate_id": null,
			"venue_id": 214,
			"referee_id": 14532,
			"localteam_id": 1,
			"visitorteam_id": 14,
			"winner_team_id": 1,
			"weather_report": {
			  "code": "rain",
			  "type": "shower rain",
			  "icon": "https:\/\/cdn.sportmonks.com\/images\/weather\/09d.png",
			  "temperature": {
				"temp": 62.96,
				"unit": "fahrenheit"
			  },
			  "temperature_celcius": {
				"temp": 17.2,
				"unit": "celcius"
			  },
			  "clouds": "75%",
			  "humidity": "82%",
			  "pressure": 1004,
			  "wind": {
				"speed": "5.82 m\/s",
				"degree": 240
			  },
			  "coordinates": {
				"lat": 51.51,
				"lon": -0.13
			  },
			  "updated_at": "2019-09-22T14:45:05.289505Z"
			},
			"commentaries": true,
			"attendance": 59936,
			"pitch": null,
			"details": null,
			"neutral_venue": false,
			"winning_odds_calculated": true,
			"formations": {
			  "localteam_formation": "4-1-4-1",
			  "visitorteam_formation": "4-2-3-1"
			},
			"scores": {
			  "localteam_score": 2,
			  "visitorteam_score": 0,
			  "localteam_pen_score": null,
			  "visitorteam_pen_score": null,
			  "ht_score": "1-0",
			  "ft_score": "2-0",
			  "et_score": null,
			  "ps_score": null
			},
			"time": {
			  "status": "FT",
			  "starting_at": {
				"date_time": "2019-09-22 13:00:00",
				"date": "2019-09-22",
				"time": "13:00:00",
				"timestamp": 1569157200,
				"timezone": "UTC"
			  },
			  "minute": 90,
			  "second": null,
			  "added_time": null,
			  "extra_minute": null,
			  "injury_time": null
			},
			"coaches": {
			  "localteam_coach_id": 523898,
			  "visitorteam_coach_id": 524307
			},
			"standings": {
			  "localteam_position": 11,
			  "visitorteam_position": 6
			},
			"assistants": {
			  "first_assistant_id": 12794,
			  "second_assistant_id": 12798,
			  "fourth_official_id": 15270
			},
			"leg": "1\/1",
			"colors": {
			  "localteam": {
				"color": "#832034",
				"kit_colors": "#C0D6FE,#C0D6FE,#832034,#832034,#999999,#888888,#832034"
			  },
			  "visitorteam": {
				"color": null,
				"kit_colors": null
			  }
			},
			"deleted": false,
			"league": {
				"data": {
					"id": 82,
					"active": true,
					"type": "domestic",
					"legacy_id": 4,
					"country_id": 11,
					"logo_path": "https:\/\/cdn.sportmonks.com\/images\/soccer\/leagues\/82.png",
					"name": "Bundesliga",
					"is_cup": false,
					"current_season_id": 16264,
					"current_round_id": 174546,
					"current_stage_id": 77444845,
					"live_standings": true,
					"coverage": {
					  "predictions": true,
					  "topscorer_goals": true,
					  "topscorer_assists": true,
					  "topscorer_cards": true
					}
				}
			},
			"stage": {
				"data": {
					"id": 10,
					"name": "Regular Season",
					"type": null,
					"league_id": 8,
					"season_id": 12,
					"sort_order": null,
					"has_standings": null
				}
			},
			"goals": {
				"data": [
					{
						"id": 11867297001,
						"team_id": "78",
						"type": "goal",
						"fixture_id": 11867297,
						"player_id": 95776,
						"player_name": "N. Maupay",
						"player_assist_id": null,
						"player_assist_name": null,
						"minute": 3,
						"extra_minute": null,
						"reason": null,
						"result": "1-0"
					}
				]
			}
	    }
	]
}`

func TestLivesScores(t *testing.T) {
	t.Run("returns slice of Fixture struct", func(t *testing.T) {
		url := defaultBaseURL + "/livescores?api_token=api-key&include=&page=1"

		server := mockResponseServer(t, liveScoresResponse, 200, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.LiveScores(context.Background(), 1, []string{}, map[string][]int{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertFixture(t, &fixtures[0])
	})

	t.Run("returns slice of Fixture struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/livescores?api_token=api-key&include=league%2Cstage%2Cgoals&page=1"

		server := mockResponseServer(t, liveScoresResponse, 200, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.LiveScores(
			context.Background(),
			1,
			[]string{"league", "stage", "goals"},
			map[string][]int{},
		)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertFixture(t, &fixtures[0])
		assertLeague(t, fixtures[0].League())
		assertStage(t, fixtures[0].Stage())
		assertGoalEvent(t, &fixtures[0].Goals()[0])
	})

	t.Run("returns slice of Fixture struct with includes data and filter parameters", func(t *testing.T) {
		url := defaultBaseURL + "/livescores?api_token=api-key&include=league%2Cstage%2Cgoals&leagues=8%2C10&page=1"

		server := mockResponseServer(t, liveScoresResponse, 200, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.LiveScores(
			context.Background(),
			1,
			[]string{"league", "stage", "goals"},
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
		assertLeague(t, fixtures[0].League())
		assertStage(t, fixtures[0].Stage())
		assertGoalEvent(t, &fixtures[0].Goals()[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/livescores?api_token=api-key&include=&page=1"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.LiveScores(context.Background(), 1, []string{}, map[string][]int{})

		if fixtures != nil {
			t.Fatalf("Test failed, expected nil, got %+v", fixtures)
		}

		assertError(t, err)
	})
}

func TestLivesScoresInPlay(t *testing.T) {
	t.Run("returns slice of Fixture struct", func(t *testing.T) {
		url := defaultBaseURL + "/livescores/now?api_token=api-key&include="

		server := mockResponseServer(t, liveScoresResponse, 200, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.LiveScoresInPlay(context.Background(), []string{}, map[string][]int{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertFixture(t, &fixtures[0])
	})

	t.Run("returns slice of Fixture struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/livescores/now?api_token=api-key&include=league%2Cstage%2Cgoals"

		server := mockResponseServer(t, liveScoresResponse, 200, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.LiveScoresInPlay(
			context.Background(),
			[]string{"league", "stage", "goals"},
			map[string][]int{},
		)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertFixture(t, &fixtures[0])
		assertLeague(t, fixtures[0].League())
		assertStage(t, fixtures[0].Stage())
		assertGoalEvent(t, &fixtures[0].Goals()[0])
	})

	t.Run("returns slice of Fixture struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/livescores/now?api_token=api-key&include=league%2Cstage%2Cgoals&leagues=8%2C10"

		server := mockResponseServer(t, liveScoresResponse, 200, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.LiveScoresInPlay(
			context.Background(),
			[]string{"league", "stage", "goals"},
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
		assertLeague(t, fixtures[0].League())
		assertStage(t, fixtures[0].Stage())
		assertGoalEvent(t, &fixtures[0].Goals()[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/livescores/now?api_token=api-key&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.LiveScoresInPlay(context.Background(), []string{}, map[string][]int{})

		if fixtures != nil {
			t.Fatalf("Test failed, expected nil, got %+v", fixtures)
		}

		assertError(t, err)
	})
}
