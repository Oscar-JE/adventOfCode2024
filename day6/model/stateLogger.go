package model

import "adventofcode/geometry/matrix"

type stateLogger struct {
	visitedStates matrix.Matrix[visitedDirections]
}

type visitedDirections struct {
	up    bool
	right bool
	down  bool
	left  bool
}

func initStateLogger(rows int, cols int) stateLogger {
	values := []visitedDirections{}
	for i := 0; i < rows*cols; i++ {
		values = append(values, visitedDirections{false, false, false, false})
	}
	return stateLogger{visitedStates: matrix.Init(values, rows, cols)}
}

func (sl *stateLogger) set(g Gard) bool {
	directions := sl.visitedStates.Get(g.position.GetX(), g.position.GetY())
	if g.direction == up {
		if directions.up {
			return true
		}
		directions.up = true
	} else if g.direction == right {
		if directions.right {
			return true
		}
		directions.right = true
	} else if g.direction == down {
		if directions.down {
			return true
		}
		directions.down = true
	} else if g.direction == left {
		if directions.left {
			return true
		}
		directions.left = true
	}
	sl.visitedStates.Set(g.position.GetX(), g.position.GetY(), directions)
	return false
}

func (sl stateLogger) isStateVisitedOrOutside(g Gard) bool {
	if !sl.visitedStates.Inside(g.position.GetX(), g.position.GetY()) {
		return true
	}
	directions := sl.visitedStates.Get(g.position.GetX(), g.position.GetY())
	if g.direction == up {
		return directions.up
	} else if g.direction == right {
		return directions.right
	} else if g.direction == down {
		return directions.down
	} else if g.direction == left {
		return directions.left
	}
	panic("impossible state examined in stateLogger")
}
