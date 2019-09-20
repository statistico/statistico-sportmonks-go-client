package sportmonks

import "strconv"

const teamUri = "/api/v2.0/teams/"
const teamSeasonUri = "/api/v2.0/teams/season/"

type Team struct {
	ID           int     `json:"id"`
	LegacyID     int     `json:"legacy_id"`
	Name         string  `json:"name"`
	ShortCode    string  `json:"short_code"`
	Twitter      *string `json:"twitter"`
	CountryID    int     `json:"country_id"`
	NationalTeam bool    `json:"national_team"`
	Founded      int     `json:"founded"`
	LogoPath     *string `json:"logo_path"`
	VenueID      int     `json:"venue_id"`
	AggregatedAssistScorers struct {
		Data []AssistScorer `json:"data"`
	} `json:"aggregatedAssistscorers, omitempty"`
	AggregatedCardScorers struct {
		Data []CardScorer `json:"data"`
	} `json:"aggregatedCardscorers, omitempty"`
	AggregatedGoalScorers struct {
		Data []GoalScorer `json:"data"`
	} `json:"aggregatedGoalscorers, omitempty"`
	AssistScorers struct {
		Data []AssistScorer `json:"data"`
	} `json:"assistscorers, omitempty"`
	CardScorers struct {
		Data []CardScorer `json:"data"`
	} `json:"cardscorers, omitempty"`
	Coach struct {
		Data *Coach `json:"data"`
	} `json:"coach, omitempty"`
	Country struct {
		Data *Country `json:"data"`
	} `json:"country, omitempty"`
	FifaRanking struct {
		Data *Ranking `json:"data"`
	} `json:"fifaranking"`
	GoalScorers struct {
		Data []GoalScorer `json:"data"`
	} `json:"goalscorers, omitempty"`
	Latest struct {
		Data []Fixture `json:"data"`
	} `json:"latest, omitempty"`
	League struct {
		Data *League `json:"data"`
	} `json:"league, omitempty"`
	LocalFixtures struct {
		Data []Fixture `json:"data"`
	} `json:"visitorFixtures, omitempty"`
	Sidelined struct {
		Data []Sidelined `json:"data"`
	} `json:"sidelined, omitempty"`
	Squad struct {
		Data []Player `json:"data"`
	} `json:"squad, omitempty"`
	Stats struct {
		Data *TeamSeasonStats `json:"data"`
	} `json:"stats, omitempty"`
	Transfers struct {
		Data []Transfer `json:"data"`
	} `json:"transfers, omitempty"`
	UefaRanking struct {
		Data *Ranking `json:"data"`
	} `json:"uefaranking"`
	Upcoming struct {
		Data []Fixture `json:"data"`
	} `json:"upcoming, omitempty"`
	Venue struct {
		Data *Venue `json:"data"`
	} `json:"venue, omitempty"`
	VisitorFixtures struct {
		Data []Fixture `json:"data"`
	} `json:"visitorFixtures, omitempty"`
	VisitorResults struct {
		Data []Fixture `json:"data"`
	} `json:"visitorResults, omitempty"`
}

// Retrieve a single team resource by ID. Use the includes slice to enrich the response data.
func (c *ApiClient) TeamById(id int, includes []string) (*Team, *Meta, error) {
	url := teamUri + "/" + strconv.Itoa(id)

	response := new(TeamResponse)

	err := c.handleRequest(url, includes, response)

	if err != nil {
		return nil, nil, err
	}

	return &response.Data, &response.Meta, err
}

// Make a request to retrieve multiple team resources for a given season. The request endpoint executed within this
// method is paginated, the second argument to this method allows the consumer to specify a page to request.
// Use the includes slice to enrich the response data.
func (c *ApiClient) TeamsBySeasonId(seasonId int, page int, includes []string) ([]Team, *Meta, error) {
	url := teamSeasonUri + "/" + strconv.Itoa(seasonId)

	response := new(TeamsResponse)

	err := c.handlePaginatedRequest(url, includes, page, response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, &response.Meta, err
}