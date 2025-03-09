package claw

import (
	vec "adventofcode/geometry/vec2d"
	"testing"
)

func TestEx1(t *testing.T) {
	a := InitBut(vec.Init(94, 34), 3)
	b := InitBut(vec.Init(22, 67), 1)
	pair := ButtonPair{a, b}
	target := vec.Init(8400, 5400)
	res := minimumCost(pair, target)
	if res != 280 {
		t.Errorf("wrong number of tokens on first example")
	}
}
