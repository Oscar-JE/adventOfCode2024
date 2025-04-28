package graph

import (
	"adventofcode/integer"
	"adventofcode/set"
	"fmt"
)

type Graph struct {
	nrNodes     int
	connections set.OrderedSet[connection]
}

func Init() Graph {
	nrNodes := 0
	connections := set.InitOrdered([]connection{}, connectionComparer{})
	return Graph{nrNodes: nrNodes, connections: connections}
}

func (g *Graph) AddNode() {
	g.nrNodes++
}

func (g *Graph) AddConnection(con connection) {
	g.nrNodes = integer.Max(integer.Max(con.from, con.to)+1, g.nrNodes)
	g.connections.Insert(con)
}

func (g Graph) String() string {
	list := g.connections.AsSlice()
	nrNodes := fmt.Sprint(g.nrNodes)
	rep := "Number of nodes :" + nrNodes + "\r\n"
	rep += fmt.Sprintf("%v", list)
	return rep
}

func (g Graph) Equal(other Graph) bool {
	if g.nrNodes != other.nrNodes {
		return false
	}
	if !g.connections.Equal(other.connections) {
		return false
	}
	return true
}
