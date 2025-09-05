package main

import (
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
	fmt.Println(distsStart.Get(s.GetX(), s.GetY()))
	fmt.Println(distsEnd.Get(s.GetX(), s.GetY()))
}
