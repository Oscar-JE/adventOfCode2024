package instructions

import (
	"adventofcode/day17/cpu/outid"
	"adventofcode/integer"
)

type Bxc struct {
	op1 int
	op2 int
}

func InitBxc(register int, operand int) Bxc {
	return Bxc{op1: register, op2: operand}
}

func (a Bxc) Out() int {
	return integer.XOR(a.op1, a.op2)
}

func (a Bxc) ResultStore() outid.OutId {
	return outid.B
}
