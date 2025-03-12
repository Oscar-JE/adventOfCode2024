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

type TimeAndDistanceScore struct {
	timeSteps     int
	distanceScore int
}

func (a Arena) SearTree() []TimeAndDistanceScore {
	timeAndDistanceScores := []TimeAndDistanceScore{}
	for i := 0; i < 10; i++ {
		distScore := distanceScore(a.endPositions(i))
		timeAndDistanceScores = append(timeAndDistanceScores, TimeAndDistanceScore{i, distScore})
	}
	for i := 10; i < 100000; i++ {
		distScore := distanceScore(a.endPositions(i))
		timeAndDistanceScores = uppdateScoreList(timeAndDistanceScores, TimeAndDistanceScore{i, distScore})
	}
	return timeAndDistanceScores
}

func uppdateScoreList(list []TimeAndDistanceScore, val TimeAndDistanceScore) []TimeAndDistanceScore {
	shouldReplace := false
	replaceIndex := 0
	for i, score := range list {
		if val.distanceScore < score.distanceScore {
			shouldReplace = true
			replaceIndex = i
			break
		}
	}
	if !shouldReplace {
		return list
	} else {
		return append(append(list[0:replaceIndex], val), list[replaceIndex+1:len(list)-1]...)
	}
}

func distanceScore(positions []vec.Vec2d) int {
	sum := 0
	for i := range positions {
		v1 := positions[i]
		for j := i + 1; j < len(positions); j++ {
			v2 := positions[j]
			sum += vec.DistSquared(v1, v2)
		}
	}
	return sum
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
