package field

import (
	"adventofcode/day16/field/tile.go"
	"adventofcode/geometry/matrix"
	vec "adventofcode/geometry/vec2d"
	"strings"
)

var directions []vec.Vec2d = []vec.Vec2d{vec.Init(-1, 0), vec.Init(0, -1), vec.Init(1, 0), vec.Init(0, 1)}

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

func (f field) isFloor(point vec.Vec2d) bool {
	return f.m.Get(point.GetX(), point.GetY()) == tile.Floor
}

func (f field) isNode(point vec.Vec2d) bool {
	isFloor := f.isFloor(point)
	if !isFloor {
		return false
	}
	return !f.straight(point)
}

func (f field) nrConnections(point vec.Vec2d) int {
	nrConnectedFloor := 0
	for _, v := range directions {
		p := vec.Add(point, v)
		if f.isFloor(p) {
			nrConnectedFloor++
		}
	}
	return nrConnectedFloor
}

func (f field) ConnectionDirs(point vec.Vec2d) []vec.Vec2d {
	dirs := []vec.Vec2d{}
	for _, v := range directions {
		p := vec.Add(point, v)
		if f.isFloor(p) {
			dirs = append(dirs, v)
		}
	}
	return dirs
}

func (f field) straight(point vec.Vec2d) bool {
	nrCon := f.nrConnections(point)
	if nrCon != 2 {
		return false
	}
	var firstDirection vec.Vec2d
	for _, v := range directions {
		p := vec.Add(point, v)
		if f.isFloor(p) {
			firstDirection = v
			break
		}
	}
	flipped := vec.Flip(firstDirection)
	straightAhead := vec.Add(point, flipped)
	return f.isFloor(straightAhead)
}

type Connection struct {
	start vec.Vec2d
	dir   vec.Vec2d
	len   int
}

func (c Connection) GetStart() vec.Vec2d {
	return c.start
}

func (c Connection) GetEnd() vec.Vec2d {
	addVec := vec.ScalarMult(c.len, c.dir)
	return vec.Add(c.start, addVec)
}

func (f field) FindAllConnections() []Connection {
	nodes := f.findAllNodePositions()
	connections := []Connection{}
	for _, n := range nodes {
		dirs := f.ConnectionDirs(n)
		for _, d := range dirs {
			connections = append(connections, f.findConnection(n, d))
		}
	}
	return connections
}

func (f field) findConnection(node vec.Vec2d, direct vec.Vec2d) Connection {
	steps := 0
	currentPosition := node
	for {
		currentPosition = vec.Add(currentPosition, direct)
		if !f.isFloor(currentPosition) {
			break
		}
		steps++
		if !f.straight(currentPosition) {
			break
		}
	}
	return Connection{node, direct, steps}
}
