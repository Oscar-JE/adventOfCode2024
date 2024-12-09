package main

import (
	"adventofcode/day6/logger"
	"adventofcode/day6/model"
	"adventofcode/geometry/matrix"
	vec "adventofcode/geometry/vec2d"
	"fmt"
)

func main() {
	partTwo()
}

func partTwo() {
	mat := model.Parse("big.txt")
	mod := model.Init(mat)
	firstPatrole := firstRoute(mod, mat)
	sum := 0
	for _, point := range firstPatrole {
		if mod.IsLoop(point.GetX(), point.GetY()) {
			//fmt.Println(point)
			sum++
		}
	}
	fmt.Println(sum)
}

func partOne() {
	mat := model.Parse("big.txt")
	mod := model.Init(mat)
	visitedStates := firstRoute(mod, mat)
	fmt.Println(len(visitedStates))
}

func firstRoute(mod model.Model, mat matrix.Matrix[rune]) []vec.Vec2d {
	log := logger.Init(mat.GetNrRows(), mat.GetNrCols())
	for !mod.IsItOver() {
		log.Set(mod.GetGardPosition())
		mod.TimeStep()
	}
	return log.VisitedPoints()
}
