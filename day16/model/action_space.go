package model

type Action int

const (
	turnUp   Action = 0
	turnDown Action = 1
	forward  Action = 2
)

var ActionSpace []Action = []Action{forward, turnDown, turnUp}
