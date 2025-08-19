package field

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	t.Skip("it has served its purpose")
	input := "####\r\n#S.#\r\n#.E#"
	f := Parse(input)
	rep := fmt.Sprint(f)
	fmt.Println(rep)
}
