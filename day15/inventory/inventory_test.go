package inventory

import (
	"adventofcode/day15/directions"
	vec "adventofcode/geometry/vec2d"
	"fmt"
	"testing"
)

func TestFromString(t *testing.T) {
	rep := "###\r\n#.#\r\n#O#\r\n#@#\r\n###"
	inv := FromString(rep)
	if inv.robotPosition != vec.Init(3, 1) {
		t.Errorf("robot parsed to wrong position")
	}
	writen := inv.String()
	expected := "###\r\n#.#\r\n#[]\r\n#@#\r\n###"
	if writen != expected {
		t.Errorf("unexpected to string behaviour")
	}
	fmt.Println(inv)
}

func TestExpandInventory(t *testing.T) {
	rep := "###\r\n#.#\r\n#O#\r\n#@#\r\n###"
	expected := "######\r\n##..##\r\n##[]##\r\n##@.##\r\n######"
	inv := FromString(rep)
	inv.Expand()
	res := inv.String()
	if expected != res {
		t.Errorf("expanded string rep not what we expect")
	}
}

func TestMoveRobotLeft(t *testing.T) {
	rep := "###\r\n#.#\r\n#O#\r\n#@#\r\n###"
	expected := "######\r\n##..##\r\n##[]##\r\n##@.##\r\n######"
	inv := FromString(rep)
	inv.Expand()
	inv.MoveRobot(directions.West())
	res := inv.String()
	if expected != res {
		t.Errorf("move left string rep not what we expect")
	}
}

func TestMoveRobotNorth(t *testing.T) {
	rep := "###\r\n#.#\r\n#O#\r\n#@#\r\n###"
	expected := "######\r\n##[]##\r\n##@.##\r\n##..##\r\n######"
	inv := FromString(rep)
	inv.Expand()
	fmt.Println(inv)
	inv.MoveRobot(directions.North())
	fmt.Println(inv)
	res := inv.String()
	if expected != res {
		t.Errorf("expanded string rep not what we expect")
	}
}
