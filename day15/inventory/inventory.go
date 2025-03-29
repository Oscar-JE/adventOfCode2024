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
}

func (i Inventory) String() string {
	return i.space.String()
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
	return Inventory{space: tiling, robotPosition: robotPos}
}

func (i *Inventory) MoveRobot(direction directions.Direction) {
	if i.canRobotMoveToThe(direction) {
		i.forceMove(direction)
	}
}

func (i Inventory) canRobotMoveToThe(direction directions.Direction) bool {
	currentRow := []vec.Vec2d{i.robotPosition}
	for len(currentRow) > 0 {
		for _, pos := range currentRow {
			if i.tileAt(pos) == tile.Obstructed() {
				return false
			}
		}
		currentRow = i.advanceRow(currentRow, direction)
	}
	return true
}

func (i Inventory) advanceRow(currentRow []vec.Vec2d, dir directions.Direction) []vec.Vec2d {
	nextRow := []vec.Vec2d{}
	for _, pos := range currentRow {
		nextPositions := i.advancePosition(pos, dir)
		for _, p := range nextPositions {
			if i.tileAt(p) != tile.Free() && !contains(nextRow, p) {
				nextRow = append(nextRow, p)
			}
		}
	}
	return nextRow
}

func (i Inventory) advancePosition(position vec.Vec2d, dir directions.Direction) []vec.Vec2d {
	horizontalMovement := vec.DotProduct(vec.Vec2d(dir), vec.Init(0, 1)) == 1 || vec.DotProduct(vec.Vec2d(dir), vec.Init(0, 1)) == -1
	nextBase := vec.Add(position, vec.Vec2d(dir))
	if horizontalMovement {
		return []vec.Vec2d{nextBase}
	}
	t := i.tileAt(nextBase)
	if t == tile.BoxLeft() {
		return []vec.Vec2d{nextBase, vec.Add(nextBase, vec.Init(0, 1))}
	}
	if t == tile.BoxRight() {
		return []vec.Vec2d{nextBase, vec.Add(nextBase, vec.Init(0, -1))}
	}
	return []vec.Vec2d{nextBase}
}

func contains(list []vec.Vec2d, el vec.Vec2d) bool {
	for _, v := range list {
		if v == el {
			return true
		}
	}
	return false
}

func (i Inventory) tileAt(position vec.Vec2d) tile.Tile {
	return i.space.Get(position.GetX(), position.GetY())
}

func (i *Inventory) setTileAt(position vec.Vec2d, t tile.Tile) {
	i.space.Set(position.GetX(), position.GetY(), t)
}

type tileAndPos struct {
	t   tile.Tile
	pos vec.Vec2d
}

func (i *Inventory) forceMove(direction directions.Direction) {
	firstLine := []tileAndPos{{t: i.tileAt(i.robotPosition), pos: i.robotPosition}}
	secondLine := i.advanceTileAndPos(firstLine, direction) 
	for len(firstLine) > 0 { // blir fel
		for _, el := range firstLine {
			i.setTileAt(el.pos, tile.Free())
		}
		for _, el := range firstLine {
			pos := vec.Add(el.pos, vec.Vec2d(direction))
			i.setTileAt(pos, el.t)
		}
		firstLine = secondLine
		secondLine = i.advanceTileAndPos(firstLine, direction)
	}
}

func (i *Inventory) advanceTileAndPos(tpos []tileAndPos, dir directions.Direction) []tileAndPos {
	currentRow := []vec.Vec2d{}
	for _, el := range tpos {
		currentRow = append(currentRow, el.pos)
	}
	nextPos := i.advanceRow(currentRow, dir)
	retV := []tileAndPos{}
	for _, pos := range nextPos {
		retV = append(retV, tileAndPos{pos: pos, t: i.tileAt(pos)})
	}
	return retV
}

func (i Inventory) findNextFree(pos vec.Vec2d, dir directions.Direction) vec.Vec2d {
	currentPosition := vec.Add(i.robotPosition, vec.Vec2d(dir))
	for i.tileAt(currentPosition) != tile.Free() {
		currentPosition = vec.Add(currentPosition, vec.Vec2d(dir))
	}
	return currentPosition
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
	rowRobot, colRobot := i.space.FirstRowAndColOf(tile.Robot())
	i.robotPosition = vec.Init(rowRobot, colRobot)
}
