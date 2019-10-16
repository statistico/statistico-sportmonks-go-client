package sportmonks

type (
	// AdditionalPlayerMatchStats provides additional stats information
	AdditionalPlayerMatchStats struct {
		Offsides       *int `json:"offsides"`
		Saves          *int `json:"saves"`
		InsideBoxSaves *int `json:"inside_box_saves"`
		PenScored      *int `json:"pen_scored"`
		PenMissed      *int `json:"pen_missed"`
		PenSaved       *int `json:"pen_saved"`
		PenCommitted   *int `json:"pen_committed"`
		PenWon         *int `json:"pen_won"`
		HitWoodwork    *int `json:"hit_woodwork"`
		Tackles        *int `json:"tackles"`
		Blocks         *int `json:"blocks"`
		Interceptions  *int `json:"interceptions"`
		Clearances     *int `json:"clearances"`
		Dispossessed   *int `json:"dispossesed"`
		MinutesPlayed  *int `json:"minutes_played"`
	}

	// Assistants provides IDs of match officials
	Assistants struct {
		FirstAssistantID  *int `json:"first_assistant_id"`
		SecondAssistantID *int `json:"second_assistant_id"`
		FourthOfficialID  *int `json:"fourth_official_id"`
	}

	// CardEvent provides details of a card event
	CardEvent struct {
		ID          int64   `json:"id"`
		TeamID      string  `json:"team_id"`
		Type        string  `json:"type"`
		FixtureID   int     `json:"fixture_id"`
		PlayerID    int     `json:"player_id"`
		PlayerName  string  `json:"player_name"`
		Minute      int     `json:"minute"`
		ExtraMinute *int    `json:"extra_minute"`
		Reason      *string `json:"reason"`
	}

	// Coaches provides IDs of coaches in a fixture
	Coaches struct {
		LocalTeamCoachID   int `json:"localteam_coach_id"`
		VisitorTeamCoachID int `json:"visitorteam_coach_id"`
	}

	// CornerEvent provides details of a corner event
	CornerEvent struct {
		ID          int    `json:"id"`
		TeamID      int    `json:"team_id"`
		FixtureID   int    `json:"fixture_id"`
		Minute      int    `json:"minute"`
		ExtraMinute *int   `json:"extra_minute"`
		Comment     string `json:"comment"`
	}

	Countries struct {
		Data []Country `json:"data"`
	}

	CountryExtra struct {
		Continent   string `json:"continent"`
		SubRegion   string `json:"sub_region"`
		WorldRegion string `json:"world_region"`
		FIFA        string `json:"fifa,string"`
		ISO         string `json:"iso,string"`
		ISO2        string `json:"iso2,string"`
		Longitude   string `json:"longitude"`
		Latitude    string `json:"latitude"`
		Flag        string `json:"flag"`
	}

	Crosses struct {
		Total    *int `json:"total"`
		Accurate *int `json:"accurate"`
	}

	Dribbles struct {
		Attempts     *int `json:"attempts"`
		Success      *int `json:"success"`
		DribbledPast *int `json:"dribbled_past"`
	}

	Duels struct {
		Total *int `json:"total"`
		Won   *int `json:"won"`
	}

	FixtureTime struct {
		Status     string `json:"status"`
		StartingAt struct {
			DateTime  string `json:"date_time"`
			Date      string `json:"date"`
			Time      string `json:"time"`
			Timestamp int    `json:"timestamp"`
			Timezone  string `json:"timezone"`
		} `json:"starting_at"`
		Minute      int  `json:"minute"`
		Second      *int `json:"second"`
		AddedTime   *int `json:"added_time"`
		ExtraMinute *int `json:"extra_minute"`
		InjuryTime  *int `json:"injury_time"`
	}

	Formations struct {
		LocalTeamFormation   string `json:"localteam_formation"`
		VisitorTeamFormation string `json:"visitorteam_formation"`
	}

	Fouls struct {
		Committed int  `json:"committed"`
		Drawn     *int `json:"drawn"`
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

	Group struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		LeagueID  int    `json:"league_id"`
		SeasonID  int    `json:"season_id"`
		RoundID   *int   `json:"round_id"`
		RoundName *int   `json:"round_name"`
		StageID   int    `json:"stage_id"`
		StageName string `json:"stage_name"`
		Resource  string `json:"resource"`
	}

	KitColors struct {
		Color    *string `json:"color"`
		KitColor *string `json:"kit_colors"`
	}

	Leagues struct {
		Data []League `json:"data"`
	}

	LeagueCoverage struct {
		Predictions      int  `json:"predictions"`
		TopScorerGoals   bool `json:"topscorer_goals"`
		TopScorerAssists bool `json:"topscorer_assists"`
		TopScorerCards   bool `json:"topscorer_cards"`
	}

	Lineup struct {
		Data []PlayerStats `json:"data"`
	}

	LiveStandings struct {
		Position           int    `json:"position"`
		Played             int    `json:"played"`
		TeamID             int    `json:"team_id"`
		TeamName           string `json:"team_name"`
		ShortCode          string `json:"short_code"`
		TeamLogo           string `json:"team_logo"`
		Goals              string `json:"goals"`
		GoalDiff           int    `json:"goal_diff"`
		Wins               int    `json:"wins"`
		Lost               int    `json:"lost"`
		Draws              int    `json:"draws"`
		Points             int    `json:"points"`
		Description        string `json:"description"`
		FairPlayPointsLose int    `json:"fairplay_points_lose"`
	}

	MatchCommentary struct {
		Data []Commentary `json:"data"`
	}

	MatchEvent struct {
		ID                int64   `json:"id"`
		TeamID            string  `json:"team_id"`
		Type              string  `json:"type"`
		FixtureID         int     `json:"fixture_id"`
		PlayerID          int     `json:"player_id"`
		PlayerName        string  `json:"player_name"`
		RelatedPlayerID   int     `json:"related_player_id"`
		RelatedPlayerName string  `json:"related_player_name"`
		Minute            int     `json:"minute"`
		ExtraMinute       *int    `json:"extra_minute"`
		Reason            *string `json:"reason"`
		Injured           *bool   `json:"injuried"`
		Result            string  `json:"result"`
	}

	MatchEvents struct {
		Data []MatchEvent `json:"data"`
	}

	MatchOfficial struct {
		ID         int    `json:"id"`
		CommonName string `json:"common_name"`
		FullName   string `json:"fullname"`
		FirstName  string `json:"firstname"`
		LastName   string `json:"lastname"`
	}

	Passes struct {
		Total     int  `json:"total"`
		Accuracy  *int `json:"accuracy"`
		KeyPasses *int `json:"key_passes"`
	}

	Penalties struct {
		Won       *int `json:"won"`
		Scores    *int `json:"scores"`
		Missed    *int `json:"missed"`
		Committed *int `json:"committed"`
		Saves     *int `json:"saves"`
	}

	PlayerStats struct {
		TeamID             int              `json:"team_id"`
		FixtureID          int              `json:"fixture_id"`
		PlayerID           int              `json:"player_id"`
		PlayerName         string           `json:"player_name"`
		Number             int              `json:"number"`
		Position           *string          `json:"position,string"`
		AdditionalPosition *string          ` json:"additional_position,string"`
		FormationPosition  *int             `json:"formation_position"`
		PositionX          *string          `json:"posx"`
		PositionY          *string          `json:"posy"`
		Captain            bool             `json:"captain"`
		Stats              PlayerMatchStats `json:"stats"`
	}

	PlayerCrosses struct {
		Total    *int `json:"total"`
		Accurate *int `json:"accurate"`
	}

	PlayerDribbles struct {
		Attempts     *int `json:"attempts"`
		Success      *int `json:"success"`
		DribbledPast *int `json:"dribbled_past"`
	}

	PlayerDuels struct {
		Total *int `json:"total"`
		Won   *int `json:"won"`
	}

	PlayerCards struct {
		YellowCards *int `json:"yellowcards"`
		RedCards    *int `json:"redcards"`
	}

	PlayerFouls struct {
		Drawn     *int `json:"drawn"`
		Committed *int `json:"committed"`
	}

	PlayerGoals struct {
		Scored   *int `json:"scored"`
		Assist   *int `json:"assists"`
		Conceded *int `json:"conceded"`
	}

	PlayerPenalties struct {
		Won       *int `json:"won"`
		Scores    *int `json:"scores"`
		Missed    *int `json:"missed"`
		Committed *int `json:"committed"`
		Saves     *int `json:"saves"`
	}

	PlayerMatchPasses struct {
		TotalCrosses    *int `json:"total_crosses"`
		CrossesAccuracy *int `json:"crosses_accuracy"`
		Passes          *int `json:"passes"`
		PassesAccuracy  *int `json:"passes_accuracy"`
		KeyPasses       *int `json:"key_passes"`
	}

	PlayerSeasonPasses struct {
		Total     *int `json:"total"`
		Accuracy  *int `json:"accuracy"`
		KeyPasses *int `json:"key_passes"`
	}

	PlayerShots struct {
		ShotsTotal  *int `json:"shots_total"`
		ShotsOnGoal *int `json:"shots_on_goal"`
	}

	PlayerMatchStats struct {
		Shots    PlayerShots                `json:"shots"`
		Goals    PlayerGoals                `json:"goals"`
		Fouls    PlayerFouls                `json:"fouls"`
		Cards    PlayerCards                `json:"cards"`
		Passing  PlayerMatchPasses          `json:"passing"`
		Dribbles Player                     `json:"dribbles"`
		Duels    PlayerDuels                `json:"duels"`
		Other    AdditionalPlayerMatchStats `json:"other"`
	}

	PlayerSeasonStats struct {
		PlayerID           int        `json:"player_id"`
		TeamID             int        `json:"team_id"`
		LeagueID           int        `json:"league_id"`
		SeasonID           int        `json:"season_id"`
		Captain            int        `json:"captain"`
		Minutes            int        `json:"minutes"`
		Appearances        int        `json:"appearences"`
		Lineups            int        `json:"lineups"`
		SubstituteIn       int        `json:"substitute_in"`
		SubstituteOut      int        `json:"substitute_out"`
		SubstitutesOnBench int        `json:"substitutes_on_bench"`
		Goals              int        `json:"goals"`
		Assists            int        `json:"assists"`
		Saves              int        `json:"saves"`
		InsideBoxSaves     int        `json:"inside_box_saves"`
		Dispossessed       int        `json:"dispossesed"`
		Interceptions      int        `json:"interceptions"`
		YellowCards        int        `json:"yellowcards"`
		YellowRed          int        `json:"yellowred"`
		RedCards           int        `json:"redcards"`
		Type               string     `json:"type"`
		Tackles            *int       `json:"tackles"`
		Blocks             *int       `json:"blocks"`
		HitPost            *int       `json:"hit_post"`
		Fouls              Fouls      `json:"fouls"`
		Crosses            Crosses    `json:"crosses"`
		Dribbles           Dribbles   `json:"dribbles"`
		Duels              Duels      `json:"duels"`
		Passes             Passes     `json:"passes"`
		Penalties          Penalties  `json:"penalties"`
		PlayerData         playerData `json:"player,omitempty"`
	}

	Position struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	Ranking struct {
		TeamID            int     `json:"team_id"`
		Points            float64 `json:"points"`
		Coefficient       int     `json:"coeffiecient"`
		Position          int     `json:"position"`
		PositionStatus    string  `json:"position_status"`
		PositionWonOrLost int     `json:"position_won_or_lost"`
	}

	Seasons struct {
		Data []Season `json:"data"`
	}

	Scores struct {
		LocalTeamScore      int     `json:"localteam_score"`
		VisitorTeamScore    int     `json:"visitorteam_score"`
		LocalTeamPenScore   *int    `json:"localteam_pen_score"`
		VisitorTeamPenScore *int    `json:"visitorteam_pen_score"`
		HTScore             *string `json:"ht_score"`
		FTScore             *string `json:"ft_score"`
		ETScore             *string `json:"et_score"`
		PSScore             *string `json:"ps_score"`
	}

	Sidelined struct {
		PlayerID    int    `json:"player_id"`
		SeasonID    int    `json:"season_id"`
		TeamID      int    `json:"team_id"`
		Description string `json:"description"`
		StartDate   string `json:"start_date"`
	}

	Sport struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Current bool   `json:"current"`
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

	TeamData struct {
		Data *Team `json:"data"`
	}

	TeamLeagueStats struct {
		GamesPlayed  int `json:"games_played"`
		Won          int `json:"won"`
		Draw         int `json:"draw"`
		Lost         int `json:"lost"`
		GoalsScored  int `json:"goals_scored"`
		GoalsAgainst int `json:"goals_against"`
	}

	TeamLeagueTotalStats struct {
		GoalDifference int `json:"goal_difference"`
		Points         int `json:"points"`
	}

	TeamSeasonStats struct {
		TeamID   int  `json:"team_id"`
		SeasonID int  `json:"season_id"`
		StageID  *int `json:"stage_id"`
		Win      struct {
			Total string `json:"total"`
			Home  string `json:"home"`
			Away  string `json:"away"`
		} `json:"win"`
		Draw struct {
			Total string `json:"total"`
			Home  string `json:"home"`
			Away  string `json:"away"`
		} `json:"draw"`
		Lost struct {
			Total string `json:"total"`
			Home  string `json:"home"`
			Away  string `json:"away"`
		} `json:"lost"`
		GoalsFor struct {
			Total string `json:"total"`
			Home  string `json:"home"`
			Away  string `json:"away"`
		} `json:"goals_for"`
		GoalsAgainst struct {
			Total string `json:"total"`
			Home  string `json:"home"`
			Away  string `json:"away"`
		} `json:"goals_against"`
		CleanSheet struct {
			Total string `json:"total"`
			Home  string `json:"home"`
			Away  string `json:"away"`
		} `json:"clean_sheet"`
		FailedToScore struct {
			Total string `json:"total"`
			Home  string `json:"home"`
			Away  string `json:"away"`
		} `json:"failed_to_score"`
		ScoringMinutes []struct {
			Period []struct {
				Minute     string `json:"minute"`
				Count      string `json:"count"`
				Percentage string `json:"percentage"`
			} `json:"period"`
		} `json:"scoring_minutes"`
		AvgGoalsPerGameScored struct {
			Total string `json:"total"`
			Home  string `json:"home"`
			Away  string `json:"away"`
		} `json:"avg_goals_per_game_scored"`
		AvgGoalsPerGameConceded struct {
			Total string `json:"total"`
			Home  string `json:"home"`
			Away  string `json:"away"`
		} `json:"avg_goals_per_game_conceded"`
		AvgFirstGoalScored struct {
			Total string `json:"total"`
			Home  string `json:"home"`
			Away  string `json:"away"`
		} `json:"avg_first_goal_scored"`
		AvgFirstGoalConceded struct {
			Total string `json:"total"`
			Home  string `json:"home"`
			Away  string `json:"away"`
		} `json:"avg_first_goal_conceded"`
		Attacks                     *int `json:"attacks"`
		DangerousAttacks            *int `json:"dangerous_attacks"`
		AvgBallPossessionPercentage *int `json:"avg_ball_possession_percentage"`
		Fouls                       *int `json:"fouls"`
		AvgFoulsPerGame             *int `json:"avg_fouls_per_game"`
		Offsides                    *int `json:"offsides"`
		Redcards                    *int `json:"redcards"`
		Yellowcards                 *int `json:"yellowcards"`
		ShotsBlocked                *int `json:"shots_blocked"`
		ShotsOffTarget              *int `json:"shots_off_target"`
		AvgShotsOffTargetPerGame    *int `json:"avg_shots_off_target_per_game"`
		ShotsOnTarget               *int `json:"shots_on_target"`
		AvgShotsOnTargetPerGame     *int `json:"avg_shots_on_target_per_game"`
		Btts                        *int `json:"btts"`
		GoalLine                    *int `json:"goal_line"`
	}

	TeamStandings struct {
		LocalTeamPosition   int `json:"localteam_position"`
		VisitorTeamPosition int `json:"visitorteam_position"`
	}

	TeamAttacks struct {
		Attacks          interface{} `json:"attacks"`
		DangerousAttacks interface{} `json:"dangerous_attacks"`
	}

	TeamColors struct {
		LocalTeam   KitColors `json:"localteam"`
		VisitorTeam KitColors `json:"visitorteam"`
	}

	TeamPasses struct {
		Total      interface{} `json:"total"`
		Accurate   interface{} `json:"accurate"`
		Percentage interface{} `json:"percentage"`
	}

	TeamShots struct {
		Total      interface{} `json:"total"`
		OnGoal     interface{} `json:"ongoal"`
		OffGoal    interface{} `json:"offgoal"`
		Blocked    interface{} `json:"blocked"`
		InsideBox  interface{} `json:"insidebox"`
		OutsideBox interface{} `json:"outsidebox"`
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
		PossessionTime *int        `json:"possessiontime"`
		YellowCards    *int        `json:"yellowcards"`
		RedCards       *int        `json:"redcards"`
		Saves          *int        `json:"saves"`
		Substitutions  *int        `json:"substitutions"`
		GoalKick       *int        `json:"goal_kick"`
		GoalAttempts   *int        `json:"goal_attempts"`
		FreeKick       *int        `json:"free_kick"`
		ThrowIn        *int        `json:"throw_in"`
		BallSafe       *int        `json:"ball_safe"`
		Goals          *int        `json:"goals"`
		Penalties      *int        `json:"penalties"`
		Injuries       *int        `json:"injuries"`
	}

	TeamsStats struct {
		Data []TeamStats `json:"data"`
	}

	Transfer struct {
		PlayerID   int     `json:"player_id"`
		FromTeamID int     `json:"from_team_id"`
		ToTeamID   int     `json:"to_team_id"`
		SeasonID   *int    `json:"season_id"`
		Transfer   string  `json:"transfer"`
		Type       string  `json:"type"`
		Date       string  `json:"date"`
		Amount     *string `json:"amount"`
	}

	Trophy struct {
		PlayerID int         `json:"player_id"`
		Status   string      `json:"status"`
		Times    int         `json:"times"`
		League   string      `json:"league"`
		LeagueID int         `json:"league_id"`
		Seasons  seasonsData `json:"seasons"`
	}

	WeatherReport struct {
		Code        string `json:"code"`
		Type        string `json:"type"`
		Icon        string `json:"icon"`
		Temperature struct {
			Temp float64 `json:"temp"`
			Unit string  `json:"unit"`
		} `json:"temperature"`
		TemperatureCelcius struct {
			Temp float64 `json:"temp"`
			Unit string  `json:"unit"`
		} `json:"temperature_celcius"`
		Clouds   string `json:"clouds"`
		Humidity string `json:"humidity"`
		Pressure *int   `json:"pressure"`
		Wind     struct {
			Speed  string `json:"speed"`
			Degree int    `json:"degree"`
		} `json:"wind"`
		Coordinates struct {
			Lat *float64 `json:"lat"`
			Lon *float64 `json:"lon"`
		}
		UpdatedAt *string `json:"updated_at"`
	}
)
