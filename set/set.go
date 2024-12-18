package set

type Set[E comparable] struct {
	setRep map[E]bool
}

func Init[E comparable](elements []E) Set[E] {
	return Set[E]{map[E]bool{}}
}

func (s Set[E]) GetElements() []E {
	elements := []E{}
	for elem, in := range s.setRep {
		if in {
			elements = append(elements, elem)
		}
	}
	return elements
}

func Eq[E comparable](A Set[E], B Set[E]) bool {
	elemsA := A.GetElements()
	elemsB := B.GetElements()

	if len(elemsA) == len(elemsB) {
		equal := true
		for i := range elemsA {
			if elemsA[i] != elemsB[i] {
				equal = false
				break
			}
		}
		return equal
	}
	return false
}
