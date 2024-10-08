package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var venueResponse = `{
	"data": {
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
		"requested_entity": "Venue"
	},
	"timezone": "UTC"
}`

var venueSeasonResponse = `{
	"data": [
		{
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
	],
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
		"requested_entity": "Venue"
	},
	"timezone": "UTC"
}`

func TestVenueByID(t *testing.T) {
	t.Run("returns a single Venue struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/venues/200?api_token=api-key"

		server := mockResponseServer(t, venueResponse, 200, url)

		client := newTestHTTPClient(server)

		venue, _, err := client.VenueByID(context.Background(), 200)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertVenue(t, venue)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/venues/200?api_token=api-key"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		venue, _, err := client.VenueByID(context.Background(), 200)

		if venue != nil {
			t.Fatalf("Test failed, expected nil, got %+v", venue)
		}

		assertError(t, err)
	})

	t.Run("can handle response details response", func(t *testing.T) {
		url := defaultBaseURL + "/football/venues/200?api_token=api-key"

		server := mockResponseServer(t, venueResponse, 200, url)

		client := newTestHTTPClient(server)

		_, details, _ := client.VenueByID(context.Background(), 200)

		assertResponseDetails(t, details, "Venue")
	})
}

func TestVenueBySeasonID(t *testing.T) {
	t.Run("returns a slice of Venue struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/venues/seasons/12962?api_token=api-key"

		server := mockResponseServer(t, venueSeasonResponse, 200, url)

		client := newTestHTTPClient(server)

		venues, _, err := client.VenuesBySeasonID(context.Background(), 12962)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertVenue(t, &venues[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/venues/seasons/12962?api_token=api-key"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		venues, _, err := client.VenuesBySeasonID(context.Background(), 12962)

		if venues != nil {
			t.Fatalf("Test failed, expected nil, got %+v", venues)
		}

		assertError(t, err)
	})

	t.Run("can handle response details response", func(t *testing.T) {
		url := defaultBaseURL + "/football/venues/seasons/12962?api_token=api-key"

		server := mockResponseServer(t, venueSeasonResponse, 200, url)

		client := newTestHTTPClient(server)

		_, details, _ := client.VenuesBySeasonID(context.Background(), 12962)

		assertResponseDetails(t, details, "Venue")
	})
}

func assertVenue(t *testing.T, venue *Venue) {
	assert.Equal(t, 214, venue.ID)
	assert.Equal(t, "London Stadium", venue.Name)
	assert.Equal(t, "grass", venue.Surface)
	assert.Equal(t, "Marshgate Lane, Stratford", venue.Address)
	assert.Equal(t, "London", venue.CityName) // Updated to use CityName
	assert.Equal(t, 60000, venue.Capacity)
	assert.Equal(t, "https://cdn.sportmonks.com/images/soccer/venues/22/214.png", venue.ImagePath)
	assert.Nil(t, venue.Zipcode) // Check if Zipcode is nil
}
