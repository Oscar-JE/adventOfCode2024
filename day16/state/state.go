package state

import (
	"adventofcode/day15/directions"
	vec "adventofcode/geometry/vec2d"
)

type State struct {
	position  vec.Vec2d
	direction directions.Direction
}

func Init(position vec.Vec2d, direction directions.Direction) State {
	return State{position: position, direction: direction}
}

func InitState(position vec.Vec2d, dir directions.Direction) State {
	return State{position: position, direction: dir}
}

func (s State) GetDirection() directions.Direction {
	return s.direction
}

func (s *State) SetDirection(d directions.Direction) {
	s.direction = d
}

func (s State) GetPosition() vec.Vec2d {
	return s.position
}

func (s *State) SetPosition(v vec.Vec2d) {
	s.position = v
}
