package sportmonks

import (
	"encoding/json"
	"strconv"
)

type (
	AggregatedAssistScorerData struct {
		Data []AssistScorer `json:"data"`
	}

	AggregatedCardScorerData struct {
		Data []CardScorer `json:"data"`
	}

	AggregatedGoalScorerData struct {
		Data []GoalScorer `json:"data"`
	}

	AggregatedMatchOddsData struct {
		Data []MatchOdds `json:"data"`
	}

	AssistScorerData struct {
		Data []AssistScorer `json:"data"`
	}

	BookmakerOddsData struct {
		Data []BookmakerOdds `json:"data"`
	}

	CardEventsData struct {
		Data []CardEvent `json:"data"`
	}

	CardScorerData struct {
		Data []CardScorer `json:"data"`
	}

	CoachData struct {
		Data *Coach `json:"data"`
	}

	CornerEventsData struct {
		Data []CornerEvent `json:"data"`
	}

	CountryData struct {
		Data *Country `json:"data"`
	}

	CountriesData struct {
		Data []Country `json:"data"`
	}

	ContinentData struct {
		Data *Continent `json:"data"`
	}

	FixtureData struct {
		Data *Fixture `json:"data"`
	}

	FixturesData struct {
		Data []Fixture `json:"data"`
	}

	GoalEventsData struct {
		Data []GoalEvent `json:"data"`
	}

	GoalScorerData struct {
		Data []GoalScorer `json:"data"`
	}

	GroupsData struct {
		Data []Group `json:"data"`
	}

	LeagueData struct {
		Data *League `json:"data"`
	}

	LeaguesData struct {
		Data []League `json:"data"`
	}

	MatchCommentaryData struct {
		Data []Commentary `json:"data"`
	}

	MatchHighlightsData struct {
		Data []VideoHighlights `json:"data"`
	}

	MatchEventsData struct {
		Data []MatchEvent `json:"data"`
	}

	MatchOfficialData struct {
		Data *MatchOfficial `json:"data"`
	}

	MatchTVStationsData struct {
		Data []TVStation `json:"data"`
	}

	OddsData struct {
		Data []Odds `json:"data"`
	}

	PlayerData struct {
		Data *Player `json:"data"`
	}

	PlayerLineupData struct {
		Data []PlayerStats `json:"data"`
	}

	PlayerSeasonStatsData struct {
		Data []PlayerSeasonStats `json:"data"`
	}

	PositionData struct {
		Data *Position `json:"data"`
	}

	RankingData struct {
		Data *Ranking `json:"data"`
	}

	RoundData struct {
		Data *Round `json:"data"`
	}

	RoundsData struct {
		Data []Round `json:"data"`
	}

	SeasonData struct {
		Data *Season `json:"data"`
	}

	SeasonsData struct {
		Data []Season `json:"data"`
	}

	SidelinedData struct {
		Data []Sidelined `json:"data"`
	}

	SquadPlayerData struct {
		Data []SquadPlayer `json:"data"`
	}

	StageData struct {
		Data *Stage `json:"data"`
	}

	StagesData struct {
		Data []Stage `json:"data"`
	}

	SubstitutionEventsData struct {
		Data []SubstitutionEvent `json:"data"`
	}

	TeamData struct {
		Data *Team `json:"data"`
	}

	TeamSeasonStatsData struct {
		Data *TeamSeasonStats `json:"data"`
	}

	TeamsStatsData struct {
		Data []TeamStats `json:"data"`
	}

	TransfersData struct {
		Data []Transfer `json:"data"`
	}

	TrophyData struct {
		Data []Trophy `json:"data"`
	}

	VenueData struct {
		Data *Venue `json:"data"`
	}
)

// A FlexFloat is a float that can be un marshalled from a JSON field that has either a number or a string value.
// E.g. if the json field contains a string "4.2", the FlexFloat value will be 4.2.
type FlexFloat float32

// UnmarshalJSON implements the json.Unmarshaler interface, which allows values of any json type as an int
// and run our custom conversion
func (fi *FlexFloat) UnmarshalJSON(b []byte) error {
	if b[0] != '"' {
		return json.Unmarshal(b, (*float32)(fi))
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	i, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return err
	}
	*fi = FlexFloat(i)
	return nil
}

// A FlexInt is an int that can be un marshalled from a JSON field that has either a number or a string value.
// E.g. if the json field contains a string "42", the FlexInt value will be 42.
type FlexInt int

// UnmarshalJSON implements the json.Unmarshaler interface, which allows values of any json type as an int
// and run our custom conversion
func (fi *FlexInt) UnmarshalJSON(b []byte) error {
	if b[0] != '"' {
		return json.Unmarshal(b, (*int)(fi))
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*fi = FlexInt(i)
	return nil
}
