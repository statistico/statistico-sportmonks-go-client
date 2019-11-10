The SportMonks API provides the functionality to add 'includes' parameters to enrich the data but also
limit the includes data. 

Below is an example of adding fixture data to the season endpoint but limiting the number of fixtures and adding
pagination:

```go
package main

import (
	"context"
	"fmt"
	"github.com/statistico/statistico-sportmonks-go-client"
)

func main() {
	client := sportmonks.NewDefaultHTTPClient("YOUR_TOKEN_GOES_HERE")

	includes := []string{"fixtures:limit(5|1)"}

	season, _, err := client.SeasonByID(context.Background(), 295, includes)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	// Do something with season variable
}
```
