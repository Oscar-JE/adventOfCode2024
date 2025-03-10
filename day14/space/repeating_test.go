package space

import (
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
