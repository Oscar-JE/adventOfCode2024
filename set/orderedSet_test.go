package set

import "testing"

type intCompare struct{}

func (c intCompare) LessOrEqaul(i int, j int) bool {
	return i <= j
}

func TestInsert(t *testing.T) {
	s := InitOrdered([]int{1, 3, 2}, intCompare{})
	if s.elements[1] != 2 {
		t.Errorf("faulty initialization")
	}
}

func TestInsertLonger(t *testing.T) {
	s := InitOrdered([]int{1, 3, 2, 5, 4, 9, 7, 6, 8}, intCompare{})
	for i := range len(s.elements) {
		if s.elements[i] != i+1 {
			t.Errorf("stanger danger")
		}
	}
}

func TestInsertAtIndex1(t *testing.T) {
	s := OrderedSet[int]{}
	s.elements = []int{1, 3, 5, 7}
	s.insertAtIndex(9, 1)
	if s.elements[1] != 9 {
		t.Errorf("wtf 9")
	}
}

func TestInsertAtIndex2(t *testing.T) {
	s := OrderedSet[int]{}
	s.elements = []int{1, 3, 5, 7}
	s.insertAtIndex(9, 0)
	if s.elements[0] != 9 {
		t.Errorf("wtf 9")
	}
}
