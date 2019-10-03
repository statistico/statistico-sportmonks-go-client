package sportmonks

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestNewHTTPClient(t *testing.T) {
	t.Run("instantiates with default properties", func(t *testing.T) {
		client := NewHTTPClient("api-key")

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
}

func newTestHTTPClient(server *http.Client) *HTTPClient {
	return &HTTPClient{
		HTTPClient: server,
		BaseURL:    defaultBaseUrl,
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

func mockResponseServer(body string, code int) *http.Client {
	return newTestClient(func (req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: code,
			Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
		}
	})
}
