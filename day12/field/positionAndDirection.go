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

func positionOrder(pDir []positionAndDirection) []positionAndDirection {
	sorted := []positionAndDirection{}
	if len(pDir) == 0 {
		panic("should always be a complete loop at this point")
	}
	position := pDir[0].position
	// kör det långsamaste sättet först
	// här tappar jag antagligen punkter
	for i := 0; i < len(pDir); i++ {
		for j := 0; j < len(pDir); j++ {
			posDir := pDir[j]
			if posDir.position == position {
				sorted = append(sorted, posDir)
				position = posDir.nextPosition()
				break
			}
		}
	}
	return sorted
}
