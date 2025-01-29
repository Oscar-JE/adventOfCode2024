package field

import (
	vec "adventofcode/geometry/vec2d"
)

type positionAndDirection struct {
	position  vec.Vec2d
	direction vec.Vec2d
}

func nrSides(positionsDirs []positionAndDirection) int {
	loops := findLoops(positionsDirs)
	nrSides := 0
	for _, l := range loops {
		nrSides += l.nrSides()
	}
	return nrSides
}

func findLoops(positionDirs []positionAndDirection) []loop {
	loops := []loop{}
	unused := positionDirs
	for len(unused) > 0 {
		u, newLoop := findLoop(unused)
		unused = u
		loops = append(loops, newLoop)
	}
	return loops
}

func findLoop(positionDir []positionAndDirection) ([]positionAndDirection, loop) {
	unused := positionDir
	unusedNext := []positionAndDirection{}
	lp := loop{}
	for !lp.isLoop() {
		for _, posDir := range positionDir {
			lk := link{posDir.position, vec.Add(posDir.position, posDir.direction)}
			added := lp.append(lk)
			if !added {
				unusedNext = append(unusedNext, posDir)
			}
		}
		unused = unusedNext
		unusedNext = []positionAndDirection{}
	}
	return unused, lp
}

type link struct {
	start vec.Vec2d
	end   vec.Vec2d
}

func (l link) direction() vec.Vec2d {
	return vec.Subtract(l.end, l.start)
}
