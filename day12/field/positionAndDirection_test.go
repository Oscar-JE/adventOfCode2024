package field

import (
	vec "adventofcode/geometry/vec2d"
	"testing"
)

func TestReduce(t *testing.T) {
	loop := []positionAndDirection{{vec.Init(0, 0), vec.Init(1, 0)},
		{vec.Init(1, 0), vec.Init(0, 1)},
		{vec.Init(1, 1), vec.Init(-1, 0)},
		{vec.Init(0, 1), vec.Init(0, -1)}}
	res := nrSides(loop)
	if res != 4 {
		t.Errorf("unexpected number of sides")
	}
}

func TestNotInOrder(t *testing.T) {
	loop := []positionAndDirection{{vec.Init(0, 0), vec.Init(1, 0)},
		{vec.Init(1, 1), vec.Init(-1, 0)},
		{vec.Init(0, 1), vec.Init(0, -1)},
		{vec.Init(1, 0), vec.Init(0, 1)}}
	res := nrSides(loop)
	if res != 4 {
		t.Errorf("unexpected number of sides")
	}
}

func TestWithRedundantLength(t *testing.T) {
	loop := []positionAndDirection{
		{vec.Init(0, 0), vec.Init(1, 0)},
		{vec.Init(1, 0), vec.Init(1, 0)},
		{vec.Init(2, 0), vec.Init(0, 1)},
		{vec.Init(2, 1), vec.Init(-1, 0)},
		{vec.Init(1, 1), vec.Init(-1, 0)},
		{vec.Init(0, 1), vec.Init(0, -1)}}
	res := nrSides(loop)
	if res != 4 {
		t.Errorf("Redundant length not handled correct in nrSides")
	}
}
