package instructions

import ("adventofcode/day17/cpu/outid"
		"adventofcode/integer")

type Bdv struct {
	op1 int
	op2 int
}

func InitBdv(register int, operand int) Bdv {
	return Bdv{op1: register, op2: operand}
}

func (a Bdv) Out() int {
	denominator := integer.ToThePowerOf(2,a.op2)
	return a.op1/denominator
}

func (a Bdv) ResultStore() outid.OutId {
	return outid.B
}