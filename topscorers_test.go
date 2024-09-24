package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var topScorersResponse = `{
	"data": [
		{
		  "id": 1540266,
		  "season_id": 19734,
		  "player_id": 159700,
		  "type_id": 83,
		  "position": 1,
		  "total": 2,
		  "participant_id": 14
		}
	]
}`

var topScorersIncludesResponse = `{
	"data": [
		{
		  "id": 1540266,
		  "season_id": 19734,
		  "player_id": 159700,
		  "type_id": 83,
		  "position": 1,
		  "total": 2,
		  "participant_id": 14
		}
	]
}`

func TestTopScorersBySeasonID(t *testing.T) {
	t.Run("returns a TopScorers struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/topscorers/seasons/12962?api_token=api-key&include="

		server := mockResponseServer(t, topScorersResponse, 200, url)

		client := newTestHTTPClient(server)

		topscorers, _, err := client.TopScorersBySeasonID(context.Background(), 12962, []string{}, map[string][]int{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertTopScorer(t, topscorers[0])
	})

	t.Run("returns a TopScorers struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/football/topscorers/seasons/12962?api_token=api-key&include=goalscorers.team%3Bcardscorers.player"

		server := mockResponseServer(t, topScorersIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		topscorers, _, err := client.TopScorersBySeasonID(
			context.Background(),
			12962,
			[]string{"goalscorers.team", "cardscorers.player"},
			map[string][]int{},
		)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertTopScorer(t, topscorers[0])
	})

	t.Run("returns a TopScorers struct with includes data and filter parameters", func(t *testing.T) {
		url := defaultBaseURL + "/football/topscorers/seasons/12962?api_token=api-key&include=goalscorers.team%3Bcardscorers.player&stage_ids=4%2C33"

		server := mockResponseServer(t, topScorersIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		topscorers, _, err := client.TopScorersBySeasonID(
			context.Background(),
			12962,
			[]string{"goalscorers.team", "cardscorers.player"},
			map[string][]int{
				"stage_ids": {
					4,
					33,
				},
			},
		)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertTopScorer(t, topscorers[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/topscorers/seasons/12962?api_token=api-key&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		topscorers, _, err := client.TopScorersBySeasonID(context.Background(), 12962, []string{}, map[string][]int{})

		if topscorers != nil {
			t.Fatalf("Test failed, expected nil, got %+v", topscorers)
		}

		assertError(t, err)
	})
}

func assertTopScorer(t *testing.T, scorer TopScorer) {
	assert.Equal(t, 1540266, scorer.ID)
	assert.Equal(t, 19734, scorer.SeasonID)  // Assuming SeasonID is meant to represent the season
	assert.Equal(t, 159700, scorer.PlayerID) // Add PlayerID assertion if needed
	assert.Equal(t, 83, scorer.TypeID)       // Assuming TypeID is relevant
	assert.Equal(t, 1, scorer.Position)
	assert.Equal(t, 2, scorer.Total)
	assert.Equal(t, 14, scorer.ParticipantID) // Add ParticipantID assertion if needed
}

func assertGoalScorer(t *testing.T, scorer *GoalScorer) {
	assert.Equal(t, 1, scorer.Position)
	assert.Equal(t, 7907, scorer.SeasonID)
	assert.Equal(t, 210173, scorer.PlayerID)
	assert.Equal(t, 579, scorer.TeamID)
	assert.Equal(t, 56494, scorer.StageID)
	assert.Equal(t, 3, scorer.Goals)
	assert.Equal(t, 0, scorer.PenaltyGoals)
	assert.Equal(t, "goals", scorer.Type)
}

func assertAssistScorer(t *testing.T, scorer *AssistScorer) {
	assert.Equal(t, 1, scorer.Position)
	assert.Equal(t, 12962, scorer.SeasonID)
	assert.Equal(t, 1331, scorer.PlayerID)
	assert.Equal(t, 18, scorer.TeamID)
	assert.Equal(t, 7456626, scorer.StageID)
	assert.Equal(t, 15, scorer.Assists)
	assert.Equal(t, "assists", scorer.Type)
}

func assertCardScorer(t *testing.T, scorer *CardScorer) {
	assert.Equal(t, 1, scorer.Position)
	assert.Equal(t, 12962, scorer.SeasonID)
	assert.Equal(t, 1276, scorer.PlayerID)
	assert.Equal(t, 25, scorer.TeamID)
	assert.Equal(t, 7456626, scorer.StageID)
	assert.Equal(t, 14, scorer.YellowCards)
	assert.Equal(t, 1, scorer.RedCards)
	assert.Equal(t, "cards", scorer.Type)
}
