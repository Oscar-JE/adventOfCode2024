package cpu

import (
	"adventofcode/day17/cpu/instructions"
	"fmt"
	"strconv"
	"strings"
)

type Registers struct {
	regs []int
}

func InitRegisters(A int, B int, C int) Registers {
	regs := []int{A, B, C}
	return Registers{regs: regs}
}

func (r Registers) GetA() int {
	return r.regs[0]
}

func (r *Registers) SetA(a int) {
	r.regs[0] = a
}

func (r Registers) GetB() int {
	return r.regs[1]
}

func (r *Registers) SetB(b int) {
	r.regs[1] = b
}

func (r Registers) GetC() int {
	return r.regs[2]
}

func (r *Registers) SetC(c int) {
	r.regs[2] = c
}

func (r Registers) String() string {
	return fmt.Sprintf("A : %d, B: %d , C: %d", r.regs[0], r.regs[1], r.regs[2])
}

func ParseRegisters(input string) Registers {
	rows := strings.Split(input, "\r\n")
	reg := []int{0, 0, 0}
	for i, row := range rows {
		rep := strings.Split(row, ": ")[1]
		val, err := strconv.Atoi(rep)
		reg[i] = val
		if err != nil {
			panic("error in parsing initial registers")
		}
	}

	return Registers{reg}
}

func (r Registers) ComboOperandToValue(combo int) int {
	if combo < 0 {
		panic("no combo operand below 0 is allowed")
	}
	if combo < 4 {
		return combo
	}
	if combo >= 4 && combo < 7 {
		return r.regs[combo-4]
	}
	panic("no combo operand above 6 is allowed")

}

func (r Registers) Eq(other Registers) bool {
	for i, v := range r.regs {
		if v != other.regs[i] {
			return false
		}
	}
	return true
}

func (r Registers) CreateInstructionAndInput(opcode int, operand int) InstructionAndInput {
	if opcode == 0 {
		return instructions.InitAdv(r.GetA(), r.ComboOperandToValue(operand))
	}
	if opcode == 1 {
		return instructions.InitBxl(r.GetB(), operand)
	}
	if opcode == 2 {
		return instructions.InitBst(r.ComboOperandToValue(operand))
	}
	if opcode == 3 {
		j := instructions.InitJnz(r.GetA(), operand)
		return &j
	}
	if opcode == 4 {
		return instructions.InitBxc(r.GetB(), r.GetC())
	}
	if opcode == 5 {
		return instructions.InitOut(r.ComboOperandToValue(operand))
	}
	if opcode == 6 {
		return instructions.InitBdv(r.GetA(), r.ComboOperandToValue(operand))
	}
	if opcode == 7 {
		return instructions.InitCdv(r.GetA(), r.ComboOperandToValue(operand))
	}
	panic("no matching instructions")
}
