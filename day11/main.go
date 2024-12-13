package main

import (
	"adventofcode/day11/stones"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic("file was not found")
	}
	content := string(bytes)
	content = strings.Replace(content, "\r\n", "", 1)
	numReps := strings.Split(content, " ")
	stonesList := []int{}
	for _, rep := range numReps {
		stone, err := strconv.Atoi(rep)
		if err != nil {
			panic("conversion to int failed")
		}
		stonesList = append(stonesList, stone)
	}
	blinker := stones.InitBlinker()
	fmt.Println(blinker.NrOfStonesAfterBlinks(stonesList, 75))
}
