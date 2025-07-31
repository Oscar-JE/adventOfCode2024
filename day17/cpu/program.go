package cpu

import (
	"strconv"
	"strings"
)

type InstructionAndOperand struct {
	InstructionCode int
	Operand         int
}

type Program struct {
	code           []int
	programPointer int
}

func InitProgram(code []InstructionAndOperand) Program {
	codeS := []int{}
	for _, row := range code {
		codeS = append(codeS, row.InstructionCode)
		codeS = append(codeS, row.Operand)
	}
	return Program{code: codeS, programPointer: 0}
}

func InitProgramRaw(codes []int) Program {
	return Program{code: codes, programPointer: 0}
}

func (p Program) HasInstructionAndOperand() bool {
	return p.programPointer < len(p.code)-1
}

func (p *Program) GetInstructionAndOperand() InstructionAndOperand {
	instruction := p.code[p.programPointer]
	p.programPointer++
	input := p.code[p.programPointer]
	p.programPointer++
	return InstructionAndOperand{InstructionCode: instruction, Operand: input}
}

func (p *Program) SetProgramPointer(pointer int) {
	p.programPointer = pointer
}

func ParseProgram(rep string) Program {
	cleaned := strings.Replace(rep, "\r\n", "", 1)
	split := strings.Split(cleaned, ": ")
	numbersRep := split[1]
	numberReps := strings.Split(numbersRep, ",")
	numbers := []int{}
	for _, rep := range numberReps {
		a, err := strconv.Atoi(rep)
		if err != nil {
			panic("failed to parse individual number in parse program")
		}
		numbers = append(numbers, a)
	}
	instructionAndOperands := []InstructionAndOperand{}
	var i int = 0
	for i < len(numbers)-1 {
		instructionAndOperands = append(instructionAndOperands, InstructionAndOperand{numbers[i], numbers[i+1]})
		i += 2
	}
	return InitProgram(instructionAndOperands)
}

func (p Program) Eq(other Program) bool {
	if p.programPointer != other.programPointer || len(p.code) != len(other.code) {
		return false
	}
	for i, row := range p.code {
		if row != other.code[i] {
			return false
		}
	}
	return true
}
