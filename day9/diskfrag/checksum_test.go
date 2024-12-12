package diskfrag

import "testing"

func TestCheckSum(t *testing.T) {
	input := []int{0, 0, 9, 9, 8, 1, 1, 1, 8, 8, 8, 2, 7, 7, 7, 3, 3, 3, 6, 4, 4, 6, 5, 5, 5, 5, 6, 6}
	sum := CheckSum(input)
	if sum != 1928 {
		t.Errorf("huston we got a problem ")
	}
}
