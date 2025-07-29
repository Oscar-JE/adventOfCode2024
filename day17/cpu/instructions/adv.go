package instructions

import "adventofcode/day17/cpu/outid"

type Adv struct {
	op1 int
	op2 int
}

func InitAdv(register int, operand int) Adv {
	return Adv{op1: register, op2: operand}
}

func (a Adv) Out() int {
	return 0
}

func (a Adv) ResultStore() outid.OutId {
	return outid.A
}
