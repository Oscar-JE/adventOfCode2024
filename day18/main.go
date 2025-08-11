package main

import (
	"adventofcode/containers"
	"adventofcode/day18/ram"
	"adventofcode/geometry/matrix"
	vec "adventofcode/geometry/vec2d"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	Part2()
}

func Part1() {
	bytes, fileErr := os.ReadFile("input.txt")
	if fileErr != nil {
		panic("input file was not found")
	}
	coordinatesRep := string(bytes)
	coordinates := ParseCoordList(coordinatesRep)
	shortest := ShortestPathCornerToCorner(70, coordinates[:1024])
	fmt.Println(shortest)
}

func Part2() {
	bytes, fileErr := os.ReadFile("input.txt")
	if fileErr != nil {
		panic("input file was not found")
	}
	coordinatesRep := string(bytes)
	coordinates := ParseCoordList(coordinatesRep)
	var upper int = len(coordinates)
	var lower int = 0
	for lower != upper {
		mid := (upper + lower) / 2
		shortest := ShortestPathCornerToCorner(70, coordinates[:mid])
		if shortest == -1 {
			upper = mid
		} else {
			lower = mid + 1
		}
	}
	var byte vec.Vec2d = coordinates[lower-1]
	fmt.Printf("%d,%d \n", byte.GetY(), byte.GetX())
}

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
		coords = append(coords, vec.Init(second, first))
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
		for _, dir := range directions {
			newDistance := dist + 1
			newPoint := vec.Add(shortest, dir)
			if !field.IsByteOk(newPoint) {
				continue
			}
			if shortestPaths.Get(newPoint.GetX(), newPoint.GetY()) != -1 {
				continue
			}
			priority, seen := pQueue.GetPriority(newPoint)
			if !seen {
				pQueue.Add(newPoint, newDistance)
			} else if newDistance < priority {
				pQueue.UppdatePriority(newPoint, newDistance)
			}
		}
	}
	return shortestPaths.Get(maxCoord, maxCoord)

}
