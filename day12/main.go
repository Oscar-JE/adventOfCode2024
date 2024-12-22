package main

import (
	"adventofcode/day12/field"
	"fmt"
	"os"
	"strings"
)

func main() {
	f := parse("short.txt")
	fmt.Println(f.Score())
}

func parse(fileName string) field.Field {
	bytes, err := os.ReadFile("short.txt")
	if err != nil {
		panic("error wile reading file")
	}
	content := string(bytes)
	lines := strings.Split(content, "\n")
	f := field.ParseFromLines(lines)
	return f
}
