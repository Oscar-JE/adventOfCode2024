package claw

import (
	"adventofcode/geometry/matrix2x2"
	vec "adventofcode/geometry/vec2d"
)

type Button struct {
	movement vec.Vec2d
	cost     int
}

func InitBut(movement vec.Vec2d, cost int) Button {
	return Button{movement: movement, cost: cost}
}

type ButtonPair struct {
	a Button
	b Button
}

func InitButtonPai(a Button, b Button) ButtonPair {
	return ButtonPair{a: a, b: b}
}

func MinimumCost(buttons ButtonPair, prize vec.Vec2d) int {
	movMatrix := matrix2x2.Init(buttons.a.movement, buttons.b.movement)
	hasSolution := movMatrix.HasSolution(prize)
	if !hasSolution {
		panic("Encountered a problem without solution")
	}
	solSet, _ := movMatrix.Solve(prize)
	if movMatrix.VecMul(solSet.GetASolution()) != prize {
		return 0
	}

	solutionSingle := solSet.GetASolution()
	costVec := vec.Init(buttons.a.cost, buttons.b.cost)
	if len(solSet.GetDirection()) == 1 {
		dir := solSet.GetDirection()[0]
		if vec.DotProduct(costVec, dir) < 0 {
			for vec.Add(solutionSingle, dir).FirstQuadrant() {
				solutionSingle = vec.Add(solutionSingle, dir)
			}
		} else if vec.DotProduct(costVec, dir) > 0 {
			for vec.Subtract(solutionSingle, dir).FirstQuadrant() {
				solutionSingle = vec.Subtract(solutionSingle, dir)
			}
		}
	}
	singleCost := buttons.a.cost*solutionSingle.GetX() + buttons.b.cost*solutionSingle.GetY()
	return singleCost
}
