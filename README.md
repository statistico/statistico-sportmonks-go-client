# Statistico SportMonks Go Client 

[![Documentation](https://godoc.org/github.com/statistico/statistico-sportmonks-go-client?status.svg)](http://godoc.org/github.com/statistico/statistico-sportmonks-go-client)
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/statistico/statistico-sportmonks-go-client/tree/master.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/statistico/statistico-sportmonks-go-client/tree/master)

This library is a Golang wrapper around version 2.0 of the SportMonks soccer API. Full documentation and API reference can be found here:

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
    
    league, _, err := client.LeagueByID(context.Background(), 10, []string{}) 

    if err != nil {
        fmt.Printf("%s\n", err.Error())
        return
    }

    // Do something with league variable
}
```
The SportMonks soccer API provides powerful 'includes' and 'filtering' features that allow you to enrich data requests. Full
documentation on API flexibility, relationships and includes functionality can be found 
[here](https://www.sportmonks.com/docs/football/2.0/getting-started/a/api-flexibility-and-relationships/88). Instructions
on how to use the filtering parameters and sort functionality can be found [here](https://www.sportmonks.com/docs/football/2.0/getting-started/a/api-filtering-sorting-and-pagination/90).

For more detailed examples on how to use this library please see the [docs](/docs) directory
## Contributing
You are more than welcome to contribute to this project. Fork and make a Pull Request, or create an Issue if you notice 
any problems or would like to suggest improvements.