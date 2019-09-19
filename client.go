package sportmonks_client

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
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
	Logger  *log.Logger
}

func NewClient(baseURL, apiKey string, logger *log.Logger) (*Client, error) {
	if baseURL == "" || apiKey == "" {
		return &Client{}, clientCreationError
	}

	trans := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 60 * time.Second,
	}

	return &Client{
		Client: &http.Client{
			Transport: trans,
		},
		BaseURL: baseURL,
		ApiKey:  apiKey,
		Logger:  logger,
	}, nil
}

func (c *Client) SetHTTPClient(client *http.Client) {
	c.Client = client
}

func (c *Client) sendRequest(req *http.Request, v interface{}, retries int) error {
	res, err := c.Client.Do(req)

	if err != nil {
		if retries > 0 {
			backOff()
			return c.sendRequest(req, v, retries-1)
		}
		c.Logger.Printf("Error when sending request in client: Message %s", err.Error())
		return err
	}

	if err := parseRequestBody(res.Body, &v); err != nil {
		c.Logger.Printf("Error when parsing response in client: Message %s", err.Error())
		return err
	}

	resetBackOff()

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

	req.Close = true

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

func backOff() {
	time.Sleep(time.Duration(backOffTotal))
	backOffTotal = backOffTotal * backOffExponent

	if backOffTotal > backOffLimit {
		backOffTotal = backOffLimit
	}
}

func resetBackOff() {
	backOffTotal = backOffStart
}
