The SportMonks API provides the functionality to add 'filter' parameters to filter the data returned
by a specific endpoint. 

Below is an example of retrieving fixtures by a specific date but reducing the returned data to specific leagues:

```go
package main

import (
	"context"
	"fmt"
	"github.com/statistico/statistico-sportmonks-go-client"
	"time"
)

const dateFormat = "2006-01-02"

func main() {
	client := sportmonks.NewDefaultHTTPClient("YOUR_TOKEN_GOES_HERE")

	var includes []string
	filters := make(map[string][]int)
	filters["leagues"] = []int{8, 10}

	date, _ := time.Parse(dateFormat, "2019-03-12")

	fixtures, _, err := client.FixturesByDate(context.Background(), date, includes, filters)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	for _, fixture := range fixtures {
		// Do something with fixture variable
	}
}
```
