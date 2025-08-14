package towels

import "testing"

func TestNoSolution(t *testing.T) {
	rack := ParseTowelRack("bs, br, rr")
	pattern := ParsePattern("brrrbb")
	if rack.SolvablePattern(pattern) {
		t.Errorf("double b with no following r is impossible in the given case")
	}
}

func TestSolution(t *testing.T) {
	rack := ParseTowelRack("b, br, rr")
	pattern := ParsePattern("brrrbb")
	if !rack.SolvablePattern(pattern) {
		t.Errorf("three rs is possible to create")
	}
}
