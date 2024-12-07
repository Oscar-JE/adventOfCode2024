package vec

type Vec2d struct {
	x int
	y int
}

func Init(x int, y int) Vec2d {
	return Vec2d{x, y}
}

func (v Vec2d) GetX() int {
	return v.x
}

func (v Vec2d) GetY() int {
	return v.y
}

func Add(v1 Vec2d, v2 Vec2d) Vec2d {
	return Init(v1.x+v2.x, v1.y+v2.y)
}
