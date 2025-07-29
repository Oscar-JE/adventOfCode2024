package instructions

import "adventofcode/day17/cpu/outid"

type Out struct {
	op1 int
}

func InitOut(operand int) Out {
	return Out{op1: operand}
}

func (a Out) Out() int {
	return 0
}

func (a Out) ResultStore() outid.OutId {
	return outid.Out
}
