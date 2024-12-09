package model

import (
	"adventofcode/geometry/matrix"
	vec "adventofcode/geometry/vec2d"
)

func Init(input matrix.Matrix[rune]) Model {
	gard := findGard(input)
	field := convertToField(input)
	stateLogger := initStateLogger(input.GetNrRows(), input.GetNrCols())
	return Model{gard: gard, field: field, prevStates: stateLogger}
}

type Model struct {
	gard       Gard
	field      matrix.Matrix[tile]
	prevStates stateLogger
}

type tile string

const (
	empty       tile = "."
	obstruction tile = "#"
)

func (m Model) GetGardPosition() vec.Vec2d {
	return m.gard.position
}

func (m Model) IsLoop(x int, y int) bool {
	m.field.Set(x, y, obstruction)
	m.prevStates = initStateLogger(m.field.GetNrRows(), m.field.GetNrCols())
	isLoop := false
	for !isLoop && m.field.Inside(m.gard.position.GetX(), m.gard.position.GetY()) {
		isLoop = m.prevStates.set(m.gard)
		m.TimeStep()
	}
	m.field.Set(x, y, empty)
	return m.field.Inside(m.gard.position.GetX(), m.gard.position.GetY())
}

func (m *Model) TimeStep() {
	positionCandidate := m.gard.nextPosition()
	for !m.validNextGardPosition(positionCandidate) {
		m.gard.turn()
		positionCandidate = m.gard.nextPosition()
	}
	m.gard.position = positionCandidate
}

func (m Model) IsItOver() bool {
	return !m.field.Inside(m.gard.position.GetX(), m.GetGardPosition().GetY())
}

func (m Model) validNextGardPosition(positionCandidate vec.Vec2d) bool {
	onField := m.field.Inside(positionCandidate.GetX(), positionCandidate.GetY())
	if !onField {
		return true
	}
	obstructed := m.field.Get(positionCandidate.GetX(), positionCandidate.GetY()) == obstruction
	return !obstructed
}

func findGard(input matrix.Matrix[rune]) Gard {
	position := findGardPosition(input)
	direction := vec.Init(-1, 0)
	return Gard{position: position, direction: direction}
}

func findGardPosition(input matrix.Matrix[rune]) vec.Vec2d {
	for i := range input.GetNrRows() {
		for j := range input.GetNrCols() {
			sign := input.Get(i, j)
			if isGard(sign) {
				return vec.Init(i, j)
			}
		}
	}
	panic(" No guardsman was found !")
}

func isGard(sign rune) bool {
	return sign == '^'
}

func convertToField(input matrix.Matrix[rune]) matrix.Matrix[tile] {
	tiles := []tile{}
	for i := range input.GetNrRows() {
		for j := range input.GetNrCols() {
			sign := input.Get(i, j)
			if sign == '#' {
				tiles = append(tiles, obstruction)
			} else {
				tiles = append(tiles, empty)
			}
		}
	}
	return matrix.Init[tile](tiles, input.GetNrRows(), input.GetNrCols())
}

func (m Model) String() string {
	retString := ""
	for i := range m.gard.position.GetX() {
		line := ""
		for j := range m.field.GetNrCols() {
			line += string(m.field.Get(i, j))
		}
		retString += line + "\r\n"
	}
	gardLine := m.gardLine()
	retString += gardLine
	for i := m.gard.position.GetX() + 1; i < m.field.GetNrRows(); i++ {
		line := ""
		for j := range m.field.GetNrCols() {
			line += string(m.field.Get(i, j))
		}
		retString += line + "\r\n"
	}
	return retString
}

func (m Model) gardLine() string {
	if !m.field.Inside(m.gard.position.GetX(), m.gard.position.GetX()) {
		return ""
	}
	xIndex := m.gard.position.GetX()
	yIndex := m.gard.position.GetY()
	retLine := ""
	for j := range yIndex {
		retLine += string(m.field.Get(xIndex, j))
	}
	retLine += m.gard.getDirektionSymbol()
	for j := yIndex + 1; j < m.field.GetNrCols(); j++ {
		retLine += string(m.field.Get(xIndex, yIndex))
	}
	return retLine + "\r\n"
}

type Gard struct {
	position  vec.Vec2d
	direction vec.Vec2d
}

func (g Gard) nextPosition() vec.Vec2d {
	return vec.Add(g.position, g.direction)
}

func (g *Gard) turn() {
	g.direction = vec.Turn90Down(g.direction)
}

var (
	up    = vec.Init(-1, 0)
	down  = vec.Init(1, 0)
	left  = vec.Init(0, -1)
	right = vec.Init(0, 1)
)

func (g Gard) getDirektionSymbol() string {
	if g.direction == up {
		return "^"
	}
	if g.direction == down {
		return "v"
	}
	if g.direction == left {
		return "<"
	}
	if g.direction == right {
		return ">"
	}
	return "No gard direction found "
}
