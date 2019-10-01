package sportmonks

type (
	AggregatedTopScorersResponse struct {
		Data AggregatedTopScorers `json:"data"`
		Meta Meta     `json:"meta"`
	}

	BookmakerResponse struct {
		Data Bookmaker `json:"data"`
		Meta Meta `json:"meta"`
	}

	BookmakersResponse struct {
		Data []Bookmaker `json:"data"`
		Meta Meta `json:"meta"`
	}

	CoachResponse struct {
		Data Coach `json:"data"`
		Meta Meta `json:"meta"`
	}

	CommentariesResponse struct {
		Data []Commentary `json:"data"`
		Meta Meta `json:"meta"`
	}

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
		Meta Meta       `json:"meta"`
	}

	FixturesResponse struct {
		Data []Fixture `json:"data"`
		Meta Meta       `json:"meta"`
	}

	LeaguesResponse struct {
		Data []League `json:"data"`
		Meta Meta     `json:"meta"`
	}

 	LeagueResponse struct {
		Data League `json:"data"`
		Meta Meta     `json:"meta"`
	}

	LeagueStandingsResponse struct {
		Data LeagueStandings `json:"data"`
		Meta Meta     `json:"meta"`
	}

	MarketResponse struct {
		Data Market `json:"data"`
		Meta Meta     `json:"meta"`
	}

	MarketsResponse struct {
		Data []Market `json:"data"`
		Meta Meta     `json:"meta"`
	}

	MatchOddsResponse struct {
		Data []MatchOdds `json:"data"`
		Meta Meta     `json:"meta"`
	}

	PlayerResponse struct {
		Data Player `json:"data"`
		Meta Meta     `json:"meta"`
	}

	RoundResponse struct {
		Data Round `json:"data"`
		Meta Meta     `json:"meta"`
	}

	RoundsResponse struct {
		Data []Round `json:"data"`
		Meta Meta     `json:"meta"`
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

	StageResponse struct {
		Data Stage `json:"data"`
		Meta Meta     `json:"meta"`
	}

	StagesSeasonResponse struct {
		Data []Stage `json:"data"`
		Meta Meta     `json:"meta"`
	}

	TeamsResponse struct {
		Data []Team `json:"data"`
		Meta Meta     `json:"meta"`
	}

	TeamResponse struct {
		Data Team `json:"data"`
		Meta Meta     `json:"meta"`
	}

	TeamSquadResponse struct {
		Data []PlayerSeasonStats `json:"data"`
		Meta Meta     `json:"meta"`
	}

	TopScorersResponse struct {
		Data TopScorers `json:"data"`
		Meta Meta     `json:"meta"`
	}

	TVStationResponse struct {
		Data []TVStation `json:"data"`
		Meta Meta     `json:"meta"`
	}

	VenueResponse struct {
		Data Venue    `json:"data"`
		Meta Meta     `json:"meta"`
	}

	VenuesResponse struct {
		Data []Venue `json:"data"`
		Meta Meta     `json:"meta"`
	}

	VideoHighlightsResponse struct {
		Data []VideoHighlights `json:"data"`
		Meta Meta     `json:"meta"`
	}
)
