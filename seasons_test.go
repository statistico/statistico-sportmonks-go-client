package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var seasonsResponse = `{
	"data": [
		{
			"id": 16029,
			"name": "2019/2020",
			"league_id": 2,
			"is_current_season": true,
			"current_round_id": 183973,
			"current_stage_id": 77443828
		}
	]
}`

var seasonsIncludesResponse = `{
	"data": [
		{
			"id": 16029,
			"name": "2019/2020",
			"league_id": 2,
			"is_current_season": true,
			"current_round_id": 183973,
			"current_stage_id": 77443828,
			"groups": {
				"data": [
					{
						"id": 226779,
						"name": "Group A",
						"league_id": 5,
						"season_id": 5330,
						"round_id": null,
						"round_name": null,
						"stage_id": 10025,
						"stage_name": "Group stage",
						"resource": "group"
					}
				]
			},
			"goalscorers": {
				"data": [
					{
						"position": 1,
						"season_id": 7907,
						"player_id": 210173,
						"team_id": 579,
						"stage_id": 56494,
						"goals": 3,
						"penalty_goals": 0,
						"type": "goals"
					}
				]
			},
			"rounds": {
				"data": [
					{
						"id": 100,
						"name": 5,
						"league_id": 8,
						"season_id": 9,
						"stage_id": 7,
						"start": "2011-09-17",
						"end": "2011-09-18"
					}
				]
			}
		}
	]
}`

var seasonResponse = `{
	"data": {
		"id": 16029,
		"name": "2019/2020",
		"league_id": 2,
		"is_current_season": true,
		"current_round_id": 183973,
		"current_stage_id": 77443828
	}
}`

var seasonIncludesResponse = `{
	"data": {
		"id": 16029,
		"name": "2019/2020",
		"league_id": 2,
		"is_current_season": true,
		"current_round_id": 183973,
		"current_stage_id": 77443828,
		"groups": {
			"data": [
				{
					"id": 226779,
					"name": "Group A",
					"league_id": 5,
					"season_id": 5330,
					"round_id": null,
					"round_name": null,
					"stage_id": 10025,
					"stage_name": "Group stage",
					"resource": "group"
				}
			]
		},
		"goalscorers": {
			"data": [
				{
					"position": 1,
					"season_id": 7907,
					"player_id": 210173,
					"team_id": 579,
					"stage_id": 56494,
					"goals": 3,
					"penalty_goals": 0,
					"type": "goals"
				}
			]
		},
		"rounds": {
			"data": [
				{
					"id": 100,
					"name": 5,
					"league_id": 8,
					"season_id": 9,
					"stage_id": 7,
					"start": "2011-09-17",
					"end": "2011-09-18"
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

func TestSeasons(t *testing.T) {
	t.Run("returns a slice of Season struct", func(t *testing.T) {
		url := defaultBaseURL + "/seasons?api_token=api-key&include=&page=1"

		server := mockResponseServer(t, seasonsResponse, 200, url)

		client := newTestHTTPClient(server)

		seasons, _, err := client.Seasons(context.Background(), 1, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertSeason(t, &seasons[0])
	})

	t.Run("returns a slice of Season struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/seasons?api_token=api-key&include=groups%2Cgoalscorers%2Crounds&page=1"

		server := mockResponseServer(t, seasonsIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		seasons, _, err := client.Seasons(context.Background(), 1, []string{"groups", "goalscorers", "rounds"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertSeason(t, &seasons[0])
		assertGroup(t, &seasons[0].Groups()[0])
		assertGoalScorer(t, &seasons[0].GoalScorers()[0])
		assertRound(t, &seasons[0].Rounds()[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/seasons?api_token=api-key&include=&page=1"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		seasons, _, err := client.Seasons(context.Background(), 1, []string{})

		if seasons != nil {
			t.Fatalf("Test failed, expected nil, got %+v", seasons)
		}

		assertError(t, err)
	})
}

func TestSeasonByID(t *testing.T) {
	t.Run("returns a single Season struct", func(t *testing.T) {
		url := defaultBaseURL + "/seasons/55?api_token=api-key&deleted=1&include="

		server := mockResponseServer(t, seasonResponse, 200, url)

		client := newTestHTTPClient(server)

		season, _, err := client.SeasonByID(context.Background(), 55, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertSeason(t, season)
	})

	t.Run("returns a single Season struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/seasons/55?api_token=api-key&deleted=1&include=groups%2Cgoalscorers%2Crounds%2Cresults"

		server := mockResponseServer(t, seasonIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		season, _, err := client.SeasonByID(context.Background(), 55, []string{"groups", "goalscorers", "rounds", "results"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertSeason(t, season)
		assertGroup(t, &season.Groups()[0])
		assertGoalScorer(t, &season.GoalScorers()[0])
		assertRound(t, &season.Rounds()[0])
		assertFixture(t, &season.Results()[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/seasons/55?api_token=api-key&deleted=1&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		season, _, err := client.SeasonByID(context.Background(), 55, []string{})

		if season != nil {
			t.Fatalf("Test failed, expected nil, got %+v", season)
		}

		assertError(t, err)
	})
}

func assertSeason(t *testing.T, season *Season) {
	assert.Equal(t, 16029, season.ID)
	assert.Equal(t, "2019/2020", season.Name)
	assert.Equal(t, 2, season.LeagueID)
	assert.Equal(t, true, season.IsCurrentSeason)
	assert.Equal(t, 183973, *season.CurrentRoundID)
	assert.Equal(t, 77443828, *season.CurrentStageID)
}

func assertGroup(t *testing.T, group *Group) {
	assert.Equal(t, 226779, group.ID)
	assert.Equal(t, "Group A", group.Name)
	assert.Equal(t, 5, group.LeagueID)
	assert.Equal(t, 5330, group.SeasonID)
	assert.Nil(t, group.RoundID)
	assert.Nil(t, group.RoundName)
	assert.Equal(t, 10025, group.StageID)
	assert.Equal(t, "Group stage", group.StageName)
	assert.Equal(t, "group", group.Resource)
}
