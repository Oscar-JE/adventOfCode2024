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

func (d Distance) IsFinite() bool {
	return d.finite
}

func (d Distance) Val() int {
	return d.dist
}

func Init(val int) Distance {
	return Distance{dist: val, finite: true}
}

func (d Distance) String() string {
	if !d.finite {
		return "i"
	}
	return strconv.Itoa(d.dist)
}

func (d Distance) Finite() bool {
	return d.finite
}

func Infinite() Distance {
	return Distance{false, -1}
}

func Finite(d int) Distance {
	return Distance{true, d}
}

func Add(a Distance, b Distance) Distance {
	if !a.finite || !b.finite {
		return Infinite()
	}
	return Distance{dist: a.dist + b.dist, finite: true}
}

func DistanceToEnd(f field.Field) matrix.Matrix[Distance] {
	return distanceToPoint(f, f.FindEnd())
}

func DistanceToStart(f field.Field) matrix.Matrix[Distance] {
	return distanceToPoint(f, f.FindStart())
}

func distanceToPoint(f field.Field, p vec.Vec2d) matrix.Matrix[Distance] {
	nrRows := f.GetNrRows()
	nrCols := f.GetNrCols()
	distances := matrix.InitSame(nrRows, nrCols, Infinite())
	dist := 0
	directions := []vec.Vec2d{vec.Init(1, 0), vec.Init(0, 1), vec.Init(-1, 0), vec.Init(0, -1)}
	epoch := []vec.Vec2d{p}
	for len(epoch) > 0 {
		for _, d := range epoch {
			distance := Finite(dist)
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
