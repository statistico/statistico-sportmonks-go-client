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

	AggregatedMatchOdds struct {
		Data []MatchOdds `json:"data"`
	}

	AssistScorerData struct {
		Data []AssistScorer `json:"data"`
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

	CardScorerData struct {
		Data []CardScorer `json:"data"`
	}

	CoachData struct {
		Data *Coach `json:"data"`
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

	CountryData struct {
		Data *Country `json:"data"`
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

	GoalScorerData struct {
		Data []GoalScorer `json:"data"`
	}

	GroupsData struct {
		Data []Group `json:"data"`
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

	FixtureData struct {
		Data *Fixture `json:"data"`
	}

	FixturesData struct {
		Data []Fixture `json:"data"`
	}

	Fouls struct {
		Committed int         `json:"committed"`
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
		RoundID   *int    `json:"round_id"`
		RoundName *int    `json:"round_name"`
		StageID   int    `json:"stage_id"`
		StageName string `json:"stage_name"`
		Resource  string `json:"resource"`
 	}

	LeagueData struct {
		Data *League `json:"data"`
	}

	Lineup struct {
		Data []PlayerMatchStats `json:"data"`
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

	Passes struct {
		Total     int         `json:"total"`
		Accuracy  *int         `json:"accuracy"`
		KeyPasses *int `json:"key_passes"`
	}

	Penalties struct {
		Won       *int `json:"won"`
		Scores    *int `json:"scores"`
		Missed    *int `json:"missed"`
		Committed *int `json:"committed"`
		Saves     *int `json:"saves"`
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

	PlayerData struct {
		Data *Player `json:"data"`
	}

	PlayersData struct {
		Data []Player `json:"data"`
	}

	PlayerLineupData struct {
		Data []PlayerMatchStats `json:"data"`
	}

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
		Fouls              Fouls `json:"fouls"`
		Crosses Crosses `json:"crosses"`
		Dribbles Dribbles `json:"dribbles"`
		Duels Duels `json:"duels"`
		Passes Passes `json:"passes"`
		Penalties Penalties`json:"penalties"`
		Player struct {
			Data Player `json:"data"`
		} `json:"player, omitempty"`
	}

	PlayerSeasonStatsData struct {
		Data []PlayerSeasonStats `json:"data"`
	}

 	Position struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	PositionData struct {
		Data *Position `json:"data"`
	}

	Ranking struct {
		TeamID            int     `json:"team_id"`
		Points            float64 `json:"points"`
		Coeffiecient      int     `json:"coeffiecient"`
		Position          int     `json:"position"`
		PositionStatus    string  `json:"position_status"`
		PositionWonOrLost int     `json:"position_won_or_lost"`
	}

	RankingData struct {
		Data *Ranking `json:"data"`
	}

 	Referee struct {
		ID         int    `json:"id"`
		CommonName string `json:"common_name"`
		FullName   string `json:"fullname"`
		FirstName  string `json:"firstname"`
		LastName   string `json:"lastname"`
	}

	RoundData struct {
		Data *Round `json:"data"`
	}

	RoundsData struct {
		Data []Round `json:"data"`
	}

	Seasons struct {
		Data []Season `json:"data"`
	}

	SeasonData struct {
		Data *Season `json:"data"`
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

	SquadPlayerData struct {
		Data []SquadPlayer `json:"data"`
	}

	StagesData struct {
		Data []Stage `json:"data"`
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

	LeagueStandingData struct {
		Data []LeagueStanding `json:"data"`
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

	TeamSeasonStatsData struct {
		Data *TeamSeasonStats `json:"data"`
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

	TransfersData struct {
		Data []Transfer `json:"data"`
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

	TrophyData struct {
		Data []Trophy `json:"data"`
	}

	VenueData struct {
		Data *Venue `json:"data"`
	}
)
