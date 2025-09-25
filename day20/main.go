package main

import (
	"adventofcode/day20/cheat"
	"adventofcode/day20/distance"
	"adventofcode/day20/field"
	"fmt"
	"os"
)

func main() {
	bytes, fileErr := os.ReadFile("input.txt")
	if fileErr != nil {
		panic("Are you sure that the file is spelled correct")
	}
	content := string(bytes)
	f := field.Parse(content)
	s := f.FindStart()
	distsEnd := distance.DistanceToEnd(f)
	distsStart := distance.DistanceToStart(f)
	orginalSolve := distsEnd.Get(s.GetX(), s.GetY()).Val()
	cheats := cheat.AllPossibleCheats(distsEnd.GetNrRows(), distsEnd.GetNrCols())
	count := 0
	for _, c := range cheats {
		d := cheat.DistanceOfCheat(distsStart, distsEnd, c)
		if d.Finite() {
			savedSteps := orginalSolve - d.Val()
			if savedSteps >= 100 {
				count++
			}
		}
	}
	fmt.Println(count)
}
