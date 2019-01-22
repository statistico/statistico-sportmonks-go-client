package sportmonks

import (
	"errors"
	"net/http"
	"io"
	"io/ioutil"
	"encoding/json"
	"strings"
	"strconv"
)

var clientCreationError = errors.New("base URL and API Key are both required to create a Client")

const IncludeLeague = "league"

type Client struct {
	Client   *http.Client
	BaseURL  string
	ApiKey   string
}

func NewClient(baseURL string, apiKey string) (*Client, error) {
	if baseURL == "" || apiKey == "" {
		return &Client{}, clientCreationError
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


func (c *Client) sendRequest(req *http.Request, v interface{}) (error) {
	res, err := c.Client.Do(req)

	if err != nil {
		return err
	}

	if err := parseRequestBody(res.Body, &v); err != nil {
		return err
	}

	return nil
}

func buildRequest(method string, url string, body io.Reader, page int, includes []string) (*http.Request, error) {
	if page > 0 {
		url = url + "&page=" + strconv.Itoa(page)
	}

	if len(includes) > 0 {
		url = url + "&include=" + strings.Join(includes, ",")
	}

	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func parseRequestBody(body io.ReadCloser, v interface{}) error {
	defer body.Close()

	b, err := ioutil.ReadAll(body)

	err = json.Unmarshal([]byte(b), &v)

	if err != nil {
		return err
	}

	return nil
}
