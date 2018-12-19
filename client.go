package sportmonks

import (
	"errors"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type Client struct {
	Client   *http.Client
	BaseURL  string
	ApiKey   string
}

func NewClient(baseURL string, apiKey string) (*Client, error) {
	if baseURL == "" || apiKey == "" {
		return &Client{}, errors.New("base URL and API Key are both required to create a Client")
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

func (c *Client) doRequest(method string, url string, body interface{}) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	res, err := c.Client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// Add includes struct and pull body parsing into struct into own private method
func (c *Client) Countries(page int) (*CountryResponse, error) {
	url := c.BaseURL + "/api/v2.0/countries?api_token=" + c.ApiKey + "&include=continent,leagues"

	res, err := c.doRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	countries := new(CountryResponse)

	body, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal([]byte(body), &countries)

	if err != nil {
		return nil, errors.New("error when parsing request body")
	}

	return countries, err
}
