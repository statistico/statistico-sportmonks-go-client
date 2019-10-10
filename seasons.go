package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const seasonsUri = "/seasons"

// Season struct
type Season struct {
	ID                         int                        `json:"id"`
	Name                       string                     `json:"name"`
	LeagueID                   int                        `json:"league_id"`
	IsCurrentSeason            bool                       `json:"is_current_season"`
	CurrentRoundID             *int                       `json:"current_round_id"`
	CurrentStageID             *int                       `json:"current_stage_id"`
	AggregatedAssistScorerData AggregatedAssistScorerData `json:"aggregatedAssistscorers, omitempty"`
	AggregatedCardScorerData   AggregatedCardScorerData   `json:"aggregatedCardscorers, omitempty"`
	AggregatedGoalScorerData   AggregatedGoalScorerData   `json:"aggregatedGoalscorers, omitempty"`
	AssistScorerData           AssistScorerData           `json:"assistscorers, omitempty"`
	CardScorerData             CardScorerData             `json:"cardscorers, omitempty"`
	FixturesData               FixturesData               `json:"fixtures, omitempty"`
	GoalScorerData             GoalScorerData             `json:"goalscorers, omitempty"`
	GroupsData                 GroupsData                 `json:"groups, omitempty"`
	LeagueData                 LeagueData                 `json:"league, omitempty"`
	ResultsData                FixturesData               `json:"results, omitempty"`
	RoundsData                 RoundsData                 `json:"rounds, omitempty"`
	StagesData                 StagesData                 `json:"stages, omitempty"`
	Upcoming                   FixturesData               `json:"upcoming, omitempty"`
}

func (s *Season) AggregatedAssistScorers() []AssistScorer {
	return s.AggregatedAssistScorerData.Data
}

func (s *Season) AggregatedCardScorers() []CardScorer {
	return s.AggregatedCardScorerData.Data
}

func (s *Season) AggregatedGoalScorers() []GoalScorer {
	return s.AggregatedGoalScorerData.Data
}

func (s *Season) AssistScorers() []AssistScorer {
	return s.AssistScorerData.Data
}

func (s *Season) CardScorers() []CardScorer {
	return s.CardScorerData.Data
}

func (s *Season) Fixtures() []Fixture {
	return s.FixturesData.Data
}

func (s *Season) GoalScorers() []GoalScorer {
	return s.GoalScorerData.Data
}

func (s *Season) Groups() []Group {
	return s.GroupsData.Data
}

func (s *Season) League() *League {
	return s.LeagueData.Data
}

func (s *Season) Results() []Fixture {
	return s.FixturesData.Data
}

func (s *Season) Rounds() []Round {
	return s.RoundsData.Data
}

func (s *Season) Stages() []Stage {
	return s.StagesData.Data
}

func (s *Season) UpcomingFixtures() []Fixture {
	return s.Upcoming.Data
}

// Seasons returns a slice of Season struct and supporting meta data. The endpoint used within this method
// is paginated, to select the required page use the 'page' method argument. Page information including current page
// and total page are included within the Meta response.

// Use the includes slice of string to enrich the response data.
func (c *HTTPClient) Seasons(ctx context.Context, page int, includes []string) ([]Season, *Meta, error) {
	values := url.Values{
		"page":    {strconv.Itoa(page)},
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data []Season 	 `json:"data"`
		Meta *Meta       `json:"meta"`
	}{}

	err := c.getResource(ctx, seasonsUri, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// SeasonById returns a single Season struct with support meta data. Use the includes slice of string to enrich
// the response data.
func (c *HTTPClient) SeasonById(ctx context.Context, id int, includes []string) (*Season, *Meta, error) {
	path := fmt.Sprintf(seasonsUri+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data *Season    `json:"data"`
		Meta *Meta      `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
