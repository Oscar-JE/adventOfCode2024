package model

import (
	"adventofcode/day15/directions"
	"adventofcode/day16/field"
	"adventofcode/day16/state"
	vec "adventofcode/geometry/vec2d"
)

type Model struct {
	internalRep field.Field
}

func Init(raw string) Model {
	f := field.Parse(raw)
	return Model{internalRep: f}
}

func (m *Model) StateTransition(s state.State, action Action) state.State {
	if action == turnDown {
		dir := s.GetDirection().TurnDown()
		s.SetDirection(dir)
		return s
	}
	if action == turnUp {
		dir := s.GetDirection().TurnUp()
		s.SetDirection(dir)
		return s
	}
	if action == forward {
		position := m.moveForward(s)
		s.SetPosition(position)
		return s
	}
	panic("Action outside the action space have been taken as input")
}

func (m *Model) moveForward(s state.State) vec.Vec2d {
	currentPosition := s.GetPosition()
	if !m.internalRep.IsFloor(currentPosition) {
		panic("impossible state passed to move forward")
	}
	dir := s.GetDirection()
	next := vec.Add(currentPosition, vec.Vec2d(dir))
	if m.internalRep.IsFloor(next) {
		currentPosition = next
	}
	return currentPosition
}

func (m *Model) Cost(s state.State, action Action) float64 {
	if m.Winning(s) {
		return 0
	}
	if action == turnUp || action == turnDown {
		return 1000
	} else if action == forward {
		return m.forewardCost(s)
	}
	panic("should necer reach this state")
}

func (m *Model) forewardCost(s state.State) float64 {
	nextPos := m.moveForward(s)
	firstPos := s.GetPosition()
	return float64(vec.L1Dist(vec.Subtract(nextPos, firstPos)))
}

func (m *Model) Winning(state state.State) bool {
	return m.internalRep.IsPositionEnd(state.GetPosition())
}

func (m *Model) StartState() state.State { //todo skriv denna på riktigt
	return state.Init(vec.Init(9, 0), directions.North())
}
