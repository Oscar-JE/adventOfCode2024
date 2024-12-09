package logger

import (
	"adventofcode/geometry/matrix"
	vec "adventofcode/geometry/vec2d"
)

type Logger struct {
	log matrix.Matrix[bool]
}

func Init(nrRows int, nrCols int) Logger {
	vals := []bool{}
	for i := 0; i < nrCols*nrRows; i++ {
		vals = append(vals, false)
	}
	return Logger{log: matrix.Init(vals, nrRows, nrCols)}
}

func (l *Logger) Set(v vec.Vec2d) {
	l.log.Set(v.GetX(), v.GetY(), true)
}

func (l Logger) Sum() int {
	sum := 0
	for i := range l.log.GetNrRows() {
		for j := range l.log.GetNrCols() {
			if l.log.Get(i, j) {
				sum++
			}
		}
	}
	return sum
}

func (l Logger) VisitedPoints() []vec.Vec2d {
	retList := []vec.Vec2d{}
	for i := range l.log.GetNrRows() {
		for j := range l.log.GetNrCols() {
			if l.log.Get(i, j) {
				retList = append(retList, vec.Init(i, j))
			}
		}
	}
	return retList
}
