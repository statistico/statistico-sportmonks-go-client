package sportmonks

type (
	// Pagination struct.
	Pagination struct {
		Count       int     `json:"count"`
		PerPage     int     `json:"per_page"`
		CurrentPage int     `json:"current_page"`
		NextPage    *string `json:"next_page"`
		HasMore     bool    `json:"has_more"`
	}

	Subscription struct {
		Meta  Meta   `json:"meta"`
		Plans []Plan `json:"plans"`
	}

	Meta struct {
		TrialEndsAt      *string `json:"trial_ends_at"`
		EndsAt           string  `json:"ends_at"`
		CurrentTimestamp int64   `json:"current_timestamp"`
	}

	Plan struct {
		Plan     string `json:"plan"`
		Sport    string `json:"sport"`
		Category string `json:"category"`
	}

	// DateTime struct.
	DateTime struct {
		Date         string `json:"date"`
		TimezoneType int    `json:"timezone_type"`
		Timezone     string `json:"timezone"`
	}
)
