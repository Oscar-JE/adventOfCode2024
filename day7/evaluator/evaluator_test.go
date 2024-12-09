package evaluator

import "testing"

func TestEvale1(t *testing.T) {
	if !Eval(190, []int{10, 19}) {
		t.Errorf("not expected")
	}
}

func TestEvale3(t *testing.T) {
	if Eval(192, []int{17, 8, 14}) {
		t.Errorf("wtf")
	}
}

func TestAppendOperator(t *testing.T) {
	if 7834 != appendOperator(78, 34) {
		t.Errorf("append operator needs more love")
	}
}
