package vec

import (
	"adventofcode/integer"
	"strconv"
	"strings"
)

type Vec2d struct {
	x int
	y int
}

func Init(x int, y int) Vec2d {
	return Vec2d{x, y}
}

func (v Vec2d) GetX() int {
	return v.x
}

func (v Vec2d) GetY() int {
	return v.y
}

func Add(v1 Vec2d, v2 Vec2d) Vec2d {
	return Init(v1.x+v2.x, v1.y+v2.y)
}

func Subtract(v1 Vec2d, v2 Vec2d) Vec2d {
	return Init(v1.x-v2.x, v1.y-v2.y)
}

func ScalarMult(a int, v Vec2d) Vec2d {
	return Init(v.x*a, v.y*a)
}

func ScalarDiv(a int, v Vec2d) Vec2d {
	return Init(v.x/a, v.y/a)
}

func DotProduct(a Vec2d, b Vec2d) int {
	return a.x*b.x + a.y*b.y
}

func DistSquared(a Vec2d, b Vec2d) int {
	return AbsSquared(Subtract(a, b))
}

func Turn90Down(v Vec2d) Vec2d {
	return Init(v.y, -v.x)
}

func AbsSquared(v Vec2d) int {
	return v.x*v.x + v.y*v.y
}

func Parallel(v1 Vec2d, v2 Vec2d) bool {
	dotProd := DotProduct(v1, v2)
	projectedLen := integer.Positive(dotProd)
	return projectedLen*projectedLen == AbsSquared(v1)*AbsSquared(v2)
}

func (v Vec2d) NonZero() bool {
	return integer.Positive(v.x) > 0 || integer.Positive(v.y) > 0
}

func (v Vec2d) ScaledTo(other Vec2d) int {
	if !v.NonZero() {
		return 0
	}
	if v.x != 0 {
		return other.x / v.x
	}
	return other.y / v.y
}

func (v Vec2d) FirstQuadrant() bool {
	return v.x >= 0 && v.y >= 0
}

func Parse(rep string, delimitor string) Vec2d {
	numbersAndExtras := strings.Split(rep, delimitor)
	firstCoordAndExtras := numbersAndExtras[0]
	secondCoordAndExtras := numbersAndExtras[1]
	firstCoordrep := stripAwayNotNumbers(firstCoordAndExtras)
	secondCoordrep := stripAwayNotNumbers(secondCoordAndExtras)
	num0, err0 := strconv.Atoi(firstCoordrep)
	num1, err1 := strconv.Atoi(secondCoordrep)
	if err0 != nil || err1 != nil {
		panic("conversion of coordinates failed")
	}
	return Init(num0, num1)
}

func stripAwayNotNumbers(in string) string {
	outStr := ""
	for _, ch := range in {
		_, err := strconv.Atoi(string(ch))
		if err == nil || ch == '-' {
			outStr = outStr + string(ch)
		}
	}
	return outStr
}
