package towels

import (
	"fmt"
	"testing"
)

func TestMatchBeginningMatch(t *testing.T) {
	p1 := ParsePattern("brwrr")
	p2 := ParsePattern("br")
	index, match := p1.matchBeginning(p2)
	if !match || index != 2 {
		t.Errorf("Match beginning on success does not match expectations")
	}
}

func TestMatchBeginningNoMatch(t *testing.T) {
	p1 := ParsePattern("vvvddss")
	p2 := ParsePattern("vvvdda")
	_, match := p1.matchBeginning(p2)
	if match {
		t.Errorf("Match beginning false positive")
	}
}

func TestParseAndStringify(t *testing.T) {
	p := ParsePattern("hej")
	rep := fmt.Sprint(p)
	if rep != "hej" {
		t.Errorf("whalabubu")
	}
}
