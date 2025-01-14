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
