package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var standingsResponse = `{
	"data": [
		{
			"name": "Regular Season",
			"league_id": 8,
			"season_id": 12962,
			"round_id": 147706,
			"round_name": 1,
			"type": "Group Stage",
			"stage_id": 7456626,
			"stage_name": "Regular Season",
			"resource": "stage",
     		"standings": {
				"data": [
					{
						"position": 1,
						"team_id": 9,
						"team_name": "Manchester City",
						"round_id": 147706,
						"round_name": 1,
						"group_id": null,
						"group_name": null,
						"overall": {
						  "games_played": 38,
						  "won": 32,
						  "draw": 2,
						  "lost": 4,
						  "goals_scored": 95,
						  "goals_against": 23
						},
						"home": {
						  "games_played": 19,
						  "won": 18,
						  "draw": 0,
						  "lost": 1,
						  "goals_scored": 57,
						  "goals_against": 12
						},
						"away": {
						  "games_played": 19,
						  "won": 14,
						  "draw": 2,
						  "lost": 3,
						  "goals_scored": 38,
						  "goals_against": 11
						},
						"total": {
						  "goal_difference": 72,
						  "points": 98
						},
						"result": "UEFA Champions League",
						"points": 98,
						"recent_form": "WWWWW",
						"status": "same"
					  }
				]
			}
		}
	]
}`

var standingsIncludesResponse = `{
	"data": [
		{
			"name": "Regular Season",
			"league_id": 8,
			"season_id": 12962,
			"round_id": 147706,
			"round_name": 1,
			"type": "Group Stage",
			"stage_id": 7456626,
			"stage_name": "Regular Season",
			"resource": "stage",
     		"standings": {
				"data": [
					{
						"position": 1,
						"team_id": 9,
						"team_name": "Manchester City",
						"round_id": 147706,
						"round_name": 1,
						"group_id": null,
						"group_name": null,
						"overall": {
						  "games_played": 38,
						  "won": 32,
						  "draw": 2,
						  "lost": 4,
						  "goals_scored": 95,
						  "goals_against": 23
						},
						"home": {
						  "games_played": 19,
						  "won": 18,
						  "draw": 0,
						  "lost": 1,
						  "goals_scored": 57,
						  "goals_against": 12
						},
						"away": {
						  "games_played": 19,
						  "won": 14,
						  "draw": 2,
						  "lost": 3,
						  "goals_scored": 38,
						  "goals_against": 11
						},
						"total": {
						  "goal_difference": 72,
						  "points": 98
						},
						"result": "UEFA Champions League",
						"points": 98,
						"recent_form": "WWWWW",
						"status": "same",
						"round": {
							"data": {
								"id": 100,
								"name": 5,
								"league_id": 8,
								"season_id": 9,
								"stage_id": 7,
								"start": "2011-09-17",
								"end": "2011-09-18"
							}
						},
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
	]
}`

var liveStandingsResponse = `{
	"data": [
		{
			"position": 1,
			"played": 17,
			"team_id": 1,
			"team_name": "København",
			"short_code": "COP",
			"team_logo": "https://cdn.sportmonks.com/images/path-to-team-1.png",
			"goals": "37:17",
			"goal_diff": 20,
			"wins": 12,
			"lost": 3,
			"draws": 2,
			"points": 38,
			"description": "Next Round",
			"fairplay_points_lose": 20
		}
	]
}`

