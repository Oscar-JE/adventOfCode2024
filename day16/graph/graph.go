package graph

import vec "adventofcode/geometry/vec2d"

type GeoGraph struct {
	nodes []GeoNode
}

type GeoNode struct {
	position vec.Vec2d
	north    connection
	west     connection
	south    connection
	east     connection
}

type connection struct {
	to   *GeoNode
	cost int
}

func (g GeoGraph) String() string {
	nrOfNodes := len(g.nodes)
	offsetForleftColumn := " " + " "
	charWidth := " "
	topBar := offsetForleftColumn
	for i := range(nrOfNodes) {
		for
	}
}
