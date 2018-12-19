package sportmonks

type (
	CountryResponse struct {
		Data []Country `json:"data"`
		Meta Meta `json:"meta"`
	}

	Meta struct {
		Pagination struct {
			Total int `json:"total"`
			Count int `json:"count"`
			PerPage int `json:"per_page"`
			CurrentPage int `json:"current_page"`
			TotalPages int `json:"total_pages"`
			Links struct {
				Next string `json:"next"`
			}
		}
	}
)
