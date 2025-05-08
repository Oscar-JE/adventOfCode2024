package main

import (
	"adventofcode/day16/field"
	"adventofcode/day16/graph"
	vec "adventofcode/geometry/vec2d"
)

func createGraph(fieldRep string) graph.Graph { // nej det här blir bara jobbigt
	field := field.Parse(fieldRep)
	cons := field.FindAllConnections()
	return graphFromConnections(cons)
}

func graphFromConnections(cons []field.Connection) graph.Graph {
	nodePoints := []vec.Vec2d{}
	for _, c := range cons {
		nodePoints = append(nodePoints, c.GetStart())
	}
	g := graph.Init()
	for i, _ := range nodePoints {
		startNumber := 8 * i
		g.AddConnection(graph.InitConnection(startNumber, startNumber+1, 1000))
		g.AddConnection(graph.InitConnection(startNumber, startNumber+3, 1000))
		g.AddConnection(graph.InitConnection(startNumber+1, startNumber, 1000))
		g.AddConnection(graph.InitConnection(startNumber+1, startNumber+2, 1000))
		g.AddConnection(graph.InitConnection(startNumber+2, startNumber+1, 1000))
		g.AddConnection(graph.InitConnection(startNumber+2, startNumber+3, 1000))
		g.AddConnection(graph.InitConnection(startNumber+3, startNumber+2, 1000))
		g.AddConnection(graph.InitConnection(startNumber+3, startNumber, 1000))

	}

	return g
}
