package towels

type Pattern struct {
	colors []rune
}

func ParsePattern(rep string) Pattern {
	runes := []rune(rep)
	return Pattern{colors: runes}
}

func (p Pattern) matchBeginning(other Pattern) (int, bool) {
	if len(p.colors) < len(other.colors) {
		return -1, false
	}
	var match bool = true
	for i, r := range other.colors {
		if r != p.colors[i] {
			match = false
		}
	}
	return len(other.colors), match
}

func (p Pattern) subPattern(index int) Pattern {
	newColors := []rune{}
	newColors = append(newColors, p.colors[index:]...)
	return Pattern{colors: newColors}
}

func (p Pattern) Empty() bool {
	return len(p.colors) == 0
}

func (p Pattern) String() string {
	return string(p.colors)
}
