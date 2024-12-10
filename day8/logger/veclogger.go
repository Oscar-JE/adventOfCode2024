package logger

import (
	"adventofcode/geometry/matrix"
	vec "adventofcode/geometry/vec2d"
)

type UniquePositions struct {
	seen matrix.Matrix[bool]
}

func InitLogger(rows int, cols int) UniquePositions {
	values := []bool{}
	for i := 0; i < rows*cols; i++ {
		values = append(values, false)
	}
	mat := matrix.Init(values, rows, cols)
	return UniquePositions{seen: mat}
}

func (u *UniquePositions) Set(v vec.Vec2d) {
	if u.seen.Inside(v.GetX(), v.GetY()) {
		u.seen.Set(v.GetX(), v.GetY(), true)
	}
}

func (u UniquePositions) Positions() []vec.Vec2d {
	seen := []vec.Vec2d{}
	for i := range u.seen.GetNrRows() {
		for j := range u.seen.GetNrCols() {
			if u.seen.Get(i, j) {
				seen = append(seen, vec.Init(i, j))
			}
		}
	}
	return seen
}
