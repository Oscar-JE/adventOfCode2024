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
	// här kanske vi behöver tänka till kring eftersom  vi använder slices
	last := l.lastLink()
	if last.end == next.start {
		l.links = append(l.links, next)
		return true
	}
	return false
}

func (l loop) antalSidor() int { // jo men dethär blir nog rimligt
	return 5
}
