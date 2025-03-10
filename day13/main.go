package main

import (
	"adventofcode/day13/claw"
	vec "adventofcode/geometry/vec2d"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	problems := parse("long.txt")
	sum := 0
	for _, bp := range problems {
		sum += claw.MinimumCost(bp.buttons, bp.prize)
	}
	fmt.Println(sum)
}

func parse(fileName string) []problem {
	content, filerr := os.ReadFile(fileName)
	if filerr != nil {
		panic("File was not found")
	}
	strRep := string(content)
	blocks := strings.Split(strRep, "\r\n\r\n")
	problems := []problem{}
	for _, block := range blocks {
		problems = append(problems, parseProb(block))
	}
	return problems
}

type problem struct {
	buttons claw.ButtonPair
	prize   vec.Vec2d
}

func parseProb(rep string) problem {
	lines := strings.Split(rep, "\r\n")
	movementA := parseVec(lines[0])
	movementB := parseVec(lines[1])
	prize := parseVec(lines[2])
	bp := claw.InitButtonPai(claw.InitBut(movementA, 3), claw.InitBut(movementB, 1))
	return problem{buttons: bp, prize: vec.Add(vec.Init(10000000000000, 10000000000000), prize)}
}

func parseVec(rep string) vec.Vec2d {
	numbersAndExtras := strings.Split(rep, ",")
	firstCoordAndExtras := numbersAndExtras[0]
	secondCoordAndExtras := numbersAndExtras[1]
	firstCoordrep := stripAwayNotNumbers(firstCoordAndExtras)
	secodndCoordrep := stripAwayNotNumbers(secondCoordAndExtras)
	num0, err0 := strconv.Atoi(firstCoordrep)
	num1, err1 := strconv.Atoi(secodndCoordrep)
	if err0 != nil || err1 != nil {
		panic("conversion of coordinates failed")
	}
	return vec.Init(num0, num1)
}

func stripAwayNotNumbers(in string) string {
	outStr := ""
	for _, ch := range in {
		_, err := strconv.Atoi(string(ch))
		if err == nil {
			outStr = outStr + string(ch)
		}
	}
	return outStr
}
