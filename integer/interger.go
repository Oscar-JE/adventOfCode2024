package integer

func NumberOfDigits(num int) int {
	numDevisionBy10 := 0
	for num/10 != 0 {
		numDevisionBy10++
		num = num / 10
	}
	return numDevisionBy10 + 1
}

func Max(a int, b int) int {
	if a <= b {
		return b
	}
	return a
}

func Positive(a int) int {
	return Max(a, -a)
}

func remainder(dividend int, divider int) int {
	if dividend < 0 || divider < 0 {
		return 0
	}
	if divider == 0 {
		return dividend
	}
	n := dividend / divider
	return dividend - n*divider
}

func GCD(a int, b int) int {
	if a < b {
		return GCD(b, a)
	}
	if b == 0 {
		return a
	}
	remainder := remainder(a, b)
	return GCD(b, remainder)
}

func ABS(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
