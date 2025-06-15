package model

import (
	"adventofcode/day15/directions"
	vec "adventofcode/geometry/vec2d"
)

type State struct {
	position  vec.Vec2d
	direction directions.Direction
}

func InitState(position vec.Vec2d, dir directions.Direction) State {
	return State{position: position, direction: dir}
}
