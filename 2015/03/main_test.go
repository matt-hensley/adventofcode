package main

import (
	"testing"
)


type TestCase struct {
  input string
  want int
}

func Test_part1(t *testing.T) {
  test_cases := []TestCase {
    {">", 2},
    {"^>v<", 4},
    {"^v^v^v^v^v", 2},
  }

  for _, test_case := range test_cases {
    want := test_case.want
    got := part1(test_case.input)

    if got != want {
		  t.Errorf("part1_sample %s = %d; want %d", test_case.input, got, want)
	  }    
  }
}

func Test_part2(t *testing.T) {
  test_cases := []TestCase {
    {"^v", 3},
    {"^>v<", 3},
    {"^v^v^v^v^v", 11},
  }

  for _, test_case := range test_cases {
    want := test_case.want
    got := part2(test_case.input)

    if got != want {
		  t.Errorf("part2_sample %s = %d; want %d", test_case.input, got, want)
	  }    
  }
}