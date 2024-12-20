package set

import "testing"

func TestDiff(t *testing.T) {
	A := Init([]int{1, 2, 3})
	B := Init([]int{2})
	res := Diff(A, B)
	expected := Init([]int{1, 3})
	if !Eq(res, expected) {
		t.Errorf("diff not removing elements")
	}
}

func TestDiffRemoveNotMutualElement(t *testing.T) {
	A := Init([]int{1})
	B := Init([]int{2})
	res := Diff(A, B)
	expected := A
	if !Eq(res, expected) {
		t.Errorf("removal of none mutual element failed")
	}
}
