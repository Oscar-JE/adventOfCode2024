package main

import (
	"adventofcode/day19/towels"
	"fmt"
	"os"
	"strings"
)

func main() {
	rack, patterns := parseInputFile("input.txt")
	sum := 0
	for _, p := range patterns {
		if rack.SolvablePattern(p) {
			sum++
		}
	}
	fmt.Println(sum)
}

func parseInputFile(name string) (towels.TowelRack, []towels.Pattern) {
	bytes, err := os.ReadFile(name)
	if err != nil {
		panic("cannot find file")
	}
	rep := string(bytes)
	return parseInput(rep)
}

func parseInput(rep string) (towels.TowelRack, []towels.Pattern) {
	splittedRackAndPatterns := strings.Split(rep, "\r\n\r\n")
	return towels.ParseTowelRack(splittedRackAndPatterns[0]), parsePatterns(splittedRackAndPatterns[1])
}

func parsePatterns(rep string) []towels.Pattern {
	ret := []towels.Pattern{}
	patternReps := strings.Split(rep, "\r\n")
	for _, rep := range patternReps {
		ret = append(ret, towels.ParsePattern(rep))
	}
	return ret
}
