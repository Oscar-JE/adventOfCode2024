package instructions

import "adventofcode/day17/cpu/outid"

type Cdv struct {
	op1 int
	op2 int
}

func InitCdv(register int, operand int) Cdv {
	return Cdv{op1: register, op2: operand}
}

func (a Cdv) Out() int {
	return 0
}

func (a Cdv) ResultStore() outid.OutId {
	return outid.A
}