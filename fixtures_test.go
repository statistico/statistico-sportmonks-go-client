package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var fixtureResponse = `{
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
}`

var fixturesResponse = `{
	"data": [
		{
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
	]
}`

func TestFixtureById(t *testing.T) {
	url := defaultBaseUrl + "/fixtures/11867285?api_token=api-key&include="

	t.Run("returns a single Fixture struct", func(t *testing.T) {
		server := mockResponseServer(t, fixtureResponse, 200, url)

		client := newTestHTTPClient(server)

		fixture, _, err := client.FixtureById(context.Background(), 11867285, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertFixture(t, fixture)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		fixture, _, err := client.FixtureById(context.Background(), 11867285, []string{})

		if fixture != nil {
			t.Fatalf("Test failed, expected nil, got %+v", fixture)
		}

		assertError(t, err)
	})
}

func TestFixturesById(t *testing.T) {
	url := defaultBaseUrl + "/fixtures/multi/11867285,555?api_token=api-key&include="

	t.Run("returns slice of Fixture struct", func(t *testing.T) {
		server := mockResponseServer(t, fixturesResponse, 200, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.FixturesById(context.Background(), []int{11867285, 555}, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertFixture(t, &fixtures[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.FixturesById(context.Background(), []int{11867285, 555}, []string{})

		if fixtures != nil {
			t.Fatalf("Test failed, expected nil, got %+v", fixtures)
		}

		assertError(t, err)
	})
}

func TestFixturesByDate(t *testing.T) {
	str := "2014-11-12T11:45:26.371Z"
	d, err := time.Parse(time.RFC3339, str)

	if err != nil {
		t.Fatalf("Test failed, expected nil, got %+v", err.Error())
	}

	url := defaultBaseUrl + "/fixtures/date/2014-11-12?api_token=api-key&include="

	t.Run("returns slice of Fixture struct", func(t *testing.T) {
		server := mockResponseServer(t, fixturesResponse, 200, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.FixturesByDate(context.Background(), d, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertFixture(t, &fixtures[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.FixturesByDate(context.Background(), d, []string{})

		if fixtures != nil {
			t.Fatalf("Test failed, expected nil, got %+v", fixtures)
		}

		assertError(t, err)
	})
}

func TestFixturesBetween(t *testing.T) {
	from := "2014-11-12T11:45:26.371Z"
	dateFrom, err := time.Parse(time.RFC3339, from)

	if err != nil {
		t.Fatalf("Test failed, expected nil, got %+v", err.Error())
	}

	to := "2014-12-12T11:45:26.371Z"
	dateTo, err := time.Parse(time.RFC3339, to)

	if err != nil {
		t.Fatalf("Test failed, expected nil, got %+v", err.Error())
	}

	url := defaultBaseUrl + "/fixtures/between/2014-11-12/2014-12-12?api_token=api-key&include="

	t.Run("returns slice of Fixture struct", func(t *testing.T) {
		server := mockResponseServer(t, fixturesResponse, 200, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.FixturesBetween(context.Background(), dateFrom, dateTo, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertFixture(t, &fixtures[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.FixturesBetween(context.Background(), dateFrom, dateTo, []string{})

		if fixtures != nil {
			t.Fatalf("Test failed, expected nil, got %+v", fixtures)
		}

		assertError(t, err)
	})
}

func TestFixturesLastUpdated(t *testing.T) {
	url := defaultBaseUrl + "/fixtures/updates?api_token=api-key&include="

	t.Run("returns slice of Fixture struct", func(t *testing.T) {
		server := mockResponseServer(t, fixturesResponse, 200, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.FixturesLastUpdated(context.Background(), []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertFixture(t, &fixtures[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		fixtures, _, err := client.FixturesLastUpdated(context.Background(), []string{})

		if fixtures != nil {
			t.Fatalf("Test failed, expected nil, got %+v", fixtures)
		}

		assertError(t, err)
	})
}

func assertFixture(t *testing.T, fixture *Fixture) {
	assert.Equal(t, 11867285, fixture.ID)
	assert.Equal(t, 8, fixture.LeagueID)
	assert.Equal(t, 16036, fixture.SeasonID)
	assert.Equal(t, 77443862, *fixture.StageID)
	assert.Equal(t, 169657, *fixture.RoundID)
	assert.Nil(t,  fixture.GroupID)
	assert.Nil(t,  fixture.AggregateID)
	assert.Equal(t, 214, *fixture.VenueID)
	assert.Equal(t, 14532, *fixture.RefereeID)
	assert.Equal(t, 1, fixture.LocalTeamID)
	assert.Equal(t, 14, fixture.VisitorTeamID)
	assert.Equal(t, 1, *fixture.WinnerTeamID)
	assert.Equal(t, "rain", fixture.WeatherReport.Code)
	assert.Equal(t, "shower rain", fixture.WeatherReport.Type)
	assert.Equal(t, "https://cdn.sportmonks.com/images/weather/09d.png", fixture.WeatherReport.Icon)
	assert.Equal(t, 62.96, fixture.WeatherReport.Temperature.Temp)
	assert.Equal(t, "fahrenheit", fixture.WeatherReport.Temperature.Unit)
	assert.Equal(t, 17.2, fixture.WeatherReport.TemperatureCelcius.Temp)
	assert.Equal(t, "celcius", fixture.WeatherReport.TemperatureCelcius.Unit)
	assert.Equal(t, "75%", fixture.WeatherReport.Clouds)
	assert.Equal(t, "82%", fixture.WeatherReport.Humidity)
	assert.Equal(t, 1004, *fixture.WeatherReport.Pressure)
	assert.Equal(t, "5.82 m/s", fixture.WeatherReport.Wind.Speed)
	assert.Equal(t, 240, fixture.WeatherReport.Wind.Degree)
	assert.Equal(t, 51.51, *fixture.WeatherReport.Coordinates.Lat)
	assert.Equal(t, -0.13, *fixture.WeatherReport.Coordinates.Lon)
	assert.Equal(t, "2019-09-22T14:45:05.289505Z", *fixture.WeatherReport.UpdatedAt)
	assert.Equal(t, true, *fixture.Commentaries)
	assert.Equal(t, 59936, *fixture.Attendance)
	assert.Nil(t, fixture.Pitch)
	assert.Nil(t, fixture.Details)
	assert.Equal(t, false, *fixture.NeutralVenue)
	assert.Equal(t, true, fixture.WinningOddsCalculated)
	assert.Equal(t, "4-1-4-1", fixture.Formations.LocalTeamFormation)
	assert.Equal(t, "4-2-3-1", fixture.Formations.VisitorTeamFormation)
	assert.Equal(t, 2, fixture.Scores.LocalTeamScore)
	assert.Equal(t, 0, fixture.Scores.VisitorTeamScore)
	assert.Nil(t, fixture.Scores.LocalTeamPenScore)
	assert.Nil(t, fixture.Scores.VisitorTeamPenScore)
	assert.Equal(t, "1-0", *fixture.Scores.HTScore)
	assert.Equal(t, "2-0", *fixture.Scores.FTScore)
	assert.Nil(t, fixture.Scores.ETScore)
	assert.Nil(t, fixture.Scores.PSScore)
	assert.Equal(t, "FT", fixture.Time.Status)
	assert.Equal(t, "2019-09-22 13:00:00", fixture.Time.StartingAt.DateTime)
	assert.Equal(t, "2019-09-22", fixture.Time.StartingAt.Date)
	assert.Equal(t, "13:00:00", fixture.Time.StartingAt.Time)
	assert.Equal(t, 1569157200, fixture.Time.StartingAt.Timestamp)
	assert.Equal(t, "UTC", fixture.Time.StartingAt.Timezone)
	assert.Equal(t, 90, fixture.Time.Minute)
	assert.Nil(t, fixture.Time.Second)
	assert.Nil(t, fixture.Time.AddedTime)
	assert.Nil(t, fixture.Time.ExtraMinute)
	assert.Nil(t, fixture.Time.InjuryTime)
	assert.Equal(t, 523898, fixture.Coaches.LocalTeamCoachID)
	assert.Equal(t, 524307, fixture.Coaches.VisitorTeamCoachID)
	assert.Equal(t, 11, fixture.Standings.LocalTeamPosition)
	assert.Equal(t, 6, fixture.Standings.VisitorTeamPosition)
	assert.Equal(t, 12794, *fixture.Assistants.FirstAssistantID)
	assert.Equal(t, 12798, *fixture.Assistants.SecondAssistantID)
	assert.Equal(t, 15270, *fixture.Assistants.FourthOfficialID)
	assert.Equal(t, "1/1", *fixture.Leg)
	assert.Equal(t, "#832034", *fixture.Colors.LocalTeam.Color)
	assert.Equal(t, "#C0D6FE,#C0D6FE,#832034,#832034,#999999,#888888,#832034", *fixture.Colors.LocalTeam.KitColor)
	assert.Nil(t, fixture.Colors.VisitorTeam.Color)
	assert.Nil(t, fixture.Colors.VisitorTeam.KitColor)
	assert.Equal(t, false, fixture.Deleted)
}