package field

import (
	vec "adventofcode/geometry/vec2d"
	"testing"
)

func TestParse(t *testing.T) {
	rep := "#.\r\n.#"
	field := Parse(rep)
	reString := field.String()
	if reString != rep {
		t.Errorf("error in parse and rep of string")
	}
}

func TestFindAllNodesPositionsSimple(t *testing.T) {
	rep := "###\r\n#.#\r\n###"
	field := Parse(rep)
	points := field.findAllNodePositions()
	if len(points) != 1 || points[0] != vec.Init(1, 1) {
		t.Errorf(" Error in find all nodes simple")
	}
}

func TestFindAllNodesTunnel(t *testing.T) {
	rep := "#####\r\n#...#\r\n#####"
	field := Parse(rep)
	points := field.findAllNodePositions()
	if len(points) != 2 || points[0] != vec.Init(1, 1) || points[1] != vec.Init(1, 3) {
		t.Errorf(" Error in find all nodes simple")
	}
}

func TestFindAllNodesSquare(t *testing.T) {
	rep := "#####\r\n#...#\r\n#.#.#\r\n#...#\r\n#####"
	field := Parse(rep)
	points := field.findAllNodePositions()
	if len(points) != 4 || points[0] != vec.Init(1, 1) || points[1] != vec.Init(1, 3) || points[2] != vec.Init(3, 1) || points[3] != vec.Init(3, 3) {
		t.Errorf(" Error in find all nodes simple")
	}
}

func TestIsNodeSimple(t *testing.T) {
	rep := "###\r\n#.#\r\n###"
	field := Parse(rep)
	node := field.isNode(vec.Init(1, 1))
	if !node {
		t.Errorf("misscategorization of single eempty space")
	}
}

func TestIsNodeCorner(t *testing.T) {
	rep := "#.#\r\n#..\r\n###"
	field := Parse(rep)
	node := field.isNode(vec.Init(1, 1))
	if !node {
		t.Errorf("misscategorization of corner")
	}
}

func TestFindConnection(t *testing.T) {
	rep := "###\r\n#.#\r\n###"
	field := Parse(rep)
	node := vec.Init(1, 1)
	con := field.findConnection(node, vec.Init(0, 1))
	if con.len != 0 {
		t.Errorf("no path should have connection len zero")
	}
}

func TestFindConnectionStraight(t *testing.T) {
	rep := "#####\r\n#...#\r\n#####"
	field := Parse(rep)
	node := vec.Init(1, 1)
	con := field.findConnection(node, vec.Init(0, 1))
	if con.len != 2 {
		t.Errorf("should be %d was %d ", 2, con.len)
	}
}

func TestFindAllConnectionsSquare(t *testing.T) {
	rep := "#####\r\n#...#\r\n#.#.#\r\n#...#\r\n#####"
	field := Parse(rep)
	cons := field.FindAllConnections()
	if len(cons) != 8 {
		t.Errorf("unexpected number oc connections in square case")
	}
}
