package main

import (
	"adventofcode/containers"
	"adventofcode/day18/ram"
	"adventofcode/geometry/matrix"
	vec "adventofcode/geometry/vec2d"
	"strconv"
	"strings"
)

func ParseCoordList(rep string) []vec.Vec2d {
	rows := strings.Split(rep, "\r\n")
	coords := []vec.Vec2d{}
	for _, row := range rows {
		firstAndSecond := strings.Split(row, ",")
		first, err1 := strconv.Atoi(firstAndSecond[0])
		second, err2 := strconv.Atoi(firstAndSecond[1])
		if err1 != nil || err2 != nil {
			panic("error parsing coordinate rows")
		}
		coords = append(coords, vec.Init(first, second))
	}
	return coords
}

var directions []vec.Vec2d = []vec.Vec2d{vec.Init(-1, 0), vec.Init(0, 1), vec.Init(1, 0), vec.Init(0, -1)}

func ShortestPathCornerToCorner(maxCoord int, deadSpots []vec.Vec2d) int {
	field := ram.InitStationary(maxCoord, deadSpots)
	shortestPaths := matrix.InitSame(maxCoord+1, maxCoord+1, -1)
	pQueue := containers.InitPriorityQueue[vec.Vec2d]()
	pQueue.Add(vec.Init(0, 0), 0)
	for !pQueue.Empty() {
		shortest, dist := pQueue.PopSmallestPriority()
		shortestPaths.Set(shortest.GetX(), shortest.GetY(), dist)
		// här ska vi generera nya platser att gå till
	}

}
