package main

import (
	"adventofcode/day14/particle"
	"adventofcode/day14/roborally"
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
	timeSteps := 100
	battle := roborally.Init(11, 7, robots)
	fmt.Println(battle)
	fmt.Println(battle.SafetyScore(timeSteps))
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
