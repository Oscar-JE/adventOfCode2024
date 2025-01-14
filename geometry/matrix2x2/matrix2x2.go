package matrix2x2

import (
	vec "adventofcode/geometry/vec2d"
	"adventofcode/integer"
)

type Matrix2x2 struct {
	c0 vec.Vec2d
	c1 vec.Vec2d
}

func (m Matrix2x2) linearlyDependent() bool {
	dotProd := vec.DotProduct(m.c0, m.c1)
	projectedLen := integer.Positive(dotProd)
	return projectedLen*projectedLen == vec.AbsSquared(m.c0)*vec.AbsSquared(m.c1)
}
