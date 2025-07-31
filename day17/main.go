package main

import (
	"adventofcode/day17/cpu"
	"adventofcode/day17/cpu/outid"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputRep := ReadFile("input.txt")
	reg, prog := ParseRegistersAndProgram(inputRep)
	fmt.Println(Evaluate(reg, prog))
}

func ReadFile(name string) string {
	bytes, err := os.ReadFile(name)
	if err != nil {
		panic("failed to read file :" + name)
	}
	content := string(bytes)
	return content
}

func ParseRegistersAndProgram(rep string) (cpu.Registers, cpu.Program) {
	splited := strings.Split(rep, "\r\n\r\n")
	registerRep := splited[0]
	registers := cpu.ParseRegisters(registerRep)
	programRep := splited[1]
	program := cpu.ParseProgram(programRep)
	return registers, program
}

func Evaluate(reg cpu.Registers, prog cpu.Program) string {
	outPut := ""
	for prog.HasInstructionAndOperand() {
		codes := prog.GetInstructionAndOperand()
		inst := reg.CreateInstructionAndInput(codes.InstructionCode, codes.Operand)
		result := inst.Out()
		outputType := inst.ResultStore()
		if outputType == outid.A {
			reg.SetA(result)
		} else if outputType == outid.B {
			reg.SetB(result)
		} else if outputType == outid.C {
			reg.SetC(result)
		} else if outputType == outid.InstructionPointer {
			prog.SetProgramPointer(result)
		} else {
			outPut += strconv.Itoa(result)
		}
	}
	return outPut
}
