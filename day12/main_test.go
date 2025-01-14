package main

import "testing"

func TestMain(t *testing.T) {
	main()
}

func TestTroublingCase(t *testing.T) {
	f := parse("short.txt")
	idBlob := f.FindPlots("R")
	res := f.ScoreOfIdBlob(idBlob)
	if res != 120 {
		t.Errorf("something went wrong")
	}
}

func TestExample1(t *testing.T) {
	f := parse("short1.txt")
	res := f.Score()
	if res != 80 {
		t.Errorf("was %d , should be %d", res, 80)
	}
}

func TestExample3(t *testing.T) {
	f := parse("short3.txt")
	res := f.Score()
	if res != 368 {
		t.Errorf(" was %d , should be %d", res, 368)
	}
}

func TestReducedExample(t *testing.T) {
	f := parse("short3reduced.txt")
	res := f.Score()
	expected := 8*8 + 4
	if res != expected {
		t.Errorf(" was %d , should be %d", res, expected)
	}
}
