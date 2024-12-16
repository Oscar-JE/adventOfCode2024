package field

import "adventofcode/geometry/matrix"

func ParseFromLines(lines []string) Field {
	nrRows := len(lines)
	if nrRows == 0 {
		panic("empty input to ParseLines is not allowed")
	}
	nrCols := len([]rune(lines[0]))
	if nrCols == 0 {
		panic("empty input to ParseLines is not allowed")
	}
	plantLines := [][]plantingPlot{}
	firstLine := []plantingPlot{}
	firstLineRunes := []rune(lines[0])
	firstLine = append(firstLine, Init(string(firstLineRunes[0])))
	for j := 1; j < len(firstLineRunes); j++ {
		previous := firstLine[j-1]
		current := Init(string(firstLineRunes[j]))
		current.boarder.west = previous.boarder.east
		firstLine = append(firstLine, current)
	}
	plantLines = append(plantLines, firstLine)
	for i := 1; i < nrRows; i++ {
		previousLine := plantLines[i-1]
		newLineRunes := []rune(lines[i])
		newLine := createLine(previousLine, newLineRunes)
		plantLines = append(plantLines, newLine)
	}
	flattened := []plantingPlot{}
	for _, row := range plantLines {
		flattened = append(flattened, row...)
	}
	mat := matrix.Init(flattened, nrRows, nrCols)
	return Field{plots: mat}
}

func createLine(prevLine []plantingPlot, newLine []rune) []plantingPlot {
	line := []plantingPlot{}
	firsElement := Init(string(newLine[0]))
	firsElement.boarder.north = prevLine[0].boarder.south
	line = append(line, firsElement)
	for j := 1; j < len(newLine); j++ {
		elem := Init(string(newLine[j]))
		elem.boarder.north = prevLine[j].boarder.south
		elem.boarder.west = line[j-1].boarder.east
		line = append(line, elem)
	}
	return line
}
