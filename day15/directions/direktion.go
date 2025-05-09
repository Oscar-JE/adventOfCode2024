package directions

import vec "adventofcode/geometry/vec2d"

type Direction vec.Vec2d

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
