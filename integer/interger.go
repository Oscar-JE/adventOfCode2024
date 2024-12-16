package integer

func NumberOfDigits(num int) int {
	numDevisionBy10 := 0
	for num/10 != 0 {
		numDevisionBy10++
		num = num / 10
	}
	return numDevisionBy10 + 1
}
