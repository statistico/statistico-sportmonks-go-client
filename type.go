package sportmonks

type (
	// Continent struct
	Continent struct {
		ID int `json:"id"`
		Name string `json:"name"`
	}

	// Country struct
	Country struct {
		ID int `json:"id"`
		Name string `json:"name"`
		Extra struct {
			Continent string `json:"continent"`
			SubRegion string `json:"sub_region"`
			WorldRegion string `json:"world_region"`
			Fifa interface{} `json:"fifa,string"`
			ISO string `json:"iso"`
			Longitude string `json:"longitude"`
			Latitude string `json:"latitude"`
		} `json:"extra"`
		Continent struct {
			Data Continent `json:"data"`
		} `json:"continent, omitempty"`
		Leagues struct {
			Data []League `json:"data"`
		} `json:"leagues, omitempty"`
	}

	// Fixture struct
	Fixture struct {
		ID  			int        `json:"id"`
		LeagueID        int        `json:"league_id"`
		SeasonID        int        `json:"season_id"`
		StageID         *int       `json:"stage_id"`
		RoundID         *int       `json:"round_id"`
		GroupID         *int       `json:"group_id"`
		AggregateID     *int       `json:"aggregate_id"`
		VenueID         int        `json:"venue_id"`
		RefereeID       *int       `json:"referee_id"`
		LocalTeamID     int        `json:"localteam_id"`
		VisitorTeamID   int        `json:"visitorteam_id"`
		Attendance      *int       `json:"attendance"`
		Pitch           *string    `json:"pitch"`
		Formations *struct {
			LocalTeam      *string     `json:"localteam_formation"`
			VisitorTeam    *string     `json:"visitorteam_formation"`
		} `json:"formations"`
		Scores *struct {
			LocalTeamScore    *int     `json:"localteam_score"`
			VisitorTeamScore  *int     `json:"visitorteam_score"`
			LocalPenScore     *int     `json:"localteam_pen_score"`
			VisitorPenScore   *int     `json:"visitorteam_pen_score"`
			HalfTimeScore     *string  `json:"ht_score"`
			FullTimeScore     *string  `json:"ft_score"`
			ExtraTimeScore    *string  `json:"et_score"`
		} `json:"scores"`
		Minute          *int       `json:"minute"`
		Second          *int       `json:"second"`
		AddedTime       *int       `json:"added_time"`
		ExtraMinute     *int       `json:"extra_minute"`
		InjuryTime      *int       `json:"injury_time"`
		Standings *struct {
			LocalTeam      *int    `json:"localteam_position"`
			VisitorTeam    *int    `json:"visitorteam_position"`
		} `json:"standings"`
	}

	// League struct
	League struct {
		ID              int    `json:"id"`
		LegacyID        int    `json:"legacy_id"`
		CountryID       int    `json:"country_id"`
		Name            string `json:"name"`
		IsCup           bool   `json:"is_cup"`
		CurrentSeasonID int    `json:"current_season_id"`
		CurrentRoundID  int    `json:"current_round_id"`
		CurrentStageID  int    `json:"current_stage_id"`
		LiveStandings   bool   `json:"live_standings"`
		Coverage        struct {
			TopscorerGoals   bool `json:"topscorer_goals"`
			TopscorerAssists bool `json:"topscorer_assists"`
			TopscorerCards   bool `json:"topscorer_cards"`
		} `json:"coverage"`
		Seasons struct {
			Data []Season `json:"data"`
		} `json:"seasons"`
	}

	// Season struct
	Season struct {
		ID              int          `json:"id"`
		Name            string       `json:"name"`
		LeagueID        int          `json:"league_id"`
		CurrentSeason   bool         `json:"is_current_season"`
		CurrentRoundID  int          `json:"current_round_id"`
		CurrentStageID  int          `json:"current_stage_id"`
		Fixtures struct {
			Data []Fixture `json:"data"`
		} `json:"fixtures, omitempty"`
		Results struct {
			Data []Fixture `json:"data"`
		} `json:"results, omitempty"`
	}

	// Sport struct
	Sport struct {
		ID int `json:"id"`
		Name string `json:"name"`
		Current bool `json:"current"`
	}
)
