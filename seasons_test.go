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
		url := defaultBaseURL + "/seasons/55?api_token=api-key&deleted=1&include=groups%2Cgoalscorers%2Crounds"

		server := mockResponseServer(t, seasonIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		season, _, err := client.SeasonByID(context.Background(), 55, []string{"groups", "goalscorers", "rounds"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertSeason(t, season)
		assertGroup(t, &season.Groups()[0])
		assertGoalScorer(t, &season.GoalScorers()[0])
		assertRound(t, &season.Rounds()[0])
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
