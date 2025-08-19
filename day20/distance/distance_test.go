package distance

import (
	"adventofcode/day20/field"
	"fmt"
	"testing"
)

func TestFindDistanceMatrixSymmetric(t *testing.T) {
	f := field.Parse("####\r\n#S.#\r\n#.E#")
	dists := DistanceToEnd(f)
	if dists.Get(0, 0) != infinite() {
		t.Errorf("stopped tiles can not reach the end")
	}
	fmt.Println(dists)
}

func TestFindDistanceMatrixNotSymmetric(t *testing.T) {
	t.Skip("visual inspection look good")
	f := field.Parse("#...#\r\n#S#.#\r\n#...#\r\n#..E#")
	dists := DistanceToEnd(f)
	fmt.Println(dists)
}
