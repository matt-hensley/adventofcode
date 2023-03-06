package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1: ", part1(input))
  fmt.Println("Part 2: ", part2(input))
}

const (
  Up = rune('(')
  Down = rune(')')
)

func part1(input string) int {
  floor := 0

  for _, c := range input {
    if c == Up {
      floor++
    } else if c == Down {
      floor--
    } else {
      panic("invalid character")
    }
  }
  
  return floor
}

func part2(input string) int {
  floor := 0

  for i, c := range input {
    if c == Up {
      floor++
    } else if c == Down {
      floor--
    } else {
      panic("invalid character")
    }

    if floor == -1 {
      return i + 1
    }
  }
  
  panic("solution not found")
}
