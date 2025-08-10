package containers

type PriorityQueue[E comparable] struct {
	queue []elementAndPriority[E]
}

func InitPriorityQueue[E comparable]() PriorityQueue[E] {
	return PriorityQueue[E]{queue: []elementAndPriority[E]{}}
}

type elementAndPriority[E comparable] struct {
	element  E
	priority int
}

func (p *PriorityQueue[E]) Add(el E, priority int) {
	var internalElement elementAndPriority[E] = elementAndPriority[E]{el, priority}
	var indexOfClosestSmaller = p.findPriorityIndex(priority)
	p.insert(internalElement, indexOfClosestSmaller)
}

func (p PriorityQueue[E]) findPriorityIndex(priority int) int { // todo skriv om som binärsök
	if len(p.queue) == 0 {
		return 0
	}
	if len(p.queue) == 1 {
		if priority <= p.queue[0].priority {
			return 0
		} else {
			return 1
		}
	}
	i := 0
	for i < len(p.queue) {
		next := p.queue[i].priority
		if next >= priority {
			break
		}
		i++
	}
	return i
}

func (p *PriorityQueue[E]) insert(el elementAndPriority[E], index int) {
	if index >= len(p.queue) {
		p.queue = append(p.queue, el)
		return
	}
	before := p.queue[:index]
	after := p.queue[index:]
	newQ := []elementAndPriority[E]{}
	newQ = append(newQ, before...)
	newQ = append(newQ, el)
	newQ = append(newQ, after...)
	p.queue = newQ
}

func (p *PriorityQueue[E]) PopSmallestPriority() (E, int) {
	first := p.queue[0]
	p.queue = p.queue[1:]
	return first.element, first.priority
}

func (p PriorityQueue[E]) GetPriority(element E) (int, bool) {
	for _, v := range p.queue {
		if v.element == element {
			return v.priority, true
		}
	}
	return -1, false
}

func (p *PriorityQueue[E]) UppdatePriority(element E, priority int) {
	indexOfElement := 0
	for indexOfElement < len(p.queue) {
		if p.queue[indexOfElement].element == element {
			break
		}
		indexOfElement++
	}
	p.remove(indexOfElement)
	p.Add(element, priority)
}

func (p *PriorityQueue[E]) remove(index int) { // denna är mycket relevant att unittesta pga slice konstigheterna
	newQ := []elementAndPriority[E]{}
	newQ = append(newQ, p.queue[:index]...)
	newQ = append(newQ, p.queue[index+1:]...)
	p.queue = newQ // borde vara en otroligt långsam implementering men bör fungera
}

func (p PriorityQueue[E]) Empty() bool {
	return len(p.queue) == 0
}
