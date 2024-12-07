package finder

import (
	vec "adventofcode/geometry/vec2d"
	"testing"
)

func TestReadDirectionDown(t *testing.T) {
	mat := ParseMatrix("12\r\n34")
	startingPoint := vec.Init(0, 0)
	direction := vec.Init(1, 0)
	res := readDirection(startingPoint, direction, mat)
	if res != "13" {
		t.Errorf("error in reading down")
	}
}

func TestReadDiagonal(t *testing.T) {
	mat := ParseMatrix("12\r\n34")
	startingPoint := vec.Init(0, 0)
	direction := vec.Init(1, 1)
	res := readDirection(startingPoint, direction, mat)
	if res != "14" {
		t.Errorf("error in reading down")
	}
}

func TestFirstDIagonal(t *testing.T) {
	mat := ParseMatrix("12\r\n34")
	res := firstDiagonal(mat)
	expected := []string{"1", "32", "4", "1", "23", "4"}
	if len(res) != len(expected) {
		t.Errorf("length unexpected on first diagonal")
	}
}

func TestFirstDIagonalBigger(t *testing.T) {
	mat := ParseMatrix("123\r\n456\r\n789")
	res := firstDiagonal(mat)
	expected := []string{"1", "32", "4", "1", "23", "4"}
	if len(res) != len(expected) {
		t.Errorf("length unexpected on first diagonal")
	}
}

func TestSecondDiagonal(t *testing.T) {
	mat := ParseMatrix("12\r\n34")
	res := secondDiagonal(mat)
	expected := []string{"2", "14", "3", "2", "41", "3"}
	if len(res) != len(expected) {
		t.Errorf("length unexpected on first diagonal")
	}
}

func TestFindOcuranceInLine(t *testing.T) {
	res := findOccurrenceInLine("MMMSXXMASM", "XMAS")
	if res != 1 {
		t.Errorf("wrong")
	}
}
