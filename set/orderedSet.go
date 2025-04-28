package set

func InitOrdered[E comparable](elements []E, comparer LessThan[E]) OrderedSet[E] {
	so := OrderedSet[E]{}
	so.comp = Comparator[E]{comparer}
	for _, el := range elements {
		so.Insert(el)
	}
	return so
}

type OrderedSet[E comparable] struct {
	elements []E
	comp     Comparator[E]
}

func (o *OrderedSet[E]) Insert(element E) {
	_, insertionIndex := o.findIndex(element)
	o.insertAtIndex(element, insertionIndex)
}

func (o *OrderedSet[E]) insertAtIndex(element E, index int) {
	tail := o.elements[index:]
	head := o.elements[:index]
	res := []E{}
	res = append(res, head...)
	res = append(res, element)
	res = append(res, tail...)
	o.elements = res
}

func (o OrderedSet[E]) findIndex(element E) (bool, int) {
	if len(o.elements) == 0 {
		return false, 0
	}
	lower := 0
	upper := len(o.elements)
	mid := 0
	for lower < upper {
		mid = (lower + upper) / 2
		midEl := o.elements[mid]
		if o.comp.Equal(midEl, element) {
			lower = mid
			break
		} else if o.comp.Less(midEl, element) {
			lower = mid + 1
		} else {
			upper = mid
		}
	}
	found := false
	if 0 <= lower && lower < len(o.elements) {
		found = o.elements[lower] == element
	}
	return found, lower
}

func (o OrderedSet[E]) Has(element E) bool {
	found, _ := o.findIndex(element)
	return found
}

func (o OrderedSet[E]) AsSlice() []E {
	return o.elements
}

func (o OrderedSet[E]) Equal(other OrderedSet[E]) bool {
	if len(o.elements) != len(other.elements) {
		return false
	}
	for i, el := range o.elements {
		if !o.comp.Equal(el, other.elements[i]) {
			return false
		}
	}
	return true
}
