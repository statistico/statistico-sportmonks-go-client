package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var topScorersResponse = `{
	"data": {
		"id": 12962,
		"name": "2018\/2019",
		"league_id": 8,
		"is_current_season": false,
		"current_round_id": null,
		"current_stage_id": null,
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
		"assistscorers": {
			"data": [
				{
					"position": 1,
					"season_id": 12962,
					"player_id": 1331,
					"team_id": 18,
					"stage_id": 7456626,
					"assists": 15,
					"type": "assists"
				}
			]
		},
		"cardscorers": {
			"data": [
				{
					"position": 1,
					"season_id": 12962,
					"player_id": 1276,
					"team_id": 25,
					"stage_id": 7456626,
					"yellowcards": 14,
					"redcards": 1,
					"type": "cards"
				}
			]
		}
	}
}`

var topScorersIncludesResponse = `{
	"data": {
		"id": 12962,
		"name": "2018\/2019",
		"league_id": 8,
		"is_current_season": false,
		"current_round_id": null,
		"current_stage_id": null,
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
					"type": "goals",
					"team": {
						"data": {
							"id": 1,
							"legacy_id": 377,
							"name": "West Ham United",
							"short_code": "WHU",
							"twitter": "@WestHamUtd",
							"country_id": 462,
							"national_team": false,
							"founded": 1895,
							"logo_path": "https:\/\/cdn.sportmonks.com\/images\/soccer\/teams\/1\/1.png",
							"venue_id": 214,
							"current_season_id": 16036
						}
					}
				}
			]
		},
		"assistscorers": {
			"data": [
				{
					"position": 1,
					"season_id": 12962,
					"player_id": 1331,
					"team_id": 18,
					"stage_id": 7456626,
					"assists": 15,
					"type": "assists"
				}
			]
		},
		"cardscorers": {
			"data": [
				{
					"position": 1,
					"season_id": 12962,
					"player_id": 1276,
					"team_id": 25,
					"stage_id": 7456626,
					"yellowcards": 14,
					"redcards": 1,
					"type": "cards",
					"player": {
						"data": {
							"player_id": 219591,
							"team_id": 1,
							"country_id": 1190,
							"position_id": 2,
							"common_name": "F. Balbuena",
							"fullname": "F. Balbuena",
							"firstname": "Fabián Cornelio",
							"lastname": "Balbuena González",
							"nationality": "Paraguay",
							"birthdate": "23\/08\/1991",
							"birthcountry": "Paraguay",
							"birthplace": "Ciudad del Este",
							"height": "188 cm",
							"weight": "82 kg",
							"image_path": "https:\/\/cdn.sportmonks.com\/images\/soccer\/players\/7\/219591.png"
						}
					}
				}
			]
		}
	}
}`

var aggregatedTopScorersResponse = `{
	"data": {
		"id": 12962,
		"name": "2018\/2019",
		"league_id": 8,
		"is_current_season": false,
		"current_round_id": null,
		"current_stage_id": null,
		"aggregatedGoalscorers": {
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
		"aggregatedAssistscorers": {
			"data": [
				{
					"position": 1,
					"season_id": 12962,
					"player_id": 1331,
					"team_id": 18,
					"stage_id": 7456626,
					"assists": 15,
					"type": "assists"
				}
			]
		},
		"aggregatedCardscorers": {
			"data": [
				{
					"position": 1,
					"season_id": 12962,
					"player_id": 1276,
					"team_id": 25,
					"stage_id": 7456626,
					"yellowcards": 14,
					"redcards": 1,
					"type": "cards"
				}
			]
		}
	}
}`

var aggregatedTopScorersIncludesResponse = `{
	"data": {
		"id": 12962,
		"name": "2018\/2019",
		"league_id": 8,
		"is_current_season": false,
		"current_round_id": null,
		"current_stage_id": null,
		"aggregatedGoalscorers": {
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
		"aggregatedAssistscorers": {
			"data": [
				{
					"position": 1,
					"season_id": 12962,
					"player_id": 1331,
					"team_id": 18,
					"stage_id": 7456626,
					"assists": 15,
					"type": "assists",
					"player": {
						"data": {
							"player_id": 219591,
							"team_id": 1,
							"country_id": 1190,
							"position_id": 2,
							"common_name": "F. Balbuena",
							"fullname": "F. Balbuena",
							"firstname": "Fabián Cornelio",
							"lastname": "Balbuena González",
							"nationality": "Paraguay",
							"birthdate": "23\/08\/1991",
							"birthcountry": "Paraguay",
							"birthplace": "Ciudad del Este",
							"height": "188 cm",
							"weight": "82 kg",
							"image_path": "https:\/\/cdn.sportmonks.com\/images\/soccer\/players\/7\/219591.png"
						}
					}
				}
			]
		},
		"aggregatedCardscorers": {
			"data": [
				{
					"position": 1,
					"season_id": 12962,
					"player_id": 1276,
					"team_id": 25,
					"stage_id": 7456626,
					"yellowcards": 14,
					"redcards": 1,
					"type": "cards",
					"team": {
						"data": {
							"id": 1,
							"legacy_id": 377,
							"name": "West Ham United",
							"short_code": "WHU",
							"twitter": "@WestHamUtd",
							"country_id": 462,
							"national_team": false,
							"founded": 1895,
							"logo_path": "https:\/\/cdn.sportmonks.com\/images\/soccer\/teams\/1\/1.png",
							"venue_id": 214,
							"current_season_id": 16036
						}
					}
				}
			]
		}
	}
}`

