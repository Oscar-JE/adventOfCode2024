package main

import (
	"adventofcode/day12/field"
	"fmt"
	"os"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("big.txt")
	if err != nil {
		panic("error wile reading file")
	}
	content := string(bytes)
	lines := strings.Split(content, "\n")
	f := field.ParseFromLines(lines)
	fmt.Println(f.Score())
}
