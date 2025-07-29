package instructions

import "adventofcode/day17/cpu/outid"

type Jnz struct {
	op1 int
	op2 int
}

func InitJnz(register int, operand int) Jnz {
	return Jnz{op1: register, op2: operand}
}

func (a Jnz) Out() int {
	return 0
}

func (a Jnz) ResultStore() outid.OutId {
	return outid.A
}
