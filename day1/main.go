package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("big.txt")
	if err != nil {
		panic("error in file reading")
	}
	content := string(bytes)
	lines := strings.Split(content, "\r\n")
	l1 := []int{}
	l2 := []int{}
	for _, line := range lines {
		num1, num2, _ := strings.Cut(line, " ")
		n1, err1 := strconv.Atoi(strings.Trim(num1, " "))
		n2, err2 := strconv.Atoi(strings.Trim(num2, " "))
		if err1 != nil || err2 != nil {
			panic("number parse failed")
		}
		l1 = append(l1, n1)
		l2 = append(l2, n2)
	}
	sort.Ints(l1)
	sort.Ints(l2)
	sum := 0
	for i, _ := range l1 {
		multiplier := occurrenceInList(l1[i], l2)
		sum += multiplier * l1[i]
	}
	fmt.Println(sum)
}

func occurrenceInList(val int, list []int) int {
	sum := 0
	for _, el := range list {
		if el == val {
			sum++
		}
	}
	return sum
}
