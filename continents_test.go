package sportmonks

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

var continentsResponse = `{
	"data": [
		{
		  "id": 1,
		  "name": "Europe"
		},
		{
		  "id": 2,
		  "name": "Asia"
		}
	]
}`

var continentsCountriesResponse = `{
	"data": [
		{
		  "id": 1,
		  "name": "Europe",
          "countries": {
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

var continentCountriesResponse = `{
	"data": {
		"id": 1,
		"name": "Europe",
        "countries": {
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
          }
		}
	}
}`

func TestContinents(t *testing.T) {
	t.Run("returns Continent struct slice", func(t *testing.T) {
		server := mockResponseServer(continentsResponse, 200)

		client := newTestHTTPClient(server)

		continents, _, err := client.Continents(context.Background(), 1, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assert.Equal(t, 1, continents[0].ID)
		assert.Equal(t, "Europe", continents[0].Name)

		assert.Equal(t, 2, continents[1].ID)
		assert.Equal(t, "Asia", continents[1].Name)
	})

	t.Run("returns Continent struct slice with country includes data", func(t *testing.T) {
		server := mockResponseServer(continentsCountriesResponse, 200)

		client := newTestHTTPClient(server)

		continents, _, err := client.Continents(context.Background(), 1, []string{"countries"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		country := continents[0].CountryData()[0]

		assert.Equal(t, 1, continents[0].ID)
		assert.Equal(t, "Europe", continents[0].Name)
		assert.Equal(t, 2, country.ID)
		assert.Equal(t, "Poland", country.Name)
		assert.Equal(t, "Europe", country.Extra.Continent)
		assert.Equal(t, "Eastern Europe", country.Extra.SubRegion)
		assert.Equal(t, "EMEA", country.Extra.WorldRegion)
		assert.Equal(t, "POL", country.Extra.FIFA)
		assert.Equal(t, "POL", country.Extra.ISO)
		assert.Equal(t, "PL", country.Extra.ISO2)
		assert.Equal(t, "19.37775993347168", country.Extra.Longitude)
		assert.Equal(t, "52.147850036621094", country.Extra.Latitude)
		assert.Equal(t, "http://www.w3.org/2000/svg", country.Extra.Flag)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		server := mockResponseServer(errorResponse, 400)

		client := newTestHTTPClient(server)

		continents, _, err := client.Continents(context.Background(), 1, []string{})

		if continents != nil {
			t.Fatalf("Test failed, expected nil, got %+v", continents)
		}

		assert.Equal(
			t,
			"Request failed with message: The requested endpoint does not exists!, code: 404",
			err.Error(),
		)
	})

	t.Run("url is parsed as expected", func(t *testing.T) {
		server := newTestClient(func(req *http.Request) *http.Response {
			assert.Equal(
				t,
				req.URL.String(),
				"https://soccer.sportmonks.com/api/v2.0/continents?api_token=api-key&include=countries&page=1",
			)

			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(""))),
			}
		})

		client := newTestHTTPClient(server)

		_, _, _ = client.Continents(context.Background(), 1, []string{"countries"})
	})
}

func TestContinentById(t *testing.T) {
	t.Run("returns a single Continent struct", func(t *testing.T) {
		server := mockResponseServer(continentResponse, 200)

		client := newTestHTTPClient(server)

		continent, _, err := client.ContinentById(context.Background(), 1, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assert.Equal(t, 1, continent.ID)
		assert.Equal(t, "Europe", continent.Name)
	})

	t.Run("returns Continent struct with country includes data", func(t *testing.T) {
		server := mockResponseServer(continentCountriesResponse, 200)

		client := newTestHTTPClient(server)

		continent, _, err := client.ContinentById(context.Background(), 1, []string{"countries"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		country := continent.CountryData()[0]

		assert.Equal(t, 1, continent.ID)
		assert.Equal(t, "Europe", continent.Name)
		assert.Equal(t, 2, country.ID)
		assert.Equal(t, "Poland", country.Name)
		assert.Equal(t, "Europe", country.Extra.Continent)
		assert.Equal(t, "Eastern Europe", country.Extra.SubRegion)
		assert.Equal(t, "EMEA", country.Extra.WorldRegion)
		assert.Equal(t, "POL", country.Extra.FIFA)
		assert.Equal(t, "POL", country.Extra.ISO)
		assert.Equal(t, "PL", country.Extra.ISO2)
		assert.Equal(t, "19.37775993347168", country.Extra.Longitude)
		assert.Equal(t, "52.147850036621094", country.Extra.Latitude)
		assert.Equal(t, "http://www.w3.org/2000/svg", country.Extra.Flag)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		server := mockResponseServer(errorResponse, 400)

		client := newTestHTTPClient(server)

		continent, _, err := client.ContinentById(context.Background(), 1, []string{})

		if continent != nil {
			t.Fatalf("Test failed, expected nil, got %+v", continent)
		}

		assert.Equal(
			t,
			"Request failed with message: The requested endpoint does not exists!, code: 404",
			err.Error(),
		)
	})

	t.Run("url is parsed as expected", func(t *testing.T) {
		server := newTestClient(func(req *http.Request) *http.Response {
			assert.Equal(
				t,
				req.URL.String(),
				"https://soccer.sportmonks.com/api/v2.0/continents/10?api_token=api-key&include=countries",
			)

			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(""))),
			}
		})

		client := newTestHTTPClient(server)

		_, _, _ = client.ContinentById(context.Background(), 10, []string{"countries"})
	})
}
