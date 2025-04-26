package field

import (
	"adventofcode/day16/field/tile.go"
	"adventofcode/geometry/matrix"
	vec "adventofcode/geometry/vec2d"
	"strings"
)

type field struct {
	m matrix.Matrix[tile.Tile]
}

func Parse(rep string) field {
	lines := strings.Split(rep, "\r\n")
	nrRows := len(lines)
	if nrRows == 0 {
		panic("unparsable string")
	}
	nrCols := len([]rune(lines[0]))
	values := []tile.Tile{}
	valueRep := strings.Replace(rep, "\r\n", "", -1)
	for _, r := range valueRep {
		t := tile.ParseTile(r)
		values = append(values, t)
	}
	m := matrix.Init(values, nrRows, nrCols)
	return field{m: m}
}

func (f field) String() string {
	ret := ""
	nrRows := f.m.GetNrRows()
	if nrRows == 0 {
		return ret
	}
	for i := range nrRows - 1 {
		row := f.m.GetRow(i)
		rowRep := convertLineToString(row)
		ret += rowRep + "\r\n"
	}
	lastRow := f.m.GetRow(nrRows - 1)
	ret += convertLineToString(lastRow)
	return ret
}

func convertLineToString(row []tile.Tile) string {
	ret := ""
	for _, el := range row {
		ret += el.String()
	}
	return ret
}

func (f field) findAllNodePositions() []vec.Vec2d {
	// hur gör man detta smidigt?
	// enbart intressant för interna punkter
	// sedan kolla antalet punkter som inte är väggar
	// om inte två så kan det inte vara en koridor
	// sednan måste det vara på en linje genom den aktuella punkten
	// det borde fungera
	points := []vec.Vec2d{}
	for i := 1; i < f.m.GetNrRows(); i++ {
		for j := 1; j < f.m.GetNrCols(); j++ {
			point := vec.Init(i, j)
			if f.isNode(point) {
				points = append(points, point)
			}
		}
	}
	return points
}

func (f field) isNode(point vec.Vec2d) bool {
	isFloor := f.m.Get(point.GetX(), point.GetY()) == tile.Floor
	if !isFloor {
		return false
	}
	upp := vec.Init(-1, 0)
	left := vec.Init(0, -1)
	down := vec.Init(1, 0)
	right := vec.Init(0, 1)
	riktningar := []vec.Vec2d{upp, left, down, right}
	antalÖppna := 0
	for _, v := range riktningar {
		p := vec.Add(point, v)
		if f.m.Get(p.GetX(), p.GetY()) == tile.Floor {
			antalÖppna++
		}
	}
	return antalÖppna != 2
}
