package space

import (
	"adventofcode/day14/quadrant"
	vec "adventofcode/geometry/vec2d"
)

type RepeatingSpace struct {
	lowerRightCorner vec.Vec2d
}

func (r RepeatingSpace) GetYLimit() int {
	return r.lowerRightCorner.GetY()
}

func (r RepeatingSpace) GetXLimit() int {
	return r.lowerRightCorner.GetX()
}

func InitRepeater(width int, height int) RepeatingSpace {
	return RepeatingSpace{lowerRightCorner: vec.Init(width+1, height+1)}
}

func (s RepeatingSpace) InternalPosition(unboandedPosition vec.Vec2d) vec.Vec2d {
	xPos := limitAbove(unboandedPosition.GetX(), s.lowerRightCorner.GetX())
	yPos := limitAbove(unboandedPosition.GetY(), s.lowerRightCorner.GetY())
	return vec.Init(xPos, yPos)
}

func (s RepeatingSpace) Quadrant(positions vec.Vec2d) (bool, quadrant.Quadrant) {
	overMidpoint := positions.GetX() < s.lowerRightCorner.GetX()/2-1
	belowMidPoint := s.lowerRightCorner.GetX()/2-1 < positions.GetX()
	leftSide := positions.GetY() < s.lowerRightCorner.GetY()/2-1
	rightSide := s.lowerRightCorner.GetY()/2-1 < positions.GetY()
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
	mod := toBeLimited % (upperLimit - 1)
	if mod >= 0 {
		return mod
	}
	return upperLimit - 1 + mod
}
