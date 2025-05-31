package model

import (
	"adventofcode/day15/directions"
	vec "adventofcode/geometry/vec2d"
	"testing"
)

func winningTest(t *testing.T) {
	m := Init("###\r\n#.#\r\n###")
	m.winning(State{vec.Init(1, 1), directions.Init()})
}
