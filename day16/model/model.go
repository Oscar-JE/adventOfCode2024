package model

import (
	"adventofcode/day16/field"
	vec "adventofcode/geometry/vec2d"
)

type Model struct {
	internalRep field.Field
}

func Init(raw string) Model {
	f := field.Parse(raw)
	return Model{internalRep: f}
}

func (m *Model) StateTransition(state State, action Action) State {
	if action == turnDown {
		dir := state.direction.TurnDown()
		state.direction = dir
		return state
	}
	if action == turnUp {
		dir := state.direction.TurnUp()
		state.direction = dir
		return state
	}
	if action == forward {
		position := m.moveForward(state)
		state.position = position
		return state
	}
	panic("Action outside the action space have been taken as input")
}

func (m *Model) moveForward(state State) vec.Vec2d {
	currentPosition := state.position
	if !m.internalRep.IsFloor(currentPosition) {
		panic("impossible state passed to move forward")
	}
	dir := state.direction
	next := vec.Add(currentPosition, vec.Vec2d(dir))
	for m.internalRep.IsFloor(next) {
		currentPosition = next
		next = vec.Add(next, vec.Vec2d(dir))
	}
	return currentPosition
}

func (m *Model) Cost(state State, action Action) float64 {
	if m.Winning(state) {
		return 0
	}
	return 0
}

func (m *Model) Winning(state State) bool {
	return m.internalRep.IsPositionEnd(state.position)
}
