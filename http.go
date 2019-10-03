package sportmonks

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const defaultBaseUrl = "https://soccer.sportmonks.com/api/v2.0"

// Client is a HTTP request builder and sender
type Client struct {
	HTTPClient *http.Client
	BaseURL string
	Key  string
}

// NewClient creates a new Client with default settings. A key is required to
// instantiate the Client
func NewClient(key string) *Client {
	return &Client{
		HTTPClient: &http.Client{},
		BaseURL: defaultBaseUrl,
		Key: key,
	}
}

func (c *Client) createRequest(ctx context.Context, path string, includes []string) (*http.Request, error) {
	return c.buildRequest(ctx, path, includes, 0)
}

func (c *Client) createPaginatedRequest(ctx context.Context, path string, includes []string, page int) (*http.Request, error) {
	return c.buildRequest(ctx, path, includes, page)
}

func (c *Client) Do(req *http.Request, response interface{}) error {
	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return err
	}

	if err := decodeResponseBody(resp.Body, &response); err != nil {
		return err
	}

	return nil
}

func (c *Client) buildRequest(ctx context.Context, path string, includes []string, page int) (*http.Request, error) {
	url := c.BaseURL + path + "?api_token=" + c.Key

	if page > 0 {
		url = url + "&page=" + strconv.Itoa(page)
	}

	if len(includes) > 0 {
		url = url + "&include=" + strings.Join(includes, ",")
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	return req, nil
}

func decodeResponseBody(body io.ReadCloser, response interface{}) error {
	defer body.Close()

	if err := json.NewDecoder(body).Decode(response); err != nil {
		return err
	}

	return nil
}
