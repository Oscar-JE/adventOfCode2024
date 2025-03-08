package matrix2x2

import vec "adventofcode/geometry/vec2d"

type SolutionSet struct {
	aSolPoint vec.Vec2d
	direction vec.Vec2d
}

func (s SolutionSet) GetASolution() vec.Vec2d {
	return s.aSolPoint
}

func (s SolutionSet) GetDirection() vec.Vec2d {
	return s.direction
}
