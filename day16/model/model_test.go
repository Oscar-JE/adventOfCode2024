package model

import (
	"adventofcode/day15/directions"
	"adventofcode/day16/state"
	vec "adventofcode/geometry/vec2d"
	"testing"
)

func TestLosing(t *testing.T) {
	m := Init("###\r\n#.#\r\n###")
	if !m.Winning(state.Init(vec.Init(1, 1), directions.East())) {
		t.Errorf("not winning")
	}
}

func TestWinning(t *testing.T) {
	m := Init("###\r\n#E#\r\n###")
	if m.Winning( state.Init(vec.Init(1, 1), directions.East())) {
		t.Errorf("is winning")
	}
}

func TestMoveForwardBoxed(t *testing.T) {
	m := Init("###\r\n#.#\r\n###")
	state := state.Init(vec.Init(1, 1), directions.North()) 
	pos := m.moveForward(state)
	if pos != vec.Init(1, 1) {
		t.Errorf("boxed in case should not move")
	}
}

func TestMoveForwardMoveUp(t *testing.T) {
	m := Init("###\r\n#.#\r\n#.#\r\n###")
	state := state.Init(vec.Init(2, 1), directions.North())
	pos := m.moveForward(state)
	if pos != vec.Init(1, 1) {
		t.Errorf("boxed in case should not move")
	}
}
