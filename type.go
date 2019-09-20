package statistico

type (
	// Assist Scorer struct
	AssistScorer struct {
		Position int    `json:"position"`
		SeasonID int    `json:"season_id"`
		PlayerID int    `json:"player_id"`
		TeamID   int    `json:"team_id"`
		StageID  int    `json:"stage_id"`
		Assists  int    `json:"assists"`
		Type     string `json:"type"`
	}

	// CardScorer struct
	CardScorer struct {
		Position    int    `json:"position"`
		SeasonID    int    `json:"season_id"`
		PlayerID    int    `json:"player_id"`
		TeamID      int    `json:"team_id"`
		StageID     int    `json:"stage_id"`
		YellowCards int    `json:"yellowcards"`
		RedCards    int    `json:"redcards"`
		Type        string `json:"type"`
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
		Bench                 Lineup             `json:"bench"`
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
		PlayerAssistID   *int    `json:"player_assist_id"`
		PlayerAssistName *string `json:"player_assist_name"`
		Minute           int     `json:"minute"`
		ExtraMinute      *int    `json:"extra_minute"`
		Reason           *string `json:"reason"`
		Result           string  `json:"result"`
	}

	GoalEvents struct {
		Data []GoalEvent `json:"data"`
	}

	GoalScorer struct {
		Position     int    `json:"position"`
		SeasonID     int    `json:"season_id"`
		PlayerID     int    `json:"player_id"`
		TeamID       int    `json:"team_id"`
		StageID      int    `json:"stage_id"`
		Goals        int    `json:"goals"`
		PenaltyGoals int    `json:"penalty_goals"`
		Type         string `json:"type"`
	}

 	Group struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		LeagueID  int    `json:"league_id"`
		SeasonID  int    `json:"season_id"`
		RoundID   int    `json:"round_id"`
		RoundName int    `json:"round_name"`
		StageID   int    `json:"stage_id"`
		StageName string `json:"stage_name"`
		Resource  string `json:"resource"`
 	}

	Lineup struct {
		Data []LineupPlayer `json:"data"`
	}

	LineupPlayer struct {
		TeamID             int         `json:"team_id"`
		FixtureID          int         `json:"fixture_id"`
		PlayerID           int         `json:"player_id"`
		PlayerName         string      `json:"player_name"`
		Number             *int        `json:"number"`
		Position           *string     `json:"position"`
		AdditionalPosition *string     `json:"additional_position"`
		FormationPosition  *int        `json:"formation_position"`
		Posx               *int        `json:"posx"`
		Posy               *int        `json:"posy"`
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
		Fouls          interface{} `json:"fouls"`
		Corners        interface{} `json:"corners"`
		Offsides       interface{} `json:"offsides"`
		Possessiontime interface{} `json:"possessiontime"`
		Yellowcards    interface{} `json:"yellowcards"`
		Redcards       interface{} `json:"redcards"`
		Saves          interface{} `json:"saves"`
		Substitutions  interface{} `json:"substitutions"`
		GoalKick       interface{} `json:"goal_kick"`
		GoalAttempts   interface{} `json:"goal_attempts"`
		FreeKick       interface{} `json:"free_kick"`
		ThrowIn        interface{} `json:"throw_in"`
		BallSafe       interface{} `json:"ball_safe"`
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
