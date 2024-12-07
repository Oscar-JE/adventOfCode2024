package matrix

import "testing"

func TestInitIntMatrix(t *testing.T) {
	m := Init([]int{1, 2, 3, 4, 5, 6}, 3, 2)
	res := m.Get(2, 0)
	if res != 5 {
		t.Errorf(" Initialization test unexpected res")
	}
	res = m.Get(1, 1)
	if res != 4 {
		t.Errorf("get (1,1) unexpected")
	}
}
