package claw

import (
	vec "adventofcode/geometry/vec2d"
)

type Button struct {
	movement vec.Vec2d
	cost     int
}

func minimumCost(buttons []Button, prize vec.Vec2d) int {
	currentPos := vec.Init(0, 0)
	return minimumCostFromPosition(0, currentPos, buttons, prize)
}

func minimumCostFromPosition(cost int, pos vec.Vec2d, buttons []Button, prize vec.Vec2d) int {
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
