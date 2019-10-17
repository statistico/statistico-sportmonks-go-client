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
    	}
	]
}
`

func TestCommentariesByFixtureID(t *testing.T) {
	url := defaultBaseURL + "/commentaries/fixture/11867289?api_token=api-key"

	t.Run("returns commentary struct slice", func(t *testing.T) {
		server := mockResponseServer(t, commentariesResponse, 200, url)

		client := newTestHTTPClient(server)

		commentaries, _, err := client.CommentariesByFixtureID(context.Background(), 11867289)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertCommentary(t, &commentaries[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		bookmakers, _, err := client.CommentariesByFixtureID(context.Background(), 11867289)

		if bookmakers != nil {
			t.Fatalf("Test failed, expected nil, got %+v", bookmakers)
		}

		assertError(t, err)
	})
}

func assertCommentary(t *testing.T, commentary *Commentary) {
	assert.Equal(t, 11867289, commentary.FixtureID)
	assert.Equal(t, false, commentary.Important)
	assert.Equal(t, 101, commentary.Order)
	assert.Equal(t, false, commentary.Goal)
	assert.Equal(t, 0, commentary.Minute)
	assert.Nil(t, commentary.ExtraMinute)
	assert.Equal(t, "Thats all. Game finished -  Crystal Palace 2, Norwich City 0.", commentary.Comment)
}
