package main

import (
	"adventofcode/day17/cpu"
	"testing"
)

func TestParseProblem(t *testing.T) {
	problemRep := "Register A: 729\r\nRegister B: 0\r\nRegister C: 0\r\n\r\nprogram: 0,1,5,4,3,0\r\n"
	registers, program := ParseRegistersAndProgram(problemRep)
	expectedRegisters := cpu.InitRegisters(729, 0, 0)
	if !registers.Eq(expectedRegisters) {
		t.Errorf("parse problem failed for register part")
	}
	rows := []cpu.InstructionAndOperand{{0, 1}, {5, 4}, {3, 0}}
	expectedProgram := cpu.InitProgram(rows)
	if !program.Eq(expectedProgram) {
		t.Errorf("parse problem failed for program part")
	}
}

func Test1(t *testing.T) {
	registets := cpu.InitRegisters(10, 0, 0)
	program := cpu.InitProgramRaw([]int{5, 0, 5, 1, 5, 4})
	res := Evaluate(registets, program)
	if res != "012" {
		t.Errorf("failed out example1")
	}
}

func Test2(t *testing.T) {
	reg := cpu.InitRegisters(2024, 0, 0)
	prog := cpu.InitProgramRaw([]int{0, 1, 5, 4, 3, 0})
	res := Evaluate(reg, prog)
	if res != "42567777310" {
		t.Errorf(" Evaluate failed example2")
	}
}

func TestFinalExample(t *testing.T) {
	reg := cpu.InitRegisters(729, 0, 0)
	prog := cpu.InitProgramRaw([]int{0, 1, 5, 4, 3, 0})
	res := Evaluate(reg, prog)
	if res != "4635635210" {
		t.Errorf("failed final exampel")
	}
}

func TestMain(t *testing.T) {
	main()
}
