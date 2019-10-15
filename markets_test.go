package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var marketsResponse = `{
	"data": [
		{
		  "id": 1,
		  "name": "3Way Result"
		}
	]
}`

var marketResponse = `{
	"data": {
		"id": 1,
		"name": "3Way Result"
	}
}`

func TestMarkets(t *testing.T) {
	t.Run("returns a slice of Market struct", func(t *testing.T) {
		url := ddefaultBaseURL + "/markets?api_token=api-key"

		server := mockResponseServer(t, marketsResponse, 200, url)

		client := newTestHTTPClient(server)

		markets, _, err := client.Markets(context.Background())

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertMarket(t, &markets[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := ddefaultBaseURL + "/markets?api_token=api-key"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		markets, _, err := client.Markets(context.Background())

		if markets != nil {
			t.Fatalf("Test failed, expected nil, got %+v", markets)
		}

		assertError(t, err)
	})
}

func TestMarketByID(t *testing.T) {
	t.Run("returns a single Market struct", func(t *testing.T) {
		url := ddefaultBaseURL + "/markets/1?api_token=api-key"

		server := mockResponseServer(t, marketResponse, 200, url)

		client := newTestHTTPClient(server)

		market, _, err := client.MarketByID(context.Background(), 1)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertMarket(t, market)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := ddefaultBaseURL + "/markets/1?api_token=api-key"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		market, _, err := client.MarketByID(context.Background(), 1)

		if market != nil {
			t.Fatalf("Test failed, expected nil, got %+v", market)
		}

		assertError(t, err)
	})
}

func assertMarket(t *testing.T, market *Market) {
	assert.Equal(t, 1, market.ID)
	assert.Equal(t, "3Way Result", market.Name)
}
