package sportmonks

type (
	// Continent struct
	Continent struct {
		ID int `json:"id"`
		Name string `json:"name"`
	}

	// Country struct
	Country struct {
		ID int `json:"id"`
		Name string `json:"name"`
		Data struct {
			Continent string `json:"continent"`
			SubRegion string `json:"sub_region"`
			WorldRegion string `json:"world_region"`
			Fifa string `json:"fifa"`
			ISO string `json:"iso"`
			Longitude string `json:"longitude"`
			Latitude string `json:"latitude"`
		} `json:"extra"`
	}

	CountryResponse struct {
		CountryList []Country `json:"data"`
	}

	ResponseMetaData struct {
		Meta struct {
			Pagination struct {
				Total int `json:"total"`
				Count int `json:"count"`
				PerPage int `json:"per_page"`
				CurrentPage int `json:"current_page"`
				TotalPages int `json:"total_pages"`
			} `json:"pagination"`
		} `json:"meta"`
	}


)
