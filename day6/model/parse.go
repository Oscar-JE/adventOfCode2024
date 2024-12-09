package model

import (
	"adventofcode/geometry/matrix"
	"os"
	"strings"
)

func Parse(filename string) matrix.Matrix[rune] {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		panic("file was not found")
	}
	content := string(bytes)
	lines := strings.Split(content, "\r\n")
	nrOfRows := len(lines)
	row := []rune(lines[0])
	nrOfCols := len(row)
	contentWithoutNewLines := strings.Replace(content, "\r\n", "", -1)
	mat := matrix.Init([]rune(contentWithoutNewLines), nrOfRows, nrOfCols)
	return mat
}
