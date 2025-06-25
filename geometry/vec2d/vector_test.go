package vec

import (
	"testing"
)

func TestL1Dist(t *testing.T) {
	v1 := Init(3, -2)
	if L1Dist(v1) != 5 {
		t.Errorf("examine behaviour of l1Dist")
	}
}
