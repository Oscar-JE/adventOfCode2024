package main

import (
	"adventofcode/day16/model"
	"fmt"
	"os"
)

func main() {
	println("hej")
	bytes, error := os.ReadFile("input_exe.txt")
	if error != nil {
		println("wtf filen är borta")
	}
	content := string(bytes)
	m := model.Init(content)
	fmt.Println(m.StartState())
}
