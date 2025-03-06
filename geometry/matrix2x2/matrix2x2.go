package matrix2x2

import (
	vec "adventofcode/geometry/vec2d"
	"adventofcode/integer"
)

type Matrix2x2 struct {
	c0 vec.Vec2d
	c1 vec.Vec2d
}

func Init(c0 vec.Vec2d, c1 vec.Vec2d) Matrix2x2 {
	return Matrix2x2{c0: c0, c1: c1}
}

func (m Matrix2x2) linearlyDependent() bool {
	dotProd := vec.DotProduct(m.c0, m.c1)
	projectedLen := integer.Positive(dotProd)
	return projectedLen*projectedLen == vec.AbsSquared(m.c0)*vec.AbsSquared(m.c1)
}

func (m Matrix2x2) HasSolution(prize vec.Vec2d) bool {
	return false
}
