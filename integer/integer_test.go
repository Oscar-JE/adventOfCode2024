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

func TestGCD(t *testing.T) {
	if GCD(3, 5) != 1 {
		t.Errorf("DIfferent Primes have a gcd of 1 ")
	}
	if GCD(50, 30) != 10 {
		t.Errorf(" GCD 50,30 should be 10")
	}
}
