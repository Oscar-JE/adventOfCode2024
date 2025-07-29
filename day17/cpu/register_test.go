package cpu

import (
	"testing"
)

func TestRegisterRepresentation(t *testing.T) {
	t.Skip("checked and have satisfaktory output")
	r := InitRegisters(33, 23, 1)
	t.Log(r)
}

func TestParseRegisters(t *testing.T) {
	parseString := "Register A: 1\r\nRegister B: 2\r\nRegister C: 3"
	parsed := ParseRegisters(parseString)
	expected := InitRegisters(1, 2, 3)
	if !parsed.Eq(expected) {
		t.Errorf("expected: %v  but was %v", expected, parsed)
	}
}
