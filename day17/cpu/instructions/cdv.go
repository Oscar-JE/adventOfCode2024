package instructions

import (
	"adventofcode/day17/cpu/outid"
	"adventofcode/integer"
)

type Cdv struct {
	op1 int
	op2 int
}

func InitCdv(register int, operand int) Cdv {
	return Cdv{op1: register, op2: operand}
}

func (a Cdv) Out() int {
	denominator := integer.ToThePowerOf(2, a.op2)
	return a.op1 / denominator
}

func (a Cdv) ResultStore() outid.OutId {
	return outid.C
}
