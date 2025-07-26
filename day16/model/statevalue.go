package model

import (
	"adventofcode/day15/directions"
	"adventofcode/day16/state"
	vec "adventofcode/geometry/vec2d"
	"fmt"
	"strconv"
)

type StateValue struct {
	nrRows      int
	nrCols      int
	stateValues []float64
}

func InitValueFunc(nrRows int, nrCols int) StateValue {
	stateValues := []float64{}
	entrys := nrRows * nrCols * directions.NrDirections
	for range entrys {
		stateValues = append(stateValues, 0)
	}
	return StateValue{nrRows: nrRows, nrCols: nrCols, stateValues: stateValues}
}

func (v StateValue) GetValueOf(state state.State) float64 {
	if v.indexOfState(state) >= len(v.stateValues) {
		fmt.Println(state)
	}
	return v.stateValues[v.indexOfState(state)]
}

func (v StateValue) SetValueOf(state state.State, val float64) {
	v.stateValues[v.indexOfState(state)] = val
}

func (sv StateValue) indexOfState(state state.State) int {
	positionIndex := state.GetPosition().GetX()*sv.nrCols + state.GetPosition().GetY()
	directionIndex := directions.Enumeration(state.GetDirection())
	return directions.NrDirections*positionIndex + directionIndex
}

func (sv StateValue) StateFromIndex(index int) state.State {
	var positionIndex int = index / directions.NrDirections
	var directionIndex = index % directions.NrDirections
	var position vec.Vec2d = sv.positionFromIndex(positionIndex)
	var direction directions.Direction = directions.DirectionFromIndex(directionIndex)
	return state.Init(position, direction)
}

func (sv StateValue) positionFromIndex(positionIndex int) vec.Vec2d {
	var x int = positionIndex / sv.nrCols
	var y int = positionIndex - sv.nrCols*x
	return vec.Init(x, y)
}

func (sv StateValue) indexFromPosition(position vec.Vec2d) int {
	return position.GetX()*sv.nrCols + position.GetY()
}

func (sv StateValue) FirstState() state.State {
	return sv.StateFromIndex(0)
}

func (sv StateValue) HasNext(s state.State) bool {
	return s != state.Init(vec.Init(sv.nrRows-1, sv.nrCols-1), directions.DirectionFromIndex(directions.NrDirections-1))
}

func (sv StateValue) GetNext(s state.State) state.State {
	index := sv.indexOfState(s) + 1
	return sv.StateFromIndex(index)
}

func (sv StateValue) String() string {
	
	retStr := ""
	for row := 0; row < sv.nrRows; row++ {
		lowestRowindex := sv.indexOfState(state.Init(vec.Init(row, 0), directions.DirectionFromIndex(0)))
		highestRowIndex := sv.indexOfState(state.Init(vec.Init(row, sv.nrCols-1), directions.DirectionFromIndex(directions.NrDirections)))
		for positionIndex := lowestRowindex; positionIndex <= highestRowIndex; positionIndex++ {
			retStr += strconv.FormatFloat(sv.stateValues[positionIndex], 'f', -1, 64)
		}
		retStr += "\n"
	}
	return retStr
}
