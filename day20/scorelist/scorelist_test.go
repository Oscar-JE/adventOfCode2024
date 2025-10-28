package scorelist

import (
	"testing"
)

func TestInit(t *testing.T) {
	scores := []int{3, 3, 1, 2}
	sl := Init(scores)
	slExpected := ScoreList{scores: []ScoreAndNumber{{1, 1}, {2, 1}, {3, 2}}}
	if !slExpected.Eq(sl) {
		t.Errorf("What what")
	}
}
