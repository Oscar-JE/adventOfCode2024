package diskfrag

func CheckSum(num []int) int {
	checkSum := 0
	for i, el := range num {
		checkSum += i * el
	}
	return checkSum

}
