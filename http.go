package sportmonks

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	defaultBaseUrl         = "https://soccer.sportmonks.com/api/v2.0"
	bookmakersUri          = "/bookmakers"
	coachesUri             = "/coaches"
	commentariesFixtureUri = "/commentaries/fixture"
	continentsUri          = "/continents"
	countriesUri           = "/countries"
	fixturesUri            = "/fixtures"
	fixturesDateUri        = "/fixtures/date"
	fixturesBetweenUri     = "/fixtures/between"
	fixturesLastUpdatedUri = "/fixtures/updates"
	fixturesMultiUri       = "/fixtures/multi"
	headToHeadUri          = "/head2head"
	inPlayOddsFixtureUri   = "/odds/inplay/fixture"
	leagueStandingsUri     = "/standings/season"
	leaguesUri             = "/leagues"
	liveLeagueStandingsUri = "/standings/season/live"
	liveScoresUri          = "/livescores"
	liveScoresInPlayUri    = "/livescores/now"
	marketsUri             = "/markets"
	oddsFixtureUri         = "/odds/fixture"
	playersUri             = "/players"
	roundsUri              = "/rounds"
	roundsSeasonUri        = "/rounds/season"
	seasonsUri             = "/seasons"
	stagesUri              = "/stages"
	stagesSeasonUri        = "/stages/season"
	teamSquadUri           = "/squad/season"
	teamsUri               = "/teams"
	teamsSeasonUri         = "/teams/season"
	topScorersSeasonUri    = "/topscorers/season"
	tvStationsUri          = "/tvstations/fixture"
	venuesUri              = "/venues"
	venuesSeasonUri        = "/venues/season"
	videoHighlightsUri     = "/highlights"
)

// Client is a HTTP request builder and sender
type HTTPClient struct {
	HTTPClient *http.Client
	BaseURL    string
	Key        string
}

// NewClient creates a new Client with default settings. A key is required to instantiate the Client
func NewHTTPClient(key string) *HTTPClient {
	return &HTTPClient{
		HTTPClient: &http.Client{},
		BaseURL:    defaultBaseUrl,
		Key:        key,
	}
}

func (c *HTTPClient) getResource(ctx context.Context, url string, query url.Values, response interface{}) error {
	req, err := http.NewRequest(http.MethodGet, c.BaseURL+url, nil)

	if err != nil {
		return err
	}

	query.Set("api_token", c.Key)

	req.URL.RawQuery = query.Encode()

	return c.do(ctx, req, response)
}

func (c *HTTPClient) do(ctx context.Context, req *http.Request, intf interface{}) error {
	req = req.WithContext(ctx)
	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if err = checkStatusCode(resp); err != nil {
		return err
	}

	return parseJsonResponseBody(resp.Body, intf)
}

func checkStatusCode(resp *http.Response) error {
	if resp.StatusCode != http.StatusOK {
		err := new(ErrBadStatusCode)

		e := parseJsonResponseBody(resp.Body, &err)

		if e != nil {
			return e
		}

		return err
	}

	return nil
}

func parseJsonResponseBody(body io.ReadCloser, response interface{}) error {
	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return err
	}

	return nil
}
