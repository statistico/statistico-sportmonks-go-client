package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var leaguesResponse = `{
	"data": [
		{
		  "id": 8,
		  "sport_id": 1,
		  "country_id": 462,
		  "name": "Premier League",
		  "active": true,
		  "short_code": "UK PL",
		  "image_path": "https://cdn.sportmonks.com/images/soccer/leagues/8/8.png",
		  "type": "league",
		  "sub_type": "domestic",
		  "last_played_at": "2024-09-22 15:30:00",
		  "category": 1,
		  "has_jerseys": false
		}
   ]
}`

var leaguesIncludesResponse = `{
	"data": [
		{
		  "id": 8,
		  "sport_id": 1,
		  "country_id": 462,
		  "name": "Premier League",
		  "active": true,
		  "short_code": "UK PL",
		  "image_path": "https://cdn.sportmonks.com/images/soccer/leagues/8/8.png",
		  "type": "league",
		  "sub_type": "domestic",
		  "last_played_at": "2024-09-22 15:30:00",
		  "category": 1,
		  "has_jerseys": false
		}
   ]
}`

var leagueResponse = `{
	"data": {
      "id": 8,
      "sport_id": 1,
      "country_id": 462,
      "name": "Premier League",
      "active": true,
      "short_code": "UK PL",
      "image_path": "https://cdn.sportmonks.com/images/soccer/leagues/8/8.png",
      "type": "league",
      "sub_type": "domestic",
      "last_played_at": "2024-09-22 15:30:00",
      "category": 1,
      "has_jerseys": false
    }
}`

var leagueIncludesResponse = `{
	"data": {
      "id": 8,
      "sport_id": 1,
      "country_id": 462,
      "name": "Premier League",
      "active": true,
      "short_code": "UK PL",
      "image_path": "https://cdn.sportmonks.com/images/soccer/leagues/8/8.png",
      "type": "league",
      "sub_type": "domestic",
      "last_played_at": "2024-09-22 15:30:00",
      "category": 1,
      "has_jerseys": false
    }
}`

func TestLeagues(t *testing.T) {
	t.Run("returns slice of League struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/leagues?api_token=api-key&include=&page=1"

		server := mockResponseServer(t, leaguesResponse, 200, url)

		client := newTestHTTPClient(server)

		leagues, _, err := client.Leagues(context.Background(), 1, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertLeague(t, &leagues[0])
	})

	t.Run("returns slice of League struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/football/leagues?api_token=api-key&include=country%3Bseason%3Bseasons&page=1"

		server := mockResponseServer(t, leaguesIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		leagues, _, err := client.Leagues(context.Background(), 1, []string{"country", "season", "seasons"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertLeague(t, &leagues[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/leagues?api_token=api-key&include=&page=1"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		leagues, _, err := client.Leagues(context.Background(), 1, []string{})

		if leagues != nil {
			t.Fatalf("Test failed, expected nil, got %+v", leagues)
		}

		assertError(t, err)
	})
}

func TestLeagueByID(t *testing.T) {
	t.Run("returns a single League struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/leagues/82?api_token=api-key&include="

		server := mockResponseServer(t, leagueResponse, 200, url)

		client := newTestHTTPClient(server)

		league, _, err := client.LeagueByID(context.Background(), 82, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertLeague(t, league)
	})

	t.Run("returns a League struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/football/leagues/82?api_token=api-key&include=country%3Bseason%3Bseasons"

		server := mockResponseServer(t, leagueIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		league, _, err := client.LeagueByID(context.Background(), 82, []string{"country", "season", "seasons"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertLeague(t, league)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/leagues/82?api_token=api-key&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		league, _, err := client.LeagueByID(context.Background(), 82, []string{})

		if league != nil {
			t.Fatalf("Test failed, expected nil, got %+v", league)
		}

		assertError(t, err)
	})
}

func assertLeague(t *testing.T, league *League) {
	assert.Equal(t, 8, league.ID)
	assert.Equal(t, 1, league.SportID)
	assert.Equal(t, 462, league.CountryID)
	assert.Equal(t, "Premier League", league.Name)
	assert.Equal(t, true, league.Active)
	assert.Equal(t, "UK PL", league.ShortCode)
	assert.Equal(t, "https://cdn.sportmonks.com/images/soccer/leagues/8/8.png", league.ImagePath)
	assert.Equal(t, "league", league.Type)
	assert.Equal(t, "domestic", league.SubType)
	assert.Equal(t, "2024-09-22 15:30:00", league.LastPlayedAt)
	assert.Equal(t, 1, league.Category)
	assert.Equal(t, false, league.HasJerseys)
}
