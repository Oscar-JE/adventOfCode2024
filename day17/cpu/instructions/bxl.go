package instructions



import "adventofcode/day17/cpu/outid"

type Bxl struct {
	op1 int
	op2 int
}

func InitBxl(register int, operand int) Bxl {
	return Bxl{op1: register, op2: operand}
}

func (a Bxl) Out() int {
	return 0
}

func (a Bxl) ResultStore() outid.OutId {
	return outid.A
}