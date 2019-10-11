package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var playerResponse = `{
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
}`

var playerIncludesResponse = `{
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
		"image_path": "https:\/\/cdn.sportmonks.com\/images\/soccer\/players\/7\/219591.png",
		"stats": {
			"data": [
				{
					  "player_id": 219591,
					  "team_id": 1,
					  "league_id": 8,
					  "season_id": 12962,
					  "captain": 0,
					  "minutes": 2014,
					  "appearences": 23,
					  "lineups": 23,
					  "substitute_in": 0,
					  "substitute_out": 1,
					  "substitutes_on_bench": 4,
					  "goals": 1,
					  "assists": 0,
					  "saves": 0,
					  "inside_box_saves": 0,
					  "dispossesed": 0,
					  "interceptions": 52,
					  "yellowcards": 2,
					  "yellowred": 0,
					  "redcards": 0,
					  "type": "domestic",
					  "tackles": 19,
					  "blocks": 15,
					  "hit_post": 1,
					  "fouls": {
						"committed": 19,
						"drawn": 4
					  },
					  "crosses": {
						"total": 1,
						"accurate": 0
					  },
					  "dribbles": {
						"attempts": 1,
						"success": 1,
						"dribbled_past": 26
					  },
					  "duels": {
						"total": 218,
						"won": 112
					  },
					  "passes": {
						"total": 660,
						"accuracy": 78,
						"key_passes": 3
					  },
					  "penalties": {
						"won": 0,
						"scores": 0,
						"missed": 0,
						"committed": 1,
						"saves": 0
					  }
				}
			]
		},
		"position": {
		  "data": {
			"id": 2,
			"name": "Defender"
		  }
		},
		"trophies": {
			"data": []
		}
	}
}`

func TestPlayerById(t *testing.T) {
	t.Run("return a single Player struct", func(t *testing.T) {
		url := defaultBaseUrl + "/players/219591?api_token=api-key&include="

		server := mockResponseServer(t, playerResponse, 200, url)

		client := newTestHTTPClient(server)

		player, _, err := client.PlayerById(context.Background(), 219591, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %+v", err)
		}

		assertPlayer(t, player)
	})

	t.Run("return a single Player struct with includes data", func(t *testing.T) {
		url := defaultBaseUrl + "/players/219591?api_token=api-key&include=stats%2Cposition%2Ctrophies"

		server := mockResponseServer(t, playerIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		player, _, err := client.PlayerById(context.Background(), 219591, []string{"stats", "position", "trophies"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %+v", err)
		}

		assertPlayer(t, player)
		assertPlayerSeasonStats(t, &player.Stats()[0])
		assertPosition(t, player.Position())
		assert.Equal(t, []Trophy{}, player.Trophies())
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseUrl + "/players/219591?api_token=api-key&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		player, _, err := client.PlayerById(context.Background(), 219591, []string{})

		if player != nil {
			t.Fatalf("Test failed, expected nil, got %+v", player)
		}

		assertError(t, err)
	})
}

func assertPlayer(t *testing.T, player *Player) {
	assert.Equal(t, 219591, player.ID)
	assert.Equal(t, 1, player.TeamID)
	assert.Equal(t, 1190, player.CountryID)
	assert.Equal(t, "F. Balbuena", player.CommonName)
	assert.Equal(t, "Fabián Cornelio", player.FirstName)
	assert.Equal(t, "Balbuena González", player.LastName)
	assert.Equal(t, "Paraguay", player.Nationality)
	assert.Equal(t, "23/08/1991", player.BirthDate)
	assert.Equal(t, "Paraguay", player.BirthCountry)
	assert.Equal(t, "Ciudad del Este", player.BirthPlace)
	assert.Equal(t, "188 cm", player.Height)
	assert.Equal(t, "82 kg", player.Weight)
	assert.Equal(t, "https://cdn.sportmonks.com/images/soccer/players/7/219591.png", player.ImagePath)
}

func assertPlayerSeasonStats(t *testing.T, stats *PlayerSeasonStats) {
	assert.Equal(t, 219591, stats.PlayerID)
	assert.Equal(t, 8, stats.LeagueID)
	assert.Equal(t, 12962, stats.SeasonID)
	assert.Equal(t, 0, stats.Captain)
	assert.Equal(t, 2014, stats.Minutes)
	assert.Equal(t, 23, stats.Appearances)
	assert.Equal(t, 23, stats.Lineups)
	assert.Equal(t, 0, stats.SubstituteIn)
	assert.Equal(t, 1, stats.SubstituteOut)
	assert.Equal(t, 4, stats.SubstitutesOnBench)
	assert.Equal(t, 1, stats.Goals)
	assert.Equal(t, 0, stats.Assists)
	assert.Equal(t, 0, stats.Saves)
	assert.Equal(t, 0, stats.InsideBoxSaves)
	assert.Equal(t, 0, stats.Dispossessed)
	assert.Equal(t, 52, stats.Interceptions)
	assert.Equal(t, 2, stats.YellowCards)
	assert.Equal(t, 0, stats.YellowRed)
	assert.Equal(t, 0, stats.RedCards)
	assert.Equal(t, "domestic", stats.Type)
	assert.Equal(t, 19, *stats.Tackles)
	assert.Equal(t, 15, *stats.Blocks)
	assert.Equal(t, 1, *stats.HitPost)
	assert.Equal(t, 19, stats.Fouls.Committed)
	assert.Equal(t, 4, *stats.Fouls.Drawn)
	assert.Equal(t, 1, *stats.Crosses.Total)
	assert.Equal(t, 0, *stats.Crosses.Accurate)
	assert.Equal(t, 1, *stats.Dribbles.Attempts)
	assert.Equal(t, 1, *stats.Dribbles.Success)
	assert.Equal(t, 26, *stats.Dribbles.DribbledPast)
	assert.Equal(t, 218, *stats.Duels.Total)
	assert.Equal(t, 112, *stats.Duels.Won)
	assert.Equal(t, 660, stats.Passes.Total)
	assert.Equal(t, 78, *stats.Passes.Accuracy)
	assert.Equal(t, 3, *stats.Passes.KeyPasses)
	assert.Equal(t, 0, *stats.Penalties.Won)
	assert.Equal(t, 0, *stats.Penalties.Scores)
	assert.Equal(t, 0, *stats.Penalties.Missed)
	assert.Equal(t, 1, *stats.Penalties.Committed)
	assert.Equal(t, 0, *stats.Penalties.Saves)
}

func assertPosition(t *testing.T, position *Position) {
	assert.Equal(t, 2, position.ID)
	assert.Equal(t, "Defender", position.Name)
}
