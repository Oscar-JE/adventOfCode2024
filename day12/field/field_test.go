package field

import (
	vec "adventofcode/geometry/vec2d"
	"adventofcode/set"
	"fmt"
	"testing"
)

func TestParsingAndPrinting(t *testing.T) {
	lines := []string{"AAA",
		"BBB"}
	field := ParseFromLines(lines)
	fmt.Println(field)
	ptr := field.plots.Get(0, 0).boarder.east
	*ptr = vec.Add(*ptr, vec.Init(1, 0))
}

func TestTotalCircumferenceSingleBlock(t *testing.T) {
	lines := []string{"A"}
	field := ParseFromLines(lines)
	circum := field.totalCircumference()
	if circum != 4 {
		t.Errorf("circumference of single block not calculated correctly")
	}
}

func TestTotalCircumference2blocks(t *testing.T) {
	lines1 := []string{"AA"}
	field1 := ParseFromLines(lines1)
	c1 := field1.totalCircumference()
	lines2 := []string{"A", "A"}
	field2 := ParseFromLines(lines2)
	c2 := field2.totalCircumference()
	if c1 != c2 || c1 != 6 {
		t.Errorf("circumference of 2 blocks failed")
	}
}

func TestMultipleCircumCalls(t *testing.T) {
	lines := []string{"A"}
	field := ParseFromLines(lines)
	circum1 := field.totalCircumference()
	circum2 := field.totalCircumference()
	if circum1 != circum2 {
		t.Errorf("circumference changes internal state")
	}
}

func TestFloodFillCorner(t *testing.T) {
	lines := []string{"ABA", "BBB", "ABA"}
	field := ParseFromLines(lines)
	res := field.floodFill(vec.Init(0, 0))
	exp := set.Init([]vec.Vec2d{vec.Init(0, 0)})
	if !set.Eq(res, exp) {
		t.Errorf("flood fill failed for closed corner case :P")
	}
}

func TestFloodFillMiddelUp(t *testing.T) {
	lines := []string{"ABA", "BBB", "ABA"}
	field := ParseFromLines(lines)
	res := field.floodFill(vec.Init(0, 1))
	exp := set.Init([]vec.Vec2d{vec.Init(0, 1), vec.Init(1, 0), vec.Init(1, 1), vec.Init(1, 2), vec.Init(2, 1)})
	if !set.Eq(exp, res) {
		t.Errorf(" flood fill failed for center case")
	}
}

func TestScoreOfAingles(t *testing.T) {
	lines := []string{"ABA", "BBB", "ABA"}
	field := ParseFromLines(lines)
	resA := field.scoreOf("A")
	if resA != 16 {
		t.Errorf("four singles not calculated correctly")
	}
}

func TestScoresOfMiddle(t *testing.T) {
	lines := []string{"ABA", "BBB", "ABA"}
	field := ParseFromLines(lines)
	resB := field.scoreOf("B")
	if resB != 60 {
		t.Errorf(" our expectations does not meet the result")
	}
}

func TestTotalScore(t *testing.T) {
	lines := []string{"ABA", "BBB", "ABA"}
	field := ParseFromLines(lines)
	res := field.Score()
	if res != 76 {
		t.Errorf("not what we expected of 3 by 3 star")
	}
}

func TestSingleNrSides(t *testing.T) {
	lines := []string{"A"}
	f := ParseFromLines(lines)
	idG := f.FindPlots("A")
	res := f.nrOfSides(idG)
	if res != 4 {
		t.Errorf("nr of sides total failure")
	}
}

func TestDoubleNrSides(t *testing.T) {
	lines := []string{"AA"}
	f := ParseFromLines(lines)
	idG := f.FindPlots("A")
	res := f.nrOfSides(idG)
	if res != 4 {
		t.Errorf("nr of sides total failure")
	}
}

func TestLshapNrSides(t *testing.T) {
	lines := []string{"BB",
		"AB"}
	field := ParseFromLines(lines)
	idGroup := field.FindPlots("B")
	res := field.nrOfSides(idGroup)
	if res != 6 {
		t.Errorf("small L shape failed for nr of sides")
	}
}
