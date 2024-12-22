package field

import vec "adventofcode/geometry/vec2d"

type plantingPlot struct {
	id      string
	boarder swirl
}

func Init(id string) plantingPlot {
	return plantingPlot{id: id, boarder: InitSwirl()}
}

func (p plantingPlot) String() string {
	return p.id
}

func (p plantingPlot) rotate() {
	p.boarder.rotate()
}

func (p plantingPlot) lengthOfRotations() int {
	return p.boarder.sumOfSquaredRotations()
}

func (p plantingPlot) reset() {
	p.boarder.reset()
}

func (p plantingPlot) verticesAndDirections(position vec.Vec2d) []positionAndDirection {
	pDir := []positionAndDirection{}
	zero := vec.Init(0, 0)
	if *p.boarder.south != zero {
		pDir = append(pDir, positionAndDirection{position: position, direction: *p.boarder.south})
	}
	if *p.boarder.west != zero {
		pos := vec.Add(position, vec.Init(-1, 0))
		pDir = append(pDir, positionAndDirection{position: pos, direction: *p.boarder.west})
	}
	if *p.boarder.north != zero {
		pos := vec.Add(position, vec.Init(-1, 1))
		pDir = append(pDir, positionAndDirection{position: pos, direction: *p.boarder.north})
	}
	if *p.boarder.east != zero {
		pos := vec.Add(position, vec.Init(0, 1))
		pDir = append(pDir, positionAndDirection{position: pos, direction: *p.boarder.east})
	}
	return pDir
}
