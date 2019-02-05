package sportmonks

type (
	// Continent struct
	Continent struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	// Country struct
	Country struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Extra struct {
			Continent   string      `json:"continent"`
			SubRegion   string      `json:"sub_region"`
			WorldRegion string      `json:"world_region"`
			Fifa        interface{} `json:"fifa,string"`
			ISO         string      `json:"iso"`
			Longitude   string      `json:"longitude"`
			Latitude    string      `json:"latitude"`
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
		ID                    int           `json:"id"`
		LeagueID              int           `json:"league_id"`
		SeasonID              int           `json:"season_id"`
		StageID               *int          `json:"stage_id"`
		RoundID               *int          `json:"round_id"`
		GroupID               *int          `json:"group_id"`
		AggregateID           *int          `json:"aggregate_id"`
		VenueID               *int          `json:"venue_id"`
		RefereeID             *int          `json:"referee_id"`
		LocalteamID           int           `json:"localteam_id"`
		VisitorteamID         int           `json:"visitorteam_id"`
		WeatherReport         WeatherReport `json:"weather_report"`
		Commentaries          *bool         `json:"commentaries"`
		Attendance            *int          `json:"attendance"`
		Pitch                 *string       `json:"pitch"`
		WinningOddsCalculated bool          `json:"winning_odds_calculated"`
		Formations            Formations    `json:"formations"`
		Scores                Scores        `json:"scores"`
		Time                  FixtureTime   `json:"time"`
		Coaches               Coaches       `json:"coaches"`
		Standings             Standings     `json:"standings"`
		Assistants            Assistants    `json:"assistants"`
		Leg                   *string       `json:"leg"`
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

	// Player struct
	Player struct {
		PlayerID     int    `json:"player_id"`
		TeamID       int    `json:"team_id"`
		CountryID    int    `json:"country_id"`
		PositionID   int    `json:"position_id"`
		CommonName   string `json:"common_name"`
		FullName     string `json:"fullname"`
		FirstName    string `json:"firstname"`
		LastName     string `json:"lastname"`
		Nationality  string `json:"nationality"`
		BirthDate    string `json:"birthdate"`
		BirthCountry string `json:"birthcountry"`
		Birthplace   string `json:"birthplace"`
		Height       string `json:"height"`
		Weight       string `json:"weight"`
		ImagePath    string `json:"image_path"`
	}

	// Round struct {
	Round struct {
		ID       int    `json:"id"`
		Name     int    `json:"name"`
		LeagueID int    `json:"league_id"`
		SeasonID int    `json:"season_id"`
		StageID  int    `json:"stage_id"`
		Start    string `json:"start"`
		End      string `json:"end"`
	}

	// Sport struct
	Sport struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Current bool   `json:"current"`
	}

	// Season struct
	Season struct {
		ID              int    `json:"id"`
		Name            string `json:"name"`
		LeagueID        int    `json:"league_id"`
		IsCurrentSeason bool   `json:"is_current_season"`
		CurrentRoundID  *int   `json:"current_round_id"`
		CurrentStageID  *int   `json:"current_stage_id"`
		Fixtures        struct {
			Data []Fixture `json:"data"`
		} `json:"fixtures"`
		Rounds struct {
			Data []Round `json:"data"`
		} `json:"rounds"`
	}

	// Squad struct
	Squad struct {
		Data []SquadPlayer `json:"data"`
	}

	// SquadPlayer struct
	SquadPlayer struct {
		PlayerID           int  `json:"player_id"`
		PositionID         int  `json:"position_id"`
		Number             int  `json:"number"`
		Injured            bool `json:"injured"`
		Minutes            int  `json:"minutes"`
		Appearances        int  `json:"appearences"`
		Lineups            int  `json:"lineups"`
		SubstituteIn       int  `json:"substitute_in"`
		SubstituteOut      int  `json:"substitute_out"`
		SubstitutesOnBench int  `json:"substitutes_on_bench"`
		Goals              int  `json:"goals"`
		Assists            int  `json:"assists"`
		Yellowcards        int  `json:"yellowcards"`
		Yellowred          int  `json:"yellowred"`
		Redcards           int  `json:"redcards"`
	}

	// Team struct
	Team struct {
		ID           int     `json:"id"`
		LegacyID     int     `json:"legacy_id"`
		Name         string  `json:"name"`
		ShortCode    string  `json:"short_code"`
		Twitter      *string `json:"twitter"`
		CountryID    int     `json:"country_id"`
		NationalTeam bool    `json:"national_team"`
		Founded      int     `json:"founded"`
		LogoPath     *string `json:"logo_path"`
		VenueID      int     `json:"venue_id"`
	}

	// Venue struct
	Venue struct {
		ID       int     `json:"id"`
		Name     string  `json:"name"`
		Surface  string  `json:"surface"`
		Address  *string `json:"address"`
		City     string  `json:"city"`
		Capacity int     `json:"capacity"`
		Image    string  `json:"image_path"`
	}
)
