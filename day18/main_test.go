package main

import (
	vec "adventofcode/geometry/vec2d"
	"testing"
)

func TestParseCoordList(t *testing.T) {
	input := "5,3\r\n4,2"
	expected := []vec.Vec2d{vec.Init(5, 3), vec.Init(4, 2)}
	out := ParseCoordList(input)
	if len(out) != 2 {
		t.Errorf("wrong length of parsed coordinates")
		return
	}
	if out[0] != expected[0] || out[1] != expected[1] {
		t.Errorf("elements are off")
	}
}
