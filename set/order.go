package set

type LessThan[E any] interface {
	LessOrEqaul(E, E) bool
}

type Comparator[E any] struct {
	l LessThan[E]
}

func (c Comparator[E]) Less(a E, b E) bool {
	return c.l.LessOrEqaul(a, b) && !c.l.LessOrEqaul(b, a)
}

func (c Comparator[E]) Greater(a E, b E) bool {
	return c.l.LessOrEqaul(b, a) && !c.l.LessOrEqaul(a, b)
}

func (c Comparator[E]) Equal(a E, b E) bool {
	return c.l.LessOrEqaul(a, b) && c.l.LessOrEqaul(b, a)
}
