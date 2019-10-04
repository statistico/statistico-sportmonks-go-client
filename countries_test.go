package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var countriesResponse = `{
	"data": [
		{
			"id": 2,
			"name": "Poland",
			"extra": {
			  "continent": "Europe",
			  "sub_region": "Eastern Europe",
			  "world_region": "EMEA",
			  "fifa": "POL",
			  "iso": "POL",
			  "iso2": "PL",
			  "longitude": "19.37775993347168",
			  "latitude": "52.147850036621094",
			  "flag": "http:\/\/www.w3.org\/2000\/svg"
			}
		}
	]
}`

var countriesIncludesResponse = `{
	"data": [
		{
			"id": 11,
			"name": "Germany",
			"extra": {
			"continent": "Europe",
			"sub_region": "Western Europe",
			"world_region": "EMEA",
			"fifa": "GER",
			"iso": "DEU",
			"iso2": "DE",
			"longitude": "19.37775993347168",
			"latitude": "52.147850036621094",
			"flag": "http:\/\/www.w3.org\/2000\/svg"
		},
		"continent": {
			"data": {
			  "id": 1,
			  "name": "Europe"
			}
		},
		"leagues": {
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
					  "predictions": 1,
					  "topscorer_goals": true,
					  "topscorer_assists": true,
					  "topscorer_cards": true
					}
				}
		   ]
		}
		}
	]
}`

var countryResponse = `{
	"data": {
		"id": 11,
		"name": "Germany",
		"extra": {
			"continent": "Europe",
			"sub_region": "Western Europe",
			"world_region": "EMEA",
			"fifa": "GER",
			"iso": "DEU",
			"iso2": "DE",
			"longitude": "19.37775993347168",
			"latitude": "52.147850036621094",
			"flag": "http://www.w3.org/2000/svg"
        }
	}
}
`

var countryIncludesResponse = `{
	"data": {
		"id": 11,
		"name": "Germany",
		"extra": {
			"continent": "Europe",
			"sub_region": "Western Europe",
			"world_region": "EMEA",
			"fifa": "GER",
			"iso": "DEU",
			"iso2": "DE",
			"longitude": "19.37775993347168",
			"latitude": "52.147850036621094",
			"flag": "http:\/\/www.w3.org\/2000\/svg"
        },
		"continent": {
			"data": {
			  "id": 1,
			  "name": "Europe"
			}
		},
		"leagues": {
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
					  "predictions": 1,
					  "topscorer_goals": true,
					  "topscorer_assists": true,
					  "topscorer_cards": true
					}
				}
		   ]
		}
		}
	}
}`

