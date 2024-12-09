package main

import (
	"adventofcode/day7/evaluator"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	res, operands := parse("big.txt")
	sum := 0
	for i := range res {
		if evaluator.Eval(res[i], operands[i]) {
			sum += res[i]
		}
	}
	fmt.Println(sum)
}

func parse(filename string) ([]int, [][]int) {
	byte, err := os.ReadFile(filename)
	if err != nil {
		panic("file was not found")
	}
	content := string(byte)
	lines := strings.Split(content, "\r\n")
	results := []int{}
	listOfOperands := [][]int{}
	for _, line := range lines {
		resAndNumbers := strings.Split(line, ":")
		res, resError := strconv.Atoi(resAndNumbers[0])
		if resError != nil {
			panic("at leas one result failed to parse")
		}
		results = append(results, res)
		operands := []int{}
		operandsRep := strings.Trim(resAndNumbers[1], " ")
		operandReps := strings.Split(operandsRep, " ")
		for _, opRep := range operandReps {
			operVal, err := strconv.Atoi(opRep)
			if err != nil {
				panic("failed to parse at least one operand")
			}
			operands = append(operands, operVal)
		}
		listOfOperands = append(listOfOperands, operands)
	}
	return results, listOfOperands
}
