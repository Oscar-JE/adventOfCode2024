package field

import vec "adventofcode/geometry/vec2d"

type swirl struct {
	north *vec.Vec2d
	west  *vec.Vec2d
	south *vec.Vec2d
	east  *vec.Vec2d
}

func InitSwirl() swirl {
	north := vec.Init(0, 0)
	west := vec.Init(0, 0)
	south := vec.Init(0, 0)
	east := vec.Init(0, 0)
	return swirl{north: &north, west: &west, south: &south, east: &east}
}

func (s swirl) rotate() {
	*s.north = vec.Add(*s.north, vec.Init(0, -1))
	*s.west = vec.Add(*s.west, vec.Init(1, 0))
	*s.south = vec.Add(*s.south, vec.Init(0, 1))
	*s.east = vec.Add(*s.east, vec.Init(-1, 0))
}

func (s swirl) reset() {
	*s.north = vec.Init(0, 0)
	*s.west = vec.Init(0, 0)
	*s.south = vec.Init(0, 0)
	*s.east = vec.Init(0, 0)
}

func (s swirl) sumOfSquaredRotations() int {
	return vec.AbsSquared(*s.north) + vec.AbsSquared(*s.west) + vec.AbsSquared(*s.south) + vec.AbsSquared(*s.east)
}
