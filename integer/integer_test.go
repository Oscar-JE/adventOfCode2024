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

func TestPowerOf(t *testing.T) {
	res := ToThePowerOf(2, 3)
	if res != 8 {
		t.Errorf("to the power of fails for input 2,3")
	}
}

func TestXORcancelingOut(t *testing.T) {
	inputs := []int{0, 1, 2, 3, 4, 10, 17, 185}
	for _, v := range inputs {
		if XOR(v, v) != 0 {
			t.Errorf("XOR: canceling does not work for %d", v)
		}
	}
}

func TestXORaditives(t *testing.T) {
	a := 2 + 4 + 16
	b := 2 + 4 + 8 + 32
	expected := 8 + 16 + 32
	if XOR(a, b) != expected {
		t.Errorf("examine xor additions")
	}
}
