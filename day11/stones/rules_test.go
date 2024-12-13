package stones

import "testing"

func TestEvenSplit(t *testing.T) {
	s := SplitEvenNrDigits{}
	stone := 1
	if s.applicable(stone) {
		t.Errorf("is not applicable")
	}
	if !s.applicable(40) {
		t.Errorf("should be applicable")
	}
}

func TestSplitApply(t *testing.T) {
	s := SplitEvenNrDigits{}
	stone := 2533
	list := s.apply(stone)
	i1, i2 := list[0], list[1]
	if i1 != 25 || i2 != 33 {
		t.Errorf("error in split")
	}
}
