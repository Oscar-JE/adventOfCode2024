package antennas

import (
	vec "adventofcode/geometry/vec2d"
	"testing"
)

func TestAntinodePair(t *testing.T) {
	res := antinodePair(vec.Init(3, 3), vec.Init(4, 4))
	if res[0] != vec.Init(2, 2) {
		t.Errorf("antinode pair element one missed the mark")
	}
	if res[1] != vec.Init(5, 5) {
		t.Errorf("antinode pair second element missed the mark")
	}
}

func TestAntiNodesImpossibleInput(t *testing.T) {
	res := antinodePair(vec.Init(2, 4), vec.Init(2, 4))
	if res[0] != res[1] || res[0] != vec.Init(2, 4) {
		t.Errorf("antinode pair for stupid input ")
	}
}
