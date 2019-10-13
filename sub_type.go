package sportmonks

type (
	Coaches struct {
		LocalTeamCoachID   int `json:"localteam_coach_id"`
		VisitorTeamCoachID int `json:"visitorteam_coach_id"`
	}

	ExtraPlayerStats struct {
		Assists       *int `json:"assists"`
		Offsides      *int `json:"offsides"`
		Saves         *int `json:"saves"`
		PenScored     *int `json:"pen_scored"`
		PenMissed     *int `json:"pen_missed"`
		PenSaved      *int `json:"pen_saved"`
		PenCommitted  *int `json:"pen_committed"`
		PenWon        *int `json:"pen_won"`
		HitWoodwork   *int `json:"hit_woodwork"`
		Tackles       *int `json:"tackles"`
		Blocks        *int `json:"blocks"`
		Interceptions *int `json:"interceptions"`
		Clearances    *int `json:"clearances"`
		MinutesPlayed *int `json:"minutes_played"`
	}

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
		Dispossessed    *int `json:"dispossesed"`
		MinutesPlayed  *int `json:"minutes_played"`
	}

	FixtureTime struct {
		Status     string `json:"status"`
		StartingAt struct {
			DateTime  string `json:"date_time"`
			Date      string `json:"date"`
			Time      string `json:"time"`
			Timestamp int  `json:"timestamp"`
			Timezone  string `json:"timezone"`
		} `json:"starting_at"`
		Minute      int  `json:"minute"`
		Second      *int  `json:"second"`
		AddedTime   *int `json:"added_time"`
		ExtraMinute *int `json:"extra_minute"`
		InjuryTime  *int `json:"injury_time"`
	}

	Formations struct {
		LocalTeamFormation   string `json:"localteam_formation"`
		VisitorTeamFormation string `json:"visitorteam_formation"`
	}

	KitColors struct {
		Color *string `json:"color"`
		KitColor *string `json:"kit_colors"`
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

	//PlayerStats struct {
	//	Shots             PlayerShots      `json:"shots"`
	//	Goals             PlayerGoals      `json:"goals"`
	//	Fouls             PlayerFouls      `json:"fouls"`
	//	Cards             PlayerCards      `json:"cards"`
	//	Passes            PlayerPasses     `json:"passing"`
	//	ExtraPlayersStats ExtraPlayerStats `json:"other"`
	//}

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

	TeamStandings struct {
		LocalTeamPosition   int `json:"localteam_position"`
		VisitorTeamPosition int `json:"visitorteam_position"`
	}

	TeamAttacks struct {
		Attacks          interface{} `json:"attacks"`
		DangerousAttacks interface{} `json:"dangerous_attacks"`
	}

	TeamColors struct {
		LocalTeam 	KitColors `json:"localteam"`
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

	Transfer struct {
		PlayerID   int         `json:"player_id"`
		FromTeamID int         `json:"from_team_id"`
		ToTeamID   int         `json:"to_team_id"`
		SeasonID   *int         `json:"season_id"`
		Transfer   string      `json:"transfer"`
		Type       string      `json:"type"`
		Date       string      `json:"date"`
		Amount     *string		`json:"amount"`
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
		Pressure *int `json:"pressure"`
		Wind     struct {
			Speed  string  `json:"speed"`
			Degree int `json:"degree"`
		} `json:"wind"`
		Coordinates struct {
			Lat *float64 `json:"lat"`
			Lon *float64 `json:"lon"`
		}
		UpdatedAt *string `json:"updated_at"`
	}
)
