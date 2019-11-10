The SportMonks API provides the functionality to add 'includes' parameters to enrich the data returned
by a specific endpoint but also allows nested includes. 

Below is an example of retrieving a fixture by ID  with added includes data for home team and away team including team 
specific venue data for each team:

```go
package main

import (
	"context"
	"fmt"
	"github.com/statistico/statistico-sportmonks-go-client"
)

func main() {
	client := sportmonks.NewDefaultHTTPClient("YOUR_TOKEN_GOES_HERE")

	includes := []string{"localTeam.venue", "visitorTeam.venue"}
	filters := map[string][]int{}

	fixture, _, err := client.FixtureByID(context.Background(), 10, includes, filters)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	// Do something with fixture variable
}
```
