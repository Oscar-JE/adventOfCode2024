package graph

import (
	"adventofcode/integer"
	"adventofcode/set"
	"fmt"
)

type Graph struct {
	nrNodes    int
	conections set.OrderedSet[connection]
}

func Init() Graph {
	nrNodes := 0
	connections := set.InitOrdered([]connection{}, connectionComparer{})
	return Graph{nrNodes: nrNodes, conections: connections}
}

func (g *Graph) AddNode() {
	g.nrNodes++
}

func (g *Graph) AddConnection(con connection) {
	g.nrNodes = integer.Max(integer.Max(con.from, con.to)+1, g.nrNodes)
	g.conections.Insert(con)
}

func (g Graph) String() string {
	list := g.conections.AsSlice()
	nrNodes := fmt.Sprint(g.nrNodes)
	rep := "Number of nodes :" + nrNodes + "\r\n"
	rep += fmt.Sprintf("%v", list)
	return rep
}
