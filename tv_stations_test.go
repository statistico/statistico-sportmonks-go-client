package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tvStationsResponse = `{
	"data": [
		{
		  "fixture_id": 11867285,
		  "tvstation": "CANAL+"
		}
	]
}`

func TestTVStationsByFixtureId(t *testing.T) {
	t.Run("returns a slice of TVStation struct", func(t *testing.T) {
		url := defaultBaseUrl + "/tvstations/fixture/11867285?api_token=api-key"

		server := mockResponseServer(t, tvStationsResponse, 200, url)

		client := newTestHTTPClient(server)

		tv, _, err := client.TVStationsByFixtureId(context.Background(), 11867285)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertTVStation(t, &tv[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseUrl + "/tvstations/fixture/11867285?api_token=api-key"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		tv, _, err := client.TVStationsByFixtureId(context.Background(), 11867285)

		if tv != nil {
			t.Fatalf("Test failed, expected nil, got %+v", tv)
		}

		assertError(t, err)
	})
}

func assertTVStation(t *testing.T, tv *TVStation) {
	assert.Equal(t, 11867285, tv.FixtureID)
	assert.Equal(t, "CANAL+", tv.TVStation)
}
