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
		ID            int  `json:"id"`
		LeagueID      int  `json:"league_id"`
		SeasonID      int  `json:"season_id"`
		StageID       *int `json:"stage_id"`
		RoundID       *int `json:"round_id"`
		GroupID       *int `json:"group_id"`
		AggregateID   *int `json:"aggregate_id"`
		VenueID       *int `json:"venue_id"`
		RefereeID     *int `json:"referee_id"`
		LocalteamID   int  `json:"localteam_id"`
		VisitorteamID int  `json:"visitorteam_id"`
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
				Speed  string `json:"speed"`
				Degree int    `json:"degree"`
			} `json:"wind"`
		} `json:"weather_report"`
		Commentaries          bool    `json:"commentaries"`
		Attendance            int     `json:"attendance"`
		Pitch                 *string `json:"pitch"`
		WinningOddsCalculated bool    `json:"winning_odds_calculated"`
		Formations            struct {
			LocalteamFormation   string `json:"localteam_formation"`
			VisitorteamFormation string `json:"visitorteam_formation"`
		} `json:"formations"`
		Scores struct {
			LocalteamScore      int     `json:"localteam_score"`
			VisitorteamScore    int     `json:"visitorteam_score"`
			LocalteamPenScore   *int    `json:"localteam_pen_score"`
			VisitorteamPenScore *int    `json:"visitorteam_pen_score"`
			HtScore             *string `json:"ht_score"`
			FtScore             *string `json:"ft_score"`
			EtScore             *string `json:"et_score"`
		} `json:"scores"`
		Time struct {
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
		} `json:"time"`
		Coaches struct {
			LocalteamCoachID   int `json:"localteam_coach_id"`
			VisitorteamCoachID int `json:"visitorteam_coach_id"`
		} `json:"coaches"`
		Standings struct {
			LocalteamPosition   int `json:"localteam_position"`
			VisitorteamPosition int `json:"visitorteam_position"`
		} `json:"standings"`
		Assistants struct {
			FirstAssistantID  *int `json:"first_assistant_id"`
			SecondAssistantID *int `json:"second_assistant_id"`
			FourthOfficialID  *int `json:"fourth_official_id"`
		} `json:"assistants"`
		Leg     string  `json:"leg"`
		Colors  *string `json:"colors"`
		Deleted bool    `json:"deleted"`
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

	// Sport struct
	Sport struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Current bool   `json:"current"`
	}

	FixtureTime struct {
		Status     *string `json:"status"`
		StartingAt struct {
			DateTime  string `json:"date_time"`
			Date      string `json:"date"`
			Time      string `json:"time"`
			Timestamp int64  `json:"timestamp"`
			Timezone  string `json:"timezone"`
		} `json:"starting_at"`
	}

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
	}
)
