package main

import (
	"adventofcode/day14/particle"
	"adventofcode/day14/roborally"
	vec "adventofcode/geometry/vec2d"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	part2()
}

func part2() {
	content, fileErr := os.ReadFile("long.txt")
	if fileErr != nil {
		panic("file not found")
	}
	var robots []particle.Particle = parse(string(content))
	arena := roborally.Init(101, 103, robots)
	count := 0
	for time := range 1000000 {
		if arena.MaybyTree(time) {
			fmt.Println(time)
			arena.Rep(time)
			input := bufio.NewScanner(os.Stdin)
			input.Scan()
			count++
		}
	}
	fmt.Printf("number of options: %d \r\n", count)
}

func part1() {
	content, fileErr := os.ReadFile("long.txt")
	if fileErr != nil {
		panic("file not found")
	}
	var robots []particle.Particle = parse(string(content))
	timeSteps := 100
	battle := roborally.Init(101, 103, robots)
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
