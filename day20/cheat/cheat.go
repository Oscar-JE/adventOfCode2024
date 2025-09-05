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

func distanceOfCheat(distsStart matrix.Matrix[distance.Distance], distsEnd matrix.Matrix[distance.Distance], c cheat) int {
	distanceFromStart := distsStart.Get(c.p1.GetX(), c.p1.GetY())
	distanceOfCheat := vec.L1Dist(vec.Subtract(c.p2, c.p1))
	distanceToEnd := distsEnd.Get(c.p2.GetX(), c.p2.GetY())
	// hantera den boxade typen distance
	return distanceFromStart.
}
