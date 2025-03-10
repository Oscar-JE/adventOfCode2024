package main

import (
	"adventofcode/day14/particle"
	"adventofcode/day14/space"
	vec "adventofcode/geometry/vec2d"
	"fmt"
	"os"
	"strings"
)

func main() {
	shortPart1()
}

func shortPart1() {
	content, fileErr := os.ReadFile("short.txt")
	if fileErr != nil {
		panic("file not found")
	}
	var robots []particle.Particle = parse(string(content))
	playGround := space.InitRepeater(11, 7)
	timeSteps := 100
	fmt.Println(safetyScore(playGround, robots, timeSteps))
}

func safetyScore(sp space.RepeatingSpace, robots []particle.Particle, timeSteps int) int {
	multArray := []int{0, 0, 0, 0}
	for _, rob := range robots {
		endPosition := rob.ProjectedPosition(timeSteps)
		inQuadrant, quad := sp.Quadrant(endPosition)
		if inQuadrant {
			multArray[quad.Enumerate()] += 1
		}
	}
	product := 1
	for _, m := range multArray {
		product *= m
	}
	return product
}

func parse(content string) []particle.Particle {
	particles := []particle.Particle{}
	lines := strings.Split(content, "\r\n")
	for _, line := range lines {
		position, velocity := parseLine(line)
		particles = append(particles, particle.Init(position, velocity))
	}
	return particles
}

func parseLine(line string) (vec.Vec2d, vec.Vec2d) {
	vecReps := strings.Split(line, " ")
	position := vec.Parse(vecReps[0], ",")
	velocity := vec.Parse(vecReps[1], ",")
	return position, velocity
}
