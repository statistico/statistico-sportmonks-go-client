package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var videoHighlightsResponse = `{
	"data": [
		{
			"fixture_id": 11865620,
			"event_id": null,
			"location": "https:\/\/cc.sporttube.com\/embed\/05tCCCG",
			"type": "video",
			"created_at": {
			"date": "2019-10-06 18:36:10.000000",
			"timezone_type": 3,
			"timezone": "UTC"
			}
		}
	]
}`

var videoHighlightsIncludesResponse = `{
	"data": [
		{
			"fixture_id": 11865620,
			"event_id": null,
			"location": "https:\/\/cc.sporttube.com\/embed\/05tCCCG",
			"type": "video",
			"created_at": {
			"date": "2019-10-06 18:36:10.000000",
			"timezone_type": 3,
			"timezone": "UTC"
			},
			"fixture": {
				"data": {
					"id": 11867285,
					"league_id": 8,
					"season_id": 16036,
					"stage_id": 77443862,
					"round_id": 169657,
					"group_id": null,
					"aggregate_id": null,
					"venue_id": 214,
					"referee_id": 14532,
					"localteam_id": 1,
					"visitorteam_id": 14,
					"winner_team_id": 1,
					"weather_report": {
					  "code": "rain",
					  "type": "shower rain",
					  "icon": "https:\/\/cdn.sportmonks.com\/images\/weather\/09d.png",
					  "temperature": {
						"temp": 62.96,
						"unit": "fahrenheit"
					  },
					  "temperature_celcius": {
						"temp": 17.2,
						"unit": "celcius"
					  },
					  "clouds": "75%",
					  "humidity": "82%",
					  "pressure": 1004,
					  "wind": {
						"speed": "5.82 m\/s",
						"degree": 240
					  },
					  "coordinates": {
						"lat": 51.51,
						"lon": -0.13
					  },
					  "updated_at": "2019-09-22T14:45:05.289505Z"
					},
					"commentaries": true,
					"attendance": 59936,
					"pitch": null,
					"details": null,
					"neutral_venue": false,
					"winning_odds_calculated": true,
					"formations": {
					  "localteam_formation": "4-1-4-1",
					  "visitorteam_formation": "4-2-3-1"
					},
					"scores": {
					  "localteam_score": 2,
					  "visitorteam_score": 0,
					  "localteam_pen_score": null,
					  "visitorteam_pen_score": null,
					  "ht_score": "1-0",
					  "ft_score": "2-0",
					  "et_score": null,
					  "ps_score": null
					},
					"time": {
					  "status": "FT",
					  "starting_at": {
						"date_time": "2019-09-22 13:00:00",
						"date": "2019-09-22",
						"time": "13:00:00",
						"timestamp": 1569157200,
						"timezone": "UTC"
					  },
					  "minute": 90,
					  "second": null,
					  "added_time": null,
					  "extra_minute": null,
					  "injury_time": null
					},
					"coaches": {
					  "localteam_coach_id": 523898,
					  "visitorteam_coach_id": 524307
					},
					"standings": {
					  "localteam_position": 11,
					  "visitorteam_position": 6
					},
					"assistants": {
					  "first_assistant_id": 12794,
					  "second_assistant_id": 12798,
					  "fourth_official_id": 15270
					},
					"leg": "1\/1",
					"colors": {
					  "localteam": {
						"color": "#832034",
						"kit_colors": "#C0D6FE,#C0D6FE,#832034,#832034,#999999,#888888,#832034"
					  },
					  "visitorteam": {
						"color": null,
						"kit_colors": null
					  }
					},
					"deleted": false
				}
			}
		}
	]
}`

func TestVideoHighlights(t *testing.T) {
	t.Run("returns a slice of VideoHighlight struct", func(t *testing.T) {
		url := defaultBaseURL + "/highlights?api_token=api-key&include=&page=2"

		server := mockResponseServer(t, videoHighlightsResponse, 200, url)

		client := newTestHTTPClient(server)

		highlights, _, err := client.VideoHighlights(context.Background(), 2, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertVideoHighlights(t, &highlights[0])
	})

	t.Run("returns a slice of VideoHighlight struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/highlights?api_token=api-key&include=fixture&page=2"

		server := mockResponseServer(t, videoHighlightsIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		highlights, _, err := client.VideoHighlights(context.Background(), 2, []string{"fixture"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertVideoHighlights(t, &highlights[0])
		assertFixture(t, highlights[0].Fixture())
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/highlights?api_token=api-key&include=&page=2"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		highlights, _, err := client.VideoHighlights(context.Background(), 2, []string{})

		if highlights != nil {
			t.Fatalf("Test failed, expected nil, got %+v", highlights)
		}

		assertError(t, err)
	})
}

func assertVideoHighlights(t *testing.T, video *VideoHighlights) {
	assert.Equal(t, 11865620, video.FixtureID)
	assert.Nil(t, video.EventID)
	assert.Equal(t, "https://cc.sporttube.com/embed/05tCCCG", video.Location)
	assert.Equal(t, "video", video.Type)
	assert.Equal(t, "2019-10-06 18:36:10.000000", video.CreatedAt.Date)
	assert.Equal(t, 3, video.CreatedAt.TimezoneType)
	assert.Equal(t, "UTC", video.CreatedAt.Timezone)
}
