package stones

import "adventofcode/integer"

type rule interface {
	applicable(int) bool
	apply(int) []int
}

type Zero struct {
}

func (z Zero) applicable(stone int) bool {
	return stone == 0
}

func (z Zero) apply(stone int) []int {
	return []int{stone + 1}
}

type SplitEvenNrDigits struct {
}

func (s SplitEvenNrDigits) applicable(stone int) bool {
	return integer.NumberOfDigits(stone)%2 == 0
}

func (s SplitEvenNrDigits) apply(stone int) []int {
	nrDigits := integer.NumberOfDigits(stone)
	factor := 1
	for i := 0; i < nrDigits/2; i++ {
		factor *= 10
	}
	front := stone / factor
	back := stone - front*factor
	return []int{front, back}
}

type MultiplyWith2024 struct {
}

func (m MultiplyWith2024) applicable(stone int) bool {
	return true
}

func (m MultiplyWith2024) apply(stone int) []int {
	return []int{stone * 2024}
}
