package ram

import (
	vec "adventofcode/geometry/vec2d"
	"fmt"
	"testing"
)

func TestSimpleConstruction(t *testing.T) {
	vals := []vec.Vec2d{vec.Init(0, 0)}
	dim := 1
	stat := InitStationary(dim, vals)
	rep := fmt.Sprint(stat)
	if rep != "#.\n.." {
		t.Errorf("String representation of stationary ram not as expected")
	}
}
