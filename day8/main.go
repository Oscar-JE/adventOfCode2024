package main

import (
	"adventofcode/day8/antennas"
	"adventofcode/day8/logger"
	"adventofcode/geometry/matrix"
	vec "adventofcode/geometry/vec2d"
	"fmt"
	"os"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("big.txt")
	if err != nil {
		panic("input file was not found")
	}
	content := string(bytes)
	rows, cols, antennasGroups := parse(content)
	fmt.Println(part2(rows, cols, antennasGroups))
}

func part1(rows int, cols int, antennasGroups []antennas.AntennaGroup) int {

	logger := logger.InitLogger(rows, cols)
	for _, group := range antennasGroups {
		poaistions := group.AntiNodes()
		for _, p := range poaistions {
			logger.Set(p)
		}
	}
	return len(logger.Positions())
}

func part2(rows int, cols int, antennasGroups []antennas.AntennaGroup) int {
	logger := logger.InitLogger(rows, cols)
	for _, group := range antennasGroups {
		poaistions := group.AntiNodesHarmonics(vec.Init(rows, cols))
		for _, p := range poaistions {
			logger.Set(p)
		}
	}
	positions := logger.Positions()
	return len(positions)
}

func parse(content string) (int, int, []antennas.AntennaGroup) {
	lines := strings.Split(content, "\r\n")
	rows := len(lines)
	aline := []rune(lines[0])
	cols := len(aline)
	contentWithooutNewLines := strings.Replace(content, "\r\n", "", -1)
	mat := matrix.Init([]rune(contentWithooutNewLines), rows, cols)
	antennaGroups := []antennas.AntennaGroup{}
	for i := range mat.GetNrRows() {
		for j := range mat.GetNrCols() {
			sign := string(mat.Get(i, j))
			if sign == "." {
				continue
			}
			groupExists, groupIndex := findGroup(sign, antennaGroups)
			if groupExists {
				antennaGroups[groupIndex].AddAntenna(vec.Init(i, j))
			} else {
				newGroup := antennas.InitGroup(string(sign))
				newGroup.AddAntenna(vec.Init(i, j))
				antennaGroups = append(antennaGroups, newGroup)
			}
		}
	}
	return rows, cols, antennaGroups
}

func findGroup(sign string, groups []antennas.AntennaGroup) (bool, int) {
	for i, g := range groups {
		if g.GetId() == sign {
			return true, i
		}
	}
	return false, -1
}
