package sportmonks

import (
	"net/http"
)

type (
	// Client represents a SportMonks REST API Client
	Client struct {
		Client *http.Client
		BaseURL string
		ApiKey string
	}
)
