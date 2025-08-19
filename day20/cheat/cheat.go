package cheat

import (
	"adventofcode/day20/distance"
	"adventofcode/geometry/matrix"
	vec "adventofcode/geometry/vec2d"
)

type cheat struct {
	p1 vec.Vec2d
	p2 vec.Vec2d
}

//oklart om ett hopp genom endast ett block ska räknas som ett eller två fusk

func savedTimeByeCheat(distances matrix.Matrix[distance.Distance], c cheat) int {
	return 5
}