func TestTopScorersBySeasonID(t *testing.T) {
	t.Run("returns a TopScorers struct", func(t *testing.T) {
		url := ddefaultBaseURL + "/topscorers/season/12962?api_token=api-key&include="

		server := mockResponseServer(t, topScorersResponse, 200, url)

		client := newTestHTTPClient(server)

		topscorers, _, err := client.TopScorersBySeasonID(context.Background(), 12962, []string{}, map[string][]int{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		goalscorers := topscorers.GoalScorers()
		assists := topscorers.AssistScorers()
		cards := topscorers.CardScorers()

		assertTopScorers(t, topscorers)
		assertGoalScorer(t, &goalscorers[0])
		assertAssistScorer(t, &assists[0])
		assertCardScorer(t, &cards[0])
	})

	t.Run("returns a TopScorers struct with includes data", func(t *testing.T) {
		url := ddefaultBaseURL + "/topscorers/season/12962?api_token=api-key&include=goalscorers.team%2Ccardscorers.player"

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

		goalscorers := topscorers.GoalScorers()
		assists := topscorers.AssistScorers()
		cards := topscorers.CardScorers()

		assertTopScorers(t, topscorers)
		assertGoalScorer(t, &goalscorers[0])
		assertTeam(t, goalscorers[0].Team())
		assertAssistScorer(t, &assists[0])
		assertCardScorer(t, &cards[0])
		assertPlayer(t, cards[0].Player())
	})

	t.Run("returns a TopScorers struct with includes data and filter parameters", func(t *testing.T) {
		url := ddefaultBaseURL + "/topscorers/season/12962?api_token=api-key&include=goalscorers.team%2Ccardscorers.player&stage_ids=4%2C33"

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

		goalscorers := topscorers.GoalScorers()
		assists := topscorers.AssistScorers()
		cards := topscorers.CardScorers()

		assertTopScorers(t, topscorers)
		assertGoalScorer(t, &goalscorers[0])
		assertTeam(t, goalscorers[0].Team())
		assertAssistScorer(t, &assists[0])
		assertCardScorer(t, &cards[0])
		assertPlayer(t, cards[0].Player())
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := ddefaultBaseURL + "/topscorers/season/12962?api_token=api-key&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		topscorers, _, err := client.TopScorersBySeasonID(context.Background(), 12962, []string{}, map[string][]int{})

		if topscorers != nil {
			t.Fatalf("Test failed, expected nil, got %+v", topscorers)
		}

		assertError(t, err)
	})
}

func TestAggregatedTopScorersBySeasonID(t *testing.T) {
	t.Run("returns an AggregatedTopScorers struct", func(t *testing.T) {
		url := ddefaultBaseURL + "/topscorers/season/12962/aggregated?api_token=api-key&include="

		server := mockResponseServer(t, aggregatedTopScorersResponse, 200, url)

		client := newTestHTTPClient(server)

		topscorers, _, err := client.AggregatedTopScorersBySeasonID(context.Background(), 12962, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		goalscorers := topscorers.GoalScorers()
		assists := topscorers.AssistScorers()
		cards := topscorers.CardScorers()

		assertAggregatedTopScorer(t, topscorers)
		assertGoalScorer(t, &goalscorers[0])
		assertAssistScorer(t, &assists[0])
		assertCardScorer(t, &cards[0])
	})

	t.Run("returns a AggregatedTopScorers struct with includes data", func(t *testing.T) {
		url := ddefaultBaseURL + "/topscorers/season/12962/aggregated?api_token=api-key&include=assistscorers.player%2Ccardscorers.team"

		server := mockResponseServer(t, aggregatedTopScorersIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		topscorers, _, err := client.AggregatedTopScorersBySeasonID(
			context.Background(),
			12962,
			[]string{"assistscorers.player", "cardscorers.team"},
		)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		goalscorers := topscorers.GoalScorers()
		assists := topscorers.AssistScorers()
		cards := topscorers.CardScorers()

		assertAggregatedTopScorer(t, topscorers)
		assertGoalScorer(t, &goalscorers[0])
		assertTeam(t, cards[0].Team())
		assertAssistScorer(t, &assists[0])
		assertCardScorer(t, &cards[0])
		assertPlayer(t, assists[0].Player())
	})
}

func assertTopScorers(t *testing.T, scorer *TopScorers) {
	assert.Equal(t, 12962, scorer.ID)
	assert.Equal(t, "2018/2019", scorer.Name)
	assert.Equal(t, 8, scorer.LeagueID)
	assert.Equal(t, false, scorer.IsCurrentSeason)
	assert.Nil(t, scorer.CurrentRoundID)
	assert.Nil(t, scorer.CurrentStageID)
}

func assertAggregatedTopScorer(t *testing.T, scorer *AggregatedTopScorers) {
	assert.Equal(t, 12962, scorer.ID)
	assert.Equal(t, "2018/2019", scorer.Name)
	assert.Equal(t, 8, scorer.LeagueID)
	assert.Equal(t, false, scorer.IsCurrentSeason)
	assert.Nil(t, scorer.CurrentRoundID)
	assert.Nil(t, scorer.CurrentStageID)
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
