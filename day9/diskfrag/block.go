package diskfrag

type Block struct {
	id   int
	free bool
}

func (b Block) Eq(other Block) bool {
	if b.free != other.free {
		return false
	}
	if b.free {
		return b.free
	}
	return b.id == other.id
}
