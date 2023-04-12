package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}

// TIL https://stackoverflow.com/a/49304209
func part1(input string) int {
	lines := strings.Split(input, "\n")
	str := 0
	mem := 0

	for _, line := range lines {
		s, err := strconv.Unquote(line)

		if err != nil {
			panic(err)
		}

		str += len(line)
		mem += len(s)
	}

	return str - mem
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	str := 0
	mem := 0

	for _, line := range lines {
		str += len(line)
		mem += len(strconv.Quote(line))
	}

	return mem - str
}
