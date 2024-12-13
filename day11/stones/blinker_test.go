package stones

import (
	"testing"
)

func TestExampelFromInstruction(t *testing.T) {
	b := InitBlinker()
	stones := []int{125, 17}
	res := b.NrOfStonesAfterBlinks(stones, 25)
	if res != 55312 {
		t.Errorf("examine NrOfStonesAfterBlinks")
	}
}
