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
		unscaledAns := unscaledInverse.vecMul(target)
		ans := vec.ScalarDiv(scale, unscaledAns)
		return SolutionSet{aSolPoint: ans, direction: vec.Init(0, 0)}, nil
	}
	if m.HasMultipleSolutions(target) {
		if m.c0.NonZero() {

		}
	}
	return SolutionSet{}, nil
}

func (m Matrix2x2) vecMul(v vec.Vec2d) vec.Vec2d {
	return vec.Add(vec.ScalarMult(v.GetX(), m.c0),
		vec.ScalarMult(v.GetY(), m.c1))
}

func (m Matrix2x2) inverse() (Matrix2x2, int) {
	return Init(vec.Init(m.c1.GetY(), -m.c0.GetY()), vec.Init(-m.c1.GetX(), m.c0.GetX())), m.c0.GetX()*m.c1.GetY() - m.c1.GetX()*m.c0.GetY()
}
