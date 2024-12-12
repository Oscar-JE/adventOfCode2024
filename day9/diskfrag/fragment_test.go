package diskfrag

import "testing"

func TestExpand(t *testing.T) {
	numRep := []int{1, 2, 3, 4, 5}
	expected := []Block{{0, false},
		{0, true}, {0, true},
		{1, false}, {1, false}, {1, false},
		{0, true}, {0, true}, {0, true}, {0, true},
		{2, false}, {2, false}, {2, false}, {2, false}, {2, false}}
	res := expand(numRep)
	if len(res) != len(expected) {
		t.Errorf("unexpected length")
	}
	eq := true
	for i := range expected {
		eq = eq && expected[i].Eq(res[i])
	}
	if !eq {
		t.Errorf("result of expanded not what to be expected")
	}
}

func TestContract(t *testing.T) {
	blocks := []Block{{0, false},
		{0, true}, {0, true},
		{1, false}, {1, false}, {1, false},
		{0, true}, {0, true}, {0, true}, {0, true},
		{2, false}, {2, false}, {2, false}, {2, false}, {2, false}}
	expected := []int{0, 2, 2, 1, 1, 1, 2, 2, 2}
	res := contract(blocks)
	if len(expected) != len(res) {
		t.Errorf("unexpected length after contracting")
	}
	eq := true
	for i := range expected {
		eq = eq && expected[i] == res[i]
	}
	if !eq {
		t.Errorf("contracted to something we did not expect")
	}
}

func TestFindFileBefore(t *testing.T) {
	blocks := []Block{{0, false},
		{0, true}, {0, true},
		{1, false}, {1, false}, {1, false},
		{0, true}, {0, true}, {0, true}, {0, true},
		{2, false}, {2, false}, {2, false}, {2, false}, {2, false}}
	res := findFileBefore(len(blocks), blocks)
	expected := segment{10, 5}
	if res != expected {
		t.Errorf("examine find blocks before")
	}
}

func TestFindFileBeforeMiddle(t *testing.T) {
	blocks := []Block{{0, false},
		{0, true}, {0, true},
		{1, false}, {1, false}, {1, false},
		{0, true}, {0, true}, {0, true}, {0, true},
		{2, false}, {2, false}, {2, false}, {2, false}, {2, false}}
	res := findFileBefore(7, blocks)
	expected := segment{3, 3}
	if res != expected {
		t.Errorf("blä")
	}
}

func TestFindFileBeforeStart(t *testing.T) {
	blocks := []Block{{0, false},
		{0, true}, {0, true},
		{1, false}, {1, false}, {1, false},
		{0, true}, {0, true}, {0, true}, {0, true},
		{2, false}, {2, false}, {2, false}, {2, false}, {2, false}}
	res := findFileBefore(0, blocks)
	expected := segment{0, 0}
	if res != expected {
		t.Errorf("blev det en panic?")
	}
}
