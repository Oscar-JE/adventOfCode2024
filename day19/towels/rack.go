package towels

import (
	"fmt"
	"strings"
)

type TowelRack struct {
	patterns []Pattern
}

func ParseTowelRack(rep string) TowelRack {
	splitted := strings.Split(rep, ", ")
	patterns := []Pattern{}
	for _, v := range splitted {
		patterns = append(patterns, ParsePattern(v))
	}
	return TowelRack{patterns: patterns}
}

func (t TowelRack) SolvablePattern(p Pattern) bool {
	if p.Empty() {
		return true
	}
	solvable := false
	for _, pattern := range t.patterns {
		index, match := p.matchBeginning(pattern)
		if match {
			solvable = solvable || t.SolvablePattern(p.subPattern(index))
		}
		if solvable {
			break
		}
	}
	return solvable
}

func (t TowelRack) NumberOfSolutions(p Pattern) int {
	m := make(map[string]int)
	return t.numberOfSolutions(p, &m)
}

func (t TowelRack) numberOfSolutions(p Pattern, m *map[string]int) int {
	key := fmt.Sprint(p)
	nr, has := (*m)[key]
	if has {
		return nr
	}
	if p.Empty() {
		return 1
	}
	sum := 0
	for _, pattern := range t.patterns {
		index, match := p.matchBeginning(pattern)
		if match {
			sum += t.numberOfSolutions(p.subPattern(index), m)
		}
	}
	(*m)[key] = sum
	return sum
}
