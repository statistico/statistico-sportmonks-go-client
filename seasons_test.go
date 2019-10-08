package sportmonks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func assertSeason(t *testing.T, season *Season) {
	assert.Equal(t, 16029, season.ID)
	assert.Equal(t, "2019/2020", season.Name)
	assert.Equal(t, 2, season.LeagueID)
	assert.Equal(t, true, season.IsCurrentSeason)
	assert.Equal(t, 183973, *season.CurrentRoundID)
	assert.Equal(t, 77443828, *season.CurrentStageID)
}
