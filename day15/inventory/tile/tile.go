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

func BoxLeft() Tile {
	return Tile(4)
}

func BoxRight() Tile {
	return Tile(5)
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
	if r == '[' {
		return BoxLeft()
	}
	if r == ']' {
		return BoxRight()
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
	if t == BoxLeft() {
		return "["
	}
	if t == BoxRight() {
		return "]"
	}
	return ""
}

func (t Tile) Expand() []Tile {
	if t == Free() {
		return []Tile{Free(), Free()}
	}
	if t == Movable() {
		return []Tile{BoxLeft(), BoxRight()}
	}
	if t == Obstructed() {
		return []Tile{Obstructed(), Obstructed()}
	}
	if t == Robot() {
		return []Tile{Robot(), Free()}
	}
	panic("unexpandable tile")
}
