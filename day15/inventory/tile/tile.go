package tile

type Tile int

func Free() Tile {
	return Tile(0)
}

func Movable() Tile {
	return Tile(1)
}

func Obstructed() Tile {
	return Tile(2)
}

func Robot() Tile {
	return Tile(3)
}

func FromRune(r rune) Tile {
	if r == '.' {
		return Free()
	}
	if r == 'O' {
		return Movable()
	}
	if r == '#' {
		return Obstructed()
	}
	if r == '@' {
		return Robot()
	}
	return Free()
}

func (t Tile) String() string {
	if t == Free() {
		return "."
	}
	if t == Movable() {
		return "O"
	}
	if t == Obstructed() {
		return "#"
	}
	if t == Robot() {
		return "@"
	}
	return ""
}
