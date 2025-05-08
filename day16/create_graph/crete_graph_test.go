package main

import (
	"adventofcode/day16/graph"
	"testing"
)

func TestTrivial(t *testing.T) {
	rep := "###\r\n###\r\n###"
	g := createGraph(rep)
	expected := graph.Init()
	if !g.Equal(expected) {
		t.Errorf("no nodes should not result in any but the trivial graph")
	}
}

func TestSingularNode(t *testing.T) {
	rep := "###\r\n#.#\r\n###"
	g := createGraph(rep)
	expected := graph.Init()
	if !g.Equal(expected) {
		t.Errorf("examine singulare node case")
	}
}

func TestTwoNodes(t *testing.T) {
	rep := "###\r\n#...#\r\n###"
	g := createGraph(rep)
	expected := graph.Init() // är detta hållbart?
	expected.AddConnection(graph.InitConnection(0, 1, 1000))
	expected.AddConnection(graph.InitConnection(0, 3, 1000))
	expected.AddConnection(graph.InitConnection(1, 0, 1000))
	expected.AddConnection(graph.InitConnection(1, 2, 1000))
	expected.AddConnection(graph.InitConnection(2, 1, 1000))
	expected.AddConnection(graph.InitConnection(2, 3, 1000))
	expected.AddConnection(graph.InitConnection(3, 2, 1000))
	expected.AddConnection(graph.InitConnection(3, 0, 1000))

	if !g.Equal(expected) {
		t.Errorf("examine two node case")
	}
}
