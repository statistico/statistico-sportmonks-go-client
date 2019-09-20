package sportmonks

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

	PlayerResponse struct {
		Data Player `json:"data"`
	}

 	SeasonsResponse struct {
		Data []Season `json:"data"`
		Meta Meta     `json:"meta"`
	}

 	SeasonResponse struct {
		Data Season `json:"data"`
		Meta Meta     `json:"meta"`
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
