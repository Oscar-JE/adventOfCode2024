package field

type loop struct {
	links []link
}

func (l loop) isLoop() bool {
	if len(l.links) == 0 {
		return false
	}
	firstLink := l.links[0]
	lastLink := l.lastLink()
	return firstLink.start == lastLink.end
}

func (l loop) lastLink() link {
	return l.links[len(l.links)-1]
}

func (l *loop) append(next link) bool {
	last := l.lastLink()
	if last.end == next.start {
		l.links = append(l.links, next)
		return true
	}
	return false
}

func (l loop) nrSides() int {
	if len(l.links) == 0 {
		return 0
	}
	nrSides := 0
	last := l.links[len(l.links)-1].direction()
	first := l.links[0].direction()
	if last != first {
		nrSides++
	}
	for i := 0; i < len(l.links)-1; i++ {
		current := l.links[i].direction()
		next := l.links[i+1].direction()
		if current != next {
			nrSides++
		}
	}
	return nrSides
}
