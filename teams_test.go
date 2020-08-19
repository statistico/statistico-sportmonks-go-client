package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var teamResponse = `{
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
}`

var teamIncludesResponse = `{
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
		"current_season_id": 16036,
		"squad": {
			"data": [
				{
					  "player_id": 219591,
					  "position_id": 2,
					  "number": 4,
					  "captain": 0,
					  "injured": false,
					  "minutes": 90,
					  "appearences": 2,
					  "lineups": 1,
					  "substitute_in": 1,
					  "substitute_out": 0,
					  "substitutes_on_bench": 7,
					  "goals": 0,
					  "assists": 0,
					  "saves": null,
					  "inside_box_saves": null,
					  "dispossesed": null,
					  "interceptions": null,
					  "yellowcards": 1,
					  "yellowred": 0,
					  "redcards": 0,
					  "tackles": null,
					  "blocks": null,
					  "hit_post": null,
					  "fouls": {
						"committed": 2,
						"drawn": null
					  },
					  "crosses": {
						"total": null,
						"accurate": null
					  },
					  "dribbles": {
						"attempts": null,
						"success": null,
						"dribbled_past": null
					  },
					  "duels": {
						"total": null,
						"won": null
					  },
					  "passes": {
						"total": 30,
						"accuracy": 75,
						"key_passes": null
					  },
					  "penalties": {
						"won": null,
						"scores": null,
						"missed": null,
						"committed": null,
						"saves": null
					  }
				}
			]
		},
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
		}
	}
}`

var teamsResponse = `{
	"data": [
		{
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
	]
}`

var teamsIncludesResponse = `{
	"data": [
		{
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
			"current_season_id": 16036,
			"squad": {
				"data": [
					{
						  "player_id": 219591,
						  "position_id": 2,
						  "number": 4,
						  "captain": 0,
						  "injured": false,
						  "minutes": 90,
						  "appearences": 2,
						  "lineups": 1,
						  "substitute_in": 1,
						  "substitute_out": 0,
						  "substitutes_on_bench": 7,
						  "goals": 0,
						  "assists": 0,
						  "saves": null,
						  "inside_box_saves": null,
						  "dispossesed": null,
						  "interceptions": null,
						  "yellowcards": 1,
						  "yellowred": 0,
						  "redcards": 0,
						  "tackles": null,
						  "blocks": null,
						  "hit_post": null,
						  "fouls": {
							"committed": 2,
							"drawn": null
						  },
						  "crosses": {
							"total": null,
							"accurate": null
						  },
						  "dribbles": {
							"attempts": null,
							"success": null,
							"dribbled_past": null
						  },
						  "duels": {
							"total": null,
							"won": null
						  },
						  "passes": {
							"total": 30,
							"accuracy": 75,
							"key_passes": null
						  },
						  "penalties": {
							"won": null,
							"scores": null,
							"missed": null,
							"committed": null,
							"saves": null
						  }
					}
				]
			},
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
			}
		}
	]
}`

