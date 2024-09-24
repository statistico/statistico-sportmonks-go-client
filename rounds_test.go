package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var roundResponse = `{
	"data": {
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
}`

var roundIncludesResponse = `{
	"data": {
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
}`

var roundsSeasonResponse = `{
	"data": [
		{
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
	]
}`

var roundsSeasonIncludesResponse = `{
	"data": [
		{
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
	]
}`

func TestRoundByID(t *testing.T) {
	t.Run("return a single Round struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/rounds/100?api_token=api-key&include="

		server := mockResponseServer(t, roundResponse, 200, url)

		client := newTestHTTPClient(server)

		round, _, err := client.RoundByID(context.Background(), 100, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertRound(t, round)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/rounds/100?api_token=api-key&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		round, _, err := client.RoundByID(context.Background(), 100, []string{})

		if round != nil {
			t.Fatalf("Test failed, expected nil, got %+v", round)
		}

		assertError(t, err)
	})

	t.Run("return a single Round struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/football/rounds/100?api_token=api-key&include=fixtures%3Bleague%3Bresults%3Bseason"

		server := mockResponseServer(t, roundIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		round, _, err := client.RoundByID(context.Background(), 100, []string{"fixtures", "league", "results", "season"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertRound(t, round)
	})
}

func TestRoundsBySeasonID(t *testing.T) {
	t.Run("returns a slice of Round struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/rounds/seasons/16029?api_token=api-key&include="

		server := mockResponseServer(t, roundsSeasonResponse, 200, url)

		client := newTestHTTPClient(server)

		rounds, _, err := client.RoundsBySeasonID(context.Background(), 16029, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertRound(t, &rounds[0])
	})

	t.Run("returns a slice of Round struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/football/rounds/seasons/16029?api_token=api-key&include=fixtures%3Bleague%3Bresults%3Bseason"

		server := mockResponseServer(t, roundsSeasonIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		rounds, _, err := client.RoundsBySeasonID(context.Background(), 16029, []string{"fixtures", "league", "results", "season"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertRound(t, &rounds[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/rounds/seasons/16029?api_token=api-key&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		rounds, _, err := client.RoundsBySeasonID(context.Background(), 16029, []string{})

		if rounds != nil {
			t.Fatalf("Test failed, expected nil, got %+v", rounds)
		}

		assertError(t, err)
	})
}

func assertRound(t *testing.T, round *Round) {
	assert.Equal(t, 43, round.ID)
	assert.Equal(t, 1, round.SportID)
	assert.Equal(t, 8, round.LeagueID)
	assert.Equal(t, 2, round.SeasonID)
	assert.Equal(t, 2, round.StageID)
	assert.Equal(t, "1", round.Name)
	assert.Equal(t, true, round.Finished)
	assert.Equal(t, false, round.IsCurrent)
	assert.Equal(t, "2010-08-14", round.StartingAt)
	assert.Equal(t, "2010-08-16", round.EndingAt)
	assert.Equal(t, false, round.GamesInCurrentWeek)
}
