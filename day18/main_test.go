package main

import (
	vec "adventofcode/geometry/vec2d"
	"os"
	"testing"
)

func TestParseCoordList(t *testing.T) {
	input := "5,3\r\n4,2"
	expected := []vec.Vec2d{vec.Init(3, 5), vec.Init(2, 4)}
	out := ParseCoordList(input)
	if len(out) != 2 {
		t.Errorf("wrong length of parsed coordinates")
		return
	}
	if out[0] != expected[0] || out[1] != expected[1] {
		t.Errorf("elements are off")
	}
}

func TestShortExample(t *testing.T) {
	bytes, _ := os.ReadFile("example.txt")
	strRep := string(bytes)
	deadBytes := ParseCoordList(strRep)
	shortest := ShortestPathCornerToCorner(6, deadBytes[:12])
	if shortest != 22 {
		t.Errorf("example given in day18 failed")
	}
}
