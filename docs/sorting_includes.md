The SportMonks API provides the functionality to add 'includes' parameters to enrich the data but also
sort the includes data. 

Below is an example of adding fixture data to the season endpoint but sorting the fixtures by created at in 
ascending order:

```go
package main

import (
	"context"
	"fmt"
	"github.com/statistico/statistico-sportmonks-go-client"
)

func main() {
	client := sportmonks.NewDefaultHTTPClient("YOUR_TOKEN_GOES_HERE")

	includes := []string{"fixtures:order(starting_at|asc)"}

	season, _, err := client.SeasonByID(context.Background(), 100, includes)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	// Do something with season variable
}
```
