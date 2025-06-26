package model

import "adventofcode/day15/directions"

type StateValue struct {
	nrCols      int
	stateValues []float64
}

func InitValueFunc(nrRows int, nrCols int) StateValue {
	stateValues := []float64{}
	entrys := nrRows * nrCols
	for range entrys {
		stateValues = append(stateValues, 0)
	}
	return StateValue{nrCols: nrCols, stateValues: stateValues}
}

func (v StateValue) GetValueOf(state State) float64 {
	return v.stateValues[v.indexOfState(state)]
}

func (v StateValue) SetValueOf(state State, val float64) {
	v.stateValues[v.indexOfState(state)] = val
}

func (sv StateValue) indexOfState(state State) int {
	positionIndex := state.position.GetX()*sv.nrCols + state.position.GetY()
	directionIndex := directions.Enumeration(state.direction)
	return directions.NrDirections*positionIndex + directionIndex
}
