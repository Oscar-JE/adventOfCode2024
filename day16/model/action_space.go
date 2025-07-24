package model

type Action int

const (
	turnUp   Action = 0
	turnDown Action = 1
	forward  Action = 2
)

var ActionSpace []Action = []Action{forward, turnDown, turnUp}

func (a Action) Stirng() string {
	if a == turnUp {
		return "turn up"
	} else if a == turnDown {
		return "turn down"
	}
	return "forward"
}
