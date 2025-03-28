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
	currentPosition := vec.Add(i.robotPosition, vec.Vec2d(direction))
	for i.tileAt(currentPosition) != tile.Free() {
		t := i.tileAt(currentPosition)
		if t == tile.Obstructed() {
			return false
		}
		currentPosition = vec.Add(currentPosition, vec.Vec2d(direction))
	}
	return true
}

func (i Inventory) tileAt(position vec.Vec2d) tile.Tile {
	return i.space.Get(position.GetX(), position.GetY())
}

func (i *Inventory) setTileAt(position vec.Vec2d, t tile.Tile) {
	i.space.Set(position.GetX(), position.GetY(), t)
}

func (i *Inventory) forceMove(direction directions.Direction) {
	nextRobotPosition := vec.Add(i.robotPosition, vec.Vec2d(direction))
	i.setTileAt(i.robotPosition, tile.Free())
	if i.tileAt(nextRobotPosition) == tile.Movable() {
		pos := i.findNextFree(nextRobotPosition, direction)
		i.setTileAt(pos, tile.Movable())
	}
	i.setTileAt(nextRobotPosition, tile.Robot())
	i.robotPosition = nextRobotPosition
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
	//todo expandera skiten
	
}
