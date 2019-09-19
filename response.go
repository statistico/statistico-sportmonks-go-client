package sportmonks

type (
	FixtureResponse struct {
		Data Fixture `json:"data"`
	}

	// LeaguesResponse
	LeaguesResponse struct {
		Data []League `json:"data"`
		Meta Meta     `json:"meta"`
	}

	// PlayerResponse struct
	PlayerResponse struct {
		Data Player `json:"data"`
	}

	// SeasonsResponse
	SeasonsResponse struct {
		Data []Season `json:"data"`
		Meta Meta     `json:"meta"`
	}

	// SeasonResponse
	SeasonResponse struct {
		Data Season `json:"data"`
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
