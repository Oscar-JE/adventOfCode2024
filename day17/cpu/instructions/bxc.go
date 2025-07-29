package instructions

import "adventofcode/day17/cpu/outid"

type Bxc struct {
	op1 int
	op2 int
}

func InitBxc(register int, operand int) Bxc {
	return Bxc{op1: register, op2: operand}
}

func (a Bxc) Out() int {
	return 0
}

func (a Bxc) ResultStore() outid.OutId {
	return outid.A
}
