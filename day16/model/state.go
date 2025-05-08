package model

import (
	"adventofcode/day15/directions" // denna bör flyttas ut ur dag 15
	vec "adventofcode/geometry/vec2d"
)

type State struct {
	position   vec.Vec2d
	directions directions.Direction
}
