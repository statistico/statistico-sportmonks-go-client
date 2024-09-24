package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tvStationsResponse = `{
	"data": [
		{
		  "id": 33,
		  "name": "Star+",
		  "url": "https://www.starplus.com/",
		  "image_path": "https://cdn.sportmonks.com/images/core/tvstations/1/33.png",
		  "type": "tv",
		  "related_id": null
		}
	]
}`

func TestTVStationsByFixtureID(t *testing.T) {
	t.Run("returns a slice of TVStation struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/tv-stations/fixtures/11867285?api_token=api-key"

		server := mockResponseServer(t, tvStationsResponse, 200, url)

		client := newTestHTTPClient(server)

		tv, _, err := client.TVStationsByFixtureID(context.Background(), 11867285)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertTVStation(t, &tv[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/tv-stations/fixtures/11867285?api_token=api-key"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		tv, _, err := client.TVStationsByFixtureID(context.Background(), 11867285)

		if tv != nil {
			t.Fatalf("Test failed, expected nil, got %+v", tv)
		}

		assertError(t, err)
	})
}

func assertTVStation(t *testing.T, tv *TVStation) {
	assert.Equal(t, 33, tv.ID)
	assert.Equal(t, "Star+", tv.Name)
	assert.Equal(t, "https://www.starplus.com/", tv.URL)
	assert.Equal(t, "https://cdn.sportmonks.com/images/core/tvstations/1/33.png", tv.ImagePath)
	assert.Equal(t, "tv", tv.Type)
	assert.Nil(t, tv.RelatedID)
}
