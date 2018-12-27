package sportmonks

type (
	// CountriesResponse struct
	CountriesResponse struct {
		Data []Country `json:"data"`
		Meta Meta `json:"meta"`
	}

	// Meta struct
	Meta struct {
		Pagination struct {
			Total int `json:"total"`
			Count int `json:"count"`
			PerPage int `json:"per_page"`
			CurrentPage int `json:"current_page"`
			TotalPages int `json:"total_pages"`
			Links struct {
				Previous string `json:"previous"`
				Next string `json:"next"`
			}
		}
	}
)
