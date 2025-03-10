package quadrant

type Quadrant struct {
	which int
}

func (q Quadrant) Enumerate() int {
	return q.which
}

func First() Quadrant {
	return Quadrant{0}
}

func Second() Quadrant {
	return Quadrant{1}
}

func Third() Quadrant {
	return Quadrant{2}
}

func Fourth() Quadrant {
	return Quadrant{3}
}
