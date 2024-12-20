package set

type nothing struct {
}

type Set[E comparable] struct {
	setRep map[E]nothing
}

func Init[E comparable](elements []E) Set[E] {
	setRep := map[E]nothing{}
	for _, el := range elements {
		setRep[el] = nothing{}
	}
	return Set[E]{setRep: setRep}
}

func (s Set[E]) GetElements() []E {
	elements := []E{}
	for elem := range s.setRep {
		elements = append(elements, elem)
	}
	return elements
}

func Eq[E comparable](A Set[E], B Set[E]) bool {
	if len(A.setRep) == len(B.setRep) {
		equal := true
		for el := range A.setRep {
			if !B.Has(el) {
				equal = false
				break
			}
		}
		return equal
	}
	return false
}

func (A Set[E]) Has(el E) bool {
	_, has := A.setRep[el]
	return has
}

func (A *Set[E]) Add(el E) {
	A.setRep[el] = nothing{}
}

func Diff[E comparable](A Set[E], B Set[E]) Set[E] {
	diff := Init(A.GetElements())
	for toRemove := range B.setRep {
		delete(diff.setRep, toRemove)
	}
	return diff
}

func (A Set[E]) IsEmpty() bool {
	return len(A.setRep) == 0
}

func (A Set[E]) GetAnElement() E {
	if A.IsEmpty() {
		panic("Cant get elements from empty set")
	}
	elems := A.GetElements()
	return elems[0]
}

func (A Set[E]) Cardinality() int {
	return len(A.setRep)
}
