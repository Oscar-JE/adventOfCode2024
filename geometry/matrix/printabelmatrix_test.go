package matrix

import (
	"fmt"
	"testing"
)

type betterint int

func TestPrint(t *testing.T) {
	mat := Init([]int{1, 2, 3, 4}, 2, 2)
	res := fmt.Sprint(mat)
	if res != "12\r\n34" {
		t.Errorf("not correct string output")
	}
}

func TestPrintableMatrix(t *testing.T) {

}
