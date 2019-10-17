package sportmonks

import (
	"context"
	"testing"
)

var teamSquadsResponse = `{
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
}`

var teamSquadsIncludesResponse = `{
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
			},
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
}`

func TestTeamSquad(t *testing.T) {
	t.Run("returns a slice of SquadPlayer struct", func(t *testing.T) {
		url := defaultBaseURL + "/squad/season/12962/team/1?api_token=api-key&include="

		server := mockResponseServer(t, teamSquadsResponse, 200, url)

		client := newTestHTTPClient(server)

		squad, _, err := client.TeamSquad(context.Background(), 12962, 1, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertSquadPlayer(t, &squad[0])
	})

	t.Run("returns a slice of SquadPlayer struct", func(t *testing.T) {
		url := defaultBaseURL + "/squad/season/12962/team/1?api_token=api-key&include=player"

		server := mockResponseServer(t, teamSquadsIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		squad, _, err := client.TeamSquad(context.Background(), 12962, 1, []string{"player"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertSquadPlayer(t, &squad[0])
		assertPlayer(t, squad[0].Player())
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/squad/season/12962/team/1?api_token=api-key&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		squad, _, err := client.TeamSquad(context.Background(), 12962, 1, []string{})

		if squad != nil {
			t.Fatalf("Test failed, expected nil, got %+v", squad)
		}

		assertError(t, err)
	})
}
