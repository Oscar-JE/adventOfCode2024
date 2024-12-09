package evaluator

func Eval(res int, operands []int) bool {
	original := operands[0]
	return evalRecur(res, original, operands[1:])
}

func evalRecur(res int, original int, operands []int) bool {
	if len(operands) == 0 {
		return res == original
	}
	addValue := original + operands[0]
	mulValue := original * operands[0]
	appendValue := appendOperator(original, operands[0])
	return evalRecur(res, addValue, operands[1:]) || evalRecur(res, mulValue, operands[1:]) || evalRecur(res, appendValue, operands[1:])
}

func appendOperator(a int, b int) int {
	exponent := numberOfPowerOften(b) + 1
	return a*toThePower(10, exponent) + b
}

func numberOfPowerOften(val int) int {
	sum := 0
	val = val / 10
	for val > 0 {
		sum++
		val = val / 10
	}
	return sum
}

func toThePower(base int, exponent int) int {
	res := 1
	for i := 0; i < exponent; i++ {
		res *= base
	}
	return res
}
