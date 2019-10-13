package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var continentsResponse = `{
	"data": [
		{
		  "id": 1,
		  "name": "Europe"
		}
	]
}`

var continentsIncludesResponse = `{
	"data": [
		{
		  "id": 1,
		  "name": "Europe",
          "countries": {
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
          }
		}
	]
}
`

var continentResponse = `{
	"data": {
		"id": 1,
		"name": "Europe"
	}
}
`

var continentIncludesResponse = `{
	"data": {
		"id": 1,
		"name": "Europe",
        "countries": {
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
          }
		}
	}
}`

func TestContinents(t *testing.T) {
	t.Run("returns Continent struct slice", func(t *testing.T) {
		url := defaultBaseUrl + "/continents?api_token=api-key&include=&page=1"

		server := mockResponseServer(t, continentsResponse, 200, url)

		client := newTestHTTPClient(server)

		continents, _, err := client.Continents(context.Background(), 1, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertContinent(t, &continents[0])
	})

	t.Run("returns Continent struct slice with country includes data", func(t *testing.T) {
		url := defaultBaseUrl + "/continents?api_token=api-key&include=countries&page=1"

		server := mockResponseServer(t, continentsIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		continents, _, err := client.Continents(context.Background(), 1, []string{"countries"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		country := continents[0].Countries()[0]

		assertContinent(t, &continents[0])
		assertCountry(t, &country)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseUrl + "/continents?api_token=api-key&include=&page=1"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		continents, _, err := client.Continents(context.Background(), 1, []string{})

		if continents != nil {
			t.Fatalf("Test failed, expected nil, got %+v", continents)
		}

		assertError(t, err)
	})
}

func TestContinentById(t *testing.T) {
	t.Run("returns a single Continent struct", func(t *testing.T) {
		url := defaultBaseUrl + "/continents/1?api_token=api-key&include="

		server := mockResponseServer(t, continentResponse, 200, url)

		client := newTestHTTPClient(server)

		continent, _, err := client.ContinentById(context.Background(), 1, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertContinent(t, continent)
	})

	t.Run("returns Continent struct with country includes data", func(t *testing.T) {
		url := defaultBaseUrl + "/continents/1?api_token=api-key&include=countries"

		server := mockResponseServer(t, continentIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		continent, _, err := client.ContinentById(context.Background(), 1, []string{"countries"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		country := continent.Countries()[0]

		assertContinent(t, continent)
		assertCountry(t, &country)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseUrl + "/continents/1?api_token=api-key&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		continent, _, err := client.ContinentById(context.Background(), 1, []string{})

		if continent != nil {
			t.Fatalf("Test failed, expected nil, got %+v", continent)
		}

		assertError(t, err)
	})
}

func assertContinent(t *testing.T, continent *Continent) {
	assert.Equal(t, 1, continent.ID)
	assert.Equal(t, "Europe", continent.Name)
}
