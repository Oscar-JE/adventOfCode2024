package finder

import (
	"adventofcode/geometry/matrix"
	vec "adventofcode/geometry/vec2d"
	"strings"
)

type finder struct {
	word  []rune
	index int
	sum   int
}

func (f *finder) consume(r rune) {
	if r == f.word[f.index] {
		f.index++
	} else if r == f.word[0] {
		f.index = 1
	} else {
		f.index = 0
	}
	if f.index == len(f.word) {
		f.sum++
		f.index = 0
	}
}

func FindXmases(mat matrix.Matrix[rune]) int {
	sum := 0
	for i := 1; i < mat.GetNrRows()-1; i++ {
		for j := 1; j < mat.GetNrCols()-1; j++ {
			if mat.Get(i, j) != 'A' {
				continue
			}
			word1 := []rune{mat.Get(i-1, j-1), mat.Get(i, j), mat.Get(i+1, j+1)}
			word2 := []rune{mat.Get(i+1, j-1), mat.Get(i, j), mat.Get(i-1, j+1)}
			if isMas(word1) && isMas(word2) {
				sum++
			}
		}
	}
	return sum
}

func isMas(word []rune) bool {
	forward := string(word)
	if forward == "MAS" {
		return true
	}
	backwardsRunes := []rune{}
	for i := len(word) - 1; i >= 0; i-- {
		backwardsRunes = append(backwardsRunes, word[i])
	}
	backwards := string(backwardsRunes)
	return backwards == "MAS"
}

func FindOccurrenceOfWord(mat matrix.Matrix[rune], word string) int {
	lines := readingDirections(mat)
	sum := 0
	for _, line := range lines {
		sum += findOccurrenceInLine(line, word)
	}
	return sum
}

func findOccurrenceInLine(line string, word string) int {
	find := finder{[]rune(word), 0, 0}
	for _, r := range line {
		find.consume(r)
	}
	return find.sum
}

func readingDirections(mat matrix.Matrix[rune]) []string {
	retStrings := []string{}
	retStrings = append(retStrings, leftRight(mat)...)
	retStrings = append(retStrings, botTop(mat)...)
	retStrings = append(retStrings, firstDiagonal(mat)...)
	retStrings = append(retStrings, secondDiagonal(mat)...)
	return retStrings
}

func secondDiagonal(mat matrix.Matrix[rune]) []string {
	retStrings := []string{}
	topHalfForeward := createIteration(vec.Init(0, mat.GetNrCols()-1), vec.Init(0, -1), vec.Init(1, 1), mat)
	retStrings = append(retStrings, topHalfForeward...)
	botHalfForeward := createIteration(vec.Init(1, 0), vec.Init(1, 0), vec.Init(1, 1), mat)
	retStrings = append(retStrings, botHalfForeward...)
	topHalfBackwards := createIteration(vec.Init(0, mat.GetNrCols()-1), vec.Init(1, 0), vec.Init(-1, -1), mat)
	retStrings = append(retStrings, topHalfBackwards...)
	botHalfBackwards := createIteration(vec.Init(mat.GetNrRows()-1, mat.GetNrCols()-2), vec.Init(0, -1), vec.Init(-1, -1), mat)
	retStrings = append(retStrings, botHalfBackwards...)
	return retStrings
}

func firstDiagonal(mat matrix.Matrix[rune]) []string {
	retStrings := []string{}
	topHalfForeward := createIteration(vec.Init(0, 0), vec.Init(1, 0), vec.Init(-1, 1), mat)
	retStrings = append(retStrings, topHalfForeward...)
	botHalfForeward := createIteration(vec.Init(mat.GetNrRows()-1, 1), vec.Init(0, 1), vec.Init(-1, 1), mat)
	retStrings = append(retStrings, botHalfForeward...)
	topHalfBackwards := createIteration(vec.Init(0, 0), vec.Init(0, 1), vec.Init(1, -1), mat)
	retStrings = append(retStrings, topHalfBackwards...)
	botHalfBackwards := createIteration(vec.Init(1, mat.GetNrCols()-1), vec.Init(1, 0), vec.Init(1, -1), mat)
	retStrings = append(retStrings, botHalfBackwards...)
	return retStrings
}

func leftRight(mat matrix.Matrix[rune]) []string {
	retStrings := []string{}
	fromLeftToRight := createIteration(vec.Init(0, 0), vec.Init(1, 0), vec.Init(0, 1), mat)
	retStrings = append(retStrings, fromLeftToRight...)
	fromRightTOLeft := createIteration(vec.Init(0, mat.GetNrCols()-1), vec.Init(1, 0), vec.Init(0, -1), mat)
	retStrings = append(retStrings, fromRightTOLeft...)
	return retStrings
}

func botTop(mat matrix.Matrix[rune]) []string {
	retStrings := []string{}
	fromTopToBot := createIteration(vec.Init(0, 0), vec.Init(0, 1), vec.Init(1, 0), mat)
	retStrings = append(retStrings, fromTopToBot...)
	fromBotToTop := createIteration(vec.Init(mat.GetNrRows()-1, 0), vec.Init(0, 1), vec.Init(-1, 0), mat)
	retStrings = append(retStrings, fromBotToTop...)
	return retStrings
}

func createIteration(startingPoint vec.Vec2d, iterationDirection vec.Vec2d, readingDirection vec.Vec2d, mat matrix.Matrix[rune]) []string {
	lines := []string{}
	currentPoint := startingPoint
	for mat.Inside(currentPoint.GetX(), currentPoint.GetY()) {
		line := readDirection(currentPoint, readingDirection, mat)
		lines = append(lines, line)
		currentPoint = vec.Add(currentPoint, iterationDirection)
	}
	return lines
}

func readDirection(start vec.Vec2d, direction vec.Vec2d, mat matrix.Matrix[rune]) string {
	runes := []rune{}
	currentPoint := start
	for mat.Inside(currentPoint.GetX(), currentPoint.GetY()) {
		rune := mat.Get(currentPoint.GetX(), currentPoint.GetY())
		runes = append(runes, rune)
		currentPoint = vec.Add(currentPoint, direction)
	}
	return string(runes)
}

func ParseMatrix(strRep string) matrix.Matrix[rune] {
	lines := strings.Split(strRep, "\r\n")
	rows := len(lines)
	firstRow := []rune(lines[0])
	cols := len(firstRow)
	strRepWithoutNewLines := strings.Replace(strRep, "\r\n", "", -1)
	values := []rune(strRepWithoutNewLines)
	return matrix.Init(values, rows, cols)
}
