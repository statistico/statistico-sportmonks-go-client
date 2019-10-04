package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var coachResponse = `{
	"data": {
		"coach_id": 2,
		"team_id": 1238,
		"country_id": 462,
		"common_name": "R. Fowler",
		"fullname": "R. Fowler",
		"firstname": "Robbie",
		"lastname": "Fowler",
		"nationality": "England",
		"birthdate": "09\/04\/1975",
		"birthcountry": "England",
		"birthplace": null,
		"image_path": "https:\/\/cdn.sportmonks.com\/images\/soccer\/players\/2\/2.png"
  	}
}`

func TestCoachById(t *testing.T) {
	t.Run("returns a single coach struct", func(t *testing.T) {
		server := mockResponseServer(coachResponse, 200)

		client := newTestHTTPClient(server)

		coach, _, err := client.CoachById(context.Background(), 2)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assert.Equal(t, 2, coach.ID)
		assert.Equal(t, 1238, coach.TeamID)
		assert.Equal(t, 462, coach.CountryID)
		assert.Equal(t, "R. Fowler", coach.CommonName)
		assert.Equal(t, "Robbie", coach.FirstName)
		assert.Equal(t, "Fowler", coach.LastName)
		assert.Equal(t, "England", coach.Nationality)
		assert.Equal(t, "09/04/1975", coach.BirthDate)
		assert.Equal(t, "England", coach.BirthCountry)
		assert.Nil(t, coach.BirthPlace)
		assert.Equal(t, "https://cdn.sportmonks.com/images/soccer/players/2/2.png", coach.ImagePath)
	})

	t.Run("returns a bad status code error", func(t *testing.T) {
		server := mockResponseServer(errorResponse, 404)

		client := newTestHTTPClient(server)

		coach, _, err := client.CoachById(context.Background(), 2)

		if coach != nil {
			t.Fatalf("Test failed, expected nil, got %+v", coach)
		}

		assert.Equal(
			t,
			"Request failed with message: The requested endpoint does not exists!, code: 404",
			err.Error(),
		)
	})
}
