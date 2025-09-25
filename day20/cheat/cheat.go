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

func inside(point vec.Vec2d, limit vec.Vec2d) bool {
	return 0 <= point.GetX() && point.GetX() < limit.GetX() && 0 <= point.GetY() && point.GetY() < limit.GetY()
}
