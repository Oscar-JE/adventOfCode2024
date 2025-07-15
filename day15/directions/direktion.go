package directions

import vec "adventofcode/geometry/vec2d"

type Direction vec.Vec2d

const NrDirections = 4

func North() Direction {
	return Direction(vec.Init(-1, 0))
}

func East() Direction {
	return Direction(vec.Init(0, 1))
}

func South() Direction {
	return Direction(vec.Init(1, 0))
}

func West() Direction {
	return Direction(vec.Init(0, -1))
}

func Init(r rune) Direction {

	if r == '^' {
		return North()
	} else if r == '>' {
		return East()
	} else if r == 'v' {
		return South()
	} else if r == '<' {
		return West()
	}
	panic("unalowed representation")
}

func (d Direction) String() string {
	if d == North() {
		return "^"
	} else if d == East() {
		return ">"
	} else if d == South() {
		return "v"
	} else if d == West() {
		return "<"
	}
	return ""
}

func (d Direction) Index() int {
	if d == North() {
		return 0
	} else if d == East() {
		return 1
	} else if d == South() {
		return 2
	} else if d == West() {
		return 3
	}
	panic("this should not be possible to happen")
}

func DirectionFromIndex(index int) Direction {
	if index == 0 {
		return North()
	} else if index == 1 {
		return East()
	} else if index == 2 {
		return South()
	} else if index == 3 {
		return West()
	}
	panic("this should not be possible to happen")
}

func (d Direction) TurnDown() Direction {
	if d == North() {
		return East()
	}
	if d == East() {
		return South()
	}
	if d == South() {
		return West()
	}
	if d == West() {
		return North()
	}
	panic("the list above needs to be exhaustive")
}

func (d Direction) TurnUp() Direction {
	if d == North() {
		return West()
	}
	if d == West() {
		return South()
	}
	if d == South() {
		return East()
	}
	if d == East() {
		return North()
	}
	panic("the list above needs to be exhaustive")
}

func Enumeration(dir Direction) int {
	if dir == North() {
		return 0
	} else if dir == East() {
		return 1
	} else if dir == South() {
		return 2
	} else if dir == West() {
		return 3
	}
	panic("should never reachthis state")
}
