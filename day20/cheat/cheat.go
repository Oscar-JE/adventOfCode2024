package cheat

import (
	"adventofcode/day20/distance"
	"adventofcode/geometry/matrix"
	vec "adventofcode/geometry/vec2d"
	"adventofcode/integer"
)

type cheat struct {
	p1 vec.Vec2d
	p2 vec.Vec2d
}

func DistanceOfCheat(distsStart matrix.Matrix[distance.Distance], distsEnd matrix.Matrix[distance.Distance], c cheat) distance.Distance {
	distanceFromStart := distsStart.Get(c.p1.GetX(), c.p1.GetY())
	distanceOfCheat := vec.L1Dist(vec.Subtract(c.p2, c.p1))
	distanceToEnd := distsEnd.Get(c.p2.GetX(), c.p2.GetY())
	if !distanceFromStart.Finite() || !distanceToEnd.Finite() {
		return distance.Infinite()
	}
	d := distance.Add(distanceFromStart, distance.Init(distanceOfCheat))
	d = distance.Add(d, distanceToEnd)
	return d
}

var direct []vec.Vec2d = []vec.Vec2d{vec.Init(1, 0), vec.Init(0, 1), vec.Init(-1, 0), vec.Init(0, -1)}

func AllPossibleCheats(nrRows int, nrCols int) []cheat {
	cheats := []cheat{}
	limit := vec.Init(nrRows, nrCols)
	for i := range nrRows {
		for j := range nrCols {
			start := vec.Init(i, j)
			for _, v := range direct {
				diff := vec.ScalarMult(2, v)
				end := vec.Add(start, diff)
				if inside(end, limit) {
					cheats = append(cheats, cheat{p1: start, p2: end})
				}
			}
		}
	}
	return cheats
}

func AllPossibleCheats2(nrRows int, nrCols int) []cheat {
	cheats := []cheat{}
	limit := vec.Init(nrRows, nrCols)
	for i := range nrRows {
		for j := range nrCols {
			start := vec.Init(i, j)
			possibleEnds := l1RadiaWithMiddpoint(start, 6)
			for _, end := range possibleEnds {
				if inside(end, limit) {
					cheats = append(cheats, cheat{p1: start, p2: end})
				}
			}
		}
	}
	return cheats
}

func inside(point vec.Vec2d, limit vec.Vec2d) bool {
	return 0 <= point.GetX() && point.GetX() < limit.GetX() && 0 <= point.GetY() && point.GetY() < limit.GetY()
}

func l1RadiaWithMiddpoint(midPoint vec.Vec2d, l1Radia int) []vec.Vec2d {
	unMovedArea := l1Area(l1Radia)
	for i, el := range unMovedArea {
		unMovedArea[i] = vec.Add(midPoint, el)
	}
	return unMovedArea
}

func l1Area(l1Radia int) []vec.Vec2d {
	if l1Radia <= 0 {
		return []vec.Vec2d{}
	}
	figure := []vec.Vec2d{}
	row := -l1Radia
	for row <= l1Radia {
		colAbs := l1Radia - integer.ABS(row)
		col := -colAbs
		for col <= colAbs {
			figure = append(figure, vec.Init(row, col))
		}
	}
	return figure

}
