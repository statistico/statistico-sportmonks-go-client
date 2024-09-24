package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var coachResponse = `{
	"data": {
      "id": 24,
      "player_id": 24,
      "sport_id": 1,
      "country_id": 462,
      "nationality_id": 462,
      "city_id": null,
      "common_name": "D. Unsworth",
      "firstname": "David",
      "lastname": "Unsworth",
      "name": "David Unsworth",
      "display_name": "David Unsworth",
      "image_path": "https://cdn.sportmonks.com/images/soccer/placeholder.png",
      "height": null,
      "weight": null,
      "date_of_birth": "1973-10-16",
      "gender": "male"
    }
}`

func TestCoachByID(t *testing.T) {
	url := defaultBaseURL + "/football/coaches/2?api_token=api-key"

	t.Run("returns a single coach struct", func(t *testing.T) {
		server := mockResponseServer(t, coachResponse, 200, url)

		client := newTestHTTPClient(server)

		coach, _, err := client.CoachByID(context.Background(), 2)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertCoach(t, coach)
	})

	t.Run("returns a bad status code error", func(t *testing.T) {
		server := mockResponseServer(t, errorResponse, 404, url)

		client := newTestHTTPClient(server)

		coach, _, err := client.CoachByID(context.Background(), 2)

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

func assertCoach(t *testing.T, coach *Coach) {
	assert.Equal(t, 24, coach.ID)
	assert.Equal(t, 24, coach.PlayerID) // Updated from TeamID to PlayerID
	assert.Equal(t, 462, coach.CountryID)
	assert.Equal(t, "D. Unsworth", coach.CommonName)
	assert.Equal(t, "David", coach.FirstName)
	assert.Equal(t, "Unsworth", coach.LastName)
	assert.Equal(t, "David Unsworth", coach.Name)
	assert.Equal(t, "David Unsworth", coach.DisplayName)
	assert.Nil(t, coach.CityID)
	assert.Equal(t, "1973-10-16", coach.DateOfBirth)
	assert.Equal(t, "male", coach.Gender)
	assert.Nil(t, coach.Height)
	assert.Nil(t, coach.Weight)
	assert.Equal(t, "https://cdn.sportmonks.com/images/soccer/placeholder.png", coach.ImagePath)
}
