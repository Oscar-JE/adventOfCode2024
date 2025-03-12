package space

import (
	"adventofcode/day14/quadrant"
	vec "adventofcode/geometry/vec2d"
	"testing"
)

func TestQuadrantOnCross(t *testing.T) {
	var s RepeatingSpace = InitRepeater(11, 7)
	insideAQuadrand, _ := s.Quadrant(vec.Init(3, 5))
	if insideAQuadrand {
		t.Errorf("The middlepoint is excluded from all quadrants")
	}
}

func TestQuadrantFirsth(t *testing.T) {
	var s RepeatingSpace = InitRepeater(11, 7)
	insideAQuadrand, quad := s.Quadrant(vec.Init(2, 6))
	if !insideAQuadrand || quad != quadrant.First() {
		t.Errorf("(2, 6) should be in first quadrant")
	}
}

func TestQuadrantSecond(t *testing.T) {
	var s RepeatingSpace = InitRepeater(11, 7)
	insideAQuadrand, quad := s.Quadrant(vec.Init(2, 4))
	if !insideAQuadrand || quad != quadrant.Second() {
		t.Errorf("(2,4) should be in second quadrant")
	}
}

func TestQuadrantThird(t *testing.T) {
	var s RepeatingSpace = InitRepeater(11, 7)
	insideAQuadrand, quad := s.Quadrant(vec.Init(4, 4))
	if !insideAQuadrand || quad != quadrant.Third() {
		t.Errorf("(4,4) should be in third quadrant")
	}
}

func TestQuadrantForth(t *testing.T) {
	var s RepeatingSpace = InitRepeater(11, 7)
	insideAQuadrand, quad := s.Quadrant(vec.Init(4, 6))
	if !insideAQuadrand || quad != quadrant.Fourth() {
		t.Errorf("(4, 6) should be in fourth quadrant")
	}
}

func TestQuadrantTot(t *testing.T) {
	var s RepeatingSpace = InitRepeater(11, 7)
	insideAQuadrand, quad := s.Quadrant(vec.Init(0, 6))
	if !insideAQuadrand || quad != quadrant.First() {
		t.Errorf("(0,6) should be in first quadrant")
	}
}

func TestFindInternalPoint(t *testing.T) {
	var s RepeatingSpace = InitRepeater(11, 7)
	internalPoint := s.InternalPosition(vec.Init(6, -2))
	if internalPoint != vec.Init(6, 5) {
		t.Errorf("Not correct wrapping")
	}
}

func TestFindInternalPoint2(t *testing.T) {
	var s RepeatingSpace = InitRepeater(11, 7)
	internalPoint := s.InternalPosition(vec.Init(10, -1))
	if internalPoint != vec.Init(10, 6) {
		t.Errorf("Not correct wrapping")
	}
}

func TestLimitAboveTrivial(t *testing.T) {
	res := limitAbove(4, 10)
	if res != 4 {
		t.Errorf("We got obvious problems")
	}
}

func TestLimitAboveLimit(t *testing.T) {
	res := limitAbove(10, 10)
	if res != 0 {
		t.Errorf("Above limit should be exclusive")
	}
}

func TestLimitBelowLimit(t *testing.T) {
	res := limitAbove(0, 10)
	if res != 0 {
		t.Errorf("Above limit should be inclusive")
	}
}

func TestLimitBelelowNegativeOverflow(t *testing.T) {
	res := limitAbove(-1, 10)
	if res != 9 {
		t.Errorf("Negativeoverflow not working as intendend")
	}
}

func TestLimitAboveLargerNegativeOverflow(t *testing.T) {
	res := limitAbove(-25, 10)
	if res != 5 {
		t.Errorf("examin larger negative overflows")
	}
}
