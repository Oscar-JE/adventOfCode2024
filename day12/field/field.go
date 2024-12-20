package field

import (
	"adventofcode/geometry/matrix"
	vec "adventofcode/geometry/vec2d"
	"adventofcode/set"
	"fmt"
)

type Field struct {
	plots matrix.Matrix[plantingPlot]
}

func (f Field) scoreOf(id string) int {
	allWithId := f.findPlots(id)
	splitedOnGroups := f.split(allWithId)
	sumPrice := 0
	for _, blob := range splitedOnGroups {
		area := blob.Cardinality()
		circum := f.circum(blob)
		sumPrice += area * circum
	}
	return sumPrice
}

func (f Field) split(positions set.Set[vec.Vec2d]) []set.Set[vec.Vec2d] { //hoppar över att testa den direkt för att jag inte kommer på hur man gör ett set av sets
	groups := []set.Set[vec.Vec2d]{}
	for !positions.IsEmpty() {
		pos := positions.GetAnElement()
		currentSet := f.floodFill(pos)
		groups = append(groups, currentSet)
		positions = set.Diff(positions, currentSet)
	}
	return groups
}

func (f Field) floodFill(pos vec.Vec2d) set.Set[vec.Vec2d] {
	targetId := f.plots.Get(pos.GetX(), pos.GetY()).id
	var north, west, south, east vec.Vec2d = vec.Init(-1, 0), vec.Init(0, -1), vec.Init(1, 0), vec.Init(0, 1)
	visited := set.Init([]vec.Vec2d{})
	queue := []vec.Vec2d{}
	if f.plots.Inside(pos.GetX(), pos.GetY()) {
		queue = append(queue, pos)
	}
	for len(queue) > 0 {
		currentPos := queue[0]
		queue = queue[1:]
		visited.Add(currentPos)
		newCandidates := []vec.Vec2d{vec.Add(currentPos, north), vec.Add(currentPos, west), vec.Add(currentPos, south), vec.Add(currentPos, east)}
		for _, newCandidate := range newCandidates {
			if visited.Has(newCandidate) || !f.plots.Inside(newCandidate.GetX(), newCandidate.GetY()) {
				continue
			}
			candidateId := f.plots.Get(newCandidate.GetX(), newCandidate.GetY()).id
			if candidateId == targetId {
				queue = append(queue, newCandidate)
			}
		}
	}
	return visited
}

func (f Field) totalCircumference() int {
	for i := 0; i < f.plots.GetNrRows(); i++ {
		for j := 0; j < f.plots.GetNrCols(); j++ {
			f.plots.Get(i, j).rotate()
		}
	}
	sum := f.sumOfAllRotations()
	f.reset()
	return sum
}

func (f Field) circum(positions set.Set[vec.Vec2d]) int {
	for _, pos := range positions.GetElements() {
		f.plots.Get(pos.GetX(), pos.GetY()).rotate()
	}
	sum := 0
	for _, pos := range positions.GetElements() {
		sum += f.plots.Get(pos.GetX(), pos.GetY()).lengthOfRotations()
	}
	for _, pos := range positions.GetElements() {
		f.plots.Get(pos.GetX(), pos.GetY()).reset()
	}
	return sum
}

func (f Field) findPlots(id string) set.Set[vec.Vec2d] {
	plantsPos := set.Init([]vec.Vec2d{})
	for i := 0; i < f.plots.GetNrRows(); i++ {
		for j := 0; j < f.plots.GetNrCols(); j++ {
			plant := f.plots.Get(i, j)
			if plant.id == id {
				plantsPos.Add(vec.Init(i, j))
			}
		}
	}
	return plantsPos
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
	for i := 0; i < f.plots.GetNrRows(); i++ {
		for j := 0; j < f.plots.GetNrCols(); j++ {
			sum += f.plots.Get(i, j).lengthOfRotations()
		}
	}

	return sum
}

func (f Field) String() string {
	return fmt.Sprint(f.plots)
}
