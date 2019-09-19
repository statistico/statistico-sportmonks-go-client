package sportmonks_client

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	base := "https://example.com"
	api := "api-key"
	c, _ := NewClient(base, api)

	uri := c.BaseURL
	key := c.ApiKey

	if uri != base {
		t.Errorf("expected %s got %s", base, uri)
	}

	if key != api {
		t.Errorf("expected %s got %s", api, key)
	}
}
