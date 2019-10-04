package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var bookmakersResponse = `{
	"data": [
		{
		  "id": 1,
		  "name": "10Bet",
		  "logo": null
		},
		{
		  "id": 2,
		  "name": "bet365",
		  "logo": null
		}
	]
}
`

var bookmakerResponse = `{
	"data": {
		"id": 1,
    	"name": "10Bet",
    	"logo": null
	}
}
`

func TestBookmakers(t *testing.T) {
	t.Run("returns bookmaker struct slice", func(t *testing.T) {
		server := mockResponseServer(bookmakersResponse, 200)

		client := newTestHTTPClient(server)

		bookmakers, _, err := client.Bookmakers(context.Background())

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assert.Equal(t, 1, bookmakers[0].ID)
		assert.Equal(t, "10Bet", bookmakers[0].Name)
		assert.Nil(t, bookmakers[0].Logo)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		server := mockResponseServer(errorResponse, 400)

		client := newTestHTTPClient(server)

		bookmakers, _, err := client.Bookmakers(context.Background())

		if bookmakers != nil {
			t.Fatalf("Test failed, expected nil, got %+v", bookmakers)
		}

		assert.Equal(
			t,
			"Request failed with message: The requested endpoint does not exists!, code: 404",
			err.Error(),
		)
	})
}

func TestBookMakerById(t *testing.T) {
	t.Run("returns a single bookmaker struct", func(t *testing.T) {
		server := mockResponseServer(bookmakerResponse, 200)

		client := newTestHTTPClient(server)

		bookmaker, _, err := client.BookmakerById(context.Background(), 1)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assert.Equal(t, 1, bookmaker.ID)
		assert.Equal(t, "10Bet", bookmaker.Name)
		assert.Nil(t, bookmaker.Logo)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		server := mockResponseServer(errorResponse, 400)

		client := newTestHTTPClient(server)

		bookmaker, _, err := client.BookmakerById(context.Background(), 10)

		if bookmaker != nil {
			t.Fatalf("Test failed, expected nil, got %+v", bookmaker)
		}

		assert.Equal(
			t,
			"Request failed with message: The requested endpoint does not exists!, code: 404",
			err.Error(),
		)
	})
}
