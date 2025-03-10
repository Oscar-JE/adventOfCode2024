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
	res := MinimumCost(pair, target)
	if res != 280 {
		t.Errorf("wrong number of tokens on first example")
	}
}

func TestEx2(t *testing.T) {
	a := InitBut(vec.Init(26, 66), 3)
	b := InitBut(vec.Init(67, 21), 1)
	pair := ButtonPair{a, b}
	target := vec.Init(12748, 12176)
	res := MinimumCost(pair, target)
	if res != 0 {
		t.Errorf("wrong number of tokens on second example")
	}
}

func TestEx3(t *testing.T) {
	a := InitBut(vec.Init(17, 86), 3)
	b := InitBut(vec.Init(84, 37), 1)
	pair := ButtonPair{a, b}
	target := vec.Init(7870, 6450)
	res := MinimumCost(pair, target)
	if res != 200 {
		t.Errorf("wrong number of tokens on third example")
	}
}

func TestEx4(t *testing.T) {
	a := InitBut(vec.Init(69, 23), 3)
	b := InitBut(vec.Init(27, 71), 1)
	pair := ButtonPair{a, b}
	target := vec.Init(18641, 10279)
	res := MinimumCost(pair, target)
	if res != 0 {
		t.Errorf("wrong number of tokens on forth example")
	}
}

func TestLinearDependent(t *testing.T) {
	a := InitBut(vec.Init(4, 4), 3)
	b := InitBut(vec.Init(1, 1), 1)
	pair := ButtonPair{a, b}
	target := vec.Init(8, 8)
	res := MinimumCost(pair, target)
	if res != 6 {
		t.Errorf("wrong number of tokens on forth example")
	}
}

func TestLinearDependent2(t *testing.T) {
	a := InitBut(vec.Init(4, 4), 5)
	b := InitBut(vec.Init(1, 1), 1)
	pair := ButtonPair{a, b}
	target := vec.Init(4, 4)
	res := MinimumCost(pair, target)
	if res != 4 {
		t.Errorf("wrong number of tokens on linearly dependent 2 example")
	}
}
