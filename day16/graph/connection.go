package graph

type connection struct {
	from   int
	to     int
	length int
}

func (c connection) lessOrEqual(other connection) bool {
	if c.from <= other.from {
		return true
	}

}

type conenctionSet struct {
	connections []connection
}
