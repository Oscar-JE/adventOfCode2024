package inventory

import (
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
	if writen != rep {
		t.Errorf("to string not inverse of from string")
	}
	fmt.Println(inv)
}
