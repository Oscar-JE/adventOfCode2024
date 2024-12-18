package field

import (
	vec "adventofcode/geometry/vec2d"
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

func TestCircumAndAreaOfId(t *testing.T) {
	lines := []string{"AA", "BA"}
	var field Field = ParseFromLines(lines)
	areA, circumA := field.circumAndAreaOfId("A")
	if areA != 3 || circumA != 8 {
		t.Errorf("Area A in TestFloodField does not get correct area of circumference")
	}
	areB, circumB := field.circumAndAreaOfId("B")
	if areB != 1 || circumB != 4 {
		t.Errorf("Area B in TestFloodField does not get correct area of circumference")
	}
}
