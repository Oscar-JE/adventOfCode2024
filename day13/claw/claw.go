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

func minimumCost(buttons ButtonPair, prize vec.Vec2d) int {
	// behöver skissa lite med papper och penna
	// super simpelt om de inte är linjärt beroende
	movMatrix := matrix2x2.Init(buttons.a.movement, buttons.b.movement)
	hasSolution := movMatrix.HasSolution(prize)
	if !hasSolution {
		panic("Encountered a problem without solution")
	}
	solSet, _ := movMatrix.Solve(prize)
	solutionSingle := solSet.GetASolution()
	singleCost := buttons.a.cost*solutionSingle.GetX() + buttons.b.cost*solutionSingle.GetY()
	//har kan vi möjligen reducera kostnaden något om vi tänker till lite
	return singleCost
}
