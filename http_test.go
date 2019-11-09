package sportmonks

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

var errorResponse = `{
	"error": {
		"message": "The requested endpoint does not exists!",
		"code": 404
	}
}`

func TestNewHTTPClient(t *testing.T) {
	t.Run("instantiates with default properties", func(t *testing.T) {
		client := NewDefaultHTTPClient("api-key")

		assert.Equal(t, "https://soccer.sportmonks.com/api/v2.0", client.BaseURL)
		assert.Equal(t, "api-key", client.Key)
	})

	t.Run("instantiates with bespoke properties", func(t *testing.T) {
		client := HTTPClient{
			&http.Client{},
			"https://example.com",
			"new-key",
		}

		assert.Equal(t, "https://example.com", client.BaseURL)
		assert.Equal(t, "new-key", client.Key)
	})

	t.Run("url is parsed as expected", func(t *testing.T) {
		server := newTestClient(func(req *http.Request) *http.Response {
			assert.Equal(
				t,
				req.URL.String(),
				"https://soccer.sportmonks.com/api/v2.0/continents/10?api_token=api-key&include=countries",
			)

			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(""))),
			}
		})

		client := newTestHTTPClient(server)

		_, _, _ = client.ContinentByID(context.Background(), 10, []string{"countries"})
	})
}

func assertError(t *testing.T, err error) {
	assert.Equal(
		t,
		"Request failed with message: The requested endpoint does not exists!, code: 404",
		err.Error(),
	)
}

func newTestHTTPClient(server *http.Client) *HTTPClient {
	return &HTTPClient{
		HTTPClient: server,
		BaseURL:    defaultBaseURL,
		Key:        "api-key",
	}
}

// roundTripFunc .
type roundTripFunc func(req *http.Request) *http.Response

// roundTrip .
func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func newTestClient(fn roundTripFunc) *http.Client {
	return &http.Client{
		Transport: roundTripFunc(fn),
	}
}

func mockResponseServer(t *testing.T, body string, code int, url string) *http.Client {
	return newTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, url, req.URL.String())

		return &http.Response{
			StatusCode: code,
			Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
		}
	})
}
