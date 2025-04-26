package tile

var Wall Tile = Tile(0)
var Floor Tile = Tile(1)

type Tile int

func (t Tile) String() string {
	if t == 0 {
		return "#"
	}
	return "."
}

func ParseTile(r rune) Tile {
	if r == '#' {
		return Wall
	}
	return Floor
}
