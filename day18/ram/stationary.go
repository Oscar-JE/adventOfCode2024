package ram

import (
	"adventofcode/geometry/matrix"
	vec "adventofcode/geometry/vec2d"
)

type StationaryRam struct {
	m matrix.Matrix[rune]
}

func InitStationary(maxCoord int, coords []vec.Vec2d) StationaryRam {
	numberOfRunes := (maxCoord + 1) * (maxCoord + 1)
	i := 0
	rs := []rune{}
	for i < numberOfRunes {
		rs = append(rs, '.')
		i++
	}
	m := matrix.Init(rs, maxCoord+1, maxCoord+1)
	for _, cord := range coords {
		m.Set(cord.GetX(), cord.GetY(), '#')
	}
	return StationaryRam{m: m}
}

func (s StationaryRam) String() string {
	rep := ""
	for i := range s.m.GetNrRows() {
		row := s.m.GetRow(i)
		rep += string(row)
		if i != s.m.GetNrRows()-1 {
			rep += "\n"
		}
	}
	return rep
}

func (s StationaryRam) Eq(other StationaryRam) bool {
	return s.m.Eq(other.m)
}
