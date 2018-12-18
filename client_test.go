package sportmonks

import "testing"

var BaseURL = "https://soccer.sportmonks.com"
var ApiKey = "hMNoq0c2fMjipNWEeG7IMmDF9bMNKeVoRi8lJ0qZDhg125U1IormejZKfwua"

func TestNewClient(t *testing.T) {
	c, _ := NewClient(BaseURL, ApiKey)
	//

	if c.BaseURL == "" {
		t.Errorf("Unexpected error %s", c.BaseURL)
	}

	c.Countries()
}
