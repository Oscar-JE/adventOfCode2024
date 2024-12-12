package main

import (
	"adventofcode/day9/diskfrag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	numbers := parse("input.txt")
	fmt.Println(part2(numbers))
}

func part1(numbers []int) int {
	cons := diskfrag.Fragment(numbers)
	return diskfrag.CheckSum(cons)
}

func part2(numbers []int) int {
	cons := diskfrag.FragmentFiles(numbers)
	return diskfrag.CheckSum(cons)
}

func parse(filename string) []int {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		panic("fileread failed")
	}
	content := string(bytes)
	return parseFromString(content)
}

func parseFromString(content string) []int {
	numRuneRep := []rune(content)
	numbers := []int{}
	for _, r := range numRuneRep {
		n, err := strconv.Atoi(string(r))
		if err != nil {
			panic("conversion to int failed")
		}
		numbers = append(numbers, n)
	}
	return numbers
}
