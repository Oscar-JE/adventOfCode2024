package instructions

import "adventofcode/day17/cpu/outid"

type Bdv struct {
	op1 int
	op2 int
}

func InitBdv(register int, operand int) Bdv {
	return Bdv{op1: register, op2: operand}
}

func (a Bdv) Out() int {
	return 0
}

func (a Bdv) ResultStore() outid.OutId {
	return outid.A
}