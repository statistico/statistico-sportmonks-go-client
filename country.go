package statistico

import (
	"encoding/json"
	"strconv"
)

const countryUri = "/api/v2.0/countries"

// Country struct
type Country struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Extra struct {
		Continent   string      `json:"continent"`
		SubRegion   string      `json:"sub_region"`
		WorldRegion string      `json:"world_region"`
		FIFA        json.Number  `json:", string"`
		ISO         string      `json:"iso, string"`
		Longitude   string      `json:"longitude"`
		Latitude    string      `json:"latitude"`
	} `json:"extra, omitempty"`
	Continent struct {
		Data *Continent `json:"data"`
	} `json:"continent, omitempty"`
	Leagues struct {
		Data []League `json:"data"`
	} `json:"leagues, omitempty"`
}

// Make a request to retrieve multiple country resources. The request endpoint executed within this method
// is paginated, the first argument to this method allows the consumer to specify a page to request.
// Use the includes slice to enrich the response data.
func (c *SportMonksClient) Countries(page int, includes []string, ) ([]Country, *Meta, error) {
	response := new(CountriesResponse)

	err := c.handlePaginatedRequest(countryUri, includes, page, response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, &response.Meta, err
}

// Retrieve a single country resource by ID. Use the includes slice to enrich the response data.
func (c *SportMonksClient) CountryById(id int, includes []string) (*Country, *Meta, error) {
	url := continentUri + "/" + strconv.Itoa(id)

	response := new(CountryResponse)

	err := c.handleRequest(url, includes, response)

	if err != nil {
		return nil, nil, err
	}

	return &response.Data, &response.Meta, err
}
