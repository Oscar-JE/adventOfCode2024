package space

import (
	"adventofcode/day14/quadrant"
	vec "adventofcode/geometry/vec2d"
)

type RepeatingSpace struct {
	lowerRightCorner vec.Vec2d
}

func InitRepeater(width int, height int) RepeatingSpace {
	return RepeatingSpace{lowerRightCorner: vec.Init(height+1, width+1)}
}

func (s RepeatingSpace) internalPosition(unboandedPosition vec.Vec2d) vec.Vec2d {
	xPos := limitAbove(unboandedPosition.GetX(), s.lowerRightCorner.GetX())
	yPos := limitAbove(unboandedPosition.GetY(), s.lowerRightCorner.GetY())
	return vec.Init(xPos, yPos)
}

func (s RepeatingSpace) Quadrant(position vec.Vec2d) (bool, quadrant.Quadrant) {
	internalPosition := s.internalPosition(position)
	overMidpoint := internalPosition.GetY() < s.lowerRightCorner.GetY()/2-1
	belowMidPoint := s.lowerRightCorner.GetY()/2-1 < internalPosition.GetY()
	leftSide := internalPosition.GetX() < s.lowerRightCorner.GetX()/2-1
	rightSide := s.lowerRightCorner.GetX()/2-1 < internalPosition.GetX()
	if overMidpoint && rightSide {
		return true, quadrant.First()
	} else if overMidpoint && leftSide {
		return true, quadrant.Second()
	} else if belowMidPoint && leftSide {
		return true, quadrant.Third()
	} else if belowMidPoint && rightSide {
		return true, quadrant.Fourth()
	}
	return false, quadrant.First()
}

func limitAbove(toBeLimited int, upperLimit int) int {
	if upperLimit < 1 {
		panic("Funny why would you like to have an inverted intervale or empty?")
	}
	mod := toBeLimited % upperLimit
	if mod >= 0 {
		return mod
	}
	return upperLimit + mod
}
