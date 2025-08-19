package distance

import (
	"adventofcode/day20/field"
	"adventofcode/geometry/matrix"
	vec "adventofcode/geometry/vec2d"
	"strconv"
)

type Distance struct {
	finite bool
	dist   int
}

func (d Distance) String() string {
	if !d.finite {
		return "i"
	}
	return strconv.Itoa(d.dist)
}

func infinite() Distance {
	return Distance{false, -1}
}

func finite(d int) Distance {
	return Distance{true, d}
}

func DistanceToEnd(f field.Field) matrix.Matrix[Distance] {
	nrRows := f.GetNrRows()
	nrCols := f.GetNrCols()
	distances := matrix.InitSame(nrRows, nrCols, infinite())
	ending := f.FindEnd()
	dist := 0
	directions := []vec.Vec2d{vec.Init(1, 0), vec.Init(0, 1), vec.Init(-1, 0), vec.Init(0, -1)}
	epoch := []vec.Vec2d{ending}
	for len(epoch) > 0 {
		for _, d := range epoch {
			distance := finite(dist)
			distances.Set(d.GetX(), d.GetY(), distance)
		}
		nextEpoch := []vec.Vec2d{}
		for _, d := range epoch {
			for _, v := range directions {
				candidate := vec.Add(d, v)
				if !f.Inside(candidate) {
					continue
				}
				if f.IsFloor(candidate) && !distances.Get(candidate.GetX(), candidate.GetY()).finite {
					nextEpoch = append(nextEpoch, candidate)
				}
			}
		}
		epoch = nextEpoch
		dist++
	}
	return distances
}
