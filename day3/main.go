package main

import (
	"day3/parser"
	"fmt"
	"os"
)

func main() {
	bytes, err := os.ReadFile("big.txt")
	if err != nil {
		panic("file not found")
	}
	content := string(bytes)
	score := parser.ParseWithExtraCommands(content)
	fmt.Println(score)
}
