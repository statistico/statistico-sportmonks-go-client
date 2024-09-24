package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var playerResponse = `{
	"data": {
		  "id": 1,
		  "sport_id": 1,
		  "country_id": 462,
		  "nationality_id": 462,
		  "city_id": 21022,
		  "position_id": 27,
		  "detailed_position_id": null,
		  "type_id": 27,
		  "common_name": "R. Hulse",
		  "firstname": "Rob",
		  "lastname": "Hulse",
		  "name": "Rob Hulse",
		  "display_name": "Rob Hulse",
		  "image_path": "https://cdn.sportmonks.com/images/soccer/players/1/1.png",
		  "height": 187,
		  "weight": 78,
		  "date_of_birth": "1979-10-25",
		  "gender": "male"
    }
}`

var playerIncludesResponse = `{
	"data": {
		  "id": 1,
		  "sport_id": 1,
		  "country_id": 462,
		  "nationality_id": 462,
		  "city_id": 21022,
		  "position_id": 27,
		  "detailed_position_id": null,
		  "type_id": 27,
		  "common_name": "R. Hulse",
		  "firstname": "Rob",
		  "lastname": "Hulse",
		  "name": "Rob Hulse",
		  "display_name": "Rob Hulse",
		  "image_path": "https://cdn.sportmonks.com/images/soccer/players/1/1.png",
		  "height": 187,
		  "weight": 78,
		  "date_of_birth": "1979-10-25",
		  "gender": "male"
    }
}`

func TestPlayerByID(t *testing.T) {
	t.Run("return a single Player struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/players/219591?api_token=api-key&include="

		server := mockResponseServer(t, playerResponse, 200, url)

		client := newTestHTTPClient(server)

		player, _, err := client.PlayerByID(context.Background(), 219591, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %+v", err)
		}

		assertPlayer(t, player)
	})

	t.Run("return a single Player struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/football/players/219591?api_token=api-key&include=stats%3Bposition%3Btrophies"

		server := mockResponseServer(t, playerIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		player, _, err := client.PlayerByID(context.Background(), 219591, []string{"stats", "position", "trophies"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %+v", err)
		}

		assertPlayer(t, player)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/players/219591?api_token=api-key&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		player, _, err := client.PlayerByID(context.Background(), 219591, []string{})

		if player != nil {
			t.Fatalf("Test failed, expected nil, got %+v", player)
		}

		assertError(t, err)
	})
}

func assertPlayer(t *testing.T, player *Player) {
	assert.Equal(t, 1, player.ID)
	assert.Equal(t, 1, player.SportID)
	assert.Equal(t, 462, player.CountryID)
	assert.Equal(t, 462, player.NationalityID)
	assert.Equal(t, 21022, player.CityID)
	assert.Equal(t, 27, player.PositionID)
	assert.Nil(t, player.DetailedPositionID)
	assert.Equal(t, 27, player.TypeID)
	assert.Equal(t, "R. Hulse", player.CommonName)
	assert.Equal(t, "Rob", player.FirstName)
	assert.Equal(t, "Hulse", player.LastName)
	assert.Equal(t, "Rob Hulse", player.Name)
	assert.Equal(t, "Rob Hulse", player.DisplayName)
	assert.Equal(t, "https://cdn.sportmonks.com/images/soccer/players/1/1.png", player.ImagePath)
	assert.Equal(t, 187, player.Height)
	assert.Equal(t, 78, player.Weight)
	assert.Equal(t, "1979-10-25", player.DateOfBirth)
	assert.Equal(t, "male", player.Gender)
}
