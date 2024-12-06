package report

func IsSafeWithOneException(rep []int) bool {
	if IsSafe(rep) {
		return true
	}
	for i := range len(rep) {
		oneLess := oneLess(rep, i)
		if IsSafe(oneLess) {
			return true
		}
	}
	return false
}

func oneLess(list []int, index int) []int {
	start := list[0:index]
	end := list[index+1:]
	ret := []int{}
	ret = append(ret, start...)
	ret = append(ret, end...)
	return ret
}

func IsSafe(rep []int) bool {
	return isMonotone(rep) && isChangingWithinBond(rep)
}

func isMonotone(rep []int) bool {
	if len(rep) < 1 {
		return true
	}
	if rep[0] < rep[1] {
		return isIncreasing(rep[1:])
	}
	if rep[0] > rep[1] {
		return isDecreasing(rep[1:])
	}
	return false
}

func isChangingWithinBond(rep []int) bool {
	changeLimit := 4
	withinLimit := true
	for i := 0; i < len(rep)-1; i++ {
		withinLimit = withinLimit && (abs(rep[i]-rep[i+1]) < changeLimit)
	}
	return withinLimit
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isIncreasing(seq []int) bool {
	increasing := true
	for i := 0; i < len(seq)-1; i++ {
		increasing = increasing && (seq[i]-seq[i+1]) < 0
	}
	return increasing
}

func isDecreasing(seq []int) bool {
	decreasing := true
	for i := 0; i < len(seq)-1; i++ {
		decreasing = decreasing && (seq[i]-seq[i+1]) > 0
	}
	return decreasing
}
