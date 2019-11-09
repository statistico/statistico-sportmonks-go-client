# Statistico SportMonks Go Client 

[![Documentation](https://godoc.org/github.com/statistico/statistico-sportmonks-go-client?status.svg)](http://godoc.org/github.com/statistico/statistico-sportmonks-go-client)
[![CircleCI](https://circleci.com/gh/statistico/statistico-sportmonks-go-client/tree/master.svg?style=shield)](https://circleci.com/gh/statistico/statistico-sportmonks-go-client/tree/master)

This library is a Golang wrapper around version 2.0 of the SportMonks soccer API. Full documentation and API reference can be found here:

[Documentation](https://www.sportmonks.com/docs/football/2.0/prologue/a/introduction/94)

[API Reference](https://docs.sportmonks.com/football)

## Installation
```.env
$ go get -u github.com/statistico/statistico-sportmonks-go-client
```
## Usage
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
The above example displays basic usage of client creation and a single league resource retrieval. The SportMonks API provides
'includes' and 'filtering' functionality to enrich data request, detailed examples of making these 
requests can be found in the **EXAMPLES....**

## Contributing
You are more than welcome to contribute to this project. Fork and make a Pull Request, or create an Issue if you notice 
any problems or would like to suggest improvements.