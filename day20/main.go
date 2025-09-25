package main

import (
	"adventofcode/day20/cheat"
	"adventofcode/day20/distance"
	"adventofcode/day20/field"
	"fmt"
	"os"
)

func main() {
	bytes, fileErr := os.ReadFile("input_example.txt")
	if fileErr != nil {
		panic("Are you sure that the file is spelled correct")
	}
	content := string(bytes)
	f := field.Parse(content)
	s := f.FindStart()
	distsEnd := distance.DistanceToEnd(f)
	distsStart := distance.DistanceToStart(f)
	fmt.Println("start")
	fmt.Println(distsStart)
	fmt.Println("end")
	fmt.Println(distsEnd)
	fmt.Println(distsStart.Get(s.GetX(), s.GetY()))
	fmt.Println(distsEnd.Get(s.GetX(), s.GetY()))
	orginalSolve := distsEnd.Get(s.GetX(), s.GetY()).Val()
	cheats := cheat.AllPossibleCheats(distsEnd.GetNrRows(), distsEnd.GetNrCols())
	saved := []int{}
	for _, c := range cheats {
		d := cheat.DistanceOfCheat(distsStart, distsEnd, c)
		if d.Finite() {
			savedSteps := orginalSolve - d.Val()
			if savedSteps > 0 {
				saved = append(saved, savedSteps)
			}
		}
	}
	// sammanfatta saved
	printConcentrated(saved)
	fmt.Println(saved)
}

func printConcentrated(is []int) {
	seen := []int{}
	for _, el := range is {
		if in(seen, el) {
			continue
		}
		count := 0
		for _, elInner := range is {
			if elInner == el {
				count++
			}
		}
		fmt.Printf("%d occurs %d times \n", el, count)
		seen = append(seen, el)
	}
}

func in(list []int, el int) bool {
	for _, i := range list {
		if i == el {
			return true
		}
	}
	return false
}
