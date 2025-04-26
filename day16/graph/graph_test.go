package graph

import "testing"

func TestGraphAddConnection(t *testing.T) {
	g := Init()
	g.AddConnection(connection{3, 6, 2})
	if g.nrNodes != 7 {
		t.Errorf("wrong number of nodes after adding a singel connection ")
	}
}
