package instructions

import "testing"

func TestFromRep(t *testing.T) {
	rep := "^^><"
	instructor := InitFromRep(rep)
	res := instructor.String()
	if rep != res {
		t.Errorf("Parse from string not inverse of toString")
	}
}
