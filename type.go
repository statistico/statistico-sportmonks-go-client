package sportmonks

type (
	AggregatedMatchOdds struct {
		Data []MatchOdds `json:"data"`
	}

	// Assist Scorer struct
	AssistScorer struct {
		Position int    `json:"position"`
		SeasonID int    `json:"season_id"`
		PlayerID int    `json:"player_id"`
		TeamID   int    `json:"team_id"`
		StageID  int    `json:"stage_id"`
		Assists  int    `json:"assists"`
		Type     string `json:"type"`
		Player struct {
			Data Player `json:"data"`
		} `json:"player"`
		Team struct {
			Data Team `json:"data"`
		} `json:"team"`
	}

 	CardEvent struct {
		ID          int64       `json:"id"`
		TeamID      string      `json:"team_id"`
		Type        string      `json:"type"`
		FixtureID   int         `json:"fixture_id"`
		PlayerID    int         `json:"player_id"`
		PlayerName  string      `json:"player_name"`
		Minute      int         `json:"minute"`
		ExtraMinute *int `json:"extra_minute"`
		Reason      *string `json:"reason"`
	}

	CardEvents struct {
		Data []CardEvent `json:"data"`
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
		Player struct {
			Data Player `json:"data"`
		} `json:"player"`
		Team struct {
			Data Team `json:"data"`
		} `json:"team"`
	}

 	CornerEvent struct {
		ID          int         `json:"id"`
		TeamID      int         `json:"team_id"`
		FixtureID   int         `json:"fixture_id"`
		Minute      int         `json:"minute"`
		ExtraMinute *int `json:"extra_minute"`
		Comment     string      `json:"comment"`
	}

	CornerEvents struct {
		Data []CornerEvent `json:"data"`
	}

	Countries struct {
		Data []Country `json:"data"`
	}

	ContinentData struct {
		Data *Continent `json:"data"`
	}

	CountryExtra struct {
		Continent   string      `json:"continent"`
		SubRegion   string      `json:"sub_region"`
		WorldRegion string      `json:"world_region"`
		FIFA        string      `json:"fifa, string"`
		ISO         string      `json:"iso, string"`
		ISO2        string      `json:"iso2, string"`
		Longitude   string      `json:"longitude"`
		Latitude    string      `json:"latitude"`
		Flag        string      `json:"flag"`
	}

	Leagues struct {
		Data []League `json:"data"`
	}

	LeagueCoverage struct {
		Predictions      int `json:"predictions"`
		TopScorerGoals   bool `json:"topscorer_goals"`
		TopScorerAssists bool `json:"topscorer_assists"`
		TopScorerCards   bool `json:"topscorer_cards"`
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
		Player struct {
			Data Player `json:"data"`
		} `json:"player"`
		Team struct {
			Data Team `json:"data"`
		} `json:"team"`
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
		Data []PlayerMatchStats `json:"data"`
	}

	MatchCommentary struct {
		Data []Commentary `json:"data"`
	}

 	MatchEvent struct {
		ID                int64       `json:"id"`
		TeamID            string      `json:"team_id"`
		Type              string      `json:"type"`
		FixtureID         int         `json:"fixture_id"`
		PlayerID          int         `json:"player_id"`
		PlayerName        string      `json:"player_name"`
		RelatedPlayerID   int         `json:"related_player_id"`
		RelatedPlayerName string      `json:"related_player_name"`
		Minute            int         `json:"minute"`
		ExtraMinute       *int `json:"extra_minute"`
		Reason            *string `json:"reason"`
		Injured           *bool       `json:"injuried"`
		Result            string      `json:"result"`
	}

	MatchEvents struct {
		Data []MatchEvent `json:"data"`
	}

	MatchHighlights struct {
		Data []VideoHighlights `json:"data"`
	}

	MatchTVStations struct {
		Data []TVStation `json:"data"`
	}

	//LineupPlayer struct {
	//	TeamID             int         `json:"team_id"`
	//	FixtureID          int         `json:"fixture_id"`
	//	PlayerID           int         `json:"player_id"`
	//	PlayerName         string      `json:"player_name"`
	//	Number             *int        `json:"number"`
	//	Position           *string     `json:"position"`
	//	AdditionalPosition *string     `json:"additional_position"`
	//	FormationPosition  *int        `json:"formation_position"`
	//	Posx               *int        `json:"posx"`
	//	Posy               *int        `json:"posy"`
	//	Stats              PlayerStats `json:"stats"`
	//}

 	PlayerMatchStats struct {
		TeamID             int         `json:"team_id"`
		FixtureID          int         `json:"fixture_id"`
		PlayerID           int         `json:"player_id"`
		PlayerName         string      `json:"player_name"`
		Number             int         `json:"number"`
		Position           *string `json:"position, string"`
		AdditionalPosition *string` json:"additional_position, string"`
		FormationPosition  *int `json:"formation_position"`
		PositionX              *string `json:"posx"`
		PositionY               *string `json:"posy"`
		Captain            bool        `json:"captain"`
		Stats              struct {
			Shots PlayerShots`json:"shots"`
			Goals PlayerGoals `json:"goals"`
			Fouls PlayerFouls`json:"fouls"`
			Cards PlayerCards `json:"cards"`
			Passing PlayerMatchPasses `json:"passing"`
			Dribbles Player `json:"dribbles"`
			Duels PlayerDuels `json:"duels"`
			Other AdditionalPlayerMatchStats`json:"other"`
		} `json:"stats"`
 	}

 	PlayerSeasonStats struct {
		PlayerID           int         `json:"player_id"`
		TeamID             int         `json:"team_id"`
		LeagueID           int         `json:"league_id"`
		SeasonID           int         `json:"season_id"`
		Captain            int         `json:"captain"`
		Minutes            int         `json:"minutes"`
		Appearances        int         `json:"appearences"`
		Lineups            int         `json:"lineups"`
		SubstituteIn       int         `json:"substitute_in"`
		SubstituteOut      int         `json:"substitute_out"`
		SubstitutesOnBench int         `json:"substitutes_on_bench"`
		Goals              int         `json:"goals"`
		Assists            int         `json:"assists"`
		Saves              int         `json:"saves"`
		InsideBoxSaves     int         `json:"inside_box_saves"`
		Dispossessed        int         `json:"dispossesed"`
		Interceptions      int         `json:"interceptions"`
		YellowCards        int         `json:"yellowcards"`
		YellowRed          int         `json:"yellowred"`
		RedCards           int         `json:"redcards"`
		Type               string      `json:"type"`
		Tackles            *int`json:"tackles"`
		Blocks             *int `json:"blocks"`
		HitPost            *int `json:"hit_post"`
		Fouls             PlayerFouls      `json:"fouls"`
		Crosses            PlayerCrosses      `json:"crosses"`
		Dribbles  			PlayerDribbles`json:"dribbles"`
		Duels 				PlayerDuels `json:"duels"`
		Passes            PlayerSeasonPasses `json:"passes"`
		Penalties  		PlayerPenalties `json:"penalties"`
		Player struct {
			Data Player `json:"data"`
		} `json:"player, omitempty"`
	}

 	Position struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	Ranking struct {
		TeamID            int     `json:"team_id"`
		Points            float64 `json:"points"`
		Coeffiecient      int     `json:"coeffiecient"`
		Position          int     `json:"position"`
		PositionStatus    string  `json:"position_status"`
		PositionWonOrLost int     `json:"position_won_or_lost"`
	}

 	Referee struct {
		ID         int    `json:"id"`
		CommonName string `json:"common_name"`
		FullName   string `json:"fullname"`
		FirstName  string `json:"firstname"`
		LastName   string `json:"lastname"`
	}

 	Sidelined struct {
		PlayerID    int    `json:"player_id"`
		SeasonID    int    `json:"season_id"`
		TeamID      int    `json:"team_id"`
		Description string `json:"description"`
		StartDate   string `json:"start_date"`
	}

	SidelinedData struct {
		Data []Sidelined `json:"data"`
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

	TeamData struct {
		Data Team `json:"data"`
	}

	TeamLeagueStanding struct {
		Position  int         `json:"position"`
		TeamID    int         `json:"team_id"`
		TeamName  string      `json:"team_name"`
		RoundID   int         `json:"round_id"`
		RoundName int         `json:"round_name"`
		GroupID   *int 		   `json:"group_id"`
		GroupName *int `json:"group_name"`
		Overall  	TeamLeagueStats `json:"overall"`
		Home 		TeamLeagueStats `json:"home"`
		Away 		TeamLeagueStats `json:"away"`
		Total 		TeamLeagueTotalStats `json:"total"`
		Result     string `json:"result"`
		Points     int    `json:"points"`
		RecentForm string `json:"recent_form"`
		Status     string `json:"status"`
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
		TeamID   int         `json:"team_id"`
		SeasonID int         `json:"season_id"`
		StageID  *int 		 `json:"stage_id"`
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
		Attacks                     *int`json:"attacks"`
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

	TeamStats struct {
		TeamID         int         `json:"team_id"`
		FixtureID      int         `json:"fixture_id"`
		Shots          TeamShots   `json:"shots"`
		Passes         TeamPasses  `json:"passes"`
		Attacks        TeamAttacks `json:"attacks"`
		Fouls          *int `json:"fouls"`
		Corners        *int `json:"corners"`
		Offsides       *int `json:"offsides"`
		PossessionTime *int `json:"possessiontime"`
		YellowCards    *int `json:"yellowcards"`
		RedCards       *int `json:"redcards"`
		Saves          *int `json:"saves"`
		Substitutions  *int `json:"substitutions"`
		GoalKick       *int `json:"goal_kick"`
		GoalAttempts   *int `json:"goal_attempts"`
		FreeKick       *int `json:"free_kick"`
		ThrowIn        *int `json:"throw_in"`
		BallSafe       *int `json:"ball_safe"`
		Goals          *int `json:"goals"`
		Penalties          *int `json:"penalties"`
		Injuries          *int `json:"injuries"`
	}

	TeamsStats struct {
		Data []TeamStats `json:"data"`
	}

 	Trophy struct {
		PlayerID int    `json:"player_id"`
		Status   string `json:"status"`
		Times    int    `json:"times"`
		League   string `json:"league"`
		LeagueID int    `json:"league_id"`
		Seasons  struct {
			Data []Season `json:"data"`
		} `json:"seasons"`
	}
)
