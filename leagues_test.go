package sportmonks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func assertLeague(t *testing.T, league *League) {
	assert.Equal(t, 82, league.ID)
	assert.Equal(t, true, league.Active)
	assert.Equal(t, "domestic", league.Type)
	assert.Equal(t, 4, league.LegacyID)
	assert.Equal(t, 11, league.CountryID)
	assert.Equal(t, "https://cdn.sportmonks.com/images/soccer/leagues/82.png", league.LogoPath)
	assert.Equal(t, "Bundesliga", league.Name)
	assert.Equal(t, false, league.IsCup)
	assert.Equal(t, 16264, league.CurrentSeasonID)
	assert.Equal(t, 174546, league.CurrentRoundID)
	assert.Equal(t, 77444845, league.CurrentStageID)
	assert.Equal(t, true, league.LiveStandings)
	assert.Equal(t, 1, league.Coverage.Predictions)
	assert.Equal(t, true, league.Coverage.TopScorerGoals)
	assert.Equal(t, true, league.Coverage.TopScorerAssists)
	assert.Equal(t, true, league.Coverage.TopScorerCards)
}
