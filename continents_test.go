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
		  "name": "Europe",
          "code": "EU"
		}
	]
}`

var continentsIncludesResponse = `{
	"data": [
		{
		  "id": 1,
		  "name": "Europe",
          "code": "EU",
          "countries": [
				{
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
				}
		     ]
		}
	]
}
`

var continentResponse = `{
	"data": {
		"id": 1,
		"name": "Europe",
 		"code": "EU"
	}
}
`

var continentIncludesResponse = `{
	"data": {
		"id": 1,
		"name": "Europe",
		"code": "EU",
        "countries": [
				{
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
				}
		     ]
		}
	}
}`

func TestContinents(t *testing.T) {
	t.Run("returns Continent struct slice", func(t *testing.T) {
		url := defaultBaseURL + "/core/continents?api_token=api-key&include=&page=1"

		server := mockResponseServer(t, continentsResponse, 200, url)

		client := newTestHTTPClient(server)

		continents, _, err := client.Continents(context.Background(), 1, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertContinent(t, &continents[0])
	})

	t.Run("returns Continent struct slice with country includes data", func(t *testing.T) {
		url := defaultBaseURL + "/core/continents?api_token=api-key&include=countries&page=1"

		server := mockResponseServer(t, continentsIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		continents, _, err := client.Continents(context.Background(), 1, []string{"countries"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		country := continents[0].Countries[0]

		assertContinent(t, &continents[0])
		assertCountry(t, &country)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/core/continents?api_token=api-key&include=&page=1"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		continents, _, err := client.Continents(context.Background(), 1, []string{})

		if continents != nil {
			t.Fatalf("Test failed, expected nil, got %+v", continents)
		}

		assertError(t, err)
	})
}

func TestContinentByID(t *testing.T) {
	t.Run("returns a single Continent struct", func(t *testing.T) {
		url := defaultBaseURL + "/core/continents/1?api_token=api-key&include="

		server := mockResponseServer(t, continentResponse, 200, url)

		client := newTestHTTPClient(server)

		continent, _, err := client.ContinentByID(context.Background(), 1, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertContinent(t, continent)
	})

	t.Run("returns Continent struct with country includes data", func(t *testing.T) {
		url := defaultBaseURL + "/core/continents/1?api_token=api-key&include=countries"

		server := mockResponseServer(t, continentIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		continent, _, err := client.ContinentByID(context.Background(), 1, []string{"countries"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		country := continent.Countries[0]

		assertContinent(t, continent)
		assertCountry(t, &country)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/core/continents/1?api_token=api-key&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		continent, _, err := client.ContinentByID(context.Background(), 1, []string{})

		if continent != nil {
			t.Fatalf("Test failed, expected nil, got %+v", continent)
		}

		assertError(t, err)
	})
}

func assertContinent(t *testing.T, continent *Continent) {
	assert.Equal(t, 1, continent.ID)
	assert.Equal(t, "Europe", continent.Name)
	assert.Equal(t, "EU", continent.Code)
}
