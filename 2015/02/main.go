package main

import (
	_ "embed"
	"fmt"
  "regexp"
  "sort"
  "strconv"
  "strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1: ", part1(input))
  fmt.Println("Part 2: ", part2(input))
}

func part1(input string) int {
  lines := strings.Split(input, "\n")
  re := regexp.MustCompile(`(\d+)x(\d+)x(\d+)`)
  sum := 0

  for _, line := range lines {
    matches := re.FindStringSubmatch(line)
    l, _ := strconv.Atoi(matches[1])
    w, _ := strconv.Atoi(matches[2])
    h, _ := strconv.Atoi(matches[3])
    side1 := l * w
    side2 := w * h
    side3 := h * l
    sides := []int{side1, side2, side3}
    sort.Slice(sides, func(i, j int) bool {
		  return sides[i] < sides[j]
	  })
    area := 2 * side1 + 2 * side2 + 2 * side3
    extra := sides[0]
    sum += area + extra
  }
  
  return sum
}

func part2(input string) int {
  lines := strings.Split(input, "\n")
  re := regexp.MustCompile(`(\d+)x(\d+)x(\d+)`)
  sum := 0

  for _, line := range lines {
    matches := re.FindStringSubmatch(line)
    l, _ := strconv.Atoi(matches[1])
    w, _ := strconv.Atoi(matches[2])
    h, _ := strconv.Atoi(matches[3])
    sides := []int{l, w, h}
    sort.Slice(sides, func(i, j int) bool {
		  return sides[i] < sides[j]
	  })
    area := (sides[0] * 2) + (sides[1] * 2)
    extra := l * w * h
    sum += area + extra
  }
  
  return sum
}
