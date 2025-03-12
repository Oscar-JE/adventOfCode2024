package roborally

import (
	"adventofcode/day14/particle"
	"adventofcode/day14/space"
	vec "adventofcode/geometry/vec2d"
	"fmt"
)

type Arena struct {
	area   space.RepeatingSpace
	robots []particle.Particle
}

func Init(width int, height int, robots []particle.Particle) Arena {
	return Arena{area: space.InitRepeater(width, height), robots: robots}
}

func (a Arena) SafetyScore(timeSteps int) int {
	multArray := []int{0, 0, 0, 0}
	endPositions := a.endPositions(timeSteps)
	fmt.Println(endPositions)
	for _, endPosition := range endPositions {
		inQuadrant, quad := a.area.Quadrant(endPosition)
		if inQuadrant {
			multArray[quad.Enumerate()] += 1
		}
	}
	product := 1
	for _, m := range multArray {
		product *= m
	}
	return product
}

func (a Arena) endPositions(timeSteps int) []vec.Vec2d {
	endPositions := []vec.Vec2d{}
	for _, rob := range a.robots {
		endPosition := rob.ProjectedPosition(timeSteps)
		endPositions = append(endPositions, a.area.InternalPosition(endPosition))
	}
	return endPositions
}

func (a Arena) String() string {
	rep := ""
	for i := 0; i < a.area.GetYLimit(); i++ {
		lineRep := ""
		for j := 0; j < a.area.GetXLimit(); j++ {
			lineRep += "."
		}
		rep += lineRep + "\r\n"
	}
	return "The arena" + "\r\n" + rep
}
