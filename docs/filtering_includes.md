The SportMonks API provides the functionality to add 'includes' parameters to enrich the data but also
provides additional filters for the enriched data. 

Below is an example of adding odds data to the livescores endpoint but filtering the odds data to Bookmaker with ID number 2:

```go
package main

import (
    "context"
    "fmt"
    "github.com/statistico/statistico-sportmonks-go-client"
)

func main() {
    client := sportmonks.NewDefaultHTTPClient("YOUR_TOKEN_GOES_HERE")
    
    includes := []string{"flatOdds:filter(bookmaker_id|2)"}
    filters := make(map[string][]int)
    
    fixtures, _, err := client.LiveScores(context.Background(), 1, includes, filters) 

    if err != nil {
        fmt.Printf("%s\n", err)
        return
    }
    
    for _, fixture := range fixtures {
        // Do something with fixture variable   
    }
}

```