func TestStandingsBySeasonID(t *testing.T) {
	t.Run("returns a slice of Standings struct", func(t *testing.T) {
		url := ddefaultBaseURL + "/standings/season/12962?api_token=api-key&include="

		server := mockResponseServer(t, standingsResponse, 200, url)

		client := newTestHTTPClient(server)

		standings, _, err := client.StandingsBySeasonID(context.Background(), 12962, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertStandings(t, &standings[0])
		assertLeagueStanding(t, &standings[0].LeagueStandings()[0])
	})

	t.Run("returns a slice of Standings struct with includes data", func(t *testing.T) {
		t.Run("returns a slice of Standings struct", func(t *testing.T) {
			url := ddefaultBaseURL + "/standings/season/12962?api_token=api-key&include=standings.team%2Cstandings.round"

			server := mockResponseServer(t, standingsIncludesResponse, 200, url)

			client := newTestHTTPClient(server)

			standings, _, err := client.StandingsBySeasonID(
				context.Background(),
				12962,
				[]string{"standings.team", "standings.round"},
			)

			if err != nil {
				t.Fatalf("Test failed, expected nil, got %s", err.Error())
			}

			assertStandings(t, &standings[0])
			assertLeagueStanding(t, &standings[0].LeagueStandings()[0])
			assertTeam(t, standings[0].LeagueStandings()[0].Team())
			assertRound(t, standings[0].LeagueStandings()[0].Round())
		})
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := ddefaultBaseURL + "/standings/season/12962?api_token=api-key&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		standings, _, err := client.StandingsBySeasonID(context.Background(), 12962, []string{})

		if standings != nil {
			t.Fatalf("Test failed, expected nil, got %+v", standings)
		}

		assertError(t, err)
	})
}

func TestLiveStandingsBySeasonID(t *testing.T) {
	t.Run("returns a slice of LiveStandings struct", func(t *testing.T) {
		url := ddefaultBaseURL + "/standings/season/live/12962?api_token=api-key"

		server := mockResponseServer(t, liveStandingsResponse, 200, url)

		client := newTestHTTPClient(server)

		standings, _, err := client.LiveStandingsBySeasonID(context.Background(), 12962)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertLiveStandings(t, &standings[0])
	})
}

func assertStandings(t *testing.T, standings *Standings) {
	assert.Equal(t, "Regular Season", standings.Name)
	assert.Equal(t, 8, standings.LeagueID)
	assert.Equal(t, 12962, standings.SeasonID)
	assert.Equal(t, 147706, standings.RoundID)
	assert.Equal(t, 1, standings.RoundName)
	assert.Equal(t, "Group Stage", standings.Type)
	assert.Equal(t, 7456626, standings.StageID)
	assert.Equal(t, "Regular Season", standings.StageName)
	assert.Equal(t, "stage", standings.Resource)
}

func assertLeagueStanding(t *testing.T, standing *LeagueStanding) {
	assert.Equal(t, 1, standing.Position)
	assert.Equal(t, 9, standing.TeamID)
	assert.Equal(t, "Manchester City", standing.TeamName)
	assert.Equal(t, 147706, standing.RoundID)
	assert.Equal(t, 1, standing.RoundName)
	assert.Nil(t, standing.GroupID)
	assert.Nil(t, standing.GroupName)
	assert.Equal(t, 38, standing.Overall.GamesPlayed)
	assert.Equal(t, 32, standing.Overall.Won)
	assert.Equal(t, 2, standing.Overall.Draw)
	assert.Equal(t, 4, standing.Overall.Lost)
	assert.Equal(t, 95, standing.Overall.GoalsScored)
	assert.Equal(t, 23, standing.Overall.GoalsAgainst)
	assert.Equal(t, 19, standing.Home.GamesPlayed)
	assert.Equal(t, 18, standing.Home.Won)
	assert.Equal(t, 0, standing.Home.Draw)
	assert.Equal(t, 1, standing.Home.Lost)
	assert.Equal(t, 57, standing.Home.GoalsScored)
	assert.Equal(t, 12, standing.Home.GoalsAgainst)
	assert.Equal(t, 72, standing.Total.GoalDifference)
	assert.Equal(t, 98, standing.Total.Points)
	assert.Equal(t, "UEFA Champions League", *standing.Result)
	assert.Equal(t, 98, standing.Points)
	assert.Equal(t, "WWWWW", standing.RecentForm)
	assert.Equal(t, "same", standing.Status)
}

func assertLiveStandings(t *testing.T, standings *LiveStandings) {
	assert.Equal(t, 1, standings.Position)
	assert.Equal(t, 17, standings.Played)
	assert.Equal(t, "København", standings.TeamName)
	assert.Equal(t, "COP", standings.ShortCode)
	assert.Equal(t, "https://cdn.sportmonks.com/images/path-to-team-1.png", standings.TeamLogo)
	assert.Equal(t, "37:17", standings.Goals)
	assert.Equal(t, 20, standings.GoalDiff)
	assert.Equal(t, 12, standings.Wins)
	assert.Equal(t, 3, standings.Lost)
	assert.Equal(t, 2, standings.Draws)
	assert.Equal(t, 38, standings.Points)
	assert.Equal(t, "Next Round", standings.Description)
	assert.Equal(t, 20, standings.FairPlayPointsLose)
}
