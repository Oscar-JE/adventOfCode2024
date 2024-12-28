package field

import vec "adventofcode/geometry/vec2d"

type positionAndDirection struct {
	position  vec.Vec2d
	direction vec.Vec2d
}

func (p positionAndDirection) nextPosition() vec.Vec2d {
	return vec.Add(p.position, p.direction)
}

func nrSides(positionsDirs []positionAndDirection) int {
	positionsDirs = positionOrder(positionsDirs)
	reduced := reduce(positionsDirs)
	return len(reduced)
}

func prevIndex(index int, len int) int {
	prevIndex := index - 1
	if prevIndex < 0 {
		return len - 1
	}
	return prevIndex
}

func reduce(positionsDirs []positionAndDirection) []positionAndDirection {
	reduced := []positionAndDirection{}
	for i := 0; i < len(positionsDirs); i++ {
		prevIndex := prevIndex(i, len(positionsDirs))
		if positionsDirs[prevIndex].direction != positionsDirs[i].direction {
			reduced = append(reduced, positionsDirs[i])
		}
	}

	return reduced
}

type link struct {
	start vec.Vec2d
	end   vec.Vec2d
}

func positionOrder(pDir []positionAndDirection) []positionAndDirection {
	sorted := []positionAndDirection{}
	if len(pDir) == 0 {
		panic("should always be a complete loop at this point")
	}
	linkes := []link{}
	for _, el := range pDir {
		linkes = append(linkes, link{el.position, vec.Add(el.position, el.direction)})
	}
	sortedLinkes := []link{}
	sortedLinkes = append(sortedLinkes, linkes[0])
	for i := 1; i < len(linkes); i++ {
		lastLink := sortedLinkes[i-1]
		for _, linkC := range linkes {
			if lastLink.end == linkC.start {
				sortedLinkes = append(sortedLinkes, linkC)
				break
			}
		}
	}

	return sorted
}
