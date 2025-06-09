package model

import "adventofcode/day16/field"

type Model struct {
	internalRep field.Field
}

func Init(raw string) Model {
	f := field.Parse(raw)
	return Model{internalRep: f}
}

func (m *Model) StateTransition(state State, action Action) State {
	if action == turnDown {
		dir := state.directions.TurnDown()
		state.directions = dir
		return state
	}
	if action == turnUp {
		dir := state.directions.TurnUp()
		state.directions = dir
		return state
	}
	if action == forward {
		position := m.moveForeward(state)
	}
}

func (m *Model) Reward(state State, action Action) float64 {
	return 0.0
}

func (m *Model) winning(state State) bool {
	return m.internalRep.IsPositionEnd(state.position)
}
