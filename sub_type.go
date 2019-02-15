package sportmonks

type (
	Assistants struct {
		FirstAssistantID  *int `json:"first_assistant_id"`
		SecondAssistantID *int `json:"second_assistant_id"`
		FourthOfficialID  *int `json:"fourth_official_id"`
	}

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

	FixtureTime struct {
		Status     string `json:"status"`
		StartingAt struct {
			DateTime  string `json:"date_time"`
			Date      string `json:"date"`
			Time      string `json:"time"`
			Timestamp int64  `json:"timestamp"`
			Timezone  string `json:"timezone"`
		} `json:"starting_at"`
		Minute      int  `json:"minute"`
		Second      *int `json:"second"`
		AddedTime   *int `json:"added_time"`
		ExtraMinute *int `json:"extra_minute"`
		InjuryTime  *int `json:"injury_time"`
	}

	Formations struct {
		LocalteamFormation   string `json:"localteam_formation"`
		VisitorteamFormation string `json:"visitorteam_formation"`
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
		Conceded *int `json:"conceded"`
	}

	PlayerPasses struct {
		TotalCrosses    *int `json:"total_crosses"`
		CrossesAccuracy *int `json:"crosses_accuracy"`
		Passes          *int `json:"passes"`
		PassesAccuracy  *int `json:"passes_accuracy"`
	}

	PlayerShots struct {
		ShotsTotal  *int `json:"shots_total"`
		ShotsOnGoal *int `json:"shots_on_goal"`
	}

	PlayerStats struct {
		Shots             PlayerShots      `json:"shots"`
		Goals             PlayerGoals      `json:"goals"`
		Fouls             PlayerFouls      `json:"fouls"`
		Cards             PlayerCards      `json:"cards"`
		Passes            PlayerPasses     `json:"passing"`
		ExtraPlayersStats ExtraPlayerStats `json:"other"`
	}

	Scores struct {
		LocalteamScore      int     `json:"localteam_score"`
		VisitorteamScore    int     `json:"visitorteam_score"`
		LocalteamPenScore   *int    `json:"localteam_pen_score"`
		VisitorteamPenScore *int    `json:"visitorteam_pen_score"`
		HtScore             *string `json:"ht_score"`
		FtScore             *string `json:"ft_score"`
		EtScore             *string `json:"et_score"`
	}

	Standings struct {
		LocalteamPosition   int `json:"localteam_position"`
		VisitorteamPosition int `json:"visitorteam_position"`
	}

	TeamAttacks struct {
		Attacks          *int `json:"attacks,string"`
		DangerousAttacks *int `json:"dangerous_attacks,string"`
	}

	TeamPasses struct {
		Total      *int `json:"total"`
		Accurate   *int `json:"accurate"`
		Percentage *int `json:"percentage"`
	}

	TeamShots struct {
		Total      *int `json:"total"`
		Ongoal     *int `json:"ongoal"`
		Offgoal    *int `json:"offgoal"`
		Blocked    *int `json:"blocked"`
		Insidebox  *int `json:"insidebox"`
		Outsidebox *int `json:"outsidebox"`
	}

	WeatherReport struct {
		Code        string `json:"code"`
		Type        string `json:"type"`
		Icon        string `json:"icon"`
		Temperature struct {
			Temp float64 `json:"temp"`
			Unit string  `json:"unit"`
		} `json:"temperature"`
		Clouds   string `json:"clouds"`
		Humidity string `json:"humidity"`
		Wind     struct {
			Speed  string  `json:"speed"`
			Degree float64 `json:"degree"`
		} `json:"wind"`
	}
)
