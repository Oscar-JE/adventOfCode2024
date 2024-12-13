package solver

import (
	"adventofcode/geometry/matrix"
	vec "adventofcode/geometry/vec2d"
	"fmt"
)

type Solver struct {
	scores matrix.Matrix[int]
	field  matrix.Matrix[int]
}

func Init(field matrix.Matrix[int]) Solver {
	scoresVal := []int{}
	for i := 0; i < field.GetNrCols()*field.GetNrRows(); i++ {
		scoresVal = append(scoresVal, 0)
	}
	scores := matrix.Init(scoresVal, field.GetNrRows(), field.GetNrCols())
	return Solver{scores: scores, field: field}
}

func (s *Solver) Solve() int {
	nines := s.findAll(9)
	for _, pos := range nines {
		s.increaseScores(pos)
		fmt.Println(s.scores)
	}
	fmt.Println("scores" + s.scores.String())
	sum := 0
	zeros := s.findAll(0)
	for _, zero := range zeros {
		sum += s.scores.Get(zero.GetX(), zero.GetY())
	}
	return sum
}

func (s *Solver) Solve2() int {
	zeros := s.findAll(0)
	sum := 0
	for _, z := range zeros {
		sum += s.nrOfPathsTo9(z)
	}
	return sum
}

func (s Solver) nrOfPathsTo9(point vec.Vec2d) int {
	nextPosittions := s.NextPositionsUp(point)
	if len(nextPosittions) == 0 {
		height := s.field.Get(point.GetX(), point.GetY())
		if height == 9 {
			return 1
		}
		return 0
	}
	sum := 0
	for _, nextPos := range nextPosittions {
		sum += s.nrOfPathsTo9(nextPos)
	}
	return sum

}

func (s *Solver) increaseScores(pos vec.Vec2d) {
	futurePositions := []vec.Vec2d{pos}
	visited := []vec.Vec2d{}
	var loopPos vec.Vec2d
	for len(futurePositions) > 0 {
		loopPos, futurePositions = futurePositions[0], futurePositions[1:]
		s.scores.Set(loopPos.GetX(), loopPos.GetY(), s.scores.Get(loopPos.GetX(), loopPos.GetY())+1)
		visited = append(visited, loopPos)
		fmt.Println(s.scores)
		nextPositions := s.NextPositions(loopPos)
		filtered := filter(nextPositions, visited)
		futurePositions = addIfNotInList(filtered, futurePositions)
	}
}

func filter(toBeFiltered []vec.Vec2d, blackList []vec.Vec2d) []vec.Vec2d {
	retList := []vec.Vec2d{}
	for _, el := range toBeFiltered {
		if !inList(el, blackList) {
			retList = append(retList, el)
		}
	}
	return retList
}

func inList(el vec.Vec2d, list []vec.Vec2d) bool {
	for _, lel := range list {
		if lel == el {
			return true
		}
	}
	return false
}

func addIfNotInList(addList []vec.Vec2d, originalList []vec.Vec2d) []vec.Vec2d {
	for _, el := range addList {
		originalList = addIfNoyInListEl(el, originalList)
	}
	return originalList
}

func addIfNoyInListEl(el vec.Vec2d, originalList []vec.Vec2d) []vec.Vec2d {
	inList := false
	for _, org := range originalList {
		if org == el {
			inList = true
			break
		}
	}
	if inList {
		return originalList
	}
	return append(originalList, el)
}

func (s Solver) findAll(num int) []vec.Vec2d {
	ninePositions := []vec.Vec2d{}
	for i := 0; i < s.field.GetNrRows(); i++ {
		for j := 0; j < s.field.GetNrCols(); j++ {
			if s.field.Get(i, j) == num {
				ninePositions = append(ninePositions, vec.Init(i, j))
			}
		}
	}
	return ninePositions
}

func (s Solver) NextPositions(pos vec.Vec2d) []vec.Vec2d {
	nextPositions := []vec.Vec2d{}
	surrounding := s.pointsAround(pos)
	currentVal := s.field.Get(pos.GetX(), pos.GetY())
	for _, p := range surrounding {
		lVal := s.field.Get(p.GetX(), p.GetY())
		if lVal == currentVal-1 {
			nextPositions = append(nextPositions, p)
		}
	}
	return nextPositions
}

func (s Solver) NextPositionsUp(pos vec.Vec2d) []vec.Vec2d {
	nextPositions := []vec.Vec2d{}
	surrounding := s.pointsAround(pos)
	currentVal := s.field.Get(pos.GetX(), pos.GetY())
	for _, p := range surrounding {
		lVal := s.field.Get(p.GetX(), p.GetY())
		if lVal == currentVal+1 {
			nextPositions = append(nextPositions, p)
		}
	}
	return nextPositions
}

func (s Solver) pointsAround(pos vec.Vec2d) []vec.Vec2d {
	diff := []vec.Vec2d{vec.Init(-1, 0), vec.Init(0, 1), vec.Init(1, 0), vec.Init(0, -1)}
	pointsAround := []vec.Vec2d{}
	for _, d := range diff {
		p := vec.Add(pos, d)
		if s.field.Inside(p.GetX(), p.GetY()) {
			pointsAround = append(pointsAround, p)
		}
	}
	return pointsAround
}
