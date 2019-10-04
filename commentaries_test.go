package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var commentariesResponse = `{
	"data": [
		{
		  "fixture_id": 11867289,
		  "important": false,
		  "order": 101,
		  "goal": false,
		  "minute": 0,
		  "extra_minute": null,
		  "comment": "Thats all. Game finished -  Crystal Palace 2, Norwich City 0."
    	},
    	{
		  "fixture_id": 11867289,
		  "important": false,
		  "order": 100,
		  "goal": false,
		  "minute": 90,
		  "extra_minute": 6,
		  "comment": "Second Half ended - Crystal Palace 2, Norwich City 0."
    	}
	]
}
`

func TestCommentariesByFixtureId(t *testing.T) {
	t.Run("returns commentart struct slice", func(t *testing.T) {
		server := mockResponseServer(commentariesResponse, 200)

		client := newTestHTTPClient(server)

		commentaries, _, err := client.CommentariesByFixtureId(context.Background(), 11867289)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assert.Equal(t, 11867289, commentaries[0].FixtureID)
		assert.Equal(t, false, commentaries[0].Important)
		assert.Equal(t, 101, commentaries[0].Order)
		assert.Equal(t, false, commentaries[0].Goal)
		assert.Equal(t, 0, commentaries[0].Minute)
		assert.Nil(t, commentaries[0].ExtraMinute)
		assert.Equal(t, "Thats all. Game finished -  Crystal Palace 2, Norwich City 0.", commentaries[0].Comment)

		assert.Equal(t, 11867289, commentaries[1].FixtureID)
		assert.Equal(t, false, commentaries[1].Important)
		assert.Equal(t, 100, commentaries[1].Order)
		assert.Equal(t, false, commentaries[1].Goal)
		assert.Equal(t, 90, commentaries[1].Minute)
		assert.Equal(t, 6, *commentaries[1].ExtraMinute)
		assert.Equal(t, "Second Half ended - Crystal Palace 2, Norwich City 0.", commentaries[1].Comment)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		server := mockResponseServer(errorResponse, 400)

		client := newTestHTTPClient(server)

		bookmakers, _, err := client.CommentariesByFixtureId(context.Background(), 999)

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
