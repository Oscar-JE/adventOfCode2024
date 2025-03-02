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
	links := []link{}
	for _, posDir := range positionDir {
		lk := link{posDir.position, vec.Add(posDir.position, posDir.direction)}
		links = append(links, lk)

	}
	loops := []loop{}
	unused := positionDirs
	for len(unused) > 0 {
		u, newLoop := findLoop(unused)
		unused = u
		loops = append(loops, newLoop)
	}
	return loops
}

func findLoop(positionDir []link) ([]positionAndDirection, loop) {
	lp := loop{}
	links := []link{}
	for !lp.isLoop() {
		for _, lk := range links {
			lp.append(lk)
		}
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
