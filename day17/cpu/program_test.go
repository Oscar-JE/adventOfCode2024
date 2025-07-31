package cpu

import "testing"

func TestParseProgram(t *testing.T) {
	input := "Program: 2,4,1,2"
	expected := InitProgram([]InstructionAndOperand{{2,4},{1,2}})
	parsed := ParseProgram(input)
	if ! expected.Eq(parsed) {
		t.Errorf("")
	}
}