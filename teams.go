package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Team struct {
	ID                         int                        `json:"id"`
	LegacyID                   int                        `json:"legacy_id"`
	Name                       string                     `json:"name"`
	ShortCode                  string                     `json:"short_code"`
	Twitter                    *string                    `json:"twitter"`
	CountryID                  int                        `json:"country_id"`
	NationalTeam               bool                       `json:"national_team"`
	Founded                    int                        `json:"founded"`
	LogoPath                   *string                    `json:"logo_path"`
	VenueID                    int                        `json:"venue_id"`
	CurrentSeasonID            int                        `json:"current_season_id"`
	AggregatedAssistScorerData AggregatedAssistScorerData `json:"aggregatedAssistscorers, omitempty"`
	AggregatedCardScorerData   AggregatedCardScorerData   `json:"aggregatedCardscorers, omitempty"`
	AggregatedGoalScorerData   AggregatedGoalScorerData   `json:"aggregatedGoalscorers, omitempty"`
	AssistScorerData           AssistScorerData           `json:"assistscorers, omitempty"`
	CardScorerData             CardScorerData             `json:"cardscorers, omitempty"`
	CoachData                  CoachData                  `json:"coach, omitempty"`
	CountryData                CountryData                `json:"country, omitempty"`
	FIFARankingData            RankingData                `json:"fifaranking"`
	GoalScorerData             GoalScorerData             `json:"goalscorers, omitempty"`
	Latest                     FixturesData               `json:"latest, omitempty"`
	LeagueData                 LeagueData                 `json:"league, omitempty"`
	LocalFixtureData           FixturesData               `json:"visitorFixtures, omitempty"`
	SidelinedData              SidelinedData              `json:"sidelined, omitempty"`
	SquadData                  SquadPlayerData            `json:"squad, omitempty"`
	StatsData                  TeamSeasonStatsData        `json:"stats, omitempty"`
	TransfersData              TransfersData              `json:"transfers, omitempty"`
	UEFARankingData            RankingData                `json:"uefaranking. omitempty"`
	Upcoming                   FixturesData               `json:"upcoming, omitempty"`
	VenueData                  VenueData                  `json:"venue, omitempty"`
	VisitorFixtureData         FixturesData               `json:"visitorFixtures, omitempty"`
	VisitorResultData          FixturesData               `json:"visitorResults, omitempty"`
}

func (t *Team) AggregatedAssistScorers() []AssistScorer {
	return t.AggregatedAssistScorerData.Data
}

func (t *Team) AggregatedCardScorers() []CardScorer {
	return t.AggregatedCardScorerData.Data
}

func (t *Team) AggregatedGoalScorers() []CardScorer {
	return t.AggregatedCardScorerData.Data
}

func (t *Team) AssistScorers() []AssistScorer {
	return t.AssistScorerData.Data
}

func (t *Team) CardScorers() []CardScorer {
	return t.CardScorerData.Data
}

func (t *Team) Coach() *Coach {
	return t.CoachData.Data
}

func (t *Team) Country() *Country {
	return t.CountryData.Data
}

func (t *Team) FIFARanking() *Ranking {
	return t.FIFARankingData.Data
}

func (t *Team) GoalScorers() []GoalScorer {
	return t.GoalScorerData.Data
}

func (t *Team) LatestResults() []Fixture {
	return t.Latest.Data
}

func (t *Team) League() *League {
	return t.LeagueData.Data
}

func (t *Team) LocalFixtures() []Fixture {
	return t.LocalFixtureData.Data
}

func (t *Team) Sidelined() []Sidelined {
	return t.SidelinedData.Data
}

func (t *Team) Squad() []SquadPlayer {
	return t.SquadData.Data
}

func (t *Team) Stats() *TeamSeasonStats {
	return t.StatsData.Data
}

func (t *Team) Transfers() []Transfer {
	return t.TransfersData.Data
}

func (t *Team) UEFARanking() *Ranking {
	return t.UEFARankingData.Data
}

func (t *Team) UpcomingFixtures() []Fixture {
	return t.Upcoming.Data
}

func (t *Team) Venue() *Venue {
	return t.VenueData.Data
}

func (t *Team) VisitorFixtures() []Fixture {
	return t.VisitorFixtureData.Data
}

func (t *Team) VisitorResults() []Fixture {
	return t.VisitorResultData.Data
}

// TeamById returns a single Team struct with supporting meta data. Use the includes slice of string to enrich
// the response data.
func (c *HTTPClient) TeamById(ctx context.Context, id int, includes []string) (*Team, *Meta, error) {
	path := fmt.Sprintf(teamsUri+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data *Team `json:"data"`
		Meta *Meta `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// TeamsBySeasonId returns a slice of Team struct associated to a season with supporting meta data. The endpoint used
// within this method is paginated, to select the required page use the 'page' method argument. Page information
// including current page and total page are included within the Meta response.

// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) TeamsBySeasonId(ctx context.Context, seasonId, page int, includes []string) ([]Team, *Meta, error) {
	path := fmt.Sprintf(teamsSeasonUri+"/%d", seasonId)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
		"page":    {strconv.Itoa(page)},
	}

	response := struct {
		Data []Team `json:"data"`
		Meta *Meta  `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