func TestTeamByID(t *testing.T) {
	t.Run("returns a single Team struct", func(t *testing.T) {
		url := defaultBaseURL + "/teams/1?api_token=api-key&include="

		server := mockResponseServer(t, teamResponse, 200, url)

		client := newTestHTTPClient(server)

		team, _, err := client.TeamByID(context.Background(), 1, []string{}, map[string][]int{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertTeam(t, team)
	})

	t.Run("returns a single Team struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/teams/1?api_token=api-key&include=squad%2Cleague"

		server := mockResponseServer(t, teamIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		team, _, err := client.TeamByID(context.Background(), 1, []string{"squad", "league"}, map[string][]int{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertTeam(t, team)
		assertSquadPlayer(t, &team.Squad()[0])
		assertLeague(t, team.League())
	})

	t.Run("returns a single Team struct with includes data and filter parameters", func(t *testing.T) {
		url := defaultBaseURL + "/teams/1?api_token=api-key&include=squad%2Cleague&seasons=4%2C56"

		server := mockResponseServer(t, teamIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		team, _, err := client.TeamByID(
			context.Background(),
			1,
			[]string{"squad", "league"},
			map[string][]int{
				"seasons": {
					4,
					56,
				},
			},
		)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertTeam(t, team)
		assertSquadPlayer(t, &team.Squad()[0])
		assertLeague(t, team.League())
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/teams/1?api_token=api-key&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		team, _, err := client.TeamByID(context.Background(), 1, []string{}, map[string][]int{})

		if team != nil {
			t.Fatalf("Test failed, expected nil, got %+v", team)
		}

		assertError(t, err)
	})
}

func TestTeamsBySeasonID(t *testing.T) {
	t.Run("returns a slice of Team struct", func(t *testing.T) {
		url := defaultBaseURL + "/teams/season/12962?api_token=api-key&include=&page=1"

		server := mockResponseServer(t, teamsResponse, 200, url)

		client := newTestHTTPClient(server)

		teams, _, err := client.TeamsBySeasonID(context.Background(), 12962, 1, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertTeam(t, &teams[0])
	})

	t.Run("returns a slice of Team struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/teams/season/12962?api_token=api-key&include=squad%2Cleague&page=1"

		server := mockResponseServer(t, teamsIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		teams, _, err := client.TeamsBySeasonID(
			context.Background(),
			12962,
			1,
			[]string{"squad", "league"},
		)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertTeam(t, &teams[0])
		assertSquadPlayer(t, &teams[0].Squad()[0])
		assertLeague(t, teams[0].League())
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/teams/season/12962?api_token=api-key&include=&page=1"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		teams, _, err := client.TeamsBySeasonID(context.Background(), 12962, 1, []string{})

		if teams != nil {
			t.Fatalf("Test failed, expected nil, got %+v", teams)
		}

		assertError(t, err)
	})
}

func assertTeam(t *testing.T, team *Team) {
	assert.Equal(t, 1, team.ID)
	assert.Equal(t, 377, team.LegacyID)
	assert.Equal(t, "West Ham United", team.Name)
	assert.Equal(t, "WHU", team.ShortCode)
	assert.Equal(t, "@WestHamUtd", *team.Twitter)
	assert.Equal(t, 462, team.CountryID)
	assert.Equal(t, false, team.NationalTeam)
	assert.Equal(t, 1895, team.Founded)
	assert.Equal(t, "https://cdn.sportmonks.com/images/soccer/teams/1/1.png", *team.LogoPath)
	assert.Equal(t, 214, team.VenueID)
	assert.Equal(t, 16036, team.CurrentSeasonID)
}

func assertSquadPlayer(t *testing.T, player *SquadPlayer) {
	assert.Equal(t, 219591, player.PlayerID)
	assert.Equal(t, 2, player.PositionID)
	assert.Equal(t, 4, player.Number)
	assert.Equal(t, 0, player.Captain)
	assert.Equal(t, false, player.Injured)
	assert.Equal(t, 2, player.Appearances)
	assert.Equal(t, 1, player.Lineups)
	assert.Equal(t, 1, player.SubstituteIn)
	assert.Equal(t, 0, player.SubstituteOut)
	assert.Equal(t, 7, player.SubstitutesOnBench)
	assert.Equal(t, 0, player.Goals)
	assert.Equal(t, 0, player.Assists)
	assert.Nil(t, player.Saves)
	assert.Nil(t, player.InsideBoxSaves)
	assert.Nil(t, player.Dispossessed)
	assert.Equal(t, 1, player.YellowCards)
	assert.Equal(t, 0, player.YellowRed)
	assert.Nil(t, player.Tackles)
	assert.Nil(t, player.Blocks)
	assert.Nil(t, player.HitPost)
	assert.Equal(t, 2, player.Fouls.Committed)
	assert.Nil(t, player.Fouls.Drawn)
	assert.Nil(t, player.Crosses.Total)
	assert.Nil(t, player.Crosses.Accurate)
	assert.Nil(t, player.Dribbles.Attempts)
	assert.Nil(t, player.Dribbles.Success)
	assert.Nil(t, player.Dribbles.DribbledPast)
	assert.Nil(t, player.Duels.Total)
	assert.Nil(t, player.Duels.Won)
	assert.Equal(t, 30, player.Passes.Total)
	assert.Equal(t, 75, *player.Passes.Accuracy)
	assert.Nil(t, player.Passes.KeyPasses)
	assert.Nil(t, player.Penalties.Won)
	assert.Nil(t, player.Penalties.Scores)
	assert.Nil(t, player.Penalties.Missed)
	assert.Nil(t, player.Penalties.Committed)
	assert.Nil(t, player.Penalties.Saves)
}
