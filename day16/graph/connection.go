package graph

type connection struct {
	from   int
	to     int
	weight int
}

type connectionComparer struct {
}

func InitConnection(from int, to int, weight int) connection {
	return connection{from: from, to: to, weight: weight}
}

func (c connectionComparer) LessOrEqaul(c1 connection, c2 connection) bool {
	if c1.from > c2.from {
		return false
	}
	if c1.from < c2.from {
		return true
	}
	return c1.to <= c2.to
}
