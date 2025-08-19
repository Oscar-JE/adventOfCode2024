package field

import (
	"adventofcode/geometry/matrix"
	vec "adventofcode/geometry/vec2d"
	"strings"
)

type Tile string

const floor Tile = "."
const blockage Tile = "#"
const start Tile = "S"
const end Tile = "E"

type Field struct {
	m matrix.Matrix[Tile]
}

func Parse(rep string) Field {
	rows := strings.Split(rep, "\r\n")
	nrRows := len(rows)
	if nrRows < 1 {
		panic("no rows was found in string representation of field")
	}
	firstRow := rows[0]
	nrCols := len([]rune(firstRow))
	withNoNewLines := strings.Replace(rep, "\r\n", "", -1)
	possibleTiles := []Tile{floor, blockage, start, end}
	tiles := []Tile{}
	for _, r := range withNoNewLines {
		for _, t := range possibleTiles {
			if t == Tile(string(r)) {
				tiles = append(tiles, t)
				continue
			}
		}
	}
	m := matrix.Init(tiles, nrRows, nrCols)
	return Field{m: m}
}

func (f Field) String() string {
	return f.m.String()
}

func (f Field) GetTile(pos vec.Vec2d) Tile {
	return f.m.Get(pos.GetX(), pos.GetY())
}

func (f Field) GetNrCols() int {
	return f.m.GetNrCols()
}

func (f Field) GetNrRows() int {
	return f.m.GetNrRows()
}

func (f Field) FindEnd() vec.Vec2d {
	x, y := f.m.FirstRowAndColOf(end)
	return vec.Init(x, y)
}

func (f Field) FindStart() vec.Vec2d {
	x, y := f.m.FirstRowAndColOf(start)
	return vec.Init(x, y)
}

func (f Field) Inside(v vec.Vec2d) bool {
	return f.m.Inside(v.GetX(), v.GetY())
}

func (f Field) IsFloor(v vec.Vec2d) bool {
	return f.m.Get(v.GetX(), v.GetY()) == floor || f.m.Get(v.GetX(), v.GetY()) == start
}
