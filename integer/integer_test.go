package integer

import "testing"

func TestPositiveNumberOfDigits(t *testing.T) {
	in := 125
	digits := NumberOfDigits(in)
	if digits != 3 {
		t.Errorf("expected 3 digits")
	}
}

func TestNegativeNumberOfDigits(t *testing.T) {
	in := -1000
	digits := NumberOfDigits(in)
	if digits != 4 {
		t.Errorf("expected 4 digits")
	}
}
