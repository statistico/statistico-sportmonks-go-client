package sportmonks

type (
	Assistants struct {
		FirstAssistantID  *int `json:"first_assistant_id"`
		SecondAssistantID *int `json:"second_assistant_id"`
		FourthOfficialID  *int `json:"fourth_official_id"`
	}

	Coaches struct {
		LocalteamCoachID   int `json:"localteam_coach_id"`
		VisitorteamCoachID int `json:"visitorteam_coach_id"`
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
