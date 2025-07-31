package instructions

import "adventofcode/day17/cpu/outid"

type Jnz struct {
	op1     int
	op2     int
	outSite outid.OutId
}

func InitJnz(register int, operand int) Jnz {
	return Jnz{op1: register, op2: operand, outSite: outid.InstructionPointer}
}

func (a *Jnz) Out() int {
	if a.op1 == 0 {
		a.outSite = outid.A
		return 0
	}
	a.outSite = outid.InstructionPointer
	return a.op2
}

func (a *Jnz) ResultStore() outid.OutId {
	return a.outSite
}
