package parser

import "testing"

func TestPreParse(t *testing.T) {
	mem := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	res := preParse(mem)
	expected := []string{"xmul(2,4)&mul[3,7]!^",
		"?mul(8,5))"}
	if len(res) != len(expected) {
		t.Errorf("res don't have the expected length")
	}
	equal := true
	for i, _ := range expected {
		equal = equal && expected[i] == res[i]
	}
	if !equal {
		t.Errorf("result is not what is expected")
	}
}

func TestParseWithExtraCommands(t *testing.T) {
	mem := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	res := ParseWithExtraCommands(mem)
	if res != 48 {
		t.Errorf("wtf")
	}
}
