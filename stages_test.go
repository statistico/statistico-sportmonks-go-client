package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var stageResponse = `{
	"data": {
		"id": 10,
		"name": "Regular Season",
		"type": null,
		"league_id": 8,
		"season_id": 12,
		"sort_order": null,
		"has_standings": null
	}
}`

var stageIncludesResponse = `{
	"data": {
		"id": 10,
		"name": "Regular Season",
		"type": null,
		"league_id": 8,
		"season_id": 12,
		"sort_order": null,
		"has_standings": null,
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
				  "predictions": 1,
				  "topscorer_goals": true,
				  "topscorer_assists": true,
				  "topscorer_cards": true
				}
			}
		},
		"season": {
			"data": {
				"id": 16029,
				"name": "2019/2020",
				"league_id": 2,
				"is_current_season": true,
				"current_round_id": 183973,
				"current_stage_id": 77443828
			}
		},
		"fixtures": {
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
					"deleted": false
				}
			]
		},
		"results": {
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
					"deleted": false
				}
			]
		}
	}
}`

var stagesSeasonResponse = `{
	"data": [
		{
			"id": 10,
			"name": "Regular Season",
			"type": null,
			"league_id": 8,
			"season_id": 12,
			"sort_order": null,
			"has_standings": null
		}
	]
}`

var stagesSeasonIncludesResponse = `{
	"data": [
		{
			"id": 10,
			"name": "Regular Season",
			"type": null,
			"league_id": 8,
			"season_id": 12,
			"sort_order": null,
			"has_standings": null,
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
					  "predictions": 1,
					  "topscorer_goals": true,
					  "topscorer_assists": true,
					  "topscorer_cards": true
					}
				}
			},
			"season": {
				"data": {
					"id": 16029,
					"name": "2019/2020",
					"league_id": 2,
					"is_current_season": true,
					"current_round_id": 183973,
					"current_stage_id": 77443828
				}
			},
			"fixtures": {
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
						"deleted": false
					}
				]
			},
			"results": {
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
						"deleted": false
					}
				]
			}
		}
	]
}`

func TestStageById(t *testing.T) {
	t.Run("returns a Stage struct", func(t *testing.T) {
		url := defaultBaseUrl + "/stages/10?api_token=api-key&include="

		server := mockResponseServer(t, stageResponse, 200, url)

		client := newTestHTTPClient(server)

		stage, _, err := client.StageById(context.Background(), 10, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertStage(t, stage)
	})

	t.Run("return a Stage struct with includes data", func(t *testing.T) {
		url := defaultBaseUrl + "/stages/10?api_token=api-key&include=season%2Cleague%2Cfixtures%2Cresults"

		server := mockResponseServer(t, stageIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		stage, _, err := client.StageById(context.Background(), 10, []string{"season", "league", "fixtures", "results"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertStage(t, stage)
		assertLeague(t, stage.League())
		assertSeason(t, stage.Season())
		assertFixture(t, &stage.Fixtures()[0])
		assertFixture(t, &stage.Results()[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseUrl + "/stages/10?api_token=api-key&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		stage, _, err := client.StageById(context.Background(), 10, []string{})

		if stage != nil {
			t.Fatalf("Test failed, expected nil, got %+v", stage)
		}

		assertError(t, err)
	})
}

func TestStagesBySeasonId(t *testing.T) {
	t.Run("returns a slice of Stage struct", func(t *testing.T) {
		url := defaultBaseUrl + "/stages/season/10?api_token=api-key&include="

		server := mockResponseServer(t, stagesSeasonResponse, 200, url)

		client := newTestHTTPClient(server)

		stages, _, err := client.StagesBySeasonId(context.Background(), 10, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertStage(t, &stages[0])
	})

	t.Run("returns a slice of Stage struct with includes data", func(t *testing.T) {
		url := defaultBaseUrl + "/stages/season/10?api_token=api-key&include=season%2Cleague%2Cfixtures%2Cresults"

		server := mockResponseServer(t, stagesSeasonIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		stages, _, err := client.StagesBySeasonId(
			context.Background(),
			10,
			[]string{"season", "league", "fixtures", "results"},
		)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertStage(t, &stages[0])
		assertLeague(t, stages[0].League())
		assertSeason(t, stages[0].Season())
		assertFixture(t, &stages[0].Fixtures()[0])
		assertFixture(t, &stages[0].Results()[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseUrl + "/stages/season/10?api_token=api-key&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		stages, _, err := client.StagesBySeasonId(context.Background(), 10, []string{})

		if stages != nil {
			t.Fatalf("Test failed, expected nil, got %+v", stages)
		}

		assertError(t, err)
	})
}

func assertStage(t *testing.T, stage *Stage) {
	assert.Equal(t, 10, stage.ID)
	assert.Equal(t, "Regular Season", stage.Name)
	assert.Nil(t, stage.Type)
	assert.Equal(t, 8, stage.LeagueID)
	assert.Equal(t, 12, stage.SeasonID)
	assert.Nil(t, stage.SortOrder)
	assert.Nil(t, stage.HasStandings)
}
