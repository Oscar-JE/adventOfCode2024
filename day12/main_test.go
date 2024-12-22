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
