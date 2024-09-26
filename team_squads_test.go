package sportmonks

import (
	"context"
	"testing"
)

var teamSquadsResponse = `{
	"data": [
		{
		  "id": 1842,
		  "transfer_id": 38020,
		  "player_id": 1592,
		  "team_id": 1,
		  "position_id": 27,
		  "detailed_position_id": 156,
		  "start": "2020-01-31",
		  "end": "2030-06-30",
		  "captain": true,
		  "jersey_number": 20
		}
	]
}`

var teamSquadsIncludesResponse = `{
	"data": [
		{
		  "id": 1842,
		  "transfer_id": 38020,
		  "player_id": 1592,
		  "team_id": 1,
		  "position_id": 27,
		  "detailed_position_id": 156,
		  "start": "2020-01-31",
		  "end": "2030-06-30",
		  "captain": true,
		  "jersey_number": 20,
          "position": {
				"id": 27,
				"name": "Attacker",
				"code": "attacker",
				"developer_name": "ATTACKER",
				"model_type": "position",
				"stat_group": null
			},
		  "detailedposition": {
				"id": 156,
				"name": "Right Wing",
				"code": "right-wing",
				"developer_name": "RIGHT_WING",
				"model_type": "position",
				"stat_group": null
		  }
		}
	]
}`

func TestTeamSquad(t *testing.T) {
	t.Run("returns a slice of SquadPlayer struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/squads/seasons/12962/teams/1?api_token=api-key&include="

		server := mockResponseServer(t, teamSquadsResponse, 200, url)

		client := newTestHTTPClient(server)

		squad, _, err := client.TeamSquad(context.Background(), 12962, 1, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertSquadPlayer(t, &squad[0])
	})

	t.Run("returns a slice of SquadPlayer struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/squads/seasons/12962/teams/1?api_token=api-key&include=player"

		server := mockResponseServer(t, teamSquadsIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		squad, _, err := client.TeamSquad(context.Background(), 12962, 1, []string{"player"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertSquadPlayer(t, &squad[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/squads/seasons/12962/teams/1?api_token=api-key&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		squad, _, err := client.TeamSquad(context.Background(), 12962, 1, []string{})

		if squad != nil {
			t.Fatalf("Test failed, expected nil, got %+v", squad)
		}

		assertError(t, err)
	})
}
