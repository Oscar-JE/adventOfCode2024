package claw

import (
	"adventofcode/geometry/matrix2x2"
	vec "adventofcode/geometry/vec2d"
)

type Button struct {
	movement vec.Vec2d
	cost     int
}

type ButtonPair struct {
	a Button
	b Button
}

func minimumCost(buttons ButtonPair, prize vec.Vec2d) int {
	// behöver skissa lite med papper och penna
	// super simpelt om de inte är linjärt beroende
	movMatrix := matrix2x2.Init(buttons.a.movement, buttons.b.movement)
	hasSolution := movMatrix.HasSolution(prize)
	if !hasSolution {
		panic("Encountered a problem without solution")
	}
	return 4
}

func minimumCostMany(buttons []Button, prize vec.Vec2d) int {
	currentPos := vec.Init(0, 0)
	return minimumCostFromPosition(0, currentPos, buttons, prize)
}

func minimumCostFromPosition(cost int, pos vec.Vec2d, buttons []Button, prize vec.Vec2d) int {
	//vi borde kunna hitta en snabbare lösning om vi använder oss av linjär algebra
	// strunta i att gå igenom alla kombinationer fram tills målet.
	// om vi istället låser oss till att hektalslösningar av det linjära systemet och sedan röra oss mot den billigaste där båda är i första kvadranten.
	// jag tror det blir rorligare snabbare och enklare att testa
	if pos == prize {
		return cost
	}
	nextStepCosts := []int{}
	for _, b := range buttons {
		nextPos := vec.Add(pos, b.movement)
		if !impossibleRoute(nextPos, prize) {
			nextCost := cost + b.cost
			nextStepCosts = append(nextStepCosts, minimumCostFromPosition(nextCost, nextPos, buttons, prize))
		}
	}
	return multipleMin(nextStepCosts)
}

func impossibleRoute(pos vec.Vec2d, prize vec.Vec2d) bool {
	return !(vec.AbsSquared(pos) > vec.AbsSquared(prize))
}

func multipleMin(values []int) int {
	if len(values) == 0 {
		return -1
	}
	minV := values[0]
	for i := 1; i < len(values); i++ {
		minV = min(minV, values[i])
	}
	return minV
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
