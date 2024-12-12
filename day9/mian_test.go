package main

import "testing"

func TestParse(t *testing.T) {
	num := parseFromString("2333133121414131402")
	if len(num) < 10 {
		t.Errorf("wrong")
	}
}

func TestEndToEnd(t *testing.T) {
	num := parseFromString("2333133121414131402")
	res := part1(num)
	if res != 1928 {
		t.Errorf("wtf")
	}
}

func TestEndToEndPart2(t *testing.T) {
	num := parseFromString("2333133121414131402")
	res := part2(num)
	if res != 2858 {
		t.Errorf("wtf")
	}
}
