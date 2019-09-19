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

type Client struct {
	Client  *http.Client
	BaseURL string
	ApiKey  string
	Retries int
}

func NewClient(url, key string, retries int) (*Client, error) {
	if url == "" || key == "" {
		return &Client{}, clientCreationError
	}

	trans := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 60 * time.Second,
	}

	client := Client{
		Client: &http.Client{
			Transport: trans,
		},
		BaseURL: url,
		ApiKey:  key,
		Retries: retries,
	}

	return &client, nil
}

func (c *Client) SetHTTPClient(client *http.Client) {
	c.Client = client
}

func (c *Client) SetRetries(retries int) {
	c.Retries = retries
}

func (c *Client) sendRequest(url string, includes []string, page int, response interface{}) error {
	uri := buildUrl(url, includes, page)

	res, err := http.Get(url)

	if err != nil {
		if c.Retries > 0 {
			backOff()
			c.Retries--
			return c.sendRequest(uri, includes, page, response)
		}

		return err
	}

	if err := parseRequestBody(res.Body, &response); err != nil {
		return err
	}

	backOffTotal = backOffStart

	return nil
}

func buildUrl(url string, includes []string, page int) string {
	if page > 0 {
		url = url + "&page=" + strconv.Itoa(page)
	}

	if len(includes) > 0 {
		url = url + "&include=" + strings.Join(includes, ",")
	}

	return url
}

func parseRequestBody(body io.ReadCloser, response interface{}) error {
	defer body.Close()

	b, err := ioutil.ReadAll(body)

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
