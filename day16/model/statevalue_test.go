package model

import (
	vec "adventofcode/geometry/vec2d"
	"testing"
)

func TestPositionIndexAndInverse(t *testing.T) {
	sv := InitValueFunc(2, 3)
	expectedPos := []vec.Vec2d{vec.Init(0, 0), vec.Init(0, 1), vec.Init(0, 2), vec.Init(1, 0), vec.Init(1, 1), vec.Init(1, 2)}
	for i, expected := range expectedPos {
		position := sv.positionFromIndex(i)
		if expected != position {
			t.Errorf("position from index failed")
		}
		index := sv.indexFromPosition(expected)
		if index != i {
			t.Errorf("index from position failed")
		}
	}
}
