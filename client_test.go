package sportmonks

import "testing"
import "fmt"

var BaseURL = "https://soccer.sportmonks.com"
var ApiKey = "hMNoq0c2fMjipNWEeG7IMmDF9bMNKeVoRi8lJ0qZDhg125U1IormejZKfwua"

func TestNewClient(t *testing.T) {
	c, _ := NewClient(BaseURL, ApiKey)
	//

	if c.BaseURL == "" {
		t.Errorf("Unexpected error %s", c.BaseURL)
	}

	res, _ := c.Countries(1)

	for i, country := range res.Data {
		if i == 0 {
			fmt.Printf("%+v\n", country)
		}
	}
}
