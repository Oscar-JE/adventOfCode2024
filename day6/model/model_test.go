package model

import "testing"

func TestIsLoopZeroZero(t *testing.T) {
	mat := Parse("../short.txt")
	model := Init(mat)
	if model.IsLoop(0, 0) {
		t.Errorf("It does not even take part in the first walk")
	}
}

func TestIsLoopFirstLoop(t *testing.T) {
	mat := Parse("../short.txt")
	model := Init(mat)
	if !model.IsLoop(6, 3) {
		t.Errorf("This Should be a loop")
	}
}

func TestIsLoopOneFour(t *testing.T) {
	mat := Parse("../short.txt")
	model := Init(mat)
	if model.IsLoop(1, 4) {
		t.Errorf("This Should not be a loop")
	}
}
