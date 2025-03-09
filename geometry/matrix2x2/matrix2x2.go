package matrix2x2

import (
	vec "adventofcode/geometry/vec2d"
	"errors"
)

type Matrix2x2 struct {
	c0 vec.Vec2d
	c1 vec.Vec2d
}

func Init(c0 vec.Vec2d, c1 vec.Vec2d) Matrix2x2 {
	return Matrix2x2{c0: c0, c1: c1}
}

func (m Matrix2x2) linearlyDependent() bool {
	return vec.Parallel(m.c0, m.c1)
}

func (m Matrix2x2) HasSolution(target vec.Vec2d) bool {
	oneSol := m.HasSingleSolution()
	multSol := m.HasMultipleSolutions(target)
	return oneSol || multSol
}

func (m Matrix2x2) HasSingleSolution() bool {
	return !m.linearlyDependent()
}

func (m Matrix2x2) HasMultipleSolutions(target vec.Vec2d) bool {
	if !m.linearlyDependent() {
		return false
	}
	if !target.NonZero() {
		return true
	}
	if !m.c0.NonZero() && !m.c1.NonZero() {
		return false
	}
	return vec.Parallel(m.c0, target) && vec.Parallel(m.c1, target)
}

func (m Matrix2x2) Solve(target vec.Vec2d) (SolutionSet, error) {
	if !m.HasSolution(target) {
		return SolutionSet{}, errors.New("the given problem has no solution")
	}
	if m.HasSingleSolution() {
		unscaledInverse, scale := m.inverse()
		unscaledAns := unscaledInverse.VecMul(target)
		ans := vec.ScalarDiv(scale, unscaledAns) // här sker en heltalsavrundning , inte bra
		return SolutionSet{aSolPoint: ans, directions: []vec.Vec2d{}}, nil
	}
	if m.HasMultipleSolutions(target) {
		if !m.c0.NonZero() && !m.c1.NonZero() {
			return SolutionSet{vec.Init(0, 0), []vec.Vec2d{vec.Init(1, 0), vec.Init(0, 1)}}, nil // detta är igentligen fel. Men Orkar vi köra med spann?
		}
		if m.c0.NonZero() {
			scale := m.c0.ScaledTo(target)
			directions := m.ZeroSpace()
			return SolutionSet{vec.Init(scale, 0), directions}, nil
		}
		scale := m.c1.ScaledTo(target)
		directions := m.ZeroSpace()
		return SolutionSet{vec.Init(0, scale), directions}, nil
	}
	return SolutionSet{}, nil
}

func (m Matrix2x2) VecMul(v vec.Vec2d) vec.Vec2d {
	return vec.Add(vec.ScalarMult(v.GetX(), m.c0),
		vec.ScalarMult(v.GetY(), m.c1))
}

func (m Matrix2x2) inverse() (Matrix2x2, int) {
	return Init(vec.Init(m.c1.GetY(), -m.c0.GetY()), vec.Init(-m.c1.GetX(), m.c0.GetX())), m.c0.GetX()*m.c1.GetY() - m.c1.GetX()*m.c0.GetY()
}

func (m Matrix2x2) ZeroSpace() []vec.Vec2d {
	if !m.linearlyDependent() {
		return []vec.Vec2d{}
	}
	if !m.c0.NonZero() && !m.c1.NonZero() {
		return []vec.Vec2d{vec.Init(1, 0), vec.Init(0, 1)}
	}
	if !m.c0.NonZero() {
		return []vec.Vec2d{vec.Init(1, 0)}
	}
	if !m.c1.NonZero() {
		return []vec.Vec2d{vec.Init(0, 1)}
	}
	scale := m.c0.ScaledTo(m.c1)
	if scale == 0 {
		scale = m.c1.ScaledTo(m.c0)
		return []vec.Vec2d{vec.Init(1, -scale)}
	}
	return []vec.Vec2d{vec.Init(scale, -1)}
}
