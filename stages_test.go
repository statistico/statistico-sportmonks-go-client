package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var stageResponse = `{
	"data": {
      "id": 2,
      "sport_id": 1,
      "league_id": 8,
      "season_id": 2,
      "type_id": 223,
      "name": "Regular Season",
      "sort_order": 1,
      "finished": true,
      "is_current": false,
      "starting_at": "2010-08-14",
      "ending_at": "2011-05-22",
      "games_in_current_week": false,
      "tie_breaker_rule_id": null
    }
}`

var stageIncludesResponse = `{
	"data": {
      "id": 2,
      "sport_id": 1,
      "league_id": 8,
      "season_id": 2,
      "type_id": 223,
      "name": "Regular Season",
      "sort_order": 1,
      "finished": true,
      "is_current": false,
      "starting_at": "2010-08-14",
      "ending_at": "2011-05-22",
      "games_in_current_week": false,
      "tie_breaker_rule_id": null
    }
}`

var stagesSeasonResponse = `{
	"data": [
		{
		  "id": 2,
		  "sport_id": 1,
		  "league_id": 8,
		  "season_id": 2,
		  "type_id": 223,
		  "name": "Regular Season",
		  "sort_order": 1,
		  "finished": true,
		  "is_current": false,
		  "starting_at": "2010-08-14",
		  "ending_at": "2011-05-22",
		  "games_in_current_week": false,
		  "tie_breaker_rule_id": null
		}
	]
}`

var stagesSeasonIncludesResponse = `{
	"data": [
		{
			  "id": 2,
			  "sport_id": 1,
			  "league_id": 8,
			  "season_id": 2,
			  "type_id": 223,
			  "name": "Regular Season",
			  "sort_order": 1,
			  "finished": true,
			  "is_current": false,
			  "starting_at": "2010-08-14",
			  "ending_at": "2011-05-22",
			  "games_in_current_week": false,
			  "tie_breaker_rule_id": null
		}
	]
}`

func TestStageByID(t *testing.T) {
	t.Run("returns a Stage struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/stages/10?api_token=api-key&include="

		server := mockResponseServer(t, stageResponse, 200, url)

		client := newTestHTTPClient(server)

		stage, _, err := client.StageByID(context.Background(), 10, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertStage(t, stage)
	})

	t.Run("return a Stage struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/football/stages/10?api_token=api-key&include=season%3Bleague%3Bfixtures%3Bresults"

		server := mockResponseServer(t, stageIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		stage, _, err := client.StageByID(context.Background(), 10, []string{"season", "league", "fixtures", "results"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertStage(t, stage)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/stages/10?api_token=api-key&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		stage, _, err := client.StageByID(context.Background(), 10, []string{})

		if stage != nil {
			t.Fatalf("Test failed, expected nil, got %+v", stage)
		}

		assertError(t, err)
	})
}

func TestStagesBySeasonID(t *testing.T) {
	t.Run("returns a slice of Stage struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/stages/seasons/10?api_token=api-key&include="

		server := mockResponseServer(t, stagesSeasonResponse, 200, url)

		client := newTestHTTPClient(server)

		stages, _, err := client.StagesBySeasonID(context.Background(), 10, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertStage(t, &stages[0])
	})

	t.Run("returns a slice of Stage struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/football/stages/seasons/10?api_token=api-key&include=season%3Bleague%3Bfixtures%3Bresults"

		server := mockResponseServer(t, stagesSeasonIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		stages, _, err := client.StagesBySeasonID(
			context.Background(),
			10,
			[]string{"season", "league", "fixtures", "results"},
		)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertStage(t, &stages[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/stages/seasons/10?api_token=api-key&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		stages, _, err := client.StagesBySeasonID(context.Background(), 10, []string{})

		if stages != nil {
			t.Fatalf("Test failed, expected nil, got %+v", stages)
		}

		assertError(t, err)
	})
}

func assertStage(t *testing.T, stage *Stage) {
	assert.Equal(t, 2, stage.ID)
	assert.Equal(t, 1, stage.SportID)
	assert.Equal(t, 8, stage.LeagueID)
	assert.Equal(t, 2, stage.SeasonID)
	assert.Equal(t, 223, stage.TypeID)
	assert.Equal(t, "Regular Season", stage.Name)
	assert.Equal(t, 1, stage.SortOrder)
	assert.Equal(t, true, stage.Finished)
	assert.Equal(t, false, stage.IsCurrent)
	assert.Equal(t, "2010-08-14", stage.StartingAt)
	assert.Equal(t, "2011-05-22", stage.EndingAt)
	assert.Equal(t, false, stage.GamesInCurrentWeek)
	assert.Nil(t, stage.TieBreakerRuleID)
}
