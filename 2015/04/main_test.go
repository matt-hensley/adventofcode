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
    {"abcdef", 609043},
    {"pqrstuv", 1048970},
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