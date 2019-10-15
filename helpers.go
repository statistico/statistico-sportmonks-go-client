package sportmonks

type (
	AggregatedAssistScorerData struct {
		Data []AssistScorer `json:"data"`
	}

	AggregatedCardScorerData struct {
		Data []CardScorer `json:"data"`
	}

	AggregatedGoalScorerData struct {
		Data []GoalScorer `json:"data"`
	}

	AssistScorerData struct {
		Data []AssistScorer `json:"data"`
	}

	BookmakerOddsData struct {
		Data []BookmakerOdds `json:"data"`
	}

	CardScorerData struct {
		Data []CardScorer `json:"data"`
	}

	CoachData struct {
		Data *Coach `json:"data"`
	}

	CountryData struct {
		Data *Country `json:"data"`
	}

	CountriesData struct {
		Data []Country `json:"data"`
	}

	ContinentData struct {
		Data *Continent `json:"data"`
	}

	FixtureData struct {
		Data *Fixture `json:"data"`
	}

	FixturesData struct {
		Data []Fixture `json:"data"`
	}

	GoalScorerData struct {
		Data []GoalScorer `json:"data"`
	}

	GroupsData struct {
		Data []Group `json:"data"`
	}

	LeagueData struct {
		Data *League `json:"data"`
	}

	LeaguesData struct {
		Data []League `json:"data"`
	}

	MatchTVStations struct {
		Data []TVStation `json:"data"`
	}

	MatchOfficialData struct {
		Data *MatchOfficial `json:"data"`
	}

	OddsData struct {
		Data []Odds `json:"data"`
	}

	PlayerData struct {
		Data *Player `json:"data"`
	}

	PlayersData struct {
		Data []Player `json:"data"`
	}

	PlayerLineupData struct {
		Data []PlayerStats `json:"data"`
	}

	PlayerSeasonStatsData struct {
		Data []PlayerSeasonStats `json:"data"`
	}

	PositionData struct {
		Data *Position `json:"data"`
	}

	RankingData struct {
		Data *Ranking `json:"data"`
	}

	RoundData struct {
		Data *Round `json:"data"`
	}

	RoundsData struct {
		Data []Round `json:"data"`
	}

	SeasonData struct {
		Data *Season `json:"data"`
	}

	SeasonsData struct {
		Data []Season `json:"data"`
	}

	SidelinedData struct {
		Data []Sidelined `json:"data"`
	}

	SquadPlayerData struct {
		Data []SquadPlayer `json:"data"`
	}

	StageData struct {
		Data *Stage `json:"data"`
	}

	StagesData struct {
		Data []Stage `json:"data"`
	}

	TeamSeasonStatsData struct {
		Data *TeamSeasonStats `json:"data"`
	}

	TransfersData struct {
		Data []Transfer `json:"data"`
	}

	TrophyData struct {
		Data []Trophy `json:"data"`
	}

	VenueData struct {
		Data *Venue `json:"data"`
	}
)
