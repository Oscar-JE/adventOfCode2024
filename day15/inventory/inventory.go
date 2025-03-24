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
	if i.blockedRobot(direction) {
		return
	} else {
		i.forceMove(direction)
	}
}

func (i Inventory) blockedRobot(direction directions.Direction) bool {
	return false
}

func (i *Inventory) forceMove(direction directions.Direction) {
	//lite signaturer nu kan vi skriva test
}
