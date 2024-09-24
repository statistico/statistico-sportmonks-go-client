package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var teamResponse = `{
	"data": {
      "id": 1,
      "sport_id": 1,
      "country_id": 462,
      "venue_id": 214,
      "gender": "male",
      "name": "West Ham United",
      "short_code": "WHU",
      "image_path": "https://cdn.sportmonks.com/images/soccer/teams/1/1.png",
      "founded": 1895,
      "type": "domestic",
      "placeholder": false,
      "last_played_at": "2024-09-21 11:30:00"
    }
}`

var teamIncludesResponse = `{
	"data": {
		  "id": 1,
		  "sport_id": 1,
		  "country_id": 462,
		  "venue_id": 214,
		  "gender": "male",
		  "name": "West Ham United",
		  "short_code": "WHU",
		  "image_path": "https://cdn.sportmonks.com/images/soccer/teams/1/1.png",
		  "founded": 1895,
		  "type": "domestic",
		  "placeholder": false,
		  "last_played_at": "2024-09-21 11:30:00",
		  "venue":  {
			"id": 214,
			"country_id": 462,
			"city_id": 51663,
			"name": "London Stadium",
			"address": "Marshgate Lane, Stratford",
			"zipcode": null,
			"latitude": "51.538611",
			"longitude": "-0.016389",
			"capacity": 60000,
			"image_path": "https://cdn.sportmonks.com/images/soccer/venues/22/214.png",
			"city_name": "London",
			"surface": "grass",
			"national_team": false
		  }
	}
}`

var teamsResponse = `{
	"data": [
		{
		  "id": 1,
		  "sport_id": 1,
		  "country_id": 462,
		  "venue_id": 214,
		  "gender": "male",
		  "name": "West Ham United",
		  "short_code": "WHU",
		  "image_path": "https://cdn.sportmonks.com/images/soccer/teams/1/1.png",
		  "founded": 1895,
		  "type": "domestic",
		  "placeholder": false,
		  "last_played_at": "2024-09-21 11:30:00"
		}
	]
}`

var teamsIncludesResponse = `{
	"data": [
		{
		  "id": 1,
		  "sport_id": 1,
		  "country_id": 462,
		  "venue_id": 214,
		  "gender": "male",
		  "name": "West Ham United",
		  "short_code": "WHU",
		  "image_path": "https://cdn.sportmonks.com/images/soccer/teams/1/1.png",
		  "founded": 1895,
		  "type": "domestic",
		  "placeholder": false,
		  "last_played_at": "2024-09-21 11:30:00",
		  "venue":  {
			"id": 214,
			"country_id": 462,
			"city_id": 51663,
			"name": "London Stadium",
			"address": "Marshgate Lane, Stratford",
			"zipcode": null,
			"latitude": "51.538611",
			"longitude": "-0.016389",
			"capacity": 60000,
			"image_path": "https://cdn.sportmonks.com/images/soccer/venues/22/214.png",
			"city_name": "London",
			"surface": "grass",
			"national_team": false
		  }
		}
	]
}`

func TestTeamByID(t *testing.T) {
	t.Run("returns a single Team struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/teams/1?api_token=api-key&include="

		server := mockResponseServer(t, teamResponse, 200, url)

		client := newTestHTTPClient(server)

		team, _, err := client.TeamByID(context.Background(), 1, []string{}, map[string][]int{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertTeam(t, team)
	})

	t.Run("returns a single Team struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/football/teams/1?api_token=api-key&include=squad%3Bleague"

		server := mockResponseServer(t, teamIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		team, _, err := client.TeamByID(context.Background(), 1, []string{"squad", "league"}, map[string][]int{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertTeam(t, team)
	})

	t.Run("returns a single Team struct with includes data and filter parameters", func(t *testing.T) {
		url := defaultBaseURL + "/football/teams/1?api_token=api-key&include=squad%3Bleague&seasons=4%2C56"

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
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/teams/1?api_token=api-key&include="

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
		url := defaultBaseURL + "/football/teams/seasons/12962?api_token=api-key&include=&page=1"

		server := mockResponseServer(t, teamsResponse, 200, url)

		client := newTestHTTPClient(server)

		teams, _, err := client.TeamsBySeasonID(context.Background(), 12962, 1, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertTeam(t, &teams[0])
	})

	t.Run("returns a slice of Team struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/football/teams/seasons/12962?api_token=api-key&include=squad%3Bleague&page=1"

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
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/teams/seasons/12962?api_token=api-key&include=&page=1"

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
	assert.Equal(t, "West Ham United", team.Name)
	assert.Equal(t, "WHU", team.ShortCode)
	assert.Equal(t, 462, team.CountryID)
	assert.Equal(t, false, team.Placeholder)
	assert.Equal(t, 1895, team.Founded)
	assert.Equal(t, "https://cdn.sportmonks.com/images/soccer/teams/1/1.png", team.ImagePath)
	assert.Equal(t, 214, team.VenueID)
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
