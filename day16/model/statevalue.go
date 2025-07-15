package model

import (
	"adventofcode/day15/directions"
	vec "adventofcode/geometry/vec2d"
	"adventofcode/day16/state"
)

type StateValue struct {
	nrRows      int
	nrCols      int
	stateValues []float64
}

func InitValueFunc(nrRows int, nrCols int) StateValue {
	stateValues := []float64{}
	entrys := nrRows * nrCols
	for range entrys {
		stateValues = append(stateValues, 0)
	}
	return StateValue{nrRows: nrRows, nrCols: nrCols, stateValues: stateValues}
}

func (v StateValue) GetValueOf(state  state.State) float64 {
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
	return  state.Init(position, direction)
}

func (sv StateValue) positionFromIndex(positionIndex int) vec.Vec2d {
	var x int = positionIndex / sv.nrCols
	var y int = positionIndex - sv.nrCols*x
	return vec.Init(x, y)
}

func (sv StateValue) indexFromPosition(position vec.Vec2d) int {
	return position.GetX()*sv.nrCols + position.GetY()
}
