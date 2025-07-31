package cpu

import "adventofcode/day17/cpu/outid"

type InstructionAndInput interface {
	ResultStore() outid.OutId
	Out() int
}
