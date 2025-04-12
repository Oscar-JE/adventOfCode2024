package inventory

import (
	"adventofcode/day15/directions"
	"adventofcode/day15/inventory/tile"
	"adventofcode/geometry/matrix"
	vec "adventofcode/geometry/vec2d"
	"strings"
)

type Inventory struct {
	space         matrix.PrintableMatrix[tile.Tile]
	robotPosition vec.Vec2d
	leftSides     []vec.Vec2d
}

func (inv Inventory) String() string {
	rep := ""
	for i := range inv.space.GetNrRows() - 1 {
		lineRep := ""
		for j := range inv.space.GetNrCols() {
			position := vec.Init(i, j)
			if inv.amongLefts(position) {
				lineRep += tile.BoxLeft().String()
			} else if position == inv.robotPosition {
				lineRep += tile.Robot().String()
			} else {
				lineRep += inv.space.Get(i, j).String()
			}
		}
		rep += lineRep + "\r\n"
	}
	lineRep := ""
	for j := range inv.space.GetNrCols() {
		position := vec.Init(inv.space.GetNrRows()-1, j)
		if inv.amongLefts(position) {
			lineRep += tile.BoxLeft().String()
		} else {
			lineRep += inv.space.Get(inv.space.GetNrRows()-1, j).String()
		}
	}
	rep += lineRep
	return rep
}

func (inv Inventory) amongLefts(pos vec.Vec2d) bool {
	for _, leftSide := range inv.leftSides {
		if leftSide == pos {
			return true
		}
	}
	return false
}

func (inv *Inventory) convertToPosititionsFomate() {
	rows := inv.space.GetNrRows()
	cols := inv.space.GetNrCols()
	for i := range rows {
		for j := range cols {
			t := inv.space.Get(i, j)
			if t == tile.BoxLeft() || t == tile.Movable() {
				inv.space.Set(i, j, tile.Free())
				inv.leftSides = append(inv.leftSides, vec.Init(i, j))
			} else if t == tile.BoxRight() {
				inv.space.Set(i, j, tile.Free())
			}
		}
	}
}

func FromString(rep string) Inventory {
	lines := strings.Split(rep, "\r\n")
	vals := strings.Replace(rep, "\r\n", "", -1)
	nrRows := len(lines)
	nrColumns := len([]rune(lines[0]))
	tiles := []tile.Tile{}
	for _, r := range vals {
		tiles = append(tiles, tile.FromRune(r))
	}
	tiling := matrix.InitPrintable(tiles, nrRows, nrColumns)
	rowRobo, colRobo := tiling.FirstRowAndColOf(tile.Robot())
	robotPos := vec.Init(rowRobo, colRobo)
	tiling.Set(rowRobo, colRobo, tile.Free())
	inv := Inventory{space: tiling, robotPosition: robotPos}
	inv.convertToPosititionsFomate()
	return inv
}

func (i *Inventory) MoveRobot(direction directions.Direction) {
	relevantPositions := i.findRelevantPositions(direction)
	for _, mem := range relevantPositions { // nu testar vi och gråter lite
		*mem = vec.Add(vec.Vec2d(direction), *mem)
	}
}

func (i *Inventory) findRelevantPositions(dir directions.Direction) []*vec.Vec2d {
	queue := []*vec.Vec2d{&i.robotPosition}
	found := []*vec.Vec2d{}
	for len(queue) > 0 {
		el := queue[0]
		queue = queue[1:]
		obstrction, childNodes := i.findDirektCollisions(el, dir)
		if obstrction {
			return []*vec.Vec2d{}
		}
		found = append(found, el)
		queue = append(queue, childNodes...)
	}
	return found
}

func (inv *Inventory) findDirektCollisions(pos *vec.Vec2d, dir directions.Direction) (bool, []*vec.Vec2d) {
	nextPos := vec.Add(*pos, vec.Vec2d(dir))
	nextPositions := []vec.Vec2d{nextPos}
	if *pos != inv.robotPosition {
		candidate := vec.Add(nextPos, vec.Vec2d(directions.East()))
		if candidate != *pos {
			nextPositions = append(nextPositions, candidate)
		}
	}
	retVec := []*vec.Vec2d{}
	for _, nextPos := range nextPositions {
		t := inv.space.Get(nextPos.GetX(), nextPos.GetY())
		if t == tile.Obstructed() {
			return true, []*vec.Vec2d{}
		}
		isLeftOrRight, leftRef := inv.findConnectedLeftPair(nextPos)
		if isLeftOrRight && *leftRef != *pos {
			if len(retVec) > 0 && retVec[0] == leftRef {
				continue
			}
			retVec = append(retVec, leftRef)
		}
	}
	return false, retVec
}

func (inv *Inventory) findConnectedLeftPair(pos vec.Vec2d) (bool, *vec.Vec2d) {
	for index, left := range inv.leftSides {
		if left == pos || left == vec.Add(vec.Vec2d(directions.West()), pos) {
			return true, &inv.leftSides[index]
		}
	}
	return false, nil
}

func (i Inventory) PositionsOfMovabel() []vec.Vec2d {
	return i.space.PositionsOf(tile.Movable())
}

func (i *Inventory) Expand() {
	nrCols := 2 * i.space.GetNrCols()
	nrRows := i.space.GetNrRows()
	expanded := []tile.Tile{}
	for row := range nrRows {
		rowEntrys := i.space.GetRow(row)
		for _, el := range rowEntrys {
			expanded = append(expanded, el.Expand()...)
		}
	}
	i.space = matrix.InitPrintable(expanded, nrRows, nrCols)
	i.robotPosition = expandPosition(i.robotPosition)
	i.expandLeftSides()
}

func (inv *Inventory) expandLeftSides() {
	for i := range inv.leftSides {
		inv.leftSides[i] = expandPosition(inv.leftSides[i])
	}
}

func expandPosition(pos vec.Vec2d) vec.Vec2d {
	x := pos.GetX()
	y := pos.GetY() * 2
	return vec.Init(x, y)
}

func (i Inventory) Score() int {
	sum := 0
	for _, pos := range i.leftSides {
		sum += pos.GetX()*100 + pos.GetY()
	}
	return sum
}
