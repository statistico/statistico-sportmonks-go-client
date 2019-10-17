package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var countriesResponse = `{
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
		url := defaultBaseURL + "/countries?api_token=api-key&include=&page=1"

		server := mockResponseServer(t, countriesResponse, 200, url)

		client := newTestHTTPClient(server)

		countries, _, err := client.Countries(context.Background(), 1, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertCountry(t, &countries[0])
		assert.Nil(t, countries[0].Continent())
		assert.Nil(t, countries[0].Leagues())
	})

	t.Run("returns Country struct slice with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/countries?api_token=api-key&include=continent%2Cleagues&page=1"

		server := mockResponseServer(t, countriesIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		countries, _, err := client.Countries(context.Background(), 1, []string{"continent", "leagues"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		continent := countries[0].Continent()
		league := countries[0].Leagues()[0]

		assertCountry(t, &countries[0])
		assertContinent(t, continent)
		assertLeague(t, &league)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/countries?api_token=api-key&include=&page=1"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		countries, _, err := client.Countries(context.Background(), 1, []string{})

		if countries != nil {
			t.Fatalf("Test failed, expected nil, got %+v", countries)
		}

		assertError(t, err)
	})
}

func TestCountryByID(t *testing.T) {
	t.Run("returns a single Country struct", func(t *testing.T) {
		url := defaultBaseURL + "/countries/1?api_token=api-key&include="

		server := mockResponseServer(t, countryResponse, 200, url)

		client := newTestHTTPClient(server)

		country, _, err := client.CountryByID(context.Background(), 1, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertCountry(t, country)
		assert.Nil(t, country.Continent())
		assert.Nil(t, country.Leagues())
	})

	t.Run("returns Country struct with country includes data", func(t *testing.T) {
		url := defaultBaseURL + "/countries/11?api_token=api-key&include=countries%2C+leagues"

		server := mockResponseServer(t, countryIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		country, _, err := client.CountryByID(context.Background(), 11, []string{"countries, leagues"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		continent := country.Continent()
		league := country.Leagues()[0]

		assertCountry(t, country)
		assertContinent(t, continent)
		assertLeague(t, &league)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/countries/11?api_token=api-key&include=countries%2C+leagues"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		country, _, err := client.CountryByID(context.Background(), 11, []string{"countries, leagues"})

		if country != nil {
			t.Fatalf("Test failed, expected nil, got %+v", country)
		}

		assertError(t, err)
	})
}

func assertCountry(t *testing.T, country *Country) {
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
}
