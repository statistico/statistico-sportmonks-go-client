package sportmonks

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const defaultBaseUri = "https://soccer.sportmonks.com"
const defaultRetries = 5
const defaultBackOffExponent = 2
const defaultBackOffLimit = 120
const defaultBackOffTotal = 1

// ApiClient is a HTTP request builder and sender
type ApiClient struct {
	client  *http.Client
	baseURL string
	apiKey  string
	maxRetries int
	backOffExponent int
	backOffLimit int
	backOffTotal int
}

func (c *ApiClient) baseUrl() string {
	if c.baseURL == "" {
		c.baseURL = defaultBaseUri
	}

	return c.baseURL
}

func (c *ApiClient) httpClient(client *http.Client) *http.Client {
	if c.client == nil {
		c.client = http.DefaultClient
	}

	return c.client
}

func (c *ApiClient) retries() int {
	if c.maxRetries == 0 {
		c.maxRetries = defaultRetries
	}

	return c.maxRetries
}

func (c *ApiClient) createRequest(ctx context.Context, path string, includes []string) (*http.Request, error) {
	return c.buildRequest(ctx, path, includes, 0)
}

func (c *ApiClient) createPaginatedRequest(ctx context.Context, path string, includes []string, page int) (*http.Request, error) {
	return c.buildRequest(ctx, path, includes, page)
}

func (c *ApiClient) sendRequest(req *http.Request, response interface{}) error {
	res, err := c.client.Do(req)

	if err != nil {
		if c.maxRetries > 0 {
			c.backOff()
			c.maxRetries--
			return c.sendRequest(req, response)
		}

		return err
	}

	if err := decodeResponseBody(res.Body, &response); err != nil {
		return err
	}

	// Reset Client here

	return nil
}

func (c *ApiClient) buildRequest(ctx context.Context, path string, includes []string, page int) (*http.Request, error) {
	url := c.baseURL + path + "?api_token=" + c.apiKey

	if page > 0 {
		url = url + "&page=" + strconv.Itoa(page)
	}

	if len(includes) > 0 {
		url = url + "&include=" + strings.Join(includes, ",")
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	return req, nil
}

func decodeResponseBody(body io.ReadCloser, response interface{}) error {
	defer body.Close()

	b, err := ioutil.ReadAll(body)

	if err != nil {
		return nil
	}

	if err = json.Unmarshal([]byte(b), &response); err != nil {
		return err
	}

	return nil
}

func (c *ApiClient) backOff() {
	time.Sleep(time.Duration(c.backOffTotal))
	c.backOffTotal = c.backOffTotal * c.backOffExponent

	if c.backOffTotal > c.backOffLimit {
		c.backOffTotal = c.backOffLimit
	}
}
