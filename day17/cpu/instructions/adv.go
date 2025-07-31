package instructions

import (
	"adventofcode/day17/cpu/outid"
	"adventofcode/integer"
)

type Adv struct {
	op1 int
	op2 int
}

func InitAdv(register int, operand int) Adv {
	return Adv{op1: register, op2: operand}
}

func (a Adv) Out() int { // kan göra byte operationer här om man vill
	denominator := integer.ToThePowerOf(2, a.op2)
	return a.op1 / denominator
}

func (a Adv) ResultStore() outid.OutId {
	return outid.A
}
