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
