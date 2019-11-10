# Statistico SportMonks Go Client 

[![Documentation](https://godoc.org/github.com/statistico/statistico-sportmonks-go-client?status.svg)](http://godoc.org/github.com/statistico/statistico-sportmonks-go-client)
[![CircleCI](https://circleci.com/gh/statistico/statistico-sportmonks-go-client/tree/master.svg?style=shield)](https://circleci.com/gh/statistico/statistico-sportmonks-go-client/tree/master)

This library is a Golang wrapper around version 2.0 of the SportMonks soccer API. 

Full documentation and API reference can be found here:

[Documentation](https://www.sportmonks.com/docs/football/2.0/prologue/a/introduction/94)

[API Reference](https://docs.sportmonks.com/football)

## Installation
```.env
$ go get -u github.com/statistico/statistico-sportmonks-go-client
```
## Usage
To instantiate the required HTTPClient struct and retrieve a single League resource:
```go
package main

import (
    "context"
    "fmt"
    "github.com/statistico/statistico-sportmonks-go-client"
)

func main() {
    client := sportmonks.NewDefaultHTTPClient("YOUR_TOKEN_GOES_HERE")
    
    league, _, err := client.LeagueByID(context.Background(), 10, []string{"seasons"}) 

    if err != nil {
        fmt.Printf("%s\n", err)
        return
    }

    fmt.Printf("ID: %d, Name: %s", league.ID, league.Name)
}
```
The SportMonks soccer API provides powerful 'includes' and 'filtering' features that allow you to enrich data requests. Full
documentation on API flexibility, relationships and includes functionality can be found 
[here](https://www.sportmonks.com/docs/football/2.0/getting-started/a/api-flexibility-and-relationships/88). Instructions
on how to use the filtering and sort functionality can be found [here](https://www.sportmonks.com/docs/football/2.0/getting-started/a/api-filtering-sorting-and-pagination/90)

Example usage of adding 'includes' parameters to include team data when fetching fixture data
```go
package main

import (
    "context"
    "fmt"
    "github.com/statistico/statistico-sportmonks-go-client"
)

func main() {
    client := sportmonks.NewDefaultHTTPClient("YOUR_TOKEN_GOES_HERE")
    
    includes := []string{"localTeam", "visitorTeam"}
    filters := map[string][]int{}
    
    fixture, _, err := client.FixtureByID(context.Background(), 10, includes, filters) 

    if err != nil {
        fmt.Printf("%s\n", err)
        return
    }

   // Do something with fixture variable
}
```
Example usage of adding filter parameters to filter data when fetching fixture data for a specific date:
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
Example usage of adding filtering bookmaker data when fetching live score information:
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

## Contributing
You are more than welcome to contribute to this project. Fork and make a Pull Request, or create an Issue if you notice 
any problems or would like to suggest improvements.