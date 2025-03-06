package matrix2x2

import (
	vec "adventofcode/geometry/vec2d"
	"testing"
)

func TestHasSolutionZeroMatrix(t *testing.T) {
	mat := Init(vec.Init(0, 0), vec.Init(0, 0))
	if mat.HasSolution(vec.Init(1, 0)) {
		t.Errorf(" (1,0) is not in the span of (0,0) and (0,0)")
	}
	if !mat.HasSolution(vec.Init(1, 0)) {
		t.Errorf(" (0,0) is in the span of (0,0) and (0,0)")
	}
}
