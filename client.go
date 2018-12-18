package sportmonks

import (
	"errors"
	"net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
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

func (c *Client) Countries() (*CountryResponse, error) {
	url := c.BaseURL + "/api/v2.0/countries?api_token=" + c.ApiKey

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	res, err := c.Client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	r := new(ResponseMetaData)

	err = json.Unmarshal([]byte(body), &r)

	if err != nil {
		return nil, errors.New("error when parsing request body")
	}

	countries := new(CountryResponse)

	err = json.Unmarshal([]byte(body), &countries)

	if err != nil {
		return nil, errors.New("error when parsing request body")
	}

	for _, country := range countries.CountryList {
		fmt.Printf("%+v\n", country)
	}

	if r.Meta.Pagination.Count > 0 {
		page := r.Meta.Pagination.CurrentPage
		total := r.Meta.Pagination.TotalPages

		for i := page; i <= total; i++ {
			fmt.Printf("%+v\n", url + "&Page=" + strconv.Itoa(i))

			req, err := http.NewRequest("GET", url + "&page=" + strconv.Itoa(i), nil)

			if err != nil {
				return nil, err
			}

			res, err := c.Client.Do(req)

			if err != nil {
				return nil, err
			}

			body, err := ioutil.ReadAll(res.Body)

			err = json.Unmarshal([]byte(body), &countries)

			for _, country := range countries.CountryList {
				fmt.Printf("%+v\n", country)
			}
		}
	}

	return countries, err
}