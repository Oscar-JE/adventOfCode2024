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

func (f Field) Score() int {
	uncountedPlots := f.allPositions()
	sum := 0
	for !uncountedPlots.IsEmpty() {
		pos := uncountedPlots.GetAnElement()
		id := f.plots.Get(pos.GetX(), pos.GetY()).id
		fmt.Println("current id : " + id)
		fmt.Println("current position", pos)
		idBlob := f.FindPlots(id)
		fmt.Printf("length of blob set %d \n", idBlob.Cardinality())
		partSum := f.ScoreOfIdBlob(idBlob)
		fmt.Println(partSum)
		sum += partSum
		uncountedPlots = set.Diff(uncountedPlots, idBlob)
	}
	return sum
}

func (f Field) allPositions() set.Set[vec.Vec2d] {
	posSet := set.Init([]vec.Vec2d{})
	for i := 0; i < f.plots.GetNrRows(); i++ {
		for j := 0; j < f.plots.GetNrCols(); j++ {
			posSet.Add(vec.Init(i, j))
		}
	}
	return posSet
}

func (f Field) scoreOf(id string) int {
	allWithId := f.FindPlots(id)
	return f.ScoreOfIdBlob(allWithId)
}

func (f Field) ScoreOfIdBlob(blob set.Set[vec.Vec2d]) int {
	splittedOnGroups := f.split(blob)
	sumPrice := 0
	for _, blob := range splittedOnGroups {
		sumPrice += f.scoreOfBlob(blob)
	}
	return sumPrice
}

func (f Field) scoreOfBlob(blob set.Set[vec.Vec2d]) int {
	area := blob.Cardinality()
	//circum := f.circum(blob)
	nrSides := f.nrOfSides(blob)
	return area * nrSides
}

func (f Field) split(positions set.Set[vec.Vec2d]) []set.Set[vec.Vec2d] {
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
				visited.Add(newCandidate)
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
	vecPositions := positions.GetElements()
	for _, pos := range vecPositions {
		f.plots.Get(pos.GetX(), pos.GetY()).rotate()
	}
	sum := 0
	for _, pos := range vecPositions {
		sum += f.plots.Get(pos.GetX(), pos.GetY()).lengthOfRotations()
	}
	for _, pos := range vecPositions {
		f.plots.Get(pos.GetX(), pos.GetY()).reset()
	}
	return sum
}

func (f Field) nrOfSides(positions set.Set[vec.Vec2d]) int {
	vecPositions := positions.GetElements()
	for _, pos := range vecPositions {
		f.plots.Get(pos.GetX(), pos.GetY()).rotate()
	}
	positionsWithNonZeroRotations := []vec.Vec2d{}
	for _, pos := range vecPositions {
		if f.plots.Get(pos.GetX(), pos.GetY()).lengthOfRotations() > 0 {
			positionsWithNonZeroRotations = append(positionsWithNonZeroRotations, pos)
		} else {
			fmt.Println(pos)
		}
	}
	posAndDir := f.findPosAndDirections(positionsWithNonZeroRotations)
	for _, pos := range vecPositions {
		f.plots.Get(pos.GetX(), pos.GetY()).reset()
	}
	return nrSides(posAndDir)
}

func (f Field) findPosAndDirections(positions []vec.Vec2d) []positionAndDirection {
	allPossitionsAndDirections := []positionAndDirection{}
	for _, relevantPlot := range positions {
		plot := f.plots.Get(relevantPlot.GetX(), relevantPlot.GetY())
		allPossitionsAndDirections = append(allPossitionsAndDirections, plot.verticesAndDirections(relevantPlot)...)
	}
	return reduce(allPossitionsAndDirections)
}

func (f Field) FindPlots(id string) set.Set[vec.Vec2d] {
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
