package main

import (
	"adventofcode/day4/finder"
	"fmt"
	"os"
)

func main() {
	bytes, err := os.ReadFile("big.txt")
	if err != nil {
		panic("file not found")
	}
	content := string(bytes)
	runeMatrix := finder.ParseMatrix(content)
	nrRows := runeMatrix.GetNrRows()
	for i := range nrRows {
		row := runeMatrix.GetRow(i)
		fmt.Println(string(row))
	}
	fmt.Println(finder.FindXmases(runeMatrix))
}