func TestCountries(t *testing.T) {
	t.Run("returns Country struct slice", func(t *testing.T) {
		server := mockResponseServer(countriesResponse, 200)

		client := newTestHTTPClient(server)

		countries, _, err := client.Countries(context.Background(), 1, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assert.Equal(t, 2, countries[0].ID)
		assert.Equal(t, "Poland", countries[0].Name)
		assert.Equal(t, "Europe", countries[0].Extra.Continent)
		assert.Equal(t, "Eastern Europe", countries[0].Extra.SubRegion)
		assert.Equal(t, "EMEA", countries[0].Extra.WorldRegion)
		assert.Equal(t, "POL", countries[0].Extra.FIFA)
		assert.Equal(t, "POL", countries[0].Extra.ISO)
		assert.Equal(t, "PL", countries[0].Extra.ISO2)
		assert.Equal(t, "19.37775993347168", countries[0].Extra.Longitude)
		assert.Equal(t, "52.147850036621094", countries[0].Extra.Latitude)
		assert.Equal(t, "http://www.w3.org/2000/svg", countries[0].Extra.Flag)
		assert.Nil(t, countries[0].ContinentData())
		assert.Nil(t, countries[0].LeagueData())
	})

	t.Run("returns Country struct slice with includes data", func(t *testing.T) {
		server := mockResponseServer(countriesIncludesResponse, 200)

		client := newTestHTTPClient(server)

		countries, _, err := client.Countries(context.Background(), 1, []string{"continent", "leagues"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		continent := countries[0].ContinentData()
		league := countries[0].LeagueData()[0]

		assert.Equal(t, 11, countries[0].ID)
		assert.Equal(t, "Germany", countries[0].Name)
		assert.Equal(t, "Europe", countries[0].Extra.Continent)
		assert.Equal(t, "Western Europe", countries[0].Extra.SubRegion)
		assert.Equal(t, "EMEA", countries[0].Extra.WorldRegion)
		assert.Equal(t, "GER", countries[0].Extra.FIFA)
		assert.Equal(t, "DEU", countries[0].Extra.ISO)
		assert.Equal(t, "DE", countries[0].Extra.ISO2)
		assert.Equal(t, "19.37775993347168", countries[0].Extra.Longitude)
		assert.Equal(t, "52.147850036621094", countries[0].Extra.Latitude)
		assert.Equal(t, "http://www.w3.org/2000/svg", countries[0].Extra.Flag)
		assert.Equal(t, 1, continent.ID)
		assert.Equal(t, "Europe", continent.Name)
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
		assert.Equal(t, 1, league.Coverage.Predictions)
		assert.Equal(t, true, league.Coverage.TopScorerGoals)
		assert.Equal(t, true, league.Coverage.TopScorerAssists)
		assert.Equal(t, true, league.Coverage.TopScorerCards)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		server := mockResponseServer(errorResponse, 400)

		client := newTestHTTPClient(server)

		countries, _, err := client.Countries(context.Background(), 1, []string{})

		if countries != nil {
			t.Fatalf("Test failed, expected nil, got %+v", countries)
		}

		assert.Equal(
			t,
			"Request failed with message: The requested endpoint does not exists!, code: 404",
			err.Error(),
		)
	})
}

func TestCountryById(t *testing.T) {
	t.Run("returns a single Country struct", func(t *testing.T) {
		server := mockResponseServer(countryResponse, 200)

		client := newTestHTTPClient(server)

		country, _, err := client.CountryById(context.Background(), 1, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assert.Equal(t, 11, country.ID)
		assert.Equal(t, "Germany", country.Name)
		assert.Equal(t, "Europe", country.Extra.Continent)
		assert.Equal(t, "Western Europe", country.Extra.SubRegion)
		assert.Equal(t, "EMEA", country.Extra.WorldRegion)
		assert.Equal(t, "GER", country.Extra.FIFA)
		assert.Equal(t, "DEU", country.Extra.ISO)
		assert.Equal(t, "DE", country.Extra.ISO2)
		assert.Equal(t, "19.37775993347168", country.Extra.Longitude)
		assert.Equal(t, "52.147850036621094", country.Extra.Latitude)
		assert.Equal(t, "http://www.w3.org/2000/svg", country.Extra.Flag)
		assert.Nil(t, country.ContinentData())
		assert.Nil(t, country.LeagueData())
	})

	t.Run("returns Country struct with country includes data", func(t *testing.T) {
		server := mockResponseServer(countryIncludesResponse, 200)

		client := newTestHTTPClient(server)

		country, _, err := client.CountryById(context.Background(), 11, []string{"countries, leagues"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		continent := country.ContinentData()
		league := country.LeagueData()[0]

		assert.Equal(t, 11, country.ID)
		assert.Equal(t, "Germany", country.Name)
		assert.Equal(t, "Europe", country.Extra.Continent)
		assert.Equal(t, "Western Europe", country.Extra.SubRegion)
		assert.Equal(t, "EMEA", country.Extra.WorldRegion)
		assert.Equal(t, "GER", country.Extra.FIFA)
		assert.Equal(t, "DEU", country.Extra.ISO)
		assert.Equal(t, "DE", country.Extra.ISO2)
		assert.Equal(t, "19.37775993347168", country.Extra.Longitude)
		assert.Equal(t, "52.147850036621094", country.Extra.Latitude)
		assert.Equal(t, "http://www.w3.org/2000/svg", country.Extra.Flag)
		assert.Equal(t, 1, continent.ID)
		assert.Equal(t, "Europe", continent.Name)
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
		assert.Equal(t, 1, league.Coverage.Predictions)
		assert.Equal(t, true, league.Coverage.TopScorerGoals)
		assert.Equal(t, true, league.Coverage.TopScorerAssists)
		assert.Equal(t, true, league.Coverage.TopScorerCards)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		server := mockResponseServer(errorResponse, 400)

		client := newTestHTTPClient(server)

		country, _, err := client.CountryById(context.Background(), 11, []string{"countries, leagues"})

		if country != nil {
			t.Fatalf("Test failed, expected nil, got %+v", country)
		}

		assert.Equal(
			t,
			"Request failed with message: The requested endpoint does not exists!, code: 404",
			err.Error(),
		)
	})
}
