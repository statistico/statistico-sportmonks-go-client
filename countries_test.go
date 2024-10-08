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
	],
	"pagination": {
		"count": 2,
		"per_page": 25,
		"current_page": 1,
		"next_page": null,
		"has_more": false
	},
	"subscription": [
		{
			"meta": {
				"trial_ends_at": null,
				"ends_at": "2024-10-26 12:06:34",
				"current_timestamp": 1728372666
			},
			"plans": [
				{
					"plan": "Joe Sweeny Custom Plan",
					"sport": "Football",
					"category": "Advanced"
				}
			],
			"add_ons": [],
			"widgets": []
		}
	],
	"rate_limit": {
		"resets_in_seconds": 3386,
		"remaining": 2997,
		"requested_entity": "Country"
	},
	"timezone": "UTC"
}`

var countriesIncludesResponse = `{
	"data": [
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
			"image_path": "https://cdn.sportmonks.com/images/countries/png/short/pl.png",
			"continent": {
				"id": 1,
				"name": "Europe",
				"code": "EU"
			},
			"leagues": [
				{
					"id": 8,
					"sport_id": 1,
					"country_id": 462,
					"name": "Premier League",
					"active": true,
					"short_code": "UK PL",
					"image_path": "https:\/\/cdn.sportmonks.com\/images\/soccer\/leagues\/8\/8.png",
					"type": "league",
					"sub_type": "domestic",
					"last_played_at": "2024-09-22 15:30:00",
					"category": 1,
					"has_jerseys": false
				}
			]
		}
	],
	"pagination": {
		"count": 2,
		"per_page": 25,
		"current_page": 1,
		"next_page": null,
		"has_more": false
	},
	"subscription": [
		{
			"meta": {
				"trial_ends_at": null,
				"ends_at": "2024-10-26 12:06:34",
				"current_timestamp": 1728372666
			},
			"plans": [
				{
					"plan": "Joe Sweeny Custom Plan",
					"sport": "Football",
					"category": "Advanced"
				}
			],
			"add_ons": [],
			"widgets": []
		}
	],
	"rate_limit": {
		"resets_in_seconds": 3386,
		"remaining": 2997,
		"requested_entity": "Country"
	},
	"timezone": "UTC"
}`

var countryResponse = `{
	"data": {
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
	"subscription": [
		{
			"meta": {
				"trial_ends_at": null,
				"ends_at": "2024-10-26 12:06:34",
				"current_timestamp": 1728372666
			},
			"plans": [
				{
					"plan": "Joe Sweeny Custom Plan",
					"sport": "Football",
					"category": "Advanced"
				}
			],
			"add_ons": [],
			"widgets": []
		}
	],
	"rate_limit": {
		"resets_in_seconds": 3386,
		"remaining": 2997,
		"requested_entity": "Country"
	},
	"timezone": "UTC"
}
`

var countryIncludesResponse = `{
	"data": {
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
	  "image_path": "https://cdn.sportmonks.com/images/countries/png/short/pl.png",
      "continent": {
		  "id": 1,
		  "name": "Europe",
		  "code": "EU"
      },
      "leagues": [
			{
				"id": 8,
				"sport_id": 1,
				"country_id": 462,
				"name": "Premier League",
				"active": true,
				"short_code": "UK PL",
				"image_path": "https:\/\/cdn.sportmonks.com\/images\/soccer\/leagues\/8\/8.png",
				"type": "league",
				"sub_type": "domestic",
				"last_played_at": "2024-09-22 15:30:00",
				"category": 1,
				"has_jerseys": false
			}
	   ]
	}
}`

func TestCountries(t *testing.T) {
	t.Run("returns Country struct slice", func(t *testing.T) {
		url := defaultBaseURL + "/core/countries?api_token=api-key&include=&page=1"

		server := mockResponseServer(t, countriesResponse, 200, url)

		client := newTestHTTPClient(server)

		countries, _, err := client.Countries(context.Background(), 1, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertCountry(t, &countries[0])
		assert.Nil(t, countries[0].Continent)
		assert.Equal(t, 0, len(countries[0].Leagues))
	})

	t.Run("returns Country struct slice with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/core/countries?api_token=api-key&include=continent%3Bleagues&page=1"

		server := mockResponseServer(t, countriesIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		countries, _, err := client.Countries(context.Background(), 1, []string{"continent", "leagues"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		continent := countries[0].Continent
		league := countries[0].Leagues[0]

		assertCountry(t, &countries[0])
		assertContinent(t, continent)
		assertLeague(t, &league)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/core/countries?api_token=api-key&include=&page=1"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		countries, _, err := client.Countries(context.Background(), 1, []string{})

		if countries != nil {
			t.Fatalf("Test failed, expected nil, got %+v", countries)
		}

		assertError(t, err)
	})

	t.Run("can handle response details", func(t *testing.T) {
		url := defaultBaseURL + "/core/countries?api_token=api-key&include=continent%3Bleagues&page=1"

		server := mockResponseServer(t, countriesResponse, 200, url)

		client := newTestHTTPClient(server)

		_, details, _ := client.Countries(context.Background(), 1, []string{"continent", "leagues"})

		assertResponseDetails(t, details, "Country")
	})
}

func TestCountryByID(t *testing.T) {
	t.Run("returns a single Country struct", func(t *testing.T) {
		url := defaultBaseURL + "/core/countries/1?api_token=api-key&include="

		server := mockResponseServer(t, countryResponse, 200, url)

		client := newTestHTTPClient(server)

		country, _, err := client.CountryByID(context.Background(), 1, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertCountry(t, country)
		assert.Nil(t, country.Continent)
		assert.Equal(t, 0, len(country.Leagues))
	})

	t.Run("returns Country struct with country includes data", func(t *testing.T) {
		url := defaultBaseURL + "/core/countries/11?api_token=api-key&include=continent%3Bleagues"

		server := mockResponseServer(t, countryIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		country, _, err := client.CountryByID(context.Background(), 11, []string{"continent", "leagues"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		continent := country.Continent
		league := country.Leagues[0]

		assertCountry(t, country)
		assertContinent(t, continent)
		assertLeague(t, &league)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/core/countries/11?api_token=api-key&include=continent%3Bleagues"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		country, _, err := client.CountryByID(context.Background(), 11, []string{"continent", "leagues"})

		if country != nil {
			t.Fatalf("Test failed, expected nil, got %+v", country)
		}

		assertError(t, err)
	})

	t.Run("can handle response details response", func(t *testing.T) {
		url := defaultBaseURL + "/core/countries/11?api_token=api-key&include=continent%3Bleagues"

		server := mockResponseServer(t, countryResponse, 200, url)

		client := newTestHTTPClient(server)

		_, details, _ := client.CountryByID(context.Background(), 11, []string{"continent", "leagues"})

		assertResponseDetails(t, details, "Country")
	})
}

func assertCountry(t *testing.T, country *Country) {
	assert.Equal(t, 2, country.ID)
	assert.Equal(t, "Poland", country.Name)
	assert.Equal(t, "Republic of Poland", country.OfficialName)
	assert.Equal(t, "POL", country.FifaName)
	assert.Equal(t, "PL", country.Iso2)
	assert.Equal(t, "POL", country.Iso3)
	assert.Equal(t, "52.147850036621094", country.Latitude)
	assert.Equal(t, "19.37775993347168", country.Longitude)
	assert.Equal(t, "https://cdn.sportmonks.com/images/countries/png/short/pl.png", country.ImagePath)
	assert.ElementsMatch(t, []string{"BLR", "CZE", "DEU", "LTU", "RUS", "SVK", "UKR"}, country.Borders)
}

func assertResponseDetails(t *testing.T, d *ResponseDetails, entity string) {
	p := d.Pagination

	if p != nil {
		assert.Equal(t, 2, p.Count)
		assert.Equal(t, 25, p.PerPage)
		assert.Equal(t, 1, p.CurrentPage)
		assert.Nil(t, p.NextPage)
		assert.False(t, p.HasMore)
	}

	s := d.Subscription[0]

	assert.Nil(t, s.Meta.TrialEndsAt)
	assert.Equal(t, "2024-10-26 12:06:34", s.Meta.EndsAt)
	assert.Equal(t, int64(1728372666), s.Meta.CurrentTimestamp)

	assert.Equal(t, 1, len(s.Plans))
	assert.Equal(t, "Joe Sweeny Custom Plan", s.Plans[0].Plan)
	assert.Equal(t, "Football", s.Plans[0].Sport)
	assert.Equal(t, "Advanced", s.Plans[0].Category)

	r := d.RateLimit

	assert.Equal(t, 3386, r.ResetsInSeconds)
	assert.Equal(t, 2997, r.Remaining)
	assert.Equal(t, entity, r.RequestedEntity)

	assert.Equal(t, "UTC", d.TimeZone)
}
