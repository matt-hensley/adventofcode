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
  Up = rune('^')
  Right = rune('>')
  Down = rune('v')
  Left = rune('<')
)

func part1(input string) int {
  visited := make(map[string]bool)
  x, y := 0, 0

  visited[fmt.Sprint(x, y)] = true
  
  for _, move := range input {
    if move == Up {
      y--
    } else if move == Right {
      x++
    } else if move == Down {
      y++
    } else if move == Left {
      x--
    }
    
    visited[fmt.Sprint(x, y)] = true
  }
  
  return len(visited)
}

func part2(input string) int {
  visited := make([]map[string]bool, 2)
  visited[0] = make(map[string]bool)
  visited[1] = make(map[string]bool)
  x, y := []int{0,0}, []int{0,0}

  for i, v := range visited {
    v[fmt.Sprint(x[i], y[i])] = true
  }
  
  for i, move := range input {
    santa := i % 2

    if move == Up {
      y[santa]--
    } else if move == Right {
      x[santa]++
    } else if move == Down {
      y[santa]++
    } else if move == Left {
      x[santa]--
    }

    visited[santa][fmt.Sprint(x[santa], y[santa])] = true
  }

  merged := make(map[string]bool)

  for _, v := range visited {
    for key, c := range v {
      merged[key] = c
    }
  }

  return len(merged)
}
