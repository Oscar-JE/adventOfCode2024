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
	if !mat.HasSolution(vec.Init(0, 0)) {
		t.Errorf(" (0,0) is in the span of (0,0) and (0,0)")
	}
}

func TestHasMultipleSolutionsOneSameDirektion(t *testing.T) {
	mat1 := Init(vec.Init(1, 0), vec.Init(0, 0))
	mat2 := Init(vec.Init(0, 0), vec.Init(1, 0))
	target := vec.Init(2, 0)
	if !mat1.HasMultipleSolutions(target) {
		t.Errorf("thihi")
	}
	if !mat2.HasMultipleSolutions(target) {
		t.Errorf("thihihoho")
	}
}

func TestHasMultipleSolutionsBothSameDirection(t *testing.T) {
	mat := Init(vec.Init(1, 0), vec.Init(2, 0))
	target := vec.Init(2, 0)
	if !mat.HasMultipleSolutions(target) {
		t.Errorf("thihi")
	}
}

func TestHasMultipleSolutionsOneSolution(t *testing.T) {
	mat := Init(vec.Init(1, 0), vec.Init(0, 1))
	target := vec.Init(2, 0)
	if mat.HasMultipleSolutions(target) {
		t.Errorf("there can only be one")
	}
}

func TestSolveSingleSolutionI(t *testing.T) {
	mat := Init(vec.Init(1, 0), vec.Init(0, 1))
	target := vec.Init(2, 0)
	solSet, _ := mat.Solve(target)

	if solSet.direction.NonZero() || solSet.aSolPoint != target {
		t.Errorf("I maps everything to itself")
	}
}

func TestSolveSingleSolutionStretchX(t *testing.T) {
	mat := Init(vec.Init(2, 0), vec.Init(0, 1))
	target := vec.Init(2, 0)
	solSet, _ := mat.Solve(target)

	if solSet.direction.NonZero() || solSet.aSolPoint != vec.ScalarDiv(2, target) {
		t.Errorf("I maps everything to itself")
	}
}

func TestSolveSingleSolutionStretchY(t *testing.T) {
	mat := Init(vec.Init(1, 0), vec.Init(0, 2))
	target := vec.Init(2, 2)
	solSet, _ := mat.Solve(target)

	if solSet.direction.NonZero() || solSet.aSolPoint != vec.Init(2, 1) {
		t.Errorf("I maps everything to itself")
	}
}
