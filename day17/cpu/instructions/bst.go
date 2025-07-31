package instructions

import "adventofcode/day17/cpu/outid"

type Bst struct {
	op1 int
}

func InitBst(operand int) Bst {
	return Bst{op1: operand}
}

func (b Bst) Out() int {
	return b.op1 % 8
}

func (b Bst) ResultStore() outid.OutId {
	return outid.B
}
