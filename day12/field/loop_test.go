package field

import (
	vec "adventofcode/geometry/vec2d"
	"testing"
)

func TestIsLoop(t *testing.T) {
	l := loop{[]link{{vec.Init(0, 0), vec.Init(1, 1)}}}
	if l.isLoop() {
		t.Errorf("A sigle link that do not end in itself is not a loop")
	}
}

func TestIsLoopClosed(t *testing.T) {
	l := loop{[]link{{vec.Init(0, 0), vec.Init(1, 1)}, {vec.Init(1, 1), vec.Init(0, 0)}}}
	if !l.isLoop() {
		t.Errorf("should be defined as loop")
	}
}

func TestNrOfSidesSimple(t *testing.T) {
	l := loop{[]link{{vec.Init(0, 0), vec.Init(1, 1)}, {vec.Init(1, 1), vec.Init(0, 1)}, {vec.Init(0, 1), vec.Init(0, 0)}}}
	if !l.isLoop() {
		t.Errorf("triangel is a closed loop")
	}
	expected := 3
	res := l.nrSides()
	if expected != res {
		t.Errorf("A triangle has tree sides")
	}

}
