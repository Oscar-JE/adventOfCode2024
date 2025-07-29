package instructions

import "adventofcode/day17/cpu/outid"

type Bst struct {
	op1 int
}

func InitBst(operand int) Adv {
	return Adv{op1: operand}
}

func (b Bst) Out() int {
	return 0
}

func (b Bst) ResultStore() outid.OutId {
	return outid.B
}
