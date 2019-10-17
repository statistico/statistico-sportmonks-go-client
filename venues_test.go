package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var venueResponse = `{
	"data": {
		"id": 200,
		"name": "Turf Moor",
		"surface": "grass",
		"address": "Harry Potts Way",
		"city": "Burnley",
		"capacity": 22546,
		"image_path": "https:\/\/cdn.sportmonks.com\/images\/soccer\/venues\/8\/200.png",
		"coordinates": null
	}
}`

var venueSeasonResponse = `{
	"data": [
		{
			"id": 200,
			"name": "Turf Moor",
			"surface": "grass",
			"address": "Harry Potts Way",
			"city": "Burnley",
			"capacity": 22546,
			"image_path": "https:\/\/cdn.sportmonks.com\/images\/soccer\/venues\/8\/200.png",
			"coordinates": null
		}
	]
}`

func TestVenueByID(t *testing.T) {
	t.Run("returns a single Venue struct", func(t *testing.T) {
		url := defaultBaseURL + "/venues/200?api_token=api-key"

		server := mockResponseServer(t, venueResponse, 200, url)

		client := newTestHTTPClient(server)

		venue, _, err := client.VenueByID(context.Background(), 200)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertVenue(t, venue)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/venues/200?api_token=api-key"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		venue, _, err := client.VenueByID(context.Background(), 200)

		if venue != nil {
			t.Fatalf("Test failed, expected nil, got %+v", venue)
		}

		assertError(t, err)
	})
}

func TestVenueBySeasonID(t *testing.T) {
	t.Run("returns a slice of Venue struct", func(t *testing.T) {
		url := defaultBaseURL + "/venues/season/12962?api_token=api-key"

		server := mockResponseServer(t, venueSeasonResponse, 200, url)

		client := newTestHTTPClient(server)

		venues, _, err := client.VenuesBySeasonID(context.Background(), 12962)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertVenue(t, &venues[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/venues/season/12962?api_token=api-key"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		venues, _, err := client.VenuesBySeasonID(context.Background(), 12962)

		if venues != nil {
			t.Fatalf("Test failed, expected nil, got %+v", venues)
		}

		assertError(t, err)
	})
}

func assertVenue(t *testing.T, venue *Venue) {
	assert.Equal(t, 200, venue.ID)
	assert.Equal(t, "Turf Moor", venue.Name)
	assert.Equal(t, "grass", venue.Surface)
	assert.Equal(t, "Harry Potts Way", *venue.Address)
	assert.Equal(t, "Burnley", venue.City)
	assert.Equal(t, 22546, venue.Capacity)
	assert.Equal(t, "https://cdn.sportmonks.com/images/soccer/venues/8/200.png", venue.ImagePath)
	assert.Nil(t, venue.Coordinates)
}
