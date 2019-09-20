package sportmonks

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const backOffStart = 1
const backOffExponent = 2

var backOffTotal = 1
var backOffLimit = 120

var clientCreationError = errors.New("base URL and API Key are both required to create a Client")

type ApiClient struct {
	Client  *http.Client
	BaseURL string
	ApiKey  string
	Retries int
}

func NewApiClient(url, key string) (*ApiClient, error) {
	if url == "" || key == "" {
		return nil, clientCreationError
	}

	trans := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 60 * time.Second,
	}

	client := ApiClient{
		Client: &http.Client{
			Transport: trans,
		},
		BaseURL: url,
		ApiKey:  key,
		Retries: 5,
	}

	return &client, nil
}

func (c *ApiClient) SetHTTPClient(client *http.Client) {
	c.Client = client
}

func (c *ApiClient) SetRetries(retries int) {
	c.Retries = retries
}

func (c *ApiClient) handleRequest(path string, includes []string, response interface{}) error {
	built := c.buildUrl(path, includes, 0)

	return c.sendRequest(built, response)
}

func (c *ApiClient) handlePaginatedRequest(path string, includes []string, page int, response interface{}) error {
	built := c.buildUrl(path, includes, page)

	return c.sendRequest(built, response)
}

func (c *ApiClient) sendRequest(url string, response interface{}) error {
	res, err := http.Get(url)

	if err != nil {
		if c.Retries > 0 {
			backOff()
			c.Retries--
			return c.sendRequest(url, response)
		}

		return err
	}

	if err := parseResponseBody(res.Body, &response); err != nil {
		return err
	}

	backOffTotal = backOffStart

	return nil
}

func (c *ApiClient) buildUrl(path string, includes []string, page int) string {
	url := c.BaseURL + path + "?api_token=" + c.ApiKey

	if page > 0 {
		url = url + "&page=" + strconv.Itoa(page)
	}

	if len(includes) > 0 {
		url = url + "&include=" + strings.Join(includes, ",")
	}

	return url
}

func parseResponseBody(body io.ReadCloser, response interface{}) error {
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

func backOff() {
	time.Sleep(time.Duration(backOffTotal))
	backOffTotal = backOffTotal * backOffExponent

	if backOffTotal > backOffLimit {
		backOffTotal = backOffLimit
	}
}
