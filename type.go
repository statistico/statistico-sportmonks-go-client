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
		ID                    int                `json:"id"`
		LeagueID              int                `json:"league_id"`
		SeasonID              int                `json:"season_id"`
		StageID               *int               `json:"stage_id"`
		RoundID               *int               `json:"round_id"`
		GroupID               *int               `json:"group_id"`
		AggregateID           *int               `json:"aggregate_id"`
		VenueID               *int               `json:"venue_id"`
		RefereeID             *int               `json:"referee_id"`
		LocalTeamID           int                `json:"localteam_id"`
		VisitorTeamID         int                `json:"visitorteam_id"`
		WeatherReport         WeatherReport      `json:"weather_report"`
		Commentaries          *bool              `json:"commentaries"`
		Attendance            *int               `json:"attendance"`
		Pitch                 *string            `json:"pitch"`
		WinningOddsCalculated bool               `json:"winning_odds_calculated"`
		Formations            Formations         `json:"formations"`
		Scores                Scores             `json:"scores"`
		Time                  FixtureTime        `json:"time"`
		Coaches               Coaches            `json:"coaches"`
		Standings             Standings          `json:"standings"`
		Assistants            Assistants         `json:"assistants"`
		Leg                   *string            `json:"leg"`
		Lineup                Lineup             `json:"lineup"`
		TeamStats             TeamsStats         `json:"stats"`
		Goals                 GoalEvents         `json:"goals"`
		Subs                  SubstitutionEvents `json:"substitutions"`
	}

	GoalEvent struct {
		ID               int     `json:"id"`
		TeamID           string  `json:"team_id"`
		Type             string  `json:"type"`
		FixtureID        int     `json:"fixture_id"`
		PlayerID         int     `json:"player_id"`
		PlayerName       string  `json:"player_name"`
		PlayerAssistID   int     `json:"player_assist_id"`
		PlayerAssistName string  `json:"player_assist_name"`
		Minute           int     `json:"minute"`
		ExtraMinute      *int    `json:"extra_minute"`
		Reason           *string `json:"reason"`
		Result           string  `json:"result"`
	}

	GoalEvents struct {
		Data []GoalEvent `json:"data"`
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

	Lineup struct {
		Data []LineupPlayer `json:"data"`
	}

	LineupPlayer struct {
		TeamID             int         `json:"team_id"`
		FixtureID          int         `json:"fixture_id"`
		PlayerID           int         `json:"player_id"`
		PlayerName         string      `json:"player_name"`
		Number             int         `json:"number"`
		Position           string      `json:"position"`
		AdditionalPosition interface{} `json:"additional_position"`
		FormationPosition  int         `json:"formation_position"`
		Posx               interface{} `json:"posx"`
		Posy               interface{} `json:"posy"`
		Stats              PlayerStats `json:"stats"`
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

	SubstitutionEvent struct {
		ID            int    `json:"id"`
		TeamID        string `json:"team_id"`
		Type          string `json:"type"`
		FixtureID     int    `json:"fixture_id"`
		PlayerInID    int    `json:"player_in_id"`
		PlayerInName  string `json:"player_in_name"`
		PlayerOutID   int    `json:"player_out_id"`
		PlayerOutName string `json:"player_out_name"`
		Minute        int    `json:"minute"`
		ExtraMinute   *int   `json:"extra_minute"`
		Injured       *bool  `json:"injuried"`
	}

	SubstitutionEvents struct {
		Data []SubstitutionEvent `json:"data"`
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

	TeamStats struct {
		TeamID         int         `json:"team_id"`
		FixtureID      int         `json:"fixture_id"`
		Shots          TeamShots   `json:"shots"`
		Passes         TeamPasses  `json:"passes"`
		Attacks        TeamAttacks `json:"attacks"`
		Fouls          *int        `json:"fouls"`
		Corners        *int        `json:"corners"`
		Offsides       *int        `json:"offsides"`
		Possessiontime *int        `json:"possessiontime"`
		Yellowcards    *int        `json:"yellowcards"`
		Redcards       *int        `json:"redcards"`
		Saves          *int        `json:"saves"`
		Substitutions  *int        `json:"substitutions"`
		GoalKick       *int        `json:"goal_kick"`
		GoalAttempts   *int        `json:"goal_attempts"`
		FreeKick       *int        `json:"free_kick"`
		ThrowIn        *int        `json:"throw_in"`
		BallSafe       *int        `json:"ball_safe"`
	}

	TeamsStats struct {
		Data []TeamStats `json:"data"`
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
