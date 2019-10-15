package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Season provides a struct representation of a Season resource
type Season struct {
	ID                         int                        `json:"id"`
	Name                       string                     `json:"name"`
	LeagueID                   int                        `json:"league_id"`
	IsCurrentSeason            bool                       `json:"is_current_season"`
	CurrentRoundID             *int                       `json:"current_round_id"`
	CurrentStageID             *int                       `json:"current_stage_id"`
	AggregatedAssistScorerData aggregatedAssistScorerData `json:"aggregatedAssistscorers, omitempty"`
	AggregatedCardScorerData   aggregatedCardScorerData   `json:"aggregatedCardscorers, omitempty"`
	AggregatedGoalScorerData   aggregatedGoalScorerData   `json:"aggregatedGoalscorers, omitempty"`
	AssistScorerData           assistScorerData           `json:"assistscorers, omitempty"`
	CardScorerData             cardScorerData             `json:"cardscorers, omitempty"`
	FixturesData               fixturesData               `json:"fixtures, omitempty"`
	GoalScorerData             goalScorerData             `json:"goalscorers, omitempty"`
	GroupsData                 groupsData                 `json:"groups, omitempty"`
	LeagueData                 leagueData                 `json:"league, omitempty"`
	ResultsData                fixturesData               `json:"results, omitempty"`
	RoundsData                 roundsData                 `json:"rounds, omitempty"`
	StagesData                 stagesData                 `json:"stages, omitempty"`
	Upcoming                   fixturesData               `json:"upcoming, omitempty"`
}

// AggregatedAssistScorers returns aggregated assist scorer data.
func (s *Season) AggregatedAssistScorers() []AssistScorer {
	return s.AggregatedAssistScorerData.Data
}

// AggregatedCardScorers returns aggregated card scorer data.
func (s *Season) AggregatedCardScorers() []CardScorer {
	return s.AggregatedCardScorerData.Data
}

// AggregatedGoalScorers returns aggregated goal scorer data.
func (s *Season) AggregatedGoalScorers() []GoalScorer {
	return s.AggregatedGoalScorerData.Data
}

// AssistScorers returns aggregated assist data.
func (s *Season) AssistScorers() []AssistScorer {
	return s.AssistScorerData.Data
}

// CardScorers returns card scorer data.
func (s *Season) CardScorers() []CardScorer {
	return s.CardScorerData.Data
}

// Fixtures returns all fixture data.
func (s *Season) Fixtures() []Fixture {
	return s.FixturesData.Data
}

// GoalScorers returns goal scorer data.
func (s *Season) GoalScorers() []GoalScorer {
	return s.GoalScorerData.Data
}

// Groups returns all group data.
func (s *Season) Groups() []Group {
	return s.GroupsData.Data
}

// League returns league information.
func (s *Season) League() *League {
	return s.LeagueData.Data
}

// Results returns all completed fixture data.
func (s *Season) Results() []Fixture {
	return s.FixturesData.Data
}

// Rounds returns all round data.
func (s *Season) Rounds() []Round {
	return s.RoundsData.Data
}

// Stages returns all stage data.
func (s *Season) Stages() []Stage {
	return s.StagesData.Data
}

// UpcomingFixtures returns upcoming fixture data.
func (s *Season) UpcomingFixtures() []Fixture {
	return s.Upcoming.Data
}

// Seasons fetches Season resources. The endpoint used within this method is paginated, to select the required
// page use the 'page' method argument. Page information including current page and total page are included
// within the Meta response.
// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) Seasons(ctx context.Context, page int, includes []string) ([]Season, *Meta, error) {
	values := url.Values{
		"page":    {strconv.Itoa(page)},
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data []Season `json:"data"`
		Meta *Meta    `json:"meta"`
	}{}

	err := c.getResource(ctx, seasonsURI, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// SeasonByID fetches a Season resource by ID.
// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) SeasonByID(ctx context.Context, id int, includes []string) (*Season, *Meta, error) {
	path := fmt.Sprintf(seasonsURI+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data *Season `json:"data"`
		Meta *Meta   `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
