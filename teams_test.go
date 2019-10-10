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
				}
			]
		}
	}
}`

func TestTeamById(t *testing.T) {
	t.Run("returns a single Team struct", func(t *testing.T) {
		url := defaultBaseUrl + "/teams/1?api_token=api-key&include="

		server := mockResponseServer(t, teamResponse, 200, url)

		client := newTestHTTPClient(server)

		team, _, err := client.TeamById(context.Background(), 1, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertTeam(t, team)
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