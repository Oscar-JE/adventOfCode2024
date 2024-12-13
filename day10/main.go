package main

import (
	"adventofcode/day10/solver"
	"adventofcode/geometry/matrix"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("big.txt")
	if err != nil {
		panic("file reading failed")
	}
	content := string(file)
	lines := strings.Split(content, "\r\n") // om jag kan skapa en matris från en lista av strängar görs detta enklare
	nrRows := len(lines)
	aLine := []rune(lines[0])
	nrCols := len(aLine)
	contentWithoutNewLines := strings.Replace(content, "\r\n", "", -1)
	vals := []int{}
	for _, r := range contentWithoutNewLines {
		numRep := string(r)
		num, err := strconv.Atoi(numRep)
		if err != nil {
			panic("interpretation of number failed")
		}
		vals = append(vals, num)
	}
	mat := matrix.Init(vals, nrRows, nrCols)
	fmt.Println(mat)
	s := solver.Init(mat)
	fmt.Println(s.Solve2())
}
