package model

import (
	"adventofcode/day15/directions"
	vec "adventofcode/geometry/vec2d"
	"testing"
)

func TestLosing(t *testing.T) {
	m := Init("###\r\n#.#\r\n###")
	if !m.winning(State{vec.Init(1, 1), directions.East()}) {
		t.Errorf("not winning")
	}
}

func TestWinning(t *testing.T) {
	m := Init("###\r\n#E#\r\n###")
	if m.winning(State{vec.Init(1, 1), directions.East()}) {
		t.Errorf("is winning")
	}
}
