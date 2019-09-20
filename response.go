package statistico

type (
	ContinentsResponse struct {
		Data []Continent `json:"data"`
		Meta Meta `json:"meta"`
	}

	ContinentResponse struct {
		Data Continent `json:"data"`
		Meta Meta `json:"meta"`
	}

	CountriesResponse struct {
		Data []Country `json:"data"`
		Meta Meta      `json:"meta"`
	}

 	CountryResponse struct {
		Data Country 	`json:"data"`
		Meta Meta       `json:"meta"`
	}

	FixtureResponse struct {
		Data Fixture `json:"data"`
	}

	LeaguesResponse struct {
		Data []League `json:"data"`
		Meta Meta     `json:"meta"`
	}

 	LeagueResponse struct {
		Data League `json:"data"`
		Meta Meta     `json:"meta"`
	}

	// PlayerResponse struct
	PlayerResponse struct {
		Data Player `json:"data"`
	}

	//SquadResponse
	SquadResponse struct {
		Data []SquadPlayer `json:"data"`
	}

	// TeamsResponse
	TeamsResponse struct {
		Data []Team `json:"data"`
	}

	// VenueResponse
	VenuesResponse struct {
		Data []Venue `json:"data"`
	}
)
