package sportmonks

import (
	"errors"
	"net/http"
)

func NewClient(baseURL, apiKey string) (*Client, error) {
	if baseURL == "" || apiKey == "" {
		return nil, errors.New("Base URL and API Key are bothrequired to create a Client")
	}

	return &Client{
		Client:  &http.Client{},
		BaseURL: baseURL,
		ApiKey:  apiKey,
	}, nil
}

func (c *Client) SetHTTPClient(client *http.Client) {
	c.Client = client
}
