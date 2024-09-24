package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var leaguesResponse = `{
	"data": [
		{
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
   ]
}`

var leaguesIncludesResponse = `{
	"data": [
		{
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
			},
			"country": {
              "id": 2,
			  "continent_id": 1,
			  "name": "Poland",
			  "official_name": "Republic of Poland",
			  "fifa_name": "POL",
			  "iso2": "PL",
			  "iso3": "POL",
			  "latitude": "52.147850036621094",
			  "longitude": "19.37775993347168",
			  "borders": [
				"BLR",
				"CZE",
				"DEU",
				"LTU",
				"RUS",
				"SVK",
				"UKR"
			  ],
			  "image_path": "https://cdn.sportmonks.com/images/countries/png/short/pl.png"
			},
			"season": {
				"data": {
					"id": 16029,
				  	"name": "2019/2020",
				  	"league_id": 2,
				  	"is_current_season": true,
				  	"current_round_id": 183973,
				  	"current_stage_id": 77443828
				}
			},
			"seasons": {
				"data": [
					{
						"id": 16029,
					  	"name": "2019\/2020",
					  	"league_id": 2,
					  	"is_current_season": true,
					  	"current_round_id": 183973,
					  	"current_stage_id": 77443828
					}
				]
			}
		}
   ]
}`

var leagueResponse = `{
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
}`

var leagueIncludesResponse = `{
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
		},
		"country": {
          "id": 2,
		  "continent_id": 1,
		  "name": "Poland",
		  "official_name": "Republic of Poland",
		  "fifa_name": "POL",
		  "iso2": "PL",
		  "iso3": "POL",
		  "latitude": "52.147850036621094",
		  "longitude": "19.37775993347168",
		  "borders": [
			"BLR",
			"CZE",
			"DEU",
			"LTU",
			"RUS",
			"SVK",
			"UKR"
		  ],
		  "image_path": "https://cdn.sportmonks.com/images/countries/png/short/pl.png"
		},
		"season": {
			"data": {
				"id": 16029,
				"name": "2019/2020",
				"league_id": 2,
				"is_current_season": true,
				"current_round_id": 183973,
				"current_stage_id": 77443828
			}
		},
		"seasons": {
			"data": [
				{
					"id": 16029,
					"name": "2019\/2020",
					"league_id": 2,
					"is_current_season": true,
					"current_round_id": 183973,
					"current_stage_id": 77443828
				}
			]
		}
	}
}`

func TestLeagues(t *testing.T) {
	t.Run("returns slice of League struct", func(t *testing.T) {
		url := defaultBaseURL + "/leagues?api_token=api-key&include=&page=1"

		server := mockResponseServer(t, leaguesResponse, 200, url)

		client := newTestHTTPClient(server)

		leagues, _, err := client.Leagues(context.Background(), 1, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertLeague(t, &leagues[0])
	})

	t.Run("returns slice of League struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/leagues?api_token=api-key&include=country%2Cseason%2Cseasons&page=1"

		server := mockResponseServer(t, leaguesIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		leagues, _, err := client.Leagues(context.Background(), 1, []string{"country,season,seasons"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertLeague(t, &leagues[0])
		assertCountry(t, leagues[0].Country)
		assertSeason(t, leagues[0].SeasonData())
		assertSeason(t, &leagues[0].SeasonsData()[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/leagues?api_token=api-key&include=&page=1"

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
		url := defaultBaseURL + "/leagues/82?api_token=api-key&include="

		server := mockResponseServer(t, leagueResponse, 200, url)

		client := newTestHTTPClient(server)

		league, _, err := client.LeagueByID(context.Background(), 82, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertLeague(t, league)
	})

	t.Run("returns a League struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/leagues/82?api_token=api-key&include=country%2Cseason%2Cseasons"

		server := mockResponseServer(t, leagueIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		league, _, err := client.LeagueByID(context.Background(), 82, []string{"country,season,seasons"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertLeague(t, league)
		assertCountry(t, league.Country)
		assertSeason(t, league.SeasonData())
		assertSeason(t, &league.SeasonsData()[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/leagues/82?api_token=api-key&include="

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
	assert.Equal(t, 82, league.ID)
	assert.Equal(t, true, league.Active)
	assert.Equal(t, "domestic", league.Type)
	assert.Equal(t, 4, league.LegacyID)
	assert.Equal(t, 11, league.CountryID)
	assert.Equal(t, "https://cdn.sportmonks.com/images/soccer/leagues/82.png", league.LogoPath)
	assert.Equal(t, "Bundesliga", league.Name)
	assert.Equal(t, false, league.IsCup)
	assert.Equal(t, 16264, league.CurrentSeasonID)
	assert.Equal(t, 174546, league.CurrentRoundID)
	assert.Equal(t, 77444845, league.CurrentStageID)
	assert.Equal(t, true, league.LiveStandings)
	assert.Equal(t, true, league.Coverage.Predictions)
	assert.Equal(t, true, league.Coverage.TopScorerGoals)
	assert.Equal(t, true, league.Coverage.TopScorerAssists)
	assert.Equal(t, true, league.Coverage.TopScorerCards)
}
