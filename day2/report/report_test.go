package report

import (
	"fmt"
	"testing"
)

func TestOneLessFirst(t *testing.T) {
	entries := []int{1, 2, 3}
	firstGone := oneLess(entries, 0)
	fmt.Println(firstGone)
	midGone := oneLess(entries, 1)
	fmt.Println(midGone)
	lastGone := oneLess(entries, 2)
	fmt.Println(lastGone)
}
