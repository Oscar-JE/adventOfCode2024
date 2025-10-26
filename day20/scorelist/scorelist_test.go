package scorelist

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	scores := []int{3, 3, 1, 2}
	sl := Init(scores)
	fmt.Println(sl)
}
