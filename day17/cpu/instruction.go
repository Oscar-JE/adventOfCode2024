package cpu

import "adventofcode/day17/cpu/outid"

type Instruction interface {
	ResultStore() outid.OutId
	Out() int
}
