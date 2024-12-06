package main

import (
	"day2/report"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, errFile := os.ReadFile("big.txt")
	if errFile != nil {
		panic("file was not found")
	}
	content := string(bytes)
	lines := strings.Split(content, "\r\n")
	sum := 0
	for _, line := range lines {
		rep := parseRep(line)
		if report.IsSafeWithOneException(rep) {
			sum++
		}
	}
	fmt.Println(sum)
}

func parseRep(strRep string) []int {
	numReps := strings.Split(strRep, " ")
	seq := []int{}
	for _, numRep := range numReps {
		entry, err := strconv.Atoi(numRep)
		if err != nil {
			panic("entry in report could not be read")
		}
		seq = append(seq, entry)
	}
	return seq
}
