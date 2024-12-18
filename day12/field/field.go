package field

import (
	"adventofcode/geometry/matrix"
	vec "adventofcode/geometry/vec2d"
	"fmt"
)

type Field struct {
	plots matrix.Matrix[plantingPlot]
}

func (f Field) floodFill(startingPoint vec.Vec2d) (int, int) { // behöver inte flodd fill i detta läget
	directions := []vec.Vec2d{vec.Init(1, 0), vec.Init(0, 1), vec.Init(-1, 0), vec.Init(0, -1)}
	startingId := f.plots.Get(startingPoint.GetX(), startingPoint.GetY()).id
	visited := make(map[vec.Vec2d]bool)
	queue := []vec.Vec2d{startingPoint}
	area := 0
	for len(queue) > 0 {
		currentVec := queue[0]
		queue = queue[1:]
		visited[currentVec] = true
		f.plots.Get(currentVec.GetX(), currentVec.GetY()).boarder.rotate()
		area++
		for _, v := range directions {
			candidate := vec.Add(currentVec, v)
			if !visited[candidate] && f.plots.Inside(candidate.GetX(), candidate.GetY()) && f.plots.Get(candidate.GetX(), candidate.GetY()).id == startingId {
				queue = append(queue, candidate)
			}
		}
	}
	circumference := f.sumOfAllRotations()
	f.reset()
	return area, circumference
}

func (f Field) totalCircumference() int { // mer i testningssyfte
	for i := 0; i < f.plots.GetNrRows(); i++ {
		for j := 0; j < f.plots.GetNrCols(); j++ {
			f.plots.Get(i, j).boarder.rotate()
		}
	}
	sum := f.sumOfAllRotations()
	f.reset()
	return sum
}

func (f Field) circumAndAreaOfId(id string) (int, int) {
	plantsWithId := f.findPlots(id)
	area := len(plantsWithId)
	circumference := 0
	for _, plant := range plantsWithId {
		plant.rotate()
	}
	for _, plant := range plantsWithId {
		circumference += plant.lengthOfRotations()
	}
	for _, plant := range plantsWithId {
		plant.boarder.reset()
	}
	return area, circumference
}

func (f Field) findPlots(id string) []plantingPlot {
	plants := []plantingPlot{}
	for i := 0; i < f.plots.GetNrRows(); i++ {
		for j := 0; j < f.plots.GetNrCols(); j++ {
			plant := f.plots.Get(i, j)
			if plant.id == id {
				plants = append(plants, plant)
			}
		}
	}
	return plants
}

func (f Field) reset() {
	for i := 0; i < f.plots.GetNrRows(); i++ {
		for j := 0; j < f.plots.GetNrCols(); j++ {
			f.plots.Get(i, j).boarder.reset()
		}
	}
}

func (f Field) sumOfAllRotations() int {
	sum := 0
	// summera längs till höger
	for i := 0; i < f.plots.GetNrRows(); i++ {
		sum += f.plots.Get(i, f.plots.GetNrCols()-1).boarder.eastAbsSquared()
	}
	// summera längst ner
	for j := 0; j < f.plots.GetNrCols(); j++ {
		sum += f.plots.Get(f.plots.GetNrRows()-1, j).boarder.southAbsSquared()
	}
	// summera rästen
	for i := 0; i < f.plots.GetNrRows(); i++ {
		for j := 0; j < f.plots.GetNrCols(); j++ {
			sum += f.plots.Get(i, j).boarder.northAbsSquared() + f.plots.Get(i, j).boarder.westAbsSquared()
		}
	}

	return sum
}

func (f Field) String() string {
	return fmt.Sprint(f.plots)
}
