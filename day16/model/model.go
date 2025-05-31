package model

import "adventofcode/day16/field"

type Model struct {
	internalRep field.Field
}

func Init(raw string) Model { //alright då har vi något att leka med
	f := field.Parse(raw)
	return Model{internalRep: f}
}

func (m *Model) StateTransition(state State, action Action) []State {
	return []State{}
}

func (m *Model) Reward(state State, action Action) float64 {
	//ska returnera noll för allt som inte är vinnande state
	return 0.0
}

// hur var det en plan såg ut
func (m *Model) winning(state State) bool {
	return false
}
