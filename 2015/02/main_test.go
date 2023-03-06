package main

import (
  _ "embed"
	"testing"
)

//go:embed sample.txt
var sample string

func Test_part1(t *testing.T) {
	want := 58 + 43
	got := part1(sample)

	if got != want {
		t.Errorf("part1_sample = %d; want %d", got, want)
	}
}

func Test_part2(t *testing.T) {
	want := 34 + 14
	got := part2(sample)

	if got != want {
		t.Errorf("part2_sample = %d; want %d", got, want)
	}
}
