package sportmonks

import "fmt"

const fixtureOddsUri = "/api/v2.0/odds/fixture"
const inPlayFixtureOddsUri = "/api/v2.0/odds/inplay/fixture"

type (
	MatchOdds struct {
		ID    int `json:"id"`
		Name string `json:"name"`
		Suspended int `json:"suspended"`
		Bookmaker struct {
			Data [] BookmakerOdds`json:"data"`
		} `json:"bookmaker"`
	}

	BookmakerOdds struct {
		ID    int `json:"id"`
		Name string `json:"name"`
		Bookmaker struct {
			Data []Odds `json:"data"`
		} `json:"odds"`
	}

	Odds struct {
		Value            string      `json:"value"`
		Handicap         *string `json:"handicap"`
		Total            string      `json:"total"`
		Label            string      `json:"label"`
		Probability      string      `json:"probability"`
		Dp3              string      `json:"dp3"`
		American         int         `json:"american"`
		Fractional       *string `json:"factional"`
		Winning          *string `json:"winning"`
		Stop             bool        `json:"stop"`
		BookmakerEventID int         `json:"bookmaker_event_id"`
		LastUpdate       struct {
			Date         string `json:"date"`
			TimezoneType int    `json:"timezone_type"`
			Timezone     string `json:"timezone"`
		} `json:"last_update"`
	}
)


// Retrieve a multiple match odds resources by fixture ID.
func (c *ApiClient) OddsByFixtureId(fixtureId int) ([]MatchOdds, *Meta, error) {
	url := fmt.Sprintf(fixtureOddsUri + "/%d", fixtureId)

	response := new(MatchOddsResponse)

	err := c.handleRequest(url, []string{}, response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, &response.Meta, err
}

// Retrieve a multiple match odds resources by fixture ID.
func (c *ApiClient) InPlayOddsByFixtureId(fixtureId int) ([]MatchOdds, *Meta, error) {
	url := fmt.Sprintf(inPlayFixtureOddsUri + "/%d", fixtureId)

	response := new(MatchOddsResponse)

	err := c.handleRequest(url, []string{}, response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, &response.Meta, err
}

// Retrieve a multiple match odds resources by fixture and bookmaker ID.
func (c *ApiClient) OddsByFixtureAndBookmaker(fixtureId, bookmakerId int) ([]MatchOdds, *Meta, error) {
	url := fmt.Sprintf(fixtureOddsUri + "/%d/bookmaker/%d", fixtureId, bookmakerId)

	response := new(MatchOddsResponse)

	err := c.handleRequest(url, []string{}, response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, &response.Meta, err
}

// Retrieve a multiple match odds resources by fixture and bookmaker ID.
func (c *ApiClient) OddsByFixtureAndMarket(fixtureId, marketId int) ([]MatchOdds, *Meta, error) {
	url := fmt.Sprintf(fixtureOddsUri + "/%d/market/%d", fixtureId, marketId)

	response := new(MatchOddsResponse)

	err := c.handleRequest(url, []string{}, response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, &response.Meta, err
}