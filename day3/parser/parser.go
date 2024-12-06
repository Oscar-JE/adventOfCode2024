package parser

import (
	"strconv"
	"strings"
)

func ParseWithExtraCommands(mem string) int {
	//antar att man inte ska kunna kombinera en
	// rad innan dont med det efter dont
	memList := preParse(mem)
	sum := 0
	for _, section := range memList {
		sum += Parse(section)
	}
	return sum
}

func preParse(mem string) []string {
	memList := []string{}
	startsWithDo := strings.Split(mem, "do()")
	for _, candidate := range startsWithDo {
		endsWithDont := strings.Split(candidate, "don't()")[0]
		memList = append(memList, endsWithDont)
	}
	return memList
}

func Parse(mem string) int {
	sum := 0
	mulAtBeginnings := strings.Split(mem, "mul(")
	for _, mumulAtBeginning := range mulAtBeginnings {
		endWithClosing := strings.Split(mumulAtBeginning, ")")[0]
		numbers := strings.Split(endWithClosing, ",")
		if len(numbers) != 2 {
			continue
		}
		firstNum, firstError := strconv.Atoi(numbers[0])
		if firstError != nil {
			continue
		}
		secondNum, seconError := strconv.Atoi(numbers[1])
		if seconError != nil {
			continue
		}
		sum += firstNum * secondNum
	}
	return sum
}
