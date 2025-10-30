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

func TestAdd(t *testing.T) {
	a := ScoreList{scores: []ScoreAndNumber{{0, 1}, {2, 3}, {3, 2}}}
	b := ScoreList{scores: []ScoreAndNumber{{2, 4}}}
	expected := ScoreList{scores: []ScoreAndNumber{{0, 1}, {2, 7}, {3, 2}}}
	c := Add(a, b)
	if !c.Eq(expected) {
		t.Errorf("wobahoba")
	}
}
