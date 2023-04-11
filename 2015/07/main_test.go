package main

import (
	_ "embed"
	"testing"
)

//go:embed sample.txt
var sample string

type TestCase struct {
	input string
	want  int
}

func Test_part1(t *testing.T) {
	test_cases := []TestCase{
		{"d", 72},
		{"e", 507},
		{"f", 492},
		{"g", 114},
		{"h", 65412},
		{"i", 65079},
		{"x", 123},
		{"y", 456},
	}

	for _, test_case := range test_cases {
		want := test_case.want
		got := part1(sample, test_case.input)

		if got != want {
			t.Errorf("part1_sample %s = %d; want %d", test_case.input, got, want)
		}
	}
}

func Test_part2(t *testing.T) {
	test_cases := []TestCase{
		{"", -1},
	}

	for _, test_case := range test_cases {
		want := test_case.want
		got := part2(test_case.input)

		if got != want {
			t.Errorf("part2_sample %s = %d; want %d", test_case.input, got, want)
		}
	}
}
