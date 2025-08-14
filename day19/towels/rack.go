package towels

import "strings"

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
