package sportmonks

import (
	"fmt"
	"testing"
)

func TestContinents(t *testing.T) {
	client, _ := NewApiClient(
		"https://soccer.sportmonks.com",
		"M8ouHby0xb5ADsIji9EXlg3wQQqIXH1OQaYGiqZ84o0u7sehm2rnS08Y99Ag",
		)

	//includes := []string{"countries"}

	continents, _, err := client.ContinentById(1, []string{"countries"})

	if err != nil {
		fmt.Printf("Error when calling continents %s", err.Error())
	}

	//for _, c := range continents {
	//	fmt.Printf("Continent %+v\n\n", c)
	//}

	fmt.Printf("Continent %+v\n\n", continents)


	fmt.Print("Done")
}
