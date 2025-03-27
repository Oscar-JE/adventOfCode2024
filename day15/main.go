package main

import (
	"adventofcode/day15/instructions"
	"adventofcode/day15/inventory"
	vec "adventofcode/geometry/vec2d"
	"fmt"
	"os"
	"strings"
)

func main() {
	bytes, fileErr := os.ReadFile("long.txt")
	if fileErr != nil {
		panic("file was not found")
	}
	content := string(bytes)
	spaceRepAndinstructionRep := strings.Split(content, "\r\n\r\n")
	inv := inventory.FromString(spaceRepAndinstructionRep[0])
	forman := instructions.InitFromRep(spaceRepAndinstructionRep[1])
	for forman.HasNext() {
		direct := forman.GetNext()
		inv.MoveRobot(direct)
	}
	endPositions := inv.PositionsOfMovabel()
	fmt.Println(score(endPositions))
}

func score(positions []vec.Vec2d) int {
	sum := 0
	for _, v := range positions {
		sum += 100*v.GetX() + v.GetY()
	}
	return sum
}
