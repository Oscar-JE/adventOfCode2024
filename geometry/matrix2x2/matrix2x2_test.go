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

	if len(solSet.directions) > 0 || solSet.aSolPoint != target {
		t.Errorf("I maps everything to itself")
	}
}

func TestSolveSingleSolutionStretchX(t *testing.T) {
	mat := Init(vec.Init(2, 0), vec.Init(0, 1))
	target := vec.Init(2, 0)
	solSet, _ := mat.Solve(target)

	if len(solSet.directions) > 0 || solSet.aSolPoint != vec.ScalarDiv(2, target) {
		t.Errorf("I maps everything to itself a bit stretched")
	}
}

func TestSolveSingleSolutionStretchY(t *testing.T) {
	mat := Init(vec.Init(1, 0), vec.Init(0, 2))
	target := vec.Init(2, 2)
	solSet, _ := mat.Solve(target)

	if len(solSet.directions) > 0 || solSet.aSolPoint != vec.Init(2, 1) {
		t.Errorf("I maps everything a little stretched")
	}
}

func TestMultipleSolutionOneZeroVec(t *testing.T) {
	mat := Init(vec.Init(1, 0), vec.Init(0, 0))
	target := vec.Init(2, 0)
	solset, _ := mat.Solve(target)
	if len(solset.directions) != 1 {
		t.Errorf("wrong number of basis in the zeroSpace")
	}
	if solset.directions[0] != vec.Init(0, 1) {
		t.Errorf("wrong zerospace")
	}
	if solset.aSolPoint != vec.Init(2, 0) {
		t.Errorf("wrong solutionPoint")
	}
}

func TestMultipleSolutionOneZeroVecShifted(t *testing.T) {
	mat := Init(vec.Init(0, 0), vec.Init(0, 2))
	target := vec.Init(0, 8)
	solset, _ := mat.Solve(target)
	if len(solset.directions) != 1 {
		t.Errorf("wrong number of basis in the zeroSpace")
	}
	if solset.directions[0] != vec.Init(1, 0) {
		t.Errorf("wrong zerospace")
	}
	if solset.aSolPoint != vec.Init(0, 4) {
		t.Errorf("wrong solutionPoint")
	}
}

func TestMultipleSolution(t *testing.T) {
	mat := Init(vec.Init(1, 1), vec.Init(2, 2))
	target := vec.Init(1, 1)
	solSet, _ := mat.Solve(target)
	if solSet.aSolPoint != vec.Init(1, 0) {
		t.Errorf("wrong simple solution")
	}
	if len(solSet.directions) != 1 {
		t.Errorf("wrong dimensions on zerospace")
	}
	if !vec.Parallel(solSet.directions[0], vec.Init(2, -1)) {
		t.Errorf("Span of zerospace is not correct")
	}
}

func TestNoSolutionZero(t *testing.T) {
	mat := Init(vec.Init(0, 0), vec.Init(0, 0))
	target := vec.Init(1, 1)
	_, err := mat.Solve(target)
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestNoSolutionHalfZero(t *testing.T) {
	mat := Init(vec.Init(1, 0), vec.Init(0, 0))
	target := vec.Init(1, 1)
	_, err := mat.Solve(target)
	if err == nil {
		t.Errorf("expected error")
	}
}
